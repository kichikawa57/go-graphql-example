package user

import (
	"github.com/kichikawa/ent"
	"github.com/kichikawa/repository"
)

type UserInfra struct {
	db *ent.Client
}

func NewUserInfra(db *ent.Client) repository.UserRepository {
	return &UserInfra{db}
}
