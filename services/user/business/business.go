package business

import (
	"context"
	"database/sql"

	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/pkg/core"
	"github.com/cs-tungthanh/Bank_Golang/services/user/entity"
	"github.com/cs-tungthanh/Bank_Golang/token"
	"github.com/cs-tungthanh/Bank_Golang/util"

	pg "github.com/lib/pq"
)

type Business interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*entity.User, error)
	LoginUser(ctx context.Context, req LoginUserRequest) (*entity.TokenResponse, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, data *db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, username string) (db.User, error)
}

type business struct {
	userRepo   UserRepository
	tokenMaker token.Maker
	config     util.Config
}

type BusinessParams struct {
	UserRepository UserRepository
	TokenMaker     token.Maker
	Config         util.Config
}

func NewBusiness(params BusinessParams) *business {
	return &business{
		userRepo:   params.UserRepository,
		config:     params.Config,
		tokenMaker: params.TokenMaker,
	}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (biz *business) CreateUser(ctx context.Context, req CreateUserRequest) (*entity.User, error) {
	hashedPw, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, core.ErrInternalServerError.WithErrorf("failed to hash password: %s", err.Error())
	}

	user, err := biz.userRepo.CreateUser(ctx, &db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPw,
		FullName:       req.FullName,
		Email:          req.Email,
	})
	if err != nil {
		if pgError, ok := err.(*pg.Error); ok {
			switch pgError.Code.Name() {
			case "unique_violation":
				return nil, core.ErrConflict.WithErrorf("user already exists")
			}
		}
		return nil, core.ErrInternalServerError.WithErrorf("failed to create user: %s", err.Error())
	}

	return &entity.User{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func (biz *business) LoginUser(ctx context.Context, req LoginUserRequest) (*entity.TokenResponse, error) {
	user, err := biz.userRepo.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.ErrNotFound.WithError(err.Error())
		}
		return nil, core.ErrInternalServerError.WithErrorf("failed to get user: %s", err.Error())
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, core.ErrUnauthorized.WithError(err.Error())
	}

	accessToken, err := biz.tokenMaker.CreateToken(
		req.Username,
		biz.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, core.ErrInternalServerError.WithErrorf("failed to create token: %s", err.Error())
	}

	return &entity.TokenResponse{
		AccessToken: entity.Token{
			Token:     accessToken,
			ExpiredIn: int(biz.config.AccessTokenDuration.Seconds()),
		},
		User: entity.User{
			Username: user.Username,
			FullName: user.FullName,
			Email:    user.Email,
		},
	}, nil
}
