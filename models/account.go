package models

type Account struct {
    userId int
    amount float64
}

func (a *Account) UserId() int {
    return a.userId
}

func (a *Account) SetUserId(userId int) {
    a.userId = userId
}

func (a *Account) Amount() float64 {
    return a.amount
}

func (a *Account) setAmount(amount float64) {
    a.amount = amount
}

func NewAccount(amount float64) *Account {
    return &Account{amount: amount}
}


