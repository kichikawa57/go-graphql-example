package user

import (
	"context"

	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/schema"
)

func (ui UserInfra) ListId(id schema.UserId) ([]*ent.User, error) {
	ctx := context.Background()

	return ui.db.User.Query().All(ctx)
}
