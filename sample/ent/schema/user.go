package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(`groups`, Group.Type),
		edge.To(`emails`, Email.Type),
		edge.To(`name`, Names.Type),
		edge.To(`entitlements`, Entitlement.Type),
		edge.To(`roles`, Role.Type),
		edge.To(`imses`, IMS.Type),
		edge.To(`phone_numbers`, PhoneNumber.Type),
	}
}
