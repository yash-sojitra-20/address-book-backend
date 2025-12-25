package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	ID uint `gorm:"primaryKey"`
	Email string `gorm:"not null;unique"`
	Password string `gorm:"not null"`

	Contacts []Contact `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}