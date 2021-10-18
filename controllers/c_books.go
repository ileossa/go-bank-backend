package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ileossa/go-bank-backend/models"
	"net/http"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	b := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&b)
	c.JSON(http.StatusOK, gin.H{"data": b})
}

// GET /books/:id
// Find a book
func FindBookById(c *gin.Context) {
	var b models.Book

	err := models.DB.Where("id=?", c.Param("id")).First(&b).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": b})
}

// PATCH /books/:id
// Update a book
func UpdateBookById(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Exec("UPDATE books SET title = ? WHERE id = ?", input.Title, input.Author)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Deleta a book
func DeleteBookById(c *gin.Context) {
	var b models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&b).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&b)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
