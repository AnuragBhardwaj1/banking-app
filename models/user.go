package models

import (
    "math/rand"
    "strconv"
)

type User struct {
    id       int
    name     string
    email    string
    apiKey   string
    password string
}

func (u *User) Id() int {
    return u.id
}

func (u *User) WithName(name string) *User {
    u.name = name
    return u
}

func (u *User) WithEmail(email string) *User {
    u.email = email
    return u
}

func (u *User) WithApiKey(apiKey string) *User {
    u.apiKey = apiKey
    return u
}

func (u *User) WithId(id int) *User {
    u.id = id
    return u
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

func (u *User) Password() string {
    return u.name
}

func (u *User) WithGeneratedApiKey() *User {
    apiKey := strconv.Itoa(rand.Intn(99999))
    u.setApikey(apiKey)
    return u
}

func (u *User) setApikey(key string) {
    u.apiKey = key
}

func NewUser(email string, password string, userName string) *User {
    return &User{email: email, password: password, name: userName}
}
