package main

import (
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/tsivinsky/goenv"
	"github.com/tsivinsky/wishlify/db"
	"github.com/tsivinsky/wishlify/env"
)

func main() {
	flag.Parse()

	goenv.MustLoad(env.Env)

	db, err := db.Connect(env.Env.DBHost, env.Env.DBUser, env.Env.DBPassword, env.Env.DBName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}

	cmd := flag.Arg(0)
	if cmd == "" {
		log.Fatal("No command provided")
	}

	switch cmd {
	case "up":
		err = m.Up()
		break
	case "down":
		err = m.Down()
		break
	}
	if err != nil {
		panic(err)
	}
}
