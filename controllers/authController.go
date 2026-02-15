package controllers

import (
	"net/http"

	"bookstore-management-api/services"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	// bind & validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username and password are required",
		})
		return
	}

	// call service
	token, err := services.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}
