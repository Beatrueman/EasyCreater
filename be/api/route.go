package api

import (
	"demo/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{

		api.POST("/register", register)
		api.POST("/login", login)

		AdminRouter := api.Group("/auth")
		{
			AdminRouter.Use(middleware.JWTAuthMiddleware(), middleware.AdminAuthMiddleware())
			AdminRouter.GET("/search", getAllUserInfo)
			AdminRouter.DELETE("/delete/:userId", deleteUser)
			AdminRouter.POST("/role/:userId", ChangeUserRole)
		}

		UserRouter := api.Group("/user")
		{
			UserRouter.Use(middleware.JWTAuthMiddleware())
			UserRouter.POST("/change", ChangeUserPassword)
			UserRouter.GET("/info", getUserInfoFromToken) // 查询本用户信息（username和role）
			UserRouter.POST("/:userId/content", PostContent)
			UserRouter.GET("/:userId/content", GetContent)
			UserRouter.POST("/ask", QWenNormalChat)
			UserRouter.POST("/ask_base", QWenNormalChatBase)
			UserRouter.POST("/resume/save", AddResume)
			UserRouter.GET("/resume/list", GetResume)
			UserRouter.PUT("/resume/share/:resume_id", ShareResume)
			UserRouter.GET("/resume/share", GetSharedResume) // 获取用户已分享的简历
			UserRouter.DELETE("/resume/delete/:resume_id", DeleteResume)
			UserRouter.GET("/resume/:resume_id", GetResumeFromId)
			UserRouter.POST("/avatar/upload", AddUserAvatar)
			UserRouter.GET("/avatar/load", GetUserAvatar)
			UserRouter.GET("/resume/thumbnail/:resume_id", GetThumbnail)
			UserRouter.POST("/resume/upload", uploadResume)
			UserRouter.GET("/resume/get_loaded", GetLoadedResume)
			UserRouter.DELETE("/resume/delete_loaded/:resume_id", DeleteLoadedResume)
			UserRouter.GET("/resume/get_loaded_url/:resume_id", GetLoadedResumeURL)
		}
	}
	r.Run(":8888")
}
