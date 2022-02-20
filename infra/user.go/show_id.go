package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/kichikawa/ent"
)

func (ui UserInfra) ShowId(id uuid.UUID) (*ent.User, error) {
	ctx := context.Background()

	return ui.db.User.Get(ctx, id)
}
