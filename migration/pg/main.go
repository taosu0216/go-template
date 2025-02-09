package main

import (
	"context"
	_ "github.com/lib/pq"
	"go-template/data/db/ent"
	"log"
)

const (
	defaultPG = "postgres://postgres:pass@localhost:15432/template?sslmode=disable"
)

func main() {
	client, err := ent.Open("postgres", defaultPG)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
