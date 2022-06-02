package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	// TODO: would be nice to generate this automatically
	return []ent.Field{
		field.Bool("active").Default(false),
		field.Text("externalID").Optional(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("password").NotEmpty(),
		field.Text("preferredLanguage").Optional(),
		field.Text("locale").Optional(),
		field.Text("timezone").Optional(),
		field.Text("userType").Optional(),
		field.Text("userName").Unique().NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(`groups`, Group.Type),
		edge.To(`emails`, Email.Type),
		edge.To(`names`, Name.Type),
	}
}
