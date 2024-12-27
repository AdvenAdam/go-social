package main

import (
	"log"
	"time"

	"github.com/AdvenAdam/go-social/internal/db"
	"github.com/AdvenAdam/go-social/internal/env"
	"github.com/AdvenAdam/go-social/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:           env.GetString("DB_ADDR", "postgres://psuser:pspass@localhost:5432/social?sslmode=disable"),
			maxOpenConns:   env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns:   env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTimeout: time.Duration(env.GetInt("DB_MAX_IDLE_TIMEOUT", 15)) * time.Minute,
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTimeout,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Printf("Connected to database %s", cfg.db.addr)

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
