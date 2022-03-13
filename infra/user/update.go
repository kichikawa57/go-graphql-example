package user

import (
	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/predicate"
	"github.com/kichikawa/logger"
	"github.com/kichikawa/shared"
)

func (ui UserInfra) Update(user ent.User, where ...predicate.User) error {
	logger.Infof("Infra user Update")

	updateQuery := ui.db.User.Update()

	if len(where) > 0 {
		updateQuery = updateQuery.Where(where...)
	}

	if user.Email != "" {
		updateQuery = updateQuery.SetEmail(user.Email)
	}

	if user.Name != "" {
		updateQuery = updateQuery.SetName(user.Name)
	}

	if user.AccountName != "" {
		updateQuery = updateQuery.SetAccountName(user.AccountName)
	}

	if user.Password != "" {
		password, _ := shared.PasswordEncrypt(user.Password)
		updateQuery = updateQuery.SetPassword(password)
	}

	if user.Age != 0 {
		updateQuery = updateQuery.SetAge(user.Age)
	}

	if _, err := updateQuery.Save(ui.ctx); err != nil {
		logger.Errorf("Infra user Update to %s", err)
		return nil
	}

	return nil
}
