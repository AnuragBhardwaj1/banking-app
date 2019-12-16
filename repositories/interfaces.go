package repositories

import (
    "database/sql"
)

type SqlProvider interface {
    Connection() *sql.DB
    Connect() *sql.DB
    Execute(query string) error
    Close() error
}
