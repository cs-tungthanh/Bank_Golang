package business

import (
	"context"
	"fmt"

	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/pkg/core"
	"github.com/cs-tungthanh/Bank_Golang/services/account/entity"
)

type CreateTransferParams struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (b *business) CreateTransfer(ctx context.Context, req CreateTransferParams) (*entity.TransferResult, error) {
	if err := b.validAccount(ctx, req.FromAccountID, req.Currency); err != nil {
		return nil, err
	}

	if err := b.validAccount(ctx, req.ToAccountID, req.Currency); err != nil {
		return nil, err
	}

	result, err := b.accountRepo.CreateTransfer(ctx, db.CreateTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil {
		return nil, core.ErrInternalServerError.WithErrorf("failed to create transfer: %s", err.Error())
	}

	return &entity.TransferResult{
		ID:            result.ID,
		FromAccountID: result.FromAccountID,
		ToAccountID:   result.ToAccountID,
		Amount:        result.Amount,
		CreatedAt:     result.CreatedAt,
	}, nil
}

// check if an account with a specific ID really exists.
func (b *business) validAccount(ctx context.Context, accountID int64, currency string) error {
	account, err := b.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		return core.ErrBadRequest.WithError(err.Error())
	}

	return nil
}
