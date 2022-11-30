package utils

import (
	"web-service-gin/interfaces"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	var response interfaces.Response
	response.Status = uint(statusCode)
	response.Message = message
	response.Data = data
	c.IndentedJSON(statusCode, response)
}
