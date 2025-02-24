package api

import (
	"demo/api/middleware"
	"demo/dao"
	"demo/model"
	"demo/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespFail(c, "Verification failed!")
		return
	}

	username := user.Username
	password := user.Password
	email := user.Email
	phone := user.Phone

	// 打印传入的参数值
	log.Printf("Received data: username=%s, email=%s, phone=%s", username, email, phone)

	if email == "" || phone == "" {
		utils.RespFail(c, "email and phone are required!")
		return
	}

	// 验证用户名是否重复
	flag, err := dao.UserRecordExists("username", username)
	if err != nil {
		utils.RespFail(c, "Error checking user existence")
		return
	}

	if flag {
		utils.RespFail(c, "This user is already exists!")
		return
	}

	// 验证电话是否重复
	isPhoneExist, err := dao.UserRecordExists("phone", phone)
	isEmailExist, err := dao.UserRecordExists("email", email)

	if err != nil {
		utils.RespFail(c, "Error checking phone or email existence")
		return
	}

	if isPhoneExist || isEmailExist {
		utils.RespFail(c, "This phone or email is already in use!")
		return
	}

	// 设置默认角色为“user”
	role := "user"

	err = dao.AddUser(username, password, email, phone, role)
	if err != nil {
		utils.RespFail(c, "Error adding user......")
		return
	}

	utils.RespSuccess(c, "add user successfully!")
}

func login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespFail(c, "Verification failed!")
		return
	}

	username := user.Username
	password := user.Password

	flag, err := dao.UserRecordExists("username", username)
	if err != nil {
		utils.RespFail(c, "Error checking user existence")
		return
	}

	if !flag {
		utils.RespFail(c, "user doesn't exists")
		return
	}

	selectPassword, err := dao.SelectUserInfo(username, "Password")
	if err != nil {
		utils.RespFail(c, "Error retrieving password")
		return
	}
	if err := dao.CheckPasswordHash(selectPassword, password); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			utils.RespFail(c, "Wrong password!")
		default:
			utils.RespFail(c, "error comparing password!")
		}
		return
	}

	role, err := dao.SelectUserInfo(username, "Role")
	if err != nil {
		utils.RespFail(c, "Error selecting this user's role!")
		log.Println(err)
		return
	}
	log.Printf("Role for user %s: %s", username, role) // 添加日志

	claim := model.MyClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(), // jwt-token过期时间
			Issuer:    "yiiong",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	if err != nil {
		utils.RespFail(c, "Error generating token")
		return
	}
	utils.RespSuccess(c, tokenString)
}

func deleteUser(c *gin.Context) {
	userIDstr := c.Param("userId")
	log.Printf("Received user_id: %s\n", userIDstr)

	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		log.Printf("Invalid user_id: %s, error: %v\n", userIDstr, err)
		utils.RespFail(c, "Invalid user_id")
		return
	}

	err = dao.DeleteUser(userID)
	if err != nil {
		utils.RespFail(c, "error delete this user!")
		return
	}
	utils.RespSuccess(c, "delete successfully!")
}

func getUserInfoFromToken(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}
	role, ok := c.Get("role")
	if !ok {
		utils.RespFail(c, "Role not found in context")
		return
	}

	user, err := dao.SelectSingleUserInfo(username.(string))
	if err != nil {
		utils.RespFail(c, "Failed to find user!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"username": username,
		"role":     role,
		"email":    user.Email,
		"phone":    user.Phone,
	})
}

func ChangeUserPassword(c *gin.Context) {
	var req model.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespFail(c, "Invalid request body")
		log.Println("error binding JSON:", err)
		return
	}

	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}

	oldPassword := req.Password
	newPassword := req.NewPassword

	if newPassword == "" {
		utils.RespFail(c, "newPassword is missing!")
		return
	}

	selectPassword, err := dao.SelectUserInfo(username.(string), "Password")
	if err != nil {
		utils.RespFail(c, "Error retrieving password")
		log.Println("Error retrieving password:", err)
		return
	}

	if err := dao.CheckPasswordHash(selectPassword, oldPassword); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			utils.RespFail(c, "Wrong password!")
			return
		default:
			utils.RespFail(c, "error comparing password!")
		}
	}

	if oldPassword == newPassword {
		utils.RespFail(c, "new password is equal to old password!")
		return
	}

	err = dao.UpdateUserPassword(username.(string), newPassword)
	if err != nil {
		utils.RespFail(c, "Failed to update password!")
		log.Println("Failed to update password:", err)
		return
	}
	utils.RespSuccess(c, "Password updated successfully!")
}

func ChangeUserRole(c *gin.Context) {
	userIDstr := c.Param("userId")
	log.Printf("Received user_id: %s\n", userIDstr)

	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		log.Printf("Invalid user_id: %s, error: %v\n", userIDstr, err)
		utils.RespFail(c, "Invalid user_id")
		return
	}

	role := c.PostForm("role")
	if role != "user" && role != "admin" {
		utils.RespFail(c, "Invalid user role! It should be either 'user' or 'admin'.")
		return
	}

	err = dao.UpdateUserRole(userID, role)
	if err != nil {
		utils.RespFail(c, "error change user's role!")
		return
	}

	utils.RespSuccess(c, "User role updated successfully")
}

func getAllUserInfo(c *gin.Context) {
	users, err := dao.SelectAllUserInfo()
	if err != nil {
		utils.RespFail(c, "Failed to search users!")
		return
	}
	c.JSON(http.StatusOK, users)
}

func AddUserAvatar(c *gin.Context) {
	var req struct {
		Avatar string `json:"avatar"`
	}

	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespFail(c, "Invalid request body")
		log.Println("error binding JSON:", err)
		return
	}

	avatar := req.Avatar
	err := dao.AddUserAvatar(username.(string), avatar)
	if err != nil {
		utils.RespFail(c, "Failed to add user avatar!")
		log.Println("Failed to add user avatar:", err)
		return
	}
	log.Println("add avatar from", username)
	utils.RespSuccess(c, "Added user avatar successfully!")
}

func GetUserAvatar(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}
	avatar, err := dao.GetAvatarFromUsername(username.(string))
	if err != nil {
		utils.RespFail(c, "Failed to get user avatar!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "User avatar retrieved successfully!",
		"data":   avatar,
		"status": 200,
	})
}
