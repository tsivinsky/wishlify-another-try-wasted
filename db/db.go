package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(dbHost, dbUser, dbPassword, dbName string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sqlx.Connect("postgres", dsn)

	return db, err
}
