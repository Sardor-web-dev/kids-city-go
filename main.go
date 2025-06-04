package main

import (
	"kids-city-go/config"
	"kids-city-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() // Инициализируем подключение

	config.LoadEnv()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
