package providers

import(
  "banking-app/constants"
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

type sqlProvider struct {
  host     string
  port     string
  userName string
  password string
  dbCon    *sql.DB
}


func NewSqlProvider() *sqlProvider {
  return &sqlProvider{host: constants.HOST, port: constants.PORT, userName: constants.USERNAME, password: "anu_12345" }
}


func (this *sqlProvider)Connect() *sql.DB {
  connStr := fmt.Sprintf("user=anurag password=anu_12345 dbname=bankingapp sslmode=disable")
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    return nil
  }
  this.dbCon = db
  return this.dbCon
}

func (this *sqlProvider)Connection() *sql.DB {
  return this.dbCon
}

func(this *sqlProvider)Execute(query string) error {
  rows, err := this.dbCon.Query(query)
  fmt.Println(rows)
  return err
}
func (this *sqlProvider)Close() error {
  connStr := fmt.Sprintf("user=pqgotest dbname=pqgotest", this.userName, "bookingapp")
  db, err := sql.Open("postgres", connStr)
  fmt.Println(db)
  if err != nil {
    return err
  }
  return nil
}
