package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Address  string `gorm:"type:text;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Password string `gorm:"not null"`
	Token    string `gorm:"size:255"`
}
