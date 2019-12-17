package handler

import (
    "banking-app/services"
    _ "github.com/go-errors/errors"
)

type UserService interface {
    CreateUser(*services.CreateUserRequest) error
    GetUsers(apiKey string) ([]*services.User, error)
}


type SessionService interface {
    AcquireSession(services.LoginRequest) (string, error)
}

type AccountService interface {
    Create(services.CreateAccountRequest) error
}

type TransactionService interface {
    Create(string, services.CreateTransactionRequest) error
}
