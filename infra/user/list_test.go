package user_test

import (
	"testing"

	"github.com/kichikawa/ent/schema/property"
	"github.com/kichikawa/ent/user"
	_ "github.com/lib/pq"
)

func TestList(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		helper.TeadownFixture()
		users, userErr := infra.List(user.EmailEQ(property.UserEmail("hoge1@email")))

		if userErr != nil {
			t.Fatalf("failed test %s", userErr)
		}

		if len(users) != 1 {
			t.Fatalf("failed list len")
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})

	t.Run("SuccessToEmpry", func(t *testing.T) {
		helper.TeadownFixture()
		users, userErr := infra.List()

		if userErr != nil {
			t.Fatalf("failed test %s", userErr)
		}

		if len(users) != 5 {
			t.Fatalf("failed list len")
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})
}
