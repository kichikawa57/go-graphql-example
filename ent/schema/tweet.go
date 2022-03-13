package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/kichikawa/ent/schema/property"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

func (Tweet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.String("text").
			SchemaType(map[string]string{
				dialect.Postgres: "text",
			}),
		field.Enum("type").
			GoType(property.TweetType("")),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("goods", Good.Type).
			StorageKey(edge.Column("tweet_id")),
		edge.To("comments", Comment.Type).
			StorageKey(edge.Column("tweet_id")),
	}
}

func (Tweet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("text", "type"),
		index.Fields("user_id"),
	}
}
