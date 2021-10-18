package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/ileossa/go-bank-backend/controllers"
	"github.com/ileossa/go-bank-backend/database"
	_ "github.com/ileossa/go-bank-backend/docs"
)



// @title Swagger bank-backend API
// @version 0.1
// @description This is a sample server for finetech bank.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.github.com/ileossa
// @contact.email none@email.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	r := gin.Default()
	database.ConnectDB()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBookById)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBookById)
	r.DELETE("/books/:id", controllers.DeleteBookById)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}
