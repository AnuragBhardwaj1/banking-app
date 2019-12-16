package services

import (
    "banking-app/models"
    "banking-app/repositories"
    "banking-app/validators"
    "github.com/go-errors/errors"
)

type usersService struct {
    userRepo UserRepository
}

type CreateUserRequest struct {
    Email    string
    Password string
    UserName string
}

func NewUserService() *usersService {
    return &usersService{userRepo: repositories.NewUserRepository()}
}

func (this *usersService) CreateUser(createUserRequest *CreateUserRequest) error {
    user := models.NewUser(createUserRequest.Email, createUserRequest.Password, createUserRequest.UserName).WithGeneratedApiKey()
    userValidator := validators.NewUserValidator()
    err := userValidator.Validate(user)
    if err != nil {
        return errors.New("invalid user email")
    }

    err = this.userRepo.Add(user)
    if err != nil {
        return err
    }
    return nil
}

func (this *usersService) GetUsers(apiKey string) ([]*User, error) {
    //NextStep: check if user is admin
    _, err := this.userRepo.GetUser(apiKey)
    if err != nil {
        return []*User{}, err
    }
    users, err := this.userRepo.GetUsers()
    usersList := this.mapListUserDomainToDao(users)

    return usersList, nil
}

func (this *usersService) GetUser() (models.User, error) {
    return models.User{}, nil
}

//--------------------------------------------
func (this *usersService) mapListUserDomainToDao(users []*models.User) []*User {
    var usersList []*User
    for _, user := range users {
        user := User{
            Name:   user.Name(),
            Email:  user.Email(),
            ApiKey: user.ApiKey(),
        }
       usersList = append(usersList, &user)
    }
    return usersList
}
