package user

import (
	"context"

	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/schema"
)

func (ui UserInfra) ShowId(id schema.UserId) (*ent.User, error) {
	ctx := context.Background()

	return ui.db.User.Get(ctx, id)
}
