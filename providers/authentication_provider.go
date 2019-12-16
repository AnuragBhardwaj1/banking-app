package providers

import "fmt"

type authProvider struct {
    db *sqlProvider
}

func NewAuthProvider() *authProvider {
    sqlProvider := NewSqlProvider()
    return &authProvider{db: sqlProvider}
}

// Todo: Needs to send event here
func (authProvider *authProvider) AcquireSession(credentials Credentials) (error, string) {
    var apiKey string
    db := authProvider.db.Connect()
    query := fmt.Sprintf("SELECT apiKey FROM users where email='%s' and password='%s'", credentials.Email, credentials.Password)
    err := db.QueryRow(query).Scan(&apiKey)
    if err != nil {
        return err, ""
    }
    return nil, apiKey
}
