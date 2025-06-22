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

func GetClothByID(c *gin.Context) {
	id := c.Param("id")
	var cloth models.Cloth

	if err := config.DB.First(&cloth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Одежда не найдена"})
		return
	}

	c.JSON(http.StatusOK, cloth)
}
