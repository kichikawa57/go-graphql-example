package repository

import (
	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/schema"
)

type UserRepository interface {
	ListId(id schema.UserId) ([]*ent.User, error)
	ShowId(id schema.UserId) (*ent.User, error)
	Update(user ent.User) (*ent.User, error)
	Create(user ent.User) (*ent.User, error)
}
