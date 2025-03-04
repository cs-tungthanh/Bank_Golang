package repository

import (
	"context"

	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
)

type postgresRepo struct {
	store db.Store
}

func NewPostgresRepository(store db.Store) *postgresRepo {
	return &postgresRepo{store: store}
}

func (repo *postgresRepo) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	return repo.store.CreateAccount(ctx, arg)
}

func (repo *postgresRepo) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	return repo.store.GetAccount(ctx, id)
}

func (repo *postgresRepo) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error) {
	return repo.store.ListAccounts(ctx, arg)
}

func (repo *postgresRepo) CreateTransfer(ctx context.Context, arg db.CreateTransferParams) (db.Transfer, error) {
	return repo.store.CreateTransfer(ctx, arg)
}
