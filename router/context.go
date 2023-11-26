package router

import "github.com/jmoiron/sqlx"

type context struct {
	DB *sqlx.DB
}

func NewContext(db *sqlx.DB) context {
	return context{
		DB: db,
	}
}
