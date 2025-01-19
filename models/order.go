package models

import "time"

type Order struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	ProductID     uint      `gorm:"not null"`
	Product       Product   `gorm:"foreignKey:ProductID"` // Relasi ke tabel Product
	Quantity      int       `gorm:"not null"`
	Total         float64   `gorm:"not null"`
	Status        string    `gorm:"size:50;not null"` // Status order (Pending, Completed, etc.)
	PaymentStatus string    `gorm:"size:50;not null"` // Status pembayaran (Unpaid, Paid)
	OrderDate     time.Time `gorm:"autoCreateTime"`
}
