package handler

import(
  "banking-app/models"
  _ "github.com/go-errors/errors"
)

type UserService interface {
    CreateUser(user models.User) error
    GetUser() (models.User, error)
}