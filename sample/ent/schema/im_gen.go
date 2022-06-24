package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type IMS struct {
	ent.Schema
}

func (IMS) Fields() []ent.Field {
	return []ent.Field{
		field.String("value"),
	}
}

func (IMS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}
