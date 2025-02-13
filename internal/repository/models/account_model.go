package models

import (
	"time"
)

type Account struct {
	ID        int64     `gorm:"primaryKey;type:bigserial"`
	Owner     string    `gorm:"not null"`
	Currency  string    `gorm:"not null"`
	Balance   int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`

	// Relationships
	User          User       `gorm:"foreignKey:Owner;references:Username"`
	Entries       []Entry    `gorm:"foreignKey:AccountID"`
	FromTransfers []Transfer `gorm:"foreignKey:FromAccountID"`
	ToTransfers   []Transfer `gorm:"foreignKey:ToAccountID"`
}
