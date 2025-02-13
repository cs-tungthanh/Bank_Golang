package models

import (
	"time"
)

type Entry struct {
	ID        int64     `gorm:"primaryKey;type:bigserial"`
	AccountID int64     `gorm:"not null"`
	Amount    int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`

	// Relationships
	Account Account `gorm:"foreignKey:AccountID"`
}
