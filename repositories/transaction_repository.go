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

func (this *transactionRepository) Create(request CreateTransactionRequest) error {
    var id int
    dbCon := this.db.Connect()
    query := fmt.Sprintf("BEGIN;"+
        "UPDATE accounts SET amount = amount - 100.00  WHERE id= '%d';"+
        "UPDATE accounts SET amount = amount + 100.00  WHERE id= '%d';"+
        "INSERT INTO transactions (user_id, from_account_id, to_account_id, amount) VALUES ('%d', '%d', '%d', '%f') returning id;" +
        "COMMIT;",
        request.FromAccountId, request.ToAccountId, request.FromAccountId, request.FromAccountId, request.ToAccountId, request.Amount)

    row := dbCon.QueryRow(query).Scan(&id)
    if row != nil {
        return errors.New("problem in insert")
    }
    return nil
}
