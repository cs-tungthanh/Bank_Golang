package models

import (
	"time"
)

type User struct {
	Username          string    `gorm:"primaryKey"`
	HashedPassword    string    `gorm:"not null"`
	FullName          string    `gorm:"not null"`
	Email             string    `gorm:"unique;not null"`
	PasswordChangedAt time.Time `gorm:"not null;default:'0001-01-01 00:00:00Z'"`
	CreatedAt         time.Time `gorm:"not null;default:now()"`

	// Relationships
	Accounts []Account `gorm:"foreignKey:Owner"`
	Sessions []Session `gorm:"foreignKey:Username"`
}
