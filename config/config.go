package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kids-city-go/models"
)

var DB *gorm.DB

func ConnectDB() {
	LoadEnv()

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Не удалось подключиться к базе данных:", err)
	}

	// Миграции
	err = db.AutoMigrate(&models.Cloth{})
	if err != nil {
		log.Fatal("❌ Ошибка миграции:", err)
	}

	DB = db
	log.Println("✅ Успешное подключение к базе данных")
}
