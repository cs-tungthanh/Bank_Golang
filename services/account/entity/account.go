package entity

import "time"

type Account struct {
	ID       int64
	Owner    string
	Currency string
	Balance  int64
}

type TransferResult struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`

	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
