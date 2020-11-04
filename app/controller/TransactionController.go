package controller

import (
	"MVC-GOLANG-ORM/app/model"
	"MVC-GOLANG-ORM/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type TransactionController struct {
	DB *gorm.DB
}

func (ctrl TransactionController) Transfer (ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Transfer(trx)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WarpAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WarpAPISucces(ctx, "success", http.StatusOK)
	return
}

func (ctrl TransactionController) Withdraw (ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Withdraw(trx)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WarpAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WarpAPISucces(ctx, "success", http.StatusOK)
	return
}

func (ctrl TransactionController) Deposit (ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Deposit(trx)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WarpAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	return
}