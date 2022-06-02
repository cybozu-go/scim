package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Name holds the schema definition for the Name entity.
type Name struct {
	ent.Schema
}

// Fields of the Name.
func (Name) Fields() []ent.Field {
	return []ent.Field{
		field.String(`familyName`).Optional(),
		field.String(`formatted`).Optional(),
		field.String(`givenName`).Optional(),
		field.String(`honorificPrefix`).Optional(),
		field.String(`honorificSuffix`).Optional(),
		field.String(`middleName`).Optional(),
	}
}

// Edges of the Name.
func (Name) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(`users`, User.Type).
			Ref(`names`).
			Unique(),
	}
}
