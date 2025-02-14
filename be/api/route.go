package api

import (
	"demo/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register)
	r.POST("/login", login)

	AdminRouter := r.Group("/auth")
	{
		AdminRouter.Use(middleware.JWTAuthMiddleware(), middleware.AdminAuthMiddleware())
		AdminRouter.GET("/search", getAllUserInfo)
		AdminRouter.DELETE("/delete/:userId", deleteUser)
		AdminRouter.POST("/role/:userId", ChangeUserRole)
		AdminRouter.POST("/saveTemplate", SaveTemplate)
	}

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.POST("/change", ChangeUserPassword)
		UserRouter.GET("/info", getUserInfoFromToken) // 查询本用户信息（username和role）
		UserRouter.POST("/:userId/content", PostContent)
		UserRouter.GET("/:userId/content", GetContent)
		UserRouter.POST("/ask", QWenNormalChat)
		UserRouter.GET("/ask_base", QWenNormalChatBase)
		UserRouter.GET("/getTemplate/:template_id", GetTemplateFromId)
	}
	r.Run(":8888")
}
