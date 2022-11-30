package main

import (
	"fmt"
	"os"
	configs "web-service-gin/configs"
	db "web-service-gin/database"
	"web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfigEnv()
	fmt.Println("GETENV:" + os.Getenv("DATABASE_URL") + "1")
	fmt.Println("Helloworld")
	db.Sync()

	router := gin.Default()

	routes.Routes(router)

	router.Run("localhost:8000")
}
