package configmanager

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDatabse() *sql.DB {
	fmt.Println("Database connected.")
	db, err = sql.Open("mysql", "root:govind@/contactsystem")
	return db
}
