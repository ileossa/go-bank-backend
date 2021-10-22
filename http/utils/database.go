package utils

import (
	"github.com/ileossa/go-bank-backend/http/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(schema *service.CustomerSchema) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=127.0.0.1 port=5432 user=postgres dbname=bank password=changeme sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}
	db.AutoMigrate(schema)
	DB = db
}
