package controllers

import (
	"net/http"
	"kids-city-go/models"
	"github.com/gin-gonic/gin"
	"kids-city-go/config"
)

// Заблокировать/разблокировать пользователя (для админа)
func BlockUser(c *gin.Context) {
	var input struct {
		UserID uint `json:"userId"`
		Block  bool `json:"block"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if err := config.DB.Model(&models.User{}).
		Where("id = ?", input.UserID).
		Update("is_blocked", input.Block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
