package repositories

import (
    "banking-app/models"
    "banking-app/providers"
    "fmt"
    _ "github.com/go-errors/errors"
)

type userRepository struct {
    users []*models.User
    db    SqlProvider
}

func NewUserRepository() *userRepository {
    dbCon := providers.NewSqlProvider()
    return &userRepository{users: []*models.User{}, db: dbCon}
}

func (this *userRepository) Add(user *models.User) error {
    db := this.db.Connect()
    query := fmt.Sprintf("INSERT INTO users (name, email, password, api_key) VALUES ('%s', '%s', '%s', '%s')", user.Name(), user.Email(), user.Password(), user.ApiKey())
    _, err := db.Query(query)
    if err != nil {
        return err
    }
    return nil
}

func (this *userRepository) Remove(user *models.User) error {
    return nil
}

func (this *userRepository) List() []*models.User {
    return []*models.User{}
}

func (this *userRepository) Authenticate(email string, password string) error {
     db := this.db.Connect()
    query := fmt.Sprintf("SELECT * FROM users where email=%s and password%s limit 1", email, password)
    _, err := db.Query(query)
    if err != nil {
        return err
    }
    return nil
}
func (this *userRepository)GetUser(apiKey string) (*models.User, error) {
   var user readUserDao
    db := this.db.Connect()
   query := fmt.Sprintf("SELECT id, name, email, api_key FROM users where api_key='%s'", apiKey)
   err := db.QueryRow(query).Scan(&user.id, &user.name, &user.email, &user.apiKey)
   if err != nil {
       return &models.User{}, err
   }
   userModel := &models.User{}
   userModel.WithName(user.name).WithEmail(user.email).WithApiKey(user.apiKey).WithId(user.id)
    return userModel, nil
}

func (this *userRepository)GetUsers() ([]*models.User, error) {
   var user readUsersDao
   var users []readUsersDao
    db := this.db.Connect()
   query := fmt.Sprintf("SELECT name, email, api_key FROM users")
   rows, err := db.Query(query)
   if err != nil {
       return []*models.User{}, err
   }

   for rows.Next() {
       err := rows.Scan(&user.name, &user.email, &user.apiKey)
       if err != nil {
           return []*models.User{}, err
       }
       users  = append(users, user)
   }

   domainUsers := this.mapUsersDaoToModel(users)
    return domainUsers, nil
}

//---------------------------------------------------------------------
func (this *userRepository) createUserDao(user *models.User) *createUserDao {
    return &createUserDao{email: user.Email(), password: user.Password(), apiKey: user.ApiKey()}

}

func (this *userRepository) mapUsersDaoToModel(users []readUsersDao) []*models.User {
    var domainUsers []*models.User
    for _, user := range users {
        domainUser := &models.User{}
        domainUser.WithName(user.name).WithApiKey(user.apiKey).WithEmail(user.email)
        domainUsers = append(domainUsers, domainUser)
    }
    return domainUsers
}

