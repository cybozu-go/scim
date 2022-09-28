package resource

import (
	"fmt"
)

// From allows the user to add a member by specifying
// a *resource.User or *resource.Group object
//
// As of this writing, the object must have a proper Meta
// field populated
func (b *GroupMemberBuilder) FromResource(r interface{}) *GroupMemberBuilder {
	if b.err != nil {
		return b
	}

	switch r := r.(type) {
	case *Group:
		b.Value(r.ID()).
			Reference(r.Meta().Location())
	case *User:
		b.Value(r.ID()).
			Reference(r.Meta().Location())
	default:
		b.err = fmt.Errorf(`invalid object type passed to GroupMemberBuilder.From: %T`, r)
	}

	return b
}

// MembersFrom is a convenience method to directly add a SCIM
// resource to the Group object without having to construct the
// GroupMember object yourself.
//
// Currently this method only accepts `*resource.User` and
// `*resource.Group` as its input, and otherwise an error
// is stored in the builder, failing the Build() call.
//
// If you would like to otherwise construct the `members` field
// yourself, use the `Members()` method
func (b *GroupBuilder) MembersFrom(in ...interface{}) *GroupBuilder {
	if b.err != nil {
		return b
	}

	var builder Builder
LOOP:
	for _, r := range in {
		switch r := r.(type) {
		case *User:
			m, err := builder.GroupMember().FromResource(r).Build()
			if err != nil {
				b.err = err
				break LOOP
			}
			b.object.members = append(b.object.members, m)
		case *Group:
			m, err := builder.GroupMember().FromResource(r).Build()
			if err != nil {
				b.err = err
				break LOOP
			}
			b.object.members = append(b.object.members, m)
		default:
			b.err = fmt.Errorf(`invalid type passed to MembersFrom: %T`, r)
			break LOOP
		}
	}

	return b
}
