package models

type User struct {
    name   string
    email  string
    apiKey string
    password string
}

func (u *User) ApiKey() string {
    return u.apiKey
}

func (u *User) Email() string {
    return u.email
}

func (u *User) Name() string {
    return u.name
}

func NewUser(name string, email string) *User {
    return &User{name: name, email: email}
}
