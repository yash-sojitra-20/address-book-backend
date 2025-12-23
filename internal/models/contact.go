package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID           uint           `gorm:"primaryKey"`
	UserID       uint           `gorm:"not null"`
	FirstName    string         `gorm:"not null" json:"first_name"`
	LastName     string			`json:"last_name"`
	Email        string
	Phone        string			`gorm:"not null"`
	AddressLine1 string			`json:"address_line1"`
	AddressLine2 string			`json:"address_line2"`
	City         string
	State        string
	Country      string
	Pincode      string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
