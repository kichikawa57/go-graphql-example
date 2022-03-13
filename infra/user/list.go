package user

import (
	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/predicate"
	"github.com/kichikawa/logger"
)

func (ui UserInfra) List(where ...predicate.User) ([]*ent.User, error) {
	logger.Infof("Infra user List")

	db := ui.db.User.Query()

	if len(where) > 0 {
		db = db.Where(where...)
	}

	users, usersErr := db.All(ui.ctx)

	if usersErr != nil {
		logger.Errorf("Infra user List to %s", usersErr)
		return nil, usersErr
	}

	return users, nil
}
