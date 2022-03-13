package user_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kichikawa/ent/schema/property"
	"github.com/kichikawa/ent/user"
	_ "github.com/lib/pq"
)

func TestShow(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		helper.TeadownFixture()
		user, userErr := infra.Show(user.EmailEQ(property.UserEmail("hoge2@email")))

		if userErr != nil {
			t.Fatalf("failed test %s", userErr)
		}

		if user.Email != "hoge2@email" {
			t.Fatalf("failed list len")
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})

	t.Run("FailToEmpry", func(t *testing.T) {
		helper.TeadownFixture()
		_, userErr := infra.Show()

		if userErr == nil {
			t.Fatalf("failed error empty")
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})

	t.Run("FailToNotFind", func(t *testing.T) {
		helper.TeadownFixture()
		_, userErr := infra.Show(user.IDEQ(uuid.New()))

		if userErr == nil {
			t.Fatalf("failed error empty")
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})
}
