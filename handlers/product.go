package handlers

import (
	"net/http"
	"online-store/config"
	"online-store/models"

	"github.com/gin-gonic/gin"
)

// AddProductHandler untuk menambahkan produk baru
func AddProductHandler(c *gin.Context) {
	var input models.Product

	// Validasi input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan produk ke database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully", "product": input})
}

// GetAllProductsHandler untuk mendapatkan semua produk
func GetAllProductsHandler(c *gin.Context) {
	var products []models.Product

	// Ambil semua produk dari database
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductByIDHandler untuk mendapatkan produk berdasarkan ID
func GetProductByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Cari produk berdasarkan ID
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// UpdateProductHandler untuk memperbarui data produk
func UpdateProductHandler(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Cari produk berdasarkan ID
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Validasi input JSON
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbarui data produk
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stok = input.Stok

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": product})
}

// DeleteProductHandler untuk menghapus produk
func DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Cari produk berdasarkan ID
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Hapus produk
	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
