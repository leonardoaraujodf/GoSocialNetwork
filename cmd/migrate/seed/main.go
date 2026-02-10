package main

import (
	"log"

	"github.com/leonardoaraujodf/social/internal/db"
	"github.com/leonardoaraujodf/social/internal/store"
)

func main() {
	// addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	addr := "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store)
}
