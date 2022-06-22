package sample

import (
	"fmt"
	"reflect"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/group"
	"github.com/cybozu-go/scim/sample/ent/predicate"
)

func groupLoadEntFields(q *ent.GroupQuery, fields []string) {
	if len(fields) == 0 {
		fields = []string{resource.GroupDisplayNameKey, resource.GroupExternalIDKey, resource.GroupIDKey, resource.GroupMembersKey}
	}
	selectNames := make([]string, 0, len(fields))
	for _, f := range fields {
		switch f {
		case resource.GroupDisplayNameKey:
			selectNames = append(selectNames, group.FieldDisplayName)
		case resource.GroupExternalIDKey:
			selectNames = append(selectNames, group.FieldExternalID)
		case resource.GroupIDKey:
			selectNames = append(selectNames, group.FieldID)
		case resource.GroupMembersKey:
		case resource.GroupMetaKey:
		}
	}
	q.Select(selectNames...)
}

func groupLocation(id string) string {
	return "https://foobar.com/scim/v2/Groups/" + id
}

func GroupResourceFromEnt(in *ent.Group) (*resource.Group, error) {
	var b resource.Builder

	builder := b.Group()

	meta, err := b.Meta().
		ResourceType("Group").
		Location(groupLocation(in.ID.String())).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build meta information for Group")
	}

	builder.
		Meta(meta)
	if !reflect.ValueOf(in.DisplayName).IsZero() {
		builder.DisplayName(in.DisplayName)
	}
	if !reflect.ValueOf(in.ExternalID).IsZero() {
		builder.ExternalID(in.ExternalID)
	}
	builder.ID(in.ID.String())
	if err := groupResourceFromEntHelper(in, builder); err != nil {
		return nil, err
	}
	return builder.Build()
}

func GroupEntFieldFromSCIM(s string) string {
	switch s {
	case resource.GroupDisplayNameKey:
		return group.FieldDisplayName
	case resource.GroupExternalIDKey:
		return group.FieldExternalID
	case resource.GroupIDKey:
		return group.FieldID
	default:
		return s
	}
}

func groupStartsWithPredicate(q *ent.GroupQuery, scimField string, val interface{}) (predicate.Group, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func groupEndsWithPredicate(q *ent.GroupQuery, scimField string, val interface{}) (predicate.Group, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func groupContainsPredicate(q *ent.GroupQuery, scimField string, val interface{}) (predicate.Group, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func groupEqualsPredicate(q *ent.GroupQuery, scimField string, val interface{}) (predicate.Group, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func groupPresencePredicate(scimField string) predicate.Group {
	switch scimField {
	case resource.GroupDisplayNameKey:
		return group.And(group.DisplayNameNotNil(), group.DisplayNameNEQ(""))
	case resource.GroupExternalIDKey:
		return group.And(group.ExternalIDNotNil(), group.ExternalIDNEQ(""))
	default:
		return nil
	}
}
