package user_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/kichikawa/infra/user"
	"github.com/kichikawa/logger"
	"github.com/kichikawa/repository"
	"github.com/kichikawa/shared"
	helperDB "github.com/kichikawa/test/helper"
	_ "github.com/lib/pq"
)

var helper helperDB.DB
var infra repository.UserRepository

func TestMain(m *testing.M) {
	logger.SetUp()

	client, err := shared.InitDB()

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return
	}

	ctx := context.Background()

	helper = helperDB.DB{
		Ctx:    ctx,
		Client: client,
	}

	infra = user.NewUserInfra(helper.Ctx, helper.Client)

	helper.DeleteAll()

	run := m.Run()

	helperDB.Fixture(helper)

	os.Exit(run)

	client.Close()
}
