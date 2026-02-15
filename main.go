package main

import (
	"bookstore-management-api/configs"
	"bookstore-management-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
	PORT := ":8080"
	r := gin.Default()
	closeDB := configs.ConnectDB()
	defer closeDB()

	configs.RunMigrations()
	routes.StartServer(r)
	r.Run(PORT)
}
