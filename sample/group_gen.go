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
		fields = []string{"displayName", "externalID", "id", "members"}
	}
	selectNames := make([]string, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "displayName":
			selectNames = append(selectNames, group.FieldDisplayName)
		case "externalID":
			selectNames = append(selectNames, group.FieldExternalID)
		case "id":
			selectNames = append(selectNames, group.FieldID)
		case "members":
		case "meta":
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

func groupStartsWithPredicate(scimField string, val string) predicate.Group {
	switch scimField {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val))
		})
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val))
		})
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val))
		})
	default:
		return nil
	}
}

func groupEndsWithPredicate(scimField string, val string) predicate.Group {
	switch scimField {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val))
		})
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val))
		})
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val))
		})
	default:
		return nil
	}
}

func groupContainsPredicate(scimField string, val string) predicate.Group {
	switch scimField {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val))
		})
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val))
		})
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val))
		})
	default:
		return nil
	}
}

func groupEqualsPredicate(scimField string, val string) predicate.Group {
	switch scimField {
	case resource.GroupDisplayNameKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val))
		})
	case resource.GroupExternalIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val))
		})
	case resource.GroupIDKey:
		entFieldName := GroupEntFieldFromSCIM(scimField)
		return predicate.Group(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val))
		})
	default:
		return nil
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
