package routes

import (
	"web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		userRoute := v1.Group("/users")
		{
			userRoute.POST("/", controllers.CreateUser())
			userRoute.GET("/", controllers.GetUsers())
			userRoute.GET("/:id", controllers.GetUser())
		}
		v1.POST("/signup", controllers.SignUp())
		v1.POST("/signin", controllers.SignIn())
	}
}
