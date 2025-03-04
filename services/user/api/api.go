package api

import (
	"net/http"

	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/pkg/core"
	userBiz "github.com/cs-tungthanh/Bank_Golang/services/user/business"
	"github.com/cs-tungthanh/Bank_Golang/token"
	"github.com/cs-tungthanh/Bank_Golang/util"
	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

type UserAPIParams struct {
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker

	UserBiz userBiz.Business
}

type api struct {
	userBiz    userBiz.Business
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewAPI(params UserAPIParams) UserAPI {
	return &api{
		config:     params.Config,
		store:      params.Store,
		tokenMaker: params.TokenMaker,

		userBiz: params.UserBiz,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (api *api) CreateUser(ctx *gin.Context) {
	var req userBiz.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.WriteErrorResponse(ctx, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	user, err := api.userBiz.CreateUser(ctx, req)
	if err != nil {
		util.WriteErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, core.ResponseData(user))
}

func (api *api) LoginUser(ctx *gin.Context) {
	var req userBiz.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.WriteErrorResponse(ctx, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	tokeResp, err := api.userBiz.LoginUser(ctx, req)
	if err != nil {
		util.WriteErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, core.ResponseData(tokeResp))
}
