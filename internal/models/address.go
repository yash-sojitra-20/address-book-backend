package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null;index"`
	FirstName    string `gorm:"not null" json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string
	Phone        string `gorm:"not null"`
	AddressLine1 string `gorm:"not null" json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `gorm:"not null"`
	State        string `gorm:"not null"`
	Country      string `gorm:"not null"`
	Pincode      string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
