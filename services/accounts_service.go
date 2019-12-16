package services

import (
    "banking-app/models"
    "banking-app/repositories"
)

type AccountsService struct {
    userRepo    UserRepository
    accountRepo AccountRepository
}

func NewAccountsService() *AccountsService {
    return &AccountsService{userRepo: repositories.NewUserRepository(), accountRepo: repositories.NewAccountsRepository()}
}

func (this *AccountsService) Create(request CreateAccountRequest) error {
    user, err := this.userRepo.GetUser(request.ApiKey)
    if err != nil {
        return err
    }
    account := this.buildAccount(request, user)
    err = this.accountRepo.Add(account)
    if err != nil {
        return err
    }
    return nil
}

func (this *AccountsService) buildAccount(request CreateAccountRequest, user *models.User) *models.Account {
    account := models.NewAccount(request.Amount)
    account.SetUserId(user.Id())
    return account
}

//-------------------------------------------------------
