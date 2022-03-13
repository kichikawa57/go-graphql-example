package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("tweet_id", uuid.UUID{}),
		field.String("text").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("text"),
		index.Fields("user_id"),
		index.Fields("tweet_id"),
	}
}
