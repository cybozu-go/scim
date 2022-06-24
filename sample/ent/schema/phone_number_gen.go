package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PhoneNumber struct {
	ent.Schema
}

func (PhoneNumber) Fields() []ent.Field {
	return []ent.Field{
		field.String("value"),
		field.String("display"),
		field.String("type"),
		field.Bool("primary"),
	}
}

func (PhoneNumber) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}
