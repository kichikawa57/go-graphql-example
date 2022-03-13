package user

import (
	"github.com/kichikawa/ent"
	"github.com/kichikawa/logger"
	"github.com/kichikawa/shared"
)

func (ui UserInfra) Create(user ent.User) (*ent.User, error) {
	logger.Infof("Infra user Create")

	updateQuery := ui.db.User.Create()

	if user.Name != "" {
		updateQuery = updateQuery.SetName(user.Name)
	}

	if user.Email != "" {
		updateQuery = updateQuery.SetEmail(user.Email)
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

	createUser, createUserErr := updateQuery.Save(ui.ctx)
	if createUserErr != nil {
		logger.Errorf("Infra user Create to %s", createUserErr)
		return nil, createUserErr
	}

	return createUser, nil
}
