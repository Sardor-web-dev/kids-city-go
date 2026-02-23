package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kids-city-go/config"
	"kids-city-go/routes"
	"os"
    "kids-city-go/middleware"
)

func main() {
	_ = godotenv.Load() 
	config.ConnectDB()
    middleware.InitJWTSecret() 
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":" + os.Getenv("PORT"))
}
