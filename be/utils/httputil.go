package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    message,
	})
}

func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": 500,
		"msg":    message,
	})
}

func RespUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status": 401,
		"msg":    message,
	})
}

func RespForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{
		"status": 403,
		"msg":    message,
	})
}

func RespJWT(c *gin.Context, code int32, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    message,
	})
}
