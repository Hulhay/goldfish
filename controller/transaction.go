package controller

import (
	"net/http"

	"github.com/Hulhay/goldfish/shared"
	"github.com/Hulhay/goldfish/usecase"
	"github.com/Hulhay/goldfish/usecase/transaction"
	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionUC usecase.Transaction
}

type TransactionController interface {
	CreateTransaction(ctx *gin.Context)
}

func NewTransactionRepository(transactionUC usecase.Transaction) TransactionController {
	return &transactionController{
		transactionUC: transactionUC,
	}
}

func (c *transactionController) CreateTransaction(ctx *gin.Context) {

	var (
		params transaction.CreateTransactionRequest
		err    error
	)

	err = ctx.ShouldBind(&params)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.transactionUC.CreateTransaction(ctx, params)
	if err != nil {
		res := shared.BuildErrorResponse("Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Success!", nil)
	ctx.JSON(http.StatusOK, res)

}
