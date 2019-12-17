package handler

import (
    "banking-app/services"
    "github.com/gin-gonic/gin"
    "strconv"
)

type accountsHandler struct {
    accountService AccountService
}

func NewAccountsHandler() *accountsHandler {
    accountsService := services.NewAccountsService()
    return &accountsHandler{accountService: accountsService}
}

func (this *accountsHandler)GetAccount(context *gin.Context) {

}

func (this *accountsHandler) CreateAccount(context *gin.Context) {
    var apiKeyHeader ApiKeyHeader
    var createAccountDetails CreateAccountDetails

    if err := context.ShouldBindHeader(&apiKeyHeader); err != nil {
        context.JSON(200, err)
        return
    }
    if err := context.ShouldBindJSON(&createAccountDetails); err != nil {
        context.JSON(200, err)
        return
    }
    amount, _ := strconv.ParseFloat(createAccountDetails.Amount, 64)
    var createAccountRequest = services.CreateAccountRequest{Amount: amount,ApiKey: apiKeyHeader.ApiKey}
    err := this.accountService.Create(createAccountRequest)
    if err != nil {
        context.JSON(200, gin.H{"error": err.Error()})
        return
    }
    context.JSON(200, gin.H{"msg": "Account Successfully Created"})
}

func (this *accountsHandler)PatchAccount(context *gin.Context) {

}
