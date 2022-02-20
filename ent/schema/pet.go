package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return nil
}
