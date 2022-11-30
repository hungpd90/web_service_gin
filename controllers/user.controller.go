package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"web-service-gin/database"
	"web-service-gin/models"
	"web-service-gin/serializers"
	"web-service-gin/utils"
	"web-service-gin/validates"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			utils.SendResponse(c, http.StatusBadRequest, err.Error(), "")
			return
		}
		fmt.Println(user, &user)

		database.DB.Create(&user)
		utils.SendResponse(c, http.StatusCreated, "Success", true)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		database.DB.Find(&users)
		utils.SendResponse(c, http.StatusOK, "Success", users)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id := c.Param("id")
		if id == ":id" {
			utils.SendResponse(c, http.StatusBadRequest, "Id not found!", "")
			return
		}
		err := database.DB.First(&user, id).Error
		if err != nil {
			utils.SendResponse(c, http.StatusNotFound, err.Error(), "")
			return
		}
		var serializedUser = serializers.GetUserSerializer(&user)
		utils.SendResponse(c, http.StatusOK, "Success", serializedUser)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body validates.SignUpBody
		if err := c.BindJSON(&body); err != nil {
			utils.SendResponse(c, http.StatusBadRequest, err.Error(), "")
			return
		}

		var user models.User

		// Check if email existed
		result := database.DB.Where("email = ?", body.Email).First(&user)
		if result.RowsAffected > 0 {
			utils.SendResponse(c, http.StatusBadRequest, "Email existed", "")
			return
		}

		// Check if phone number existed
		result = database.DB.Where("phone_number = ?", body.PhoneNumber).First(&user)
		if result.RowsAffected > 0 {
			utils.SendResponse(c, http.StatusBadRequest, "Phone number existed", "")
			return
		}

		// Create user on DB
		user.PhoneNumber = body.PhoneNumber
		user.Email = body.Email
		user.Password = HashPassword(body.Password)
		user.FirstName = body.FirstName
		user.LastName = body.LastName
		var getDobSuccess bool
		user.Dob, getDobSuccess = GetDOB(body.Dob)
		if !getDobSuccess {
			utils.SendResponse(c, http.StatusBadRequest, "Dob wrong format!", "")
			return
		}
		user.Gender = body.Gender
		database.DB.Create(&user)

		// Create JWT
		

		utils.SendResponse(c, http.StatusCreated, "Success", true)
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func GetDOB(dob string) (time.Time, bool) {
	splitedDobArray := strings.Split(dob, "-")
	var splitDobInt []int
	for i := 0; i < len(splitedDobArray); i++ {
		tempInt, err := strconv.Atoi(splitedDobArray[i])
		if err != nil {
			return time.Time{}, false
		}
		splitDobInt = append(splitDobInt, tempInt)
	}
	return time.Date(splitDobInt[0], time.Month(splitDobInt[1]), splitDobInt[2], 0, 0, 0, 0, time.UTC), true
}
