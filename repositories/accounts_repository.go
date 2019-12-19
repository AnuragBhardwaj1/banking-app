package repositories
import (
    "banking-app/models"
    "banking-app/providers"
    "fmt"
    _ "github.com/go-errors/errors"
)

type accountsRepository struct {
    users []*models.User
    db    SqlProvider
}

func NewAccountsRepository() *accountsRepository {
    dbCon := providers.NewSqlProvider()
    return &accountsRepository{db: dbCon}
}

func (this *accountsRepository) Add(account *models.Account) error {
    var id int
    db := this.db.Connect()
    query := fmt.Sprintf("INSERT INTO accounts (amount, user_id) VALUES ('%f', '%d') RETURNING id", account.Amount, account.UserId)
    err := db.QueryRow(query).Scan(&id)
    if err != nil {
        return err
    }
    return nil
}

func (this *accountsRepository) Get(id int) (*models.Account, error) {
    var account models.Account
    db := this.db.Connect()
    query := fmt.Sprintf("SELECT id, user_id, amount FROM accounts where id='%d'", id)
    err := db.QueryRow(query).Scan(&account.Id, &account.UserId, &account.Amount)
    if err != nil {
        return &models.Account{}, err
    }
    return &account, nil
}
