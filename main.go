package main

import (
	"fmt"
	"log"
	"os"
	"server/controller"
	"server/database"
	"server/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Person{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()
	router.Use(controller.CORSMiddleware())
	publicRoutes := router.Group("/api")
	publicRoutes.POST("", controller.Create)
	publicRoutes.GET("/:userId", controller.Read)
	publicRoutes.PUT("/:userId", controller.Update)
	publicRoutes.DELETE("/:userId", controller.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

	fmt.Println("Server running on port 8000")
}
