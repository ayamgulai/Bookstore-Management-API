package controllers

import (
	"net/http"
	"strconv"

	"bookstore-management-api/services"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid category id",
		})
		return
	}

	category, err := services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func CreateCategory(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	category, err := services.CreateCategory(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "category created",
		"data":    category,
	})
}

func DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid category id",
		})
		return
	}

	err = services.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted",
	})
}

func GetBooksByCategory(c *gin.Context) {
	idParam := c.Param("id")

	categoryID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid category id",
		})
		return
	}

	books, err := services.GetBooksByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
