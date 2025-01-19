package main

import (
	"online-store/config"
	"online-store/handlers"
	"online-store/middlewares"
	"online-store/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Hubungkan ke database
	config.ConnectDatabase()

	// Lakukan migrasi
	config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	r := gin.Default()

	// Endpoint user
	r.POST("/login", handlers.LoginHandler)
	r.POST("/logout", middlewares.AuthMiddleware(), handlers.LogoutHandler)
	r.POST("/register", handlers.RegisterHandler)

	// Routing untuk produk
	r.POST("/products", handlers.AddProductHandler)          // Tambah produk baru
	r.GET("/products", handlers.GetAllProductsHandler)       // Ambil semua produk
	r.GET("/products/:id", handlers.GetProductByIDHandler)   // Ambil produk berdasarkan ID
	r.PUT("/products/:id", handlers.UpdateProductHandler)    // Perbarui produk
	r.DELETE("/products/:id", handlers.DeleteProductHandler) // Hapus produk

	// Endpoint order
	orderGroup := r.Group("/orders", middlewares.AuthMiddleware())
	{
		orderGroup.POST("", handlers.OrderHandler)     // Buat order baru
		orderGroup.GET("", handlers.ListOrdersHandler) // Lihat daftar order
	}

	r.Run(":8080")
}
