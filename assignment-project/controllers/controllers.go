package controllers

import (
	"assignment-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {

	var input struct {
		CustomerName string        `json:"customerName" binding:"required"`
		OrderedAt    time.Time     `json:"orderedAt" binding:"required"`
		Items        []models.Item `json:"items" binding:"required,dive"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	order := models.Order{
		CustomerName: input.CustomerName,
		OrderedAt:    input.OrderedAt,
		Items:        input.Items,
	}

	if err := models.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func AllOrders(c *gin.Context) {

	var orders []models.Order

	if err := models.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func OrderById(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func UpdateOrder(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	var input struct {
		CustomerName string    `json:"customerName" binding:"required"`
		OrderedAt    time.Time `json:"orderedAt" binding:"required"`
		Items        []struct {
			Name        string `json:"name" binding:"required"`
			Description string `json:"description" binding:"required"`
			Quantity    int    `json:"quantity" binding:"required"`
		} `json:"items" binding:"required,dive"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var order models.Order
	if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	// Hapus item lama
	if err := models.DB.Where("order_id = ?", order.ID).Delete(&models.Item{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Set item baru
	var newItems []models.Item
	for _, item := range input.Items {
		newItems = append(newItems, models.Item{
			Name:        item.Name,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     order.ID,
		})
	}

	// Perbarui order
	order.CustomerName = input.CustomerName
	order.OrderedAt = input.OrderedAt
	order.Items = newItems

	if err := models.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func DeleteOrder(c *gin.Context) {
	var order models.Order

	id := c.Param("id")

	if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	// Hapus item-item terkait
	if err := models.DB.Where("order_id = ?", order.ID).Delete(&models.Item{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Hapus order
	if err := models.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order telah dihapus"})
}
