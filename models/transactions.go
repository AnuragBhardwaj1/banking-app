package models

type Transaction struct {
    FromAccount *Account
    ToAccount *Account
    Amount float64
    User *User
}
