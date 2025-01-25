package middleware

import (
	"demo/model"
	"demo/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

var Secret = []byte("yiiong")

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespJWT(c, 2003, "请求头中auth为空")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.RespJWT(c, 2004, "请求头中auth格式有误")

			c.Abort()
			return
		}

		mc, err := ParseToken(parts[1])
		if err != nil {
			utils.RespJWT(c, 2005, "无效的token")
			c.Abort()
			return
		}

		log.Printf("Parsed Role: %s", mc.Role)

		c.Set("username", mc.Username)
		c.Set("role", mc.Role)

		c.Next()

	}
}

func ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
