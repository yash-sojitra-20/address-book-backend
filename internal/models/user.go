package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`

	Address []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
