package main

import (
	db "web-service-gin/database"
	"web-service-gin/routes"
	"web-service-gin/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// configs.LoadConfigEnv()
	utils.LoadConfig()

	// fmt.Println("GETENV:" + os.Getenv("DATABASE_URL") + "1")
	// fmt.Println("Helloworld")
	db.Sync()

	router := gin.Default()

	routes.Routes(router)

	router.Run("localhost:8000")
}
