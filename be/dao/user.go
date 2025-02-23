package dao

import (
	"demo/model"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// type database map[string]string

// 使用bcrypt对password加密
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error hashing password!")
		return "", err
	}
	return string(hashedPassword), nil
}

// 验证password

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
}

func AddUser(username, password, email, phone, role string) error {

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	user := model.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
		Phone:    phone,
		Role:     role,
	}

	res := db.Create(&user)
	if res.Error != nil {
		log.Printf("Error loading database: %v\n", res.Error)
		return res.Error
	}

	return nil
}

// 删除某个用户
func DeleteUser(id uint64) error {
	var user model.User
	res := db.Where("id = ?", id).Delete(&user)
	if res.Error != nil {
		log.Printf("error delete user: %v", res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.Printf("no rows were affected by the delete operation for user: %d", id)
		return errors.New("no such user found to delete")
	}
	return nil
}

// 判断相关信息是否存在
func UserRecordExists(field string, value interface{}) (bool, error) {
	var user model.User
	res := db.Model(&user).Where(fmt.Sprintf("%s = ?", field), value).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		log.Printf("error checking existence for field %s: %v", field, res.Error)
		return false, res.Error
	}
	return !errors.Is(res.Error, gorm.ErrRecordNotFound), nil
}

// 查找用户名
func SelectUsername(username string) (*model.User, error) {
	var user model.User
	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 或者返回一个具体的错误，比如 userNotFoundError
		}
		log.Printf("error select user: %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

// 查找用户信息
func SelectUserInfo(username string, field string) (string, error) {
	var user model.User
	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		log.Printf("error select user: %v", res.Error)
		return "", res.Error
	}
	switch field {
	case "Role":
		return user.Role, nil
	case "Id":
		return strconv.Itoa(user.Id), nil
	case "Password":
		return user.Password, nil
	case "Phone":
		return user.Phone, nil
	case "Email":
		return user.Email, nil
	default:
		return "", fmt.Errorf("unknow filed name: %v", field)
	}
}

func UpdateUserPassword(username, newPassword string) error {
	var user model.User
	res := db.Where("username = ?", username).First(&user)

	if res.Error != nil {
		log.Printf("Error update password: %v\n", res.Error)
		return res.Error
	}

	hashedNewPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新密码
	user.Password = hashedNewPassword
	// 保存密码
	res = db.Save(&user)
	if res.Error != nil {
		log.Printf("Error save password: %v\n", res.Error)
		return res.Error
	}
	return nil
}

func UpdateUserRole(id uint64, role string) error {
	var user model.User
	res := db.Where("id = ?", id).Find(&user)
	if res.Error != nil {
		log.Printf("error find user with id: %d, %v", id, res.Error)
		return res.Error
	}

	// 更新用户角色
	user.Role = role

	// 保存用户信息
	res = db.Save(&user)
	if res.Error != nil {
		log.Printf("Error updating user role: %v", res.Error)
		return res.Error
	}

	return nil

}

func SelectAllUserInfo() ([]model.User, error) {
	var users []model.User
	res := db.Find(&users)
	if res.Error != nil {
		log.Println("error searching users! %v", res.Error)
		return nil, res.Error
	}
	return users, nil
}

func SelectUsernameFromId(id uint64) (string, error) {
	var user model.User
	res := db.Where("id = ?", id).First(&user)
	if res.Error != nil {
		log.Printf("error select user: %v", res.Error)
		return "", res.Error
	}
	return user.Username, nil
}

func SelectSingleUserInfo(username string) (model.User, error) {
	var user model.User
	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		log.Printf("error select user: %v", res.Error)
		return user, res.Error
	}
	return user, nil
}

func AddUserAlatar(username, alatar string) error {
	var user model.User
	res := db.Model(&user).Where("username = ?", username).Update("alatar", alatar)
	if res.Error != nil {
		log.Printf("error add user: %v", res.Error)
		return res.Error
	}
	return nil
}

func GetAlatarFromUsername(username string) (string, error) {
	var user model.User
	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		log.Printf("error get user: %v", res.Error)
		return "", res.Error
	}
	return user.Alatar, nil
}
