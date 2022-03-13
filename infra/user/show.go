package user

import (
	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/predicate"
	"github.com/kichikawa/logger"
)

func (ui UserInfra) Show(where ...predicate.User) (*ent.User, error) {
	logger.Infof("Infra user Show")

	db := ui.db.User.Query()

	if len(where) > 0 {
		db = db.Where(where...)
	}

	user, userErr := db.
		Only(ui.ctx)

	if userErr != nil {
		logger.Errorf("Infra user List to %s", userErr)
		return nil, userErr
	}

	return user, nil
}
