package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid"`
	Username     string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	UserAgent    string    `gorm:"not null"`
	ClientIP     string    `gorm:"not null"`
	IsBlocked    bool      `gorm:"not null;default:false"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`

	// Relationships
	User User `gorm:"foreignKey:Username"`
}
