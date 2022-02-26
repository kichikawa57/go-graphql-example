package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/migrate"
	"github.com/kichikawa/logger"
	"github.com/kichikawa/router"
	_ "github.com/lib/pq"
)

func main() {
	err := logger.SetUp()
	if err != nil {
		panic(err)
	}

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"db", "5432", "postgres", "development", "password"))

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	ctx := context.Background()

	if err := client.Schema.Create(
		ctx,
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	// infra := user.NewUserInfra(client)

	// if _, err := client.
	// 	User.
	// 	Create().
	// 	SetAge(10).
	// 	SetEmail("test@attivita.co.jp").
	// 	SetAccountName("testA").
	// 	SetStatus(user.StatusInProgress).
	// 	Save(ctx); err != nil {
	// 	log.Fatalf("エラーだよ ===============>", err)
	// }

	// res, err := infra.Create(ent.User{
	// 	AccountName: "test02",
	// 	Age:         32,
	// 	Email:       "hoge02@gmail.com",
	// })

	// res, err := infra.Update(ent.User{
	// 	Age:   14,
	// 	Email: "hoge01@gmail.com",
	// })

	if err != nil {
		logger.Error("can't initialize zap logger:", err)
	}

	// logger.Info("mail: ", res)

	r := router.SetupRouter()
	r.Run(":8080")
}
