package repositories

import (
    "banking-app/providers"
    "errors"
    "fmt"
)

type transactionRepository struct {
    db SqlProvider
}

func NewTransactionRepository() *transactionRepository {
    dbCon := providers.NewSqlProvider()
    return &transactionRepository{db: dbCon}
}

func (this *transactionRepository)Create(request CreateTransactionRequest) error {
    dbCon := this.db.Connect()
    query := fmt.Sprintf("INSERT INTO transactions (user_id, from_account_id, to_account_id, amount) VALUES ('%d', '%d', '%d', '%f')",request.UserId, request.FromAccountId, request.ToAccountId, request.Amount)
    row := dbCon.QueryRow(query)
    if row != nil{
        return errors.New("problem in insert")
    }
    return nil
}
