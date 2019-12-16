package handler

import (
    "banking-app/services"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

type sessionsHandler struct {
    sessionService SessionService
}

func NewSessionsHandler() *sessionsHandler {
    return &sessionsHandler{sessionService: services.NewSessionsService()}
}

func (sessionsHandler *sessionsHandler) AcquireSession(ctx *gin.Context) {
    var loginDetails Login
    if err := ctx.ShouldBindJSON(&loginDetails); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
        return
    }
    loginRequest := services.LoginRequest{Email: loginDetails.Email, Password: loginDetails.Password}
    apiKey, err := sessionsHandler.sessionService.AcquireSession(loginRequest)
    if err != nil {
        ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("%+v", err)})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"api_key": fmt.Sprintf("%+v", apiKey)})
    return
}
