package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection() (*sqlx.DB, error) {
	dsn := "user=user password=password dbname=productdb host=localhost port=5432 sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}