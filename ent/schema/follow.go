package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Follow holds the schema definition for the Follow entity.
type Follow struct {
	ent.Schema
}

func (Follow) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Follow.
func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("follower_id", uuid.UUID{}),
		field.UUID("followed_id", uuid.UUID{}),
	}
}

// Edges of the Follow.
func (Follow) Edges() []ent.Edge {
	return nil
}

func (Follow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("follower_id", "followed_id").Unique(),
	}
}
