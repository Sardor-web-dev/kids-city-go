package controllers

import (
	"github.com/gin-gonic/gin"
	"kids-city-go/config"
	"kids-city-go/models"
	"net/http"
)

func GetCloths(c *gin.Context) {
	var clothes []models.Cloth

	if err := config.DB.Find(&clothes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка"})
		return
	}

	c.JSON(http.StatusOK, clothes)
}
