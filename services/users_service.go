package services

import(
  "banking-app/validators"
  "github.com/go-errors/errors"
  "banking-app/repositories"
  "banking-app/models"
)



type usersService struct {
  userRepo UserRepository
}

type CreateUserRequest struct {
  email string
  password string
}

func NewUserService() *usersService {
    return &usersService{userRepo: repositories.NewUserRepository()}
}

func (userService *usersService)CreateUser(createUserRequest *CreateUserRequest) error  {
  user := models.User{createUserRequest.Name, createUserRequest.Email, ""}
  userValidator := validators.NewUserValidator()
  err := userValidator.Validate(createUserRequest.email)
  if err != nil {
    return  errors.New("invalid user email")
  }
  return nil
}

func (userService *usersService)GetUser() (models.User, error) {
  return models.User{}, nil
}

