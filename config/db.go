package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kids-city-go/models"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false, // <— ВАЖНО: отключает подготовленные выражения
	})

	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}
	DB = db

	// Автоматическая миграция
	DB.AutoMigrate(&models.User{}, &models.Cloth{}, &models.Favorite{}, &models.Order{})
}
