package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type UserId int

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
		field.Int("id").GoType(UserId(0)),
		field.String("account_name").Unique(),
		field.String("email").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}).
			Unique(),
		field.Int("age").
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}).
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type).
			StorageKey(edge.Column("user_id")),
		edge.To("groups", Group.Type),
	}
}
