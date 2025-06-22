package routes

import (
	"github.com/gin-gonic/gin"
	"kids-city-go/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/cloths", controllers.GetCloths)
		api.GET("/cloths/:id", controllers.GetClothByID)
	}
}
