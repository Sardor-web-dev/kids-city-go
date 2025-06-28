package controllers

import (
	"github.com/gin-gonic/gin"
	"kids-city-go/config"
	"kids-city-go/models"
	"net/http"
)

// GET /api/cloths
func GetCloths(c *gin.Context) {
	var clothes []models.Cloth
	if err := config.DB.Find(&clothes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка"})
		return
	}
	c.JSON(http.StatusOK, clothes)
}

// GET /api/cloths/:id
func GetClothByID(c *gin.Context) {
	id := c.Param("id")
	var cloth models.Cloth

	if err := config.DB.First(&cloth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Одежда не найдена"})
		return
	}

	c.JSON(http.StatusOK, cloth)
}

// POST /api/cloths
func CreateCloth(c *gin.Context) {
	var input models.Cloth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// PUT /api/cloths/:id
func UpdateCloth(c *gin.Context) {
	id := c.Param("id")
	var cloth models.Cloth

	if err := config.DB.First(&cloth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Одежда не найдена"})
		return
	}

	var input models.Cloth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Обновляем поля
	cloth.Name = input.Name
	cloth.Description = input.Description
	cloth.Gender = input.Gender
	cloth.Image = input.Image
	cloth.AuthorID = input.AuthorID
	cloth.Price = input.Price
	cloth.Size = input.Size

	if err := config.DB.Save(&cloth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, cloth)
}

// DELETE /api/cloths/:id
func DeleteCloth(c *gin.Context) {
	id := c.Param("id")
	var cloth models.Cloth

	if err := config.DB.First(&cloth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Одежда не найдена"})
		return
	}

	if err := config.DB.Delete(&cloth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Одежда удалена"})
}
