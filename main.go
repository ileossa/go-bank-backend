package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ileossa/go-bank-backend/controllers"
	"github.com/ileossa/go-bank-backend/models"
)

func main() {

	router := gin.Default()
	models.ConnectDB()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBookById)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBookById)
	router.DELETE("/books/:id", controllers.DeleteBookById)

	router.Run()
}
