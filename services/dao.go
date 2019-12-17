package services

type User struct {
    Name   string
    Email  string
    ApiKey string
}

type LoginRequest struct {
    Email    string
    Password string
}

type CreateAccountRequest struct {
    ApiKey string
    Amount float64
}

type CreateTransactionRequest struct {
    ApiKey        string
    Amount        float64
    FromAccountId int
    ToAccountId   int
}
