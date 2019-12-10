package repositories

import(
    "banking-app/models"
    _ "github.com/go-errors/errors"
)

type userRepository struct {
  users []*models.User
}

func NewUserRepository() *userRepository {
 return &userRepository{users: []*models.User{}}
}

func(userRepo *userRepository) Add(user models.User) error {
 //users, ok := append(users, user); ok {
 //  return nil
 //}
 // return errors.New("User not added to userRepository")
 return nil
}
//
func(userRepo *userRepository) Remove(user models.User) error {
 //index := userRepo.findIndex(user)
 //users, ok := append(users, user); ok {
 //  return nil
 //}
 // return errors.New("User not added to userRepository")
 return nil
}

func (userRepo *userRepository) List() []models.User{
  return []models.User{}
}

// func (userRepo *userRepository) findIndex(user User) int, err {
// // To be implemented
//  return 1, nil
// }
