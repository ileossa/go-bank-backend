// @Summary Ping
// @Description We can check if application is up
// @Produce json
// @Success 200 {object} string "return UP if application is ready"
// @Router /ping [get]
package handlers

import "github.com/gin-gonic/gin"

func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	}
}
