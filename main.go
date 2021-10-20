package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/ileossa/go-bank-backend/http/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/ileossa/go-bank-backend/http/handlers"
	"github.com/ileossa/go-bank-backend/http/service"
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
	c := service.InitCustomer()
	//utils.ConnectDB(c)

	r := gin.Default()

	r.GET("/ping", handlers.PingGet())

	r.GET("/customers", handlers.GetCustomers(c))
	r.POST("/customers", handlers.NewUser(c))
	r.PATCH("/customers", handlers.Activate(c))

	r.POST("/order/emission", handlers.Order(c))

	//r.PATCH("/customers", handlers.UpdateCustomer(c))
	//r.DELETE("/customers", handlers.DisableCustomer(c))

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()

	//// Set the USERNAME environment variable to "MattDaemon"
	//os.Setenv("USERNAME", "MattDaemon")
	//
	//// Get the USERNAME environment variable
	//username := os.Getenv("USERNAME")
	//
	//// Prints out username environment variable
	//fmt.Print(username)
}
