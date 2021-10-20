//@Summary Order paiement
// @Description Get user's informations
// @Accept json
// @Produce json
// @Success 200 {object} string "Return the status of the emission"
// @Failure 400 {object} string "Error when binding JJSON"
// @Failure 400 {object} string "Not find User with our email"
// @Failure 400 {object} string "Account isn't activate"
// @Failure 400 {object} string "Account isn't activate"
// @Failure 400 {object} string "Error User haven't this card"
// @Failure 400 {object} string "User doesn't have enought money"
// @Failure 400 {object} string "Can't retrieve payment history, for user's card"
// @Router /order/emission [post]
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ileossa/go-bank-backend/http/service"
	"net/http"
	"strconv"
	"time"
)

type CardSchemaPostRequest struct {
	Number       uint      `json:"card_number"`
	Organization string    `json:"organization"`
	Validity     time.Time `json:"validity"`
	crypto       uint      `json:"crypto"`
}

type OrderPostRequest struct {
	Email      string `json:"email"`
	CardNumber CardSchemaPostRequest
	Amount     int64  `json:"amount"`
	Status     string `json:"status"`
}

func Order(s service.Customer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate and Bind Json
		var req OrderPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// check if user exist
		user, err := s.GetUser(req.Email)
		if nil != err {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// check if user is activated
		if false == user.Active {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User " + user.Email + " isn't activate"})
			return
		}
		// check if card exist and attached to user
		if false == user.CardExist(req.CardNumber.Number) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User haven't have this card: " + strconv.FormatUint(uint64(req.CardNumber.Number), 10)})
			return
		}
		// check if user have the amount on our account and minus
		if r, _ := user.Minus(req.Amount); false == r {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User doesn't enought money"})
			return
		}
		// insert transaction into user's history
		if schema, err := user.GetCard(req.CardNumber.Number); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can't retrieve payment history, for this card: " + strconv.FormatUint(uint64(req.CardNumber.Number), 10)})
			return
		} else {
			item := service.HistorySchema{
				Date:     time.Now(),
				Amount:   uint(req.Amount),
				Category: "not_implemented",
			}
			schema.NewPayment(item)
		}
		// return status
		c.JSON(http.StatusOK, gin.H{"message": "SUCCESS"})
		return
	}
}
