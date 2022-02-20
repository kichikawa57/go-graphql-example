package repository

import (
	"github.com/google/uuid"
	"github.com/kichikawa/ent"
)

type UserRepository interface {
	ListId(id uuid.UUID) ([]*ent.User, error)
	ShowId(id uuid.UUID) (*ent.User, error)
	Update(user ent.User) (*ent.User, error)
	Create(user ent.User) (*ent.User, error)
}
