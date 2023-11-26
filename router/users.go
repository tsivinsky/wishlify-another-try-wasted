package router

import (
	"net/http"

	"github.com/tsivinsky/wishlify/db"
)

var HandleGetUsers Handler = func(ctx context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := []db.User{}
		err := ctx.DB.Select(&users, "SELECT id, email, login, created_at, updated_at FROM users")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		SendJSON(w, users)
	}
}
