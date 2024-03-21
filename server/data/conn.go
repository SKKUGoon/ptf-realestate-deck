package data

import (
	"fmt"
	"log"
	"melp-back/ent"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func New() *ent.Client {
	// Create DSN Database Source Name
	dsnElement := []string{
		fmt.Sprintf("host=%s", os.Getenv("HOST")),
		fmt.Sprintf("user=%s", os.Getenv("USERNAME")),
		fmt.Sprintf("password=%s", os.Getenv("PASSWORD")),
		fmt.Sprintf("dbname=%s", os.Getenv("DATABASE")),
		fmt.Sprintf("search_path=%s", os.Getenv("SCHEMA")),
		"sslmode=disable",
	}
	dsn := strings.Join(dsnElement, " ")

	if client, err := ent.Open("postgres", dsn, ent.Debug()); err != nil {
		log.Fatalf("failed opening connection: %v", err)
		return nil
	} else {
		return client
	}
}
