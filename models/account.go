package models

type Account struct {
    Id     int
    UserId int
    Amount float64
}

func (a *Account) SetUserId(userId int) {
    a.UserId = userId
}

func (a *Account) setAmount(amount float64) {
    a.Amount = amount
}

func NewAccount(amount float64) *Account {
    return &Account{Amount: amount}
}
