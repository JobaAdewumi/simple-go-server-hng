package main

import (
    "server/database"
    "server/model"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    loadEnv()
    loadDatabase()
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}