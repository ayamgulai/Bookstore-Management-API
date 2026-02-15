package routes

import (
	"bookstore-management-api/controllers"
	"bookstore-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer(r *gin.Engine) {
	api := r.Group("/api")

	// auth
	api.POST("/users/login", controllers.Login)

	// protected routes
	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	// categories
	protected.GET("/categories", controllers.GetCategories)
	protected.POST("/categories", controllers.CreateCategory)
	protected.GET("/categories/:id", controllers.GetCategoryByID)
	protected.DELETE("/categories/:id", controllers.DeleteCategory)
	protected.GET("/categories/:id/books", controllers.GetBooksByCategory)

	// books
	protected.GET("/books", controllers.GetBooks)
	protected.POST("/books", controllers.CreateBook)
	protected.GET("/books/:id", controllers.GetBookByID)
	protected.DELETE("/books/:id", controllers.DeleteBook)
}
