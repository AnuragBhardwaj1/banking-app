package handler

import "github.com/gin-gonic/gin"

type accountsHandler struct {
}

func NewAccountsHandler() *accountsHandler {
    return &accountsHandler{}
}

func (this *accountsHandler)GetAccounts(context *gin.Context) {
    context.JSON(200, gin.H{"hello": "world"})
}

func (this *accountsHandler)GetAccount(context *gin.Context) {

}

func (this *accountsHandler)CreateAccount(context *gin.Context) {

}

func (this *accountsHandler)PatchAccount(context *gin.Context) {

}

func (this *accountsHandler)DeleteAccount(context *gin.Context) {

}
