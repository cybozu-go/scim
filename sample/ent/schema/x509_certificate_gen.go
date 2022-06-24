package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type X509Certificate struct {
	ent.Schema
}

func (X509Certificate) Fields() []ent.Field {
	return []ent.Field{
		field.String("value"),
	}
}

func (X509Certificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}
