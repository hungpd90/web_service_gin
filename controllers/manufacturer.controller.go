package controllers

import (
	"net/http"
	"web-service-gin/database"
	"web-service-gin/models"
	"web-service-gin/utils"
	"web-service-gin/validates"

	"github.com/gin-gonic/gin"
)

func CreateManufacturer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body validates.CreateManufacturerBody
		if err := c.BindJSON(&body); err != nil {
			utils.SendResponse(c, http.StatusBadRequest, err.Error(), "")
			return
		}
		var manufacturer models.Manufacturer
		manufacturer.Name = body.Name
		manufacturer.Origin = body.Origin
		database.DB.Create(&manufacturer)
		utils.SendResponse(c, http.StatusCreated, "Success", true)
	}
}

func GetManufacturers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var manufacturers []models.Manufacturer
		var pagination utils.Pagination
		utils.GetPaginationParameter(c, &pagination)
		database.DB.Scopes(utils.Paginate(manufacturers, &pagination, database.DB)).Find(&manufacturers)
		pagination.Rows = manufacturers
		utils.SendResponse(c, http.StatusOK, "Success", pagination)
	}
}
