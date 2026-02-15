package controllers

import (
	"net/http"
	"strconv"

	"bookstore-management-api/models"
	"bookstore-management-api/services"

	"github.com/gin-gonic/gin"
)

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year" binding:"required"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"total_page" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}

func GetBooks(c *gin.Context) {
	books, err := services.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	book, err := services.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	book := models.Book{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		ReleaseYear: req.ReleaseYear,
		Price:       req.Price,
		TotalPage:   req.TotalPage,
		CategoryID:  req.CategoryID,
		CreatedBy:   "system",
	}
	createdBook, err := services.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "book created",
		"data":    createdBook,
	})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := services.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}
