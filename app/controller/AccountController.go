package controller

import (
	"MVC-GOLANG-ORM/app/model"
	"MVC-GOLANG-ORM/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AccountController struct {
	DB *gorm.DB
}

func (ctrl AccountController) CreateAccount(ctx *gin.Context) {
	accountModel := model.AccountModel{
		DB: ctrl.DB,
	}

	var account model.Account
	err := ctx.Bind(&account)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	hashPassword, err := utils.HashGenerator(account.Password)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}
	account.Password = hashPassword
	flag, err := accountModel.InsertNewAccount(account)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}
	if !flag {
		utils.WarpAPIError(ctx, "unknown failed to insert account", http.StatusInternalServerError)
		return
	}

	utils.WarpAPISucces(ctx, "success", http.StatusOK)

}
func (ctrl AccountController) Login(ctx *gin.Context) {
	authModel := model.AuthModel{
		DB: ctrl.DB,
	}
	var auth model.Auth

	err := ctx.Bind(&auth)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err, token := authModel.Login(auth)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WarpAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WarpAPIData(ctx, map[string]interface{}{
		"token": token,
	}, http.StatusOK, "success")

	log.Println("Login")
}


func (ctrl AccountController) GetAccount(ctx *gin.Context) {
	idAccount := ctx.MustGet("account_number").(int)
	accountModel := model.AccountModel{
		DB: ctrl.DB,
	}
	flag, err, transactions, account := accountModel.GetAccountDetail(idAccount)
	if err != nil {
		utils.WarpAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WarpAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WarpAPIData(ctx, map[string]interface{}{
		"account": account,
		"transaction": transactions,
	}, http.StatusOK, "success")
	return
}