package user

import (
	"context"

	"github.com/kichikawa/ent"
)

func (ui UserInfra) Create(user ent.User) (*ent.User, error) {
	ctx := context.Background()

	updateQuery := ui.db.User.Create()

	if user.Email != "" {
		updateQuery.SetEmail(user.Email)
	}

	if user.AccountName != "" {
		updateQuery.SetAccountName(user.AccountName)
	}

	if user.Age != 0 {
		updateQuery.SetAge(user.Age)
	}

	return updateQuery.Save(ctx)
}
