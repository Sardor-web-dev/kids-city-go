package controllers

import (
	"kids-city-go/config"
	"kids-city-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}
	c.JSON(http.StatusOK, cart)
}