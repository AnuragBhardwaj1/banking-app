package handler

import (
    "banking-app/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

type transactionsHandler struct {
    transactionService TransactionService
}

func NewTransactionsHandler() *transactionsHandler {
    return &transactionsHandler{transactionService: services.NewTransactionsService()}
}

func (this *transactionsHandler)GetTransactions(ctx *gin.Context) {}

func (this *transactionsHandler)GetTransaction(ctx *gin.Context) {}

func (this *transactionsHandler)CreateTransaction(ctx *gin.Context) {
    var apiKeyHeader ApiKeyHeader
    var createTransactionRequest services.CreateTransactionRequest

    if err := ctx.ShouldBindHeader(&apiKeyHeader); err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
        return
    }

    if err := ctx.ShouldBindJSON(&createTransactionRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
    }

    createTransactionRequest = services.CreateTransactionRequest{
        ApiKey:        apiKeyHeader.ApiKey,
        Amount:        createTransactionRequest.Amount,
        FromAccountId: createTransactionRequest.FromAccountId,
        ToAccountId:   createTransactionRequest.ToAccountId,
    }
    if err := this.transactionService.Create(apiKeyHeader.ApiKey, createTransactionRequest); err != nil {
        ctx.JSON(http.StatusOK, gin.H{"error": err})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"msg": "Successfully Transferred amount"})
}

