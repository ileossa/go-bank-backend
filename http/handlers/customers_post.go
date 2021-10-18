// @Summary New User
// @Description Insert a new user
// @Accept json
// @Produce json
// @Success 200 {object} []UserSchema "return arry of users"
// @Failure 400 {object} string "Error when binding JSON"
// @Router /customers [post]
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ileossa/go-bank-backend/service"
	"net/http"
	"time"
)

type customerPostRequest struct {
	Email     string    `json:"email" gorm:"primary_key"`
	Name      string    `json:"name"`
	Firstname string    `json:"firstname"`
	Address   string    `json:"address"`
	Birthday  time.Time `json:"birthday"`
	Active    bool      `json:"active"`
	Money     int64     `json:"money"`
}

func NewUser(s service.Customer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate input
		var req customerPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := service.UserSchema{
			Email:          req.Email,
			Name:           req.Name,
			Firstname:      req.Firstname,
			Address:        req.Address,
			Birthday:       req.Birthday,
			Active:         req.Active,
			Card:           service.InitUser().Card,
			MoneyOnAccount: req.Money,
		}
		c.JSON(http.StatusOK, s.NewUser(user))
	}
}
