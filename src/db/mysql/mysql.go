package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Client is a pointer to sql.DB
var (
	Client *sql.DB
)

// Connect method initialize connection to DB
func Connect() {
	var (
		username = os.Getenv("MYSQL_USERNAME")
		password = os.Getenv("MYSQL_PASSWORD")
		host     = os.Getenv("MYSQL_HOST")
		schema   = os.Getenv("MYSQL_SCHEMA")

		err error
	)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to database")
}
