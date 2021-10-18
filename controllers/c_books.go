package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ileossa/go-bank-backend/database"
	"github.com/ileossa/go-bank-backend/models"
	"net/http"
)


// @Summary List of books
// @Description get list of all books
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /books [get]
func FindBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// @Summary New books
// @Description Insert a new book
// @Accept json
// @Produce json
// @Success 200 {object} models.CreateBookInput "return book created"
// @Failure 400 {object} string "The error built-in interface type is the conventional interface for representing an error condition."
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	b := models.Book{Title: input.Title, Author: input.Author}
	database.DB.Create(&b)
	c.JSON(http.StatusOK, gin.H{"data": b})
}


// @Summary Find a specific book
// @Description Find a specfic book with your identifiant
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book "return book"
// @Failure 400 {object} string "Record not found"
// @Router /books/{id} [get]
func FindBookById(c *gin.Context) {
	var b models.Book

	err := database.DB.Where("id=?", c.Param("id")).First(&b).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": b})
}


// @Summary Update a specific book
// @Description Find a specfic book with your identifiant and update it
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.UpdateBookInput "return book updated"
// @Failure 400 {object} string "Record not found"
// @Failure 400 {object} string "The error built-in interface type is the conventional interface for representing an error condition."
// @Router /books/{id} [patch]
func UpdateBookById(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Exec("UPDATE books SET title = ? WHERE id = ?", input.Title, input.Author)
	c.JSON(http.StatusOK, gin.H{"data": book})
}


// @Summary Delete a specific book
// @Description Find a specfic book with your identifiant and delete it
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} bool "return the result (boolean)"
// @Failure 400 {object} string "Record not found"
// @Failure 400 {object} string "The error built-in interface type is the conventional interface for representing an error condition."
// @Router /books/{id} [delete]
func DeleteBookById(c *gin.Context) {
	var b models.Book
	if err := database.DB.Where("id = ?", c.Param("id")).First(&b).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&b)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
