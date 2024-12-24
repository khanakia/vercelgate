package main

import (
	"context"
	"log"

	"github.com/khanakia/vercelgate/gen/ent"
	"github.com/khanakia/vercelgate/gen/ent/migrate"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	client, err := ent.Open("sqlite3", "file:vercelgate.db?cache=shared&_fk=1&_journal_mode=wal")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	err = client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false), // Disable foreign keys.
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
