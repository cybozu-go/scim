package resource

import (
	"fmt"
)

func (b *GroupBuilder) MemberFrom(r interface{}) *GroupBuilder {
	if b.err != nil {
		return b
	}

	var builder Builder
	switch r := r.(type) {
	case *User:
		m, err := builder.GroupMember().
			Value(r.ID()).
			Ref(r.Meta().Location()).
			Build()
		if err != nil {
			b.err = err
			return b
		}
		b.object.members = append(b.object.members, m)
	case *Group:
		m, err := builder.GroupMember().
			Value(r.ID()).
			Ref(r.Meta().Location()).
			Build()
		if err != nil {
			b.err = err
			return b
		}
		b.object.members = append(b.object.members, m)
	default:
		b.err = fmt.Errorf(`invalid type passed to MemberFrom: %T`, r)
	}

	return b
}
