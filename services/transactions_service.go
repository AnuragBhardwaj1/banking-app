package services

import (
    "banking-app/models"
    "banking-app/repositories"
    "banking-app/validators"
)

type transactionsService struct {
    userRepo UserRepository
    transactionRepo TransactionRepository
    accountRepo AccountRepository
}

func NewTransactionsService() *transactionsService {
    userRepo :=  repositories.NewUserRepository()
    transactionRepo  := repositories.NewTransactionRepository()
    accountsRepo  := repositories.NewAccountsRepository()
    return &transactionsService{userRepo: userRepo, transactionRepo: transactionRepo, accountRepo: accountsRepo}
}

func (this *transactionsService)Create(apiKey string, request CreateTransactionRequest) error {
    user, err := this.userRepo.GetUser(apiKey)
    if err != nil {
        return  err
    }
    createTransactionRequest := repositories.CreateTransactionRequest{
        UserId:        user.Id(),
        FromAccountId: request.FromAccountId,
        ToAccountId:   request.ToAccountId,
        Amount:        request.Amount,
    }
    transaction, err := this.buildTransaction(apiKey, createTransactionRequest)
    transactionValidator := validators.NewTransactionValidator()
    err = transactionValidator.Validate(transaction)
    if err != nil {
        return err
    }
    err = this.transactionRepo.Create(createTransactionRequest)
    if err != nil {
        return err
    }
    return nil
}

func (this *transactionsService) buildTransaction(apiKey string, request repositories.CreateTransactionRequest) (*models.Transaction, error) {
    user, err := this.userRepo.GetUser(apiKey)
    if err != nil {
        return &models.Transaction{}, err
    }
    fromAccount, err := this.accountRepo.Get(request.FromAccountId)
    if err != nil {
        return &models.Transaction{}, err
    }

    toAccount, err := this.accountRepo.Get(request.ToAccountId)
    if err != nil {
        return &models.Transaction{}, err
    }

    transaction := models.Transaction{
        FromAccount: fromAccount,
        ToAccount:   toAccount,
        Amount:      request.Amount,
        User:        user,
    }
    return &transaction, nil
}
