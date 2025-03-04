package business

import (
	"context"
	"database/sql"

	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/pkg/core"
	"github.com/cs-tungthanh/Bank_Golang/services/account/entity"

	pg "github.com/lib/pq"
)

type AccountBiz interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (*entity.Account, error)
	GetAccount(ctx context.Context, ID int64) (*entity.Account, error)
	ListAccounts(ctx context.Context, req ListAccountRequest) ([]entity.Account, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (*entity.TransferResult, error)
}

type AccountRepository interface {
	CreateAccount(ctx context.Context, req db.CreateAccountParams) (db.Account, error)
	GetAccount(ctx context.Context, id int64) (db.Account, error)
	ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error)
	CreateTransfer(ctx context.Context, arg db.CreateTransferParams) (db.Transfer, error)
}

type BusinessParams struct {
	AccountRepository AccountRepository
}

type business struct {
	accountRepo AccountRepository
}

func NewBusiness(params BusinessParams) AccountBiz {
	return &business{
		accountRepo: params.AccountRepository,
	}
}

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
	Balance  int64  `json:"balance"`
}

func (b *business) CreateAccount(ctx context.Context, req CreateAccountRequest) (*entity.Account, error) {
	account, err := b.accountRepo.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  req.Balance,
	})
	if err != nil {
		if pgError, ok := err.(*pg.Error); ok {
			switch pgError.Code.Name() {
			case "unique_violation":
				return nil, core.ErrForbidden.WithErrorf("account already exists: %s", err.Error())
			case "foreign_key_violation":
				return nil, core.ErrForbidden.WithErrorf("invalid owner: %s", err.Error())
			}
		}
		return nil, core.ErrInternalServerError.WithErrorf("failed to create account: %s", err.Error())
	}

	return &entity.Account{
		ID:       account.ID,
		Owner:    account.Owner,
		Currency: account.Currency,
	}, nil
}

func (b *business) GetAccount(ctx context.Context, ID int64) (*entity.Account, error) {
	account, err := b.accountRepo.GetAccount(ctx, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.ErrNotFound.WithError(err.Error())
		}
		return nil, core.ErrInternalServerError.WithErrorf("failed to get account: %s", err.Error())
	}

	return &entity.Account{
		ID:       account.ID,
		Owner:    account.Owner,
		Currency: account.Currency,
	}, nil
}

type ListAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

func (b *business) ListAccounts(ctx context.Context, req ListAccountRequest) ([]entity.Account, error) {
	accountModels, err := b.accountRepo.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.ErrNotFound.WithError(err.Error())
		}
		return nil, core.ErrInternalServerError.WithErrorf("failed to get account: %s", err.Error())
	}

	accounts := make([]entity.Account, len(accountModels))
	for i, account := range accountModels {
		accounts[i] = entity.Account{
			ID:       account.ID,
			Owner:    account.Owner,
			Currency: account.Currency,
		}
	}
	return accounts, nil
}
