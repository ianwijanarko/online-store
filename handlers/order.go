package handlers

import (
	"net/http"
	"online-store/config"
	"online-store/models"
	"time"

	"github.com/gin-gonic/gin"
)

// OrderHandler untuk membuat order baru
func OrderHandler(c *gin.Context) {
	userID := c.GetUint("user_id") // Ambil user_id dari middleware

	var input struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Periksa apakah produk tersedia
	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Stok < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Hitung total harga
	total := float64(input.Quantity) * product.Price

	// Kurangi stok produk
	product.Stok -= input.Quantity
	config.DB.Save(&product)

	// Buat order
	order := models.Order{
		UserID:    userID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Total:     total,
		Status:    "Pending",
		OrderDate: time.Now(),
	}
	config.DB.Create(&order)

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

// ListOrdersHandler untuk melihat semua order user
func ListOrdersHandler(c *gin.Context) {
	userID := c.GetUint("user_id") // Ambil user_id dari middleware

	var orders []models.Order
	if err := config.DB.Where("user_id = ?", userID).Preload("Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
