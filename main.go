package main

import (
	"MVC-GOLANG-ORM/app/config"
	"MVC-GOLANG-ORM/app/controller"
	"MVC-GOLANG-ORM/app/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main()  {
	//err := godotenv.Load()go
	//if err != nil {
	//	panic(err)
	//}

	db := config.DBInit()
	accountController := controller.AccountController{DB: db}
	transactionController := controller.TransactionController{DB: db}

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/api/v1/account/add", accountController.CreateAccount)
	router.POST("/api/v1/login", accountController.Login)
	router.GET("/api/v1/account", middleware.Auth, accountController.GetAccount)
	router.POST("/api/v1/transfer", middleware.Auth, transactionController.Transfer)
	router.POST("/api/v1/withdraw", middleware.Auth, transactionController.Withdraw)
	router.POST("/api/v1/deposit", middleware.Auth, transactionController.Deposit)
	router.Run(":8000")
}