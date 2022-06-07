package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(`groups`, Group.Type),
		edge.To(`emails`, Email.Type),
		edge.To(`name`, Name.Type),
	}
}
