package database

import (
	"database/sql"
	"devbook-api/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.BankConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
	}

	return db, nil
}
