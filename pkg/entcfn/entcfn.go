package entcfn

import (
	"app/gen/ent/migrate"
	"app/pkg/entdb"
	"context"
	"log"
)

func Migrate() error {
	client := entdb.Client()

	// Run the auto migration tool.
	err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false), // Disable foreign keys.
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return err
	}

	client.Close()
	return nil
}
