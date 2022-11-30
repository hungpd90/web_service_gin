package database

import (
	"fmt"
	"log"
	"os"

	configs "web-service-gin/configs"
	models "web-service-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	configs.LoadConfigEnv()
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

var DB *gorm.DB = Init()

func Sync() {
	DB.AutoMigrate(&models.User{})
	// fmt.Println(db.Find(&users))

}
