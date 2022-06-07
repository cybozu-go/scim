package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Email holds the schema definition for the Email entity.
type Email struct {
	ent.Schema
}

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(`users`, User.Type).
			Ref(`emails`).
			Unique(),
	}
}
