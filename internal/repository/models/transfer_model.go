package models

import (
	"time"
)

type Transfer struct {
	ID            int64     `gorm:"primaryKey;type:bigserial"`
	FromAccountID int64     `gorm:"not null"`
	ToAccountID   int64     `gorm:"not null"`
	Amount        int64     `gorm:"not null"` // must be positive
	CreatedAt     time.Time `gorm:"not null;default:now()"`

	// Relationships
	FromAccount Account `gorm:"foreignKey:FromAccountID"`
	ToAccount   Account `gorm:"foreignKey:ToAccountID"`
}
