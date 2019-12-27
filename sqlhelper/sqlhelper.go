package sqlhelper

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetDBConnection Initialize DB Connection
func GetDBConnection(databasename string, username string, password string, host string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, databasename)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
