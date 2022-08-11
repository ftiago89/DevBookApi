package database

import (
	"database/sql"
	"devbookapi/src/config"

	_ "github.com/go-sql-driver/mysql" //Driver mysql
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DB_URI)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
