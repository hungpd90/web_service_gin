package routes

import (
	"web-service-gin/controllers"
	"web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup", controllers.SignUp())
		v1.POST("/signin", controllers.SignIn())
		userRoute := v1.Group("/users")
		{
			userRoute.Use(middlewares.Authorization())
			userRoute.GET("/me", controllers.GetSelfInfo())
			userRoute.POST("/", controllers.CreateUser())
			userRoute.GET("/", controllers.GetUsers())
			userRoute.GET("/:id", controllers.GetUser())
		}
		adminRoute := v1.Group("/admin")
		{
			adminRoute.GET("/users", controllers.GetUsers())
			adminRoute.GET("/user", controllers.GetUser())
		}
		productRoute := v1.Group("/products")
		{
			productRoute.GET("/", controllers.GetProducts())
			//productRoute.POST("/products", middlewares.Authorization(), controllers.CreateProduct())
			productRoute.POST("/", controllers.CreateProduct())
		}
	}
}
