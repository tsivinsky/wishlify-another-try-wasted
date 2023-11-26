package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tsivinsky/goenv"
	"github.com/tsivinsky/wishlify/db"
)

type Env struct {
	DBUser     string `env:"POSTGRES_USER,required"`
	DBPassword string `env:"POSTGRES_PASSWORD,required"`
	DBName     string `env:"POSTGRES_DB,required"`
	DBHost     string `env:"DB_HOST,required"`
}

var env = new(Env)

var (
	allowedHeaders = []string{"Accept", "Content-Type", "Content-Length", "Authorization"}
	allowedOrigins = []string{"*"}
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ","))
	w.Header().Set("Access-Control-Allow-Origin", strings.Join(allowedOrigins, ","))
}

func main() {
	goenv.MustLoad(env)

	db, err := db.Connect(env.DBHost, env.DBUser, env.DBPassword, env.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	err = http.ListenAndServe(":5000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		mux.ServeHTTP(w, r)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
