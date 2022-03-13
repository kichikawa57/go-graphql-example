package user_test

import (
	"testing"

	"github.com/kichikawa/ent"
	_ "github.com/lib/pq"
)

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		_, userErr := infra.Create(ent.User{
			Name:        "test",
			Email:       "k.ichikawa@attivita.co.jp",
			AccountName: "アカウント名",
			Password:    "neon06240708",
			Age:         102,
		})

		if userErr != nil {
			t.Fatalf("failed test %s", userErr)
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})
}
