package controllers

import (
	"github.com/gin-gonic/gin"
	"kids-city-go/config"
	"kids-city-go/models"
	"net/http"
	"strconv"
)

// Получить избранные товары пользователя
func GetFavorites(c *gin.Context) {
	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения userId"})
		return
	}

	var favorites []models.Favorite
	if err := config.DB.
		Preload("Cloth").
		Preload("User").
		Where("userId = ?", userID).
		Find(&favorites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении избранного"})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

// Добавить товар в избранное
func AddToFavorites(c *gin.Context) {
	var input struct {
		ClothID uint `json:"clothId"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userId"})
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Некорректный userId"})
		return
	}

	// Проверка, что такой cloth существует
	var cloth models.Cloth
	if err := config.DB.First(&cloth, input.ClothID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Одежда не найдена"})
		return
	}

	// Проверка, нет ли уже в избранном
	var existing models.Favorite
	if err := config.DB.Where("userId = ? AND clothId = ?", userID, input.ClothID).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Уже в избранном"})
		return
	}

	fav := models.Favorite{
		UserID:  userID,
		ClothID: input.ClothID,
	}

	if err := config.DB.Create(&fav).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении в избранное"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Добавлено в избранное"})
}

// Удалить товар из избранного
func DeleteFromFavorites(c *gin.Context) {
	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения userId"})
		return
	}

	clothIDStr := c.Param("clothId")
	clothID, err := strconv.Atoi(clothIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID товара"})
		return
	}

	if err := config.DB.
		Where("userId = ? AND clothId = ?", userID, clothID).
		Delete(&models.Favorite{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении из избранного"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
