package middlewares

import (
	"net/http"
	"web-service-gin/utils"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			utils.SendResponse(c, http.StatusUnauthorized, "token not found!", "")
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			utils.SendResponse(c, http.StatusUnauthorized, "token validate error", "")
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("phone_number", claims.PhoneNumber)
		c.Next()
	}
}
