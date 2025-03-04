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

func (repo *postgresRepo) CreateUser(ctx context.Context, data *db.CreateUserParams) (db.User, error) {
	return repo.store.CreateUser(ctx, *data)
}

func (repo *postgresRepo) GetUser(ctx context.Context, username string) (db.User, error) {
	return repo.store.GetUser(ctx, username)
}
