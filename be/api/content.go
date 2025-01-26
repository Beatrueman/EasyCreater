package api

import (
	"demo/dao"
	"demo/model"
	"demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func PostContent(c *gin.Context) {
	usernameInterface, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Invalid or missing token")
		return
	}

	username := usernameInterface.(string)

	user, err := dao.SelectUsername(username)
	if err != nil {
		utils.RespFail(c, "Failed to find user!")
		return
	}

	content := c.PostForm("content")
	if content == "" {
		utils.RespFail(c, "Content is required!")
		return
	}

	msg := &model.Content{
		UserID:    user.Id,
		Content:   content,
		Timestamp: time.Now(),
	}

	err = dao.AddContent(msg)
	if err != nil {
		utils.RespFail(c, "Failed to add content!")
		return
	}

	utils.RespSuccess(c, "Message added successfully!")
}

func GetContent(c *gin.Context) {
	userIDstr := c.Param("userId")
	log.Printf("Received user_id: %s\n", userIDstr)

	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		log.Printf("Invalid user_id: %s, error: %v\n", userIDstr, err)
		utils.RespFail(c, "Invalid user_id")
		return
	}

	content, err := dao.GetContentByUserID(uint(userID))
	if err != nil {
		log.Printf("Failed to get content for user_id: %d, error: %v\n", userID, err)
		utils.RespFail(c, "Failed to get content!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    content,
	})
}
