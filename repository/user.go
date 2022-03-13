package repository

import (
	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/predicate"
)

type UserRepository interface {
	List(where ...predicate.User) ([]*ent.User, error)
	Show(where ...predicate.User) (*ent.User, error)
	Update(user ent.User, where ...predicate.User) error
	Create(user ent.User) (*ent.User, error)
}
