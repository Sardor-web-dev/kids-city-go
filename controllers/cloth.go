package controllers

import (
	"kids-city-go/config"
	"kids-city-go/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateCloth(c *gin.Context) {
	var cloth models.Cloth

	if err := c.ShouldBindJSON(&cloth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Автор ID временно хардкодим или вытаскиваем из контекста авторизации позже
	cloth.AuthorID = 1

	if err := config.DB.Create(&cloth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании"})
		return
	}

	c.JSON(http.StatusCreated, cloth)
}
