package composer

import (
	"fmt"

	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/token"
	"github.com/cs-tungthanh/Bank_Golang/util"

	accountAPI "github.com/cs-tungthanh/Bank_Golang/services/account/api"
	accountBiz "github.com/cs-tungthanh/Bank_Golang/services/account/business"
	accountRepo "github.com/cs-tungthanh/Bank_Golang/services/account/repository"
	userAPI "github.com/cs-tungthanh/Bank_Golang/services/user/api"
	userBiz "github.com/cs-tungthanh/Bank_Golang/services/user/business"
	userRepo "github.com/cs-tungthanh/Bank_Golang/services/user/repository"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ComposedService struct {
	UserAPI    userAPI.UserAPI
	AccountAPI accountAPI.AccountAPI
}

func ComposeAPIService(cfg util.Config, store db.Store) (*ComposedService, error) {
	tokenMaker, err := token.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// Register custom validator for gin Global
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	userRepo := userRepo.NewPostgresRepository(store)
	userBiz := userBiz.NewBusiness(userBiz.BusinessParams{
		UserRepository: userRepo,
	})
	userAPI := userAPI.NewAPI(userAPI.UserAPIParams{
		Store:      store,
		TokenMaker: tokenMaker,
		Config:     cfg,
		UserBiz:    userBiz,
	})

	accountRepo := accountRepo.NewPostgresRepository(store)
	accountBiz := accountBiz.NewBusiness(accountBiz.BusinessParams{
		AccountRepository: accountRepo,
	})
	accountAPI := accountAPI.NewAPI(accountAPI.AccountAPIParams{
		AccountBiz: accountBiz,
	})

	return &ComposedService{
		UserAPI:    userAPI,
		AccountAPI: accountAPI,
	}, nil
}

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportCurrency(currency)
	}
	return false
}
