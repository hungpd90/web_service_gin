package controllers

import (
	"net/http"
	"web-service-gin/database"
	"web-service-gin/models"
	"web-service-gin/utils"
	"web-service-gin/validates"

	"github.com/gin-gonic/gin"
)

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		database.DB.Find(&products)
		utils.SendResponse(c, http.StatusOK, "Success", products)
	}
}
func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body validates.CreateProductBody
		if err := c.BindJSON(&body); err != nil {
			utils.SendResponse(c, http.StatusBadRequest, err.Error(), "")
			return
		}

		var product models.Product
		product.SKU = body.SKU
		product.Name = body.Name
		product.Description = body.Description
		product.Images = body.Images
		product.Price = body.Price
		product.Quantity = body.Quantity
		product.Manufacturer.ID = body.ManufacturerId
		database.DB.Create(&product)
		utils.SendResponse(c, http.StatusCreated, "Success", true)
	}
}
