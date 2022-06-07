package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Names holds the schema definition for the Names entity.
type Names struct {
	ent.Schema
}

// Edges of the Names.
func (Names) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(`user`, User.Type).
			Ref(`name`).
			Unique(),
	}
}
