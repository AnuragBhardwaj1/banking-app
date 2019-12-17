package repositories

type createUserDao struct {
    name     string
    email    string
    password string
    apiKey   string
}

type readUserDao struct {
    id       int
    name     string
    password string
    email    string
    apiKey   string
}

type readUsersDao struct {
    name   string
    email  string
    apiKey string
}

type CreateTransactionRequest struct {
    UserId        int
    FromAccountId int
    ToAccountId   int
    Amount        float64
}
