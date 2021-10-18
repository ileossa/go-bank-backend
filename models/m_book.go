package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  *time.Time `sql:"index"`
}
type Book struct {
	gorm.Model	`swaggerignore:"true"`
	Title     string `json:"title"`
	Author    string `json:"author"`

}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
