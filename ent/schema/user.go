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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("account_name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(15)",
			}).
			GoType(property.UserAccountName("")).
			Unique(),
		field.String("email").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}).
			GoType(property.UserEmail("")).
			Unique(),
		field.String("password").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}),
		field.Int("age").
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tweet", Tweet.Type).
			StorageKey(edge.Column("user_id")).Required(),
		edge.To("good", Good.Type).
			StorageKey(edge.Column("user_id")).Required(),
		edge.To("comment", Comment.Type).
			StorageKey(edge.Column("user_id")).Required(),
		edge.To("follower", Follow.Type).
			StorageKey(edge.Column("follower_id")).Required(),
		edge.To("followed", Follow.Type).
			StorageKey(edge.Column("followed_id")).Required(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "account_name"),
	}
}
