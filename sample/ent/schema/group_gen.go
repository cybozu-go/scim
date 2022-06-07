package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("displayName").Optional(),
		field.String("externalID").Optional(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Optional(),
	}
}
