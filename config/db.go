package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kids-city-go/models"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it, trying to use env variables directly")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL not set in environment variables")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Successfully connected to the database")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
