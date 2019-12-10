package services

import(
  "banking-app/models"
)

type UserRepository interface {
 Add(user models.User) error
 Remove(user models.User) error
 List() []models.User
}

