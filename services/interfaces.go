package services

import (
    "banking-app/models"
    "banking-app/providers"
    "banking-app/repositories"
)

type UserRepository interface {
    Add(user *models.User) error
    Remove(user *models.User) error
    List() []*models.User
    GetUser(apiKey string) (*models.User, error)
    GetUsers() ([]*models.User, error)
}

type AccountRepository interface {
    Add(user *models.Account) error
    Get(id int) (*models.Account, error)
}

type AuthProvider interface {
    AcquireSession(credentials providers.Credentials) (error, string)
}

type TransactionRepository interface {
    Create(repositories.CreateTransactionRequest) error
}
