package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active").Optional(),
		field.String("displayName").Optional(),
		field.String("externalID").Optional(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("locale").Optional(),
		field.String("nickName").Optional(),
		field.String("password").Optional().NotEmpty().Sensitive(),
		field.String("preferredLanguage").Optional(),
		field.String("profileURL").Optional(),
		field.String("timezone").Optional(),
		field.String("title").Optional(),
		field.String("userName").Unique().NotEmpty(),
		field.String("userType").Optional(),
	}
}
