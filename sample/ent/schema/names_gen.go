package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Names struct {
	ent.Schema
}

func (Names) Fields() []ent.Field {
	return []ent.Field{
		field.String("familyName").Optional(),
		field.String("formatted").Optional(),
		field.String("givenName").Optional(),
		field.String("honorificPrefix").Optional(),
		field.String("honorificSuffix").Optional(),
		field.String("middleName").Optional(),
	}
}
