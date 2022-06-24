package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(`users`, User.Type).
			Ref(`emails`).
			Unique(),
	}
}
