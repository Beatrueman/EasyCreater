package middleware

import (
	"demo/dao"
	"demo/utils"
	"github.com/gin-gonic/gin"
)

// 检查是否为管理员
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, ok := c.Get("username")
		if !ok {
			utils.RespUnauthorized(c, "Unauthorized")
			c.Abort() // 停止后续处理
			return
		}

		user, err := dao.SelectUsername(username.(string))
		if err != nil {
			utils.RespFail(c, "Internal Server Error")
			c.Abort()
			return
		}

		// 检查用户的角色
		if user.Role != "admin" {
			utils.RespForbidden(c, "You are not admin! No permissions!")
			c.Abort()
			return
		}

		// 用户有权限，则继续执行
		c.Next()
	}
}
