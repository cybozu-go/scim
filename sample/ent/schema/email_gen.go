package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Email struct {
	ent.Schema
}

func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String("display").Optional(),
		field.Bool("primary").Optional(),
		field.String("type").Optional(),
		field.String("value"),
	}
}
