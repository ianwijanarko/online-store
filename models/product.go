package models

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Stok        int     `gorm:"not null"`
}
