// @Summary Activate Account
// @Description Activate account for a specific user
// @Accept json
// @Produce json
// @Success 200 {object} bool "Return status after update"
// @Failure 400 {object} string "Error when binding JSON"
// @Failure 400 {object} string "User not found by email"
// @Router /customers [get]
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ileossa/go-bank-backend/service"
	"net/http"
)

type ActivatePatchRequest struct {
	Active bool `json:"active"`
}

func Activate(s service.Customer) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		// Validate input
		var req ActivatePatchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "", "error": err.Error()})
			return
		}
		user, err := s.GetUser(email)
		if nil != err {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user.SetStatusAccount(req.Active))
	}
}
