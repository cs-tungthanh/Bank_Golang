package api

import (
	"net/http"

	"github.com/cs-tungthanh/Bank_Golang/pkg/core"
	accountBiz "github.com/cs-tungthanh/Bank_Golang/services/account/business"
	"github.com/cs-tungthanh/Bank_Golang/util"

	"github.com/gin-gonic/gin"
)

type AccountAPI interface {
	CreateAccount(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
	ListAccount(ctx *gin.Context)

	CreateTransfer(ctx *gin.Context)
}

type api struct {
	accountBiz accountBiz.AccountBiz
}

type AccountAPIParams struct {
	AccountBiz accountBiz.AccountBiz
}

func NewAPI(params AccountAPIParams) AccountAPI {
	return &api{
		accountBiz: params.AccountBiz,
	}
}

func (api *api) CreateAccount(ctx *gin.Context) {
	var req accountBiz.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.WriteErrorResponse(ctx, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	account, err := api.accountBiz.CreateAccount(ctx, accountBiz.CreateAccountRequest{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  req.Balance,
	})
	if err != nil {
		util.WriteErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, core.ResponseData(account))
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (api *api) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		util.WriteErrorResponse(ctx, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	account, err := api.accountBiz.GetAccount(ctx, req.ID)
	if err != nil {
		util.WriteErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (api *api) ListAccount(ctx *gin.Context) {
	var req accountBiz.ListAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		util.WriteErrorResponse(ctx, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	accounts, err := api.accountBiz.ListAccounts(ctx, req)
	if err != nil {
		util.WriteErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, core.ResponseData(accounts))
}

func (api *api) CreateTransfer(ctx *gin.Context) {
	var req accountBiz.CreateTransferParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.WriteErrorResponse(ctx, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	result, err := api.accountBiz.CreateTransfer(ctx, req)
	if err != nil {
		util.WriteErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, core.ResponseData(result))
}
