package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"yeebotech-auth/controllers"
	"yeebotech-auth/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS 配置，允许所有来源
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true                                           // 允许所有来源
	corsConfig.AddAllowHeaders("Authorization")                                 // 允许携带 Authorization 头
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"} // 允许多种方法

	// 设置跨域中间件
	r.Use(cors.New(corsConfig))

	// Auth routes
	r.POST("/auth/register", controllers.Register) // 注册
	r.POST("/auth/login", controllers.Login)       // 登录
	// 检查 token 和刷新 token 路由
	r.POST("/auth/check-token", controllers.CheckToken)
	r.POST("/auth/refresh-token", controllers.RefreshToken)
	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 使用 JWT 中间件保护的用户相关路由
	auth := r.Group("/user").Use(middlewares.JWTAuthMiddleware())
	{
		auth.GET("/profile", controllers.GetProfile)              // 获取用户信息
		auth.POST("/change-password", controllers.ChangePassword) // 修改密码
		//auth.GET("/", controllers.GetUserList)           // 获取用户列表
		//auth.GET("/:id", controllers.GetUserDetail)      // 获取用户详情
		//auth.POST("/", controllers.CreateUser)           // 创建用户
		//auth.PUT("/:id", controllers.UpdateUser)         // 更新用户
		//auth.DELETE("/:id", controllers.DeleteUser)      // 删除用户
	}

	return r
}
