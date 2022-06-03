package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Email holds the schema definition for the Email entity.
type Email struct {
	ent.Schema
}

// Fields of the Email.
func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String(`display`).Optional(),
		field.Bool(`primary`).Default(false),
		field.String(`type`).Optional(),
		field.String(`value`),
	}
}

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(`users`, User.Type).
			Ref(`emails`).
			Unique(),
	}
}
