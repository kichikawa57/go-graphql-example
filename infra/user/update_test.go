package user_test

import (
	"testing"

	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/user"
)

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		helper.TeadownFixture()

		first, _ := helper.Client.User.Query().First(helper.Ctx)

		if err := infra.Update(
			ent.User{
				Name: "更新",
			},
			user.IDEQ(first.ID),
		); err != nil {
			t.Fatalf("failed test %s", err)
		}

		firstCheck, _ := helper.Client.User.Query().First(helper.Ctx)

		if firstCheck.Name != "更新" {
			t.Fatalf("faile user Update")
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})

	t.Run("SuccessToEmpty", func(t *testing.T) {
		helper.TeadownFixture()

		if err := infra.Update(
			ent.User{
				Name: "更新",
			},
		); err != nil {
			t.Fatalf("failed test %s", err)
		}

		firstCheck, _ := helper.Client.User.Query().First(helper.Ctx)

		if firstCheck.Name != "更新" {
			t.Fatalf("faile user Update")
		}

		t.Cleanup(func() {
			helper.DeleteAll()
		})
	})
}
