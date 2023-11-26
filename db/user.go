package db

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Email     *string   `db:"email" json:"email"`
	Login     string    `db:"login" json:"login"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
