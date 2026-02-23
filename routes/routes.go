package routes

import (
	"github.com/gin-gonic/gin"
	"kids-city-go/controllers"
	"kids-city-go/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/cloths", controllers.GetCloths)
	r.GET("/api/cloths/:id", controllers.GetClothByID)
	r.POST("/api/login", controllers.Login)

	r.POST("/api/orders", middleware.AuthMiddleware(), controllers.CreateOrder)

	// Защищенные маршруты
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	// Favorites
	auth.GET("/favorites", controllers.GetFavorites)
	auth.PUT("/favorites/add", controllers.AddToFavorites)
	auth.DELETE("/favorites/:clothId", controllers.DeleteFromFavorites)

	// Заказы
	auth.GET("/orders/user", controllers.GetUserOrders)
	auth.PATCH("/orders/:id/status", middleware.AdminMiddleware(), controllers.UpdateOrderStatus)

	// Добавление одежды (только для авторизованных)
	auth.POST("/cloth", controllers.CreateCloth)
	auth.PUT("/cloths/:id", controllers.UpdateCloth)
	auth.DELETE("/cloths/:id", controllers.DeleteCloth)

	// Admin: блокировка пользователей
	auth.POST("/admin/block-user", middleware.AdminMiddleware(), controllers.BlockUser)
}
