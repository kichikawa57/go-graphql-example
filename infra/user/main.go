package user

import (
	"context"

	"github.com/kichikawa/ent"
	"github.com/kichikawa/repository"
)

type UserInfra struct {
	ctx context.Context
	db  *ent.Client
}

func NewUserInfra(
	ctx context.Context,
	db *ent.Client,
) repository.UserRepository {
	return &UserInfra{ctx, db}
}
