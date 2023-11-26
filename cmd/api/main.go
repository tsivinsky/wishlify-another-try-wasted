package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tsivinsky/goenv"
	"github.com/tsivinsky/wishlify/db"
	"github.com/tsivinsky/wishlify/env"
	"github.com/tsivinsky/wishlify/router"
)

var (
	allowedHeaders = []string{"Accept", "Content-Type", "Content-Length", "Authorization"}
	allowedOrigins = []string{"*"}
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ","))
	w.Header().Set("Access-Control-Allow-Origin", strings.Join(allowedOrigins, ","))
}

func main() {
	goenv.MustLoad(env.Env)

	db, err := db.Connect(env.Env.DBHost, env.Env.DBUser, env.Env.DBPassword, env.Env.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	ctx := router.NewContext(db)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			router.HandleGetUsers(ctx)(w, r)
		}
	})

	err = http.ListenAndServe(":5000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		mux.ServeHTTP(w, r)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
