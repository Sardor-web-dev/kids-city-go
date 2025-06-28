package controllers

import (
	"bytes"
	"encoding/json" // ← вот эта строка нужна
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // 🧩 добавь вот эту строку
	"io"
	"kids-city-go/config"
	"kids-city-go/models"
	"net/http"
)

// Создание заказа
func CreateOrder(c *gin.Context) {
	// Читаем тело вручную, чтобы обработать json.RawMessage
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось прочитать тело запроса"})
		return
	}

	// Восстанавливаем тело запроса для возможного повторного чтения
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Парсим JSON в структуру Order
	var input models.Order
	if err := json.Unmarshal(body, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат JSON"})
		return
	}

	// Достаём userId из контекста
	userIDVal, exists := c.Get("userId") // ← а тут ты его достаёшь

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userId"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userId неверного типа"})
		return
	}
	fmt.Println("Получен userID:", userID)

	// Распарсим Items из RawMessage
	var items []map[string]interface{}
	if err := json.Unmarshal(input.Items, &items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в формате items"})
		return
	}

	// Подсчёт total
	total := 0
	for _, item := range items {
		qty, okQty := item["quantity"].(float64)
		price, okPrice := item["price"].(float64)
		if okQty && okPrice {
			total += int(qty * price)
		}
	}
	input.Total = total

	input.ID = uuid.NewString()
	input.UserID = int(userID)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
		return
	}

	fmt.Println("RAW BODY:", string(body))
	fmt.Printf("Заказ перед сохранением: %+v\n", input)
	c.JSON(http.StatusCreated, gin.H{"success": true, "order": input})
}

// Получение заказов текущего пользователя
func GetUserOrders(c *gin.Context) {
	userIDVal, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userId"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userId неверного типа"})
		return
	}
	fmt.Println("Получен userID:", userID)

	var orders []models.Order
	if err := config.DB.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении заказов"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// Обновление статуса заказа (admin only)
func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	validStatuses := map[string]bool{
		"process":  true,
		"done":     true,
		"canceled": true,
	}

	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недопустимый статус"})
		return
	}

	if err := config.DB.Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", input.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
