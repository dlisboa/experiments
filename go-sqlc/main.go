package main

import (
	"context"
	"database/sql"
	"demo/db"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	postgres, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=sqlc_testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(postgres)

	ctx := context.Background()
	author, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "john lennon",
		Bio:  sql.NullString{String: "a singer", Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("inserted author: %v\n", author)
}
