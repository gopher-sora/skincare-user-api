package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "mysql_users_username"
	PASSWORD = "mysql_users_password"
	HOST     = "mysql_users_host"
	SCHEMA   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(USERNAME)
	password = os.Getenv(PASSWORD)
	host     = os.Getenv(HOST)
	schema   = os.Getenv(SCHEMA)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		"root", "p@ssw0rD", "127.0.0.1", "3306", "users")
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully configured")
}
