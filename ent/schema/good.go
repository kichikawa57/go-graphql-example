package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Good holds the schema definition for the Good entity.
type Good struct {
	ent.Schema
}

func (Good) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Good.
func (Good) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("tweet_id", uuid.UUID{}),
	}
}

// Edges of the Good.
func (Good) Edges() []ent.Edge {
	return nil
}

func (Good) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "tweet_id").Unique(),
	}
}
