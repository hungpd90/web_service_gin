package database

import (
	"fmt"
	"log"

	models "web-service-gin/models"
	"web-service-gin/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	utils.LoadConfig()
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.Config.DatabaseHost,
		utils.Config.DatabaseUser,
		utils.Config.DatabasePassword,
		utils.Config.DatabaseName,
		utils.Config.DatabasePort,
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

var DB *gorm.DB = Init()

func Sync() {
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Manufacturer{})
	// fmt.Println(db.Find(&users))
}
