package utils

import (
	"database/sql"
	"fmt"
)

type DB struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func (db *DB) Connect() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.Username, db.Password, db.Host, db.Port, db.Name))
	return conn, err
}
