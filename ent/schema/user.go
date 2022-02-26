package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		field.String("account_name").Unique(),
		field.String("email").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}).
			Unique(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Annotations(
				entgql.OrderField("STATUS"),
			),
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
