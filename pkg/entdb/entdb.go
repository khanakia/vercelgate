package entdb

import (
	"fmt"
	"log"

	"github.com/khanakia/vercelgate/gen/ent"
	"github.com/khanakia/vercelgate/pkg/vercelutil"
)

var client *ent.Client

func Client() *ent.Client {
	if client != nil {
		return client
	}

	dataSourceName := fmt.Sprintf("file:%s?cache=shared&_fk=1&_journal_mode=wal", DBfilePath())

	client, err := ent.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	return client
}

func DBfilePath() string {
	globalPath, _ := vercelutil.GetGlobalPathConfig()
	return globalPath + "/vercelgate.db"
}
