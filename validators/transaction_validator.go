package validators

import (
    "banking-app/models"
    "errors"
)

type transactionValidator struct {
}

func NewTransactionValidator() *transactionValidator {
    return &transactionValidator{}
}

func (this *transactionValidator) Validate(transaction *models.Transaction) error {
    if transaction.User.Id() != transaction.FromAccount.UserId {
        return errors.New("one can only transfer from own account")
    }

    if transaction.FromAccount.Amount < transaction.Amount {
        return errors.New("not sufficient balance to transfer")
    }

    return nil
}
