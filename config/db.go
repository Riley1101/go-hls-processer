package config

import _ "github.com/go-sql-driver/mysql"
import (
	"database/sql"
	"fmt"
)

type DB struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func (db *DB) Connect() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprint("root:admin@(127.0.0.1:3306)/vid"))
	return conn, err
}
