package main

import (
	"kids-city-go/config"
	"kids-city-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() // Подключаем БД и загружаем .env

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080") // Слушаем порт 8080
}
