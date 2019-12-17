package handler

import "github.com/gin-gonic/gin"

type transactionsHandler struct {

}

func NewTransactionsHandler() *transactionsHandler {
    return &transactionsHandler{}
}

func (this *transactionsHandler)GetTransactions(context *gin.Context) {

}

func (this *transactionsHandler)GetTransaction(context *gin.Context) {

}

func (this *transactionsHandler)CreateTransaction(context *gin.Context) {

}

