package shared

import (
	"fmt"
	"log"

	"github.com/kichikawa/ent"
)

func InitDB() (*ent.Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"db", "5432", "postgres", "development", "password"))

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	return client, nil
}
