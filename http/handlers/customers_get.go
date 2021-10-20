// @Summary Get User
// @Description Get user's informations
// @Accept json
// @Produce json
// @Success 200 {object} UserSchema "return informations of specific email's user passed on the requet"
// @Success 200 {object} []UserSchema "return arry of users if email is empty or no passed"
// @Failure 400 {object} string "Email not found"
// @Router /customers [get]
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ileossa/go-bank-backend/http/service"
	"net/http"
)

func GetCustomers(methods service.Customer) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if "" == email {
			c.JSON(http.StatusOK, methods.GetUsers())
		} else {
			user, err := methods.GetUser(email)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, user)
		}
	}
}
