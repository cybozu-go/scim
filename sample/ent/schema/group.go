package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("children", Group.Type).
			From("parent").
			Unique(),
	}
}
