package sample

import (
	"fmt"
	"reflect"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/predicate"
	"github.com/cybozu-go/scim/sample/ent/role"
)

func RoleResourceFromEnt(in *ent.Role) (*resource.Role, error) {
	var b resource.Builder

	builder := b.Role()
	if !reflect.ValueOf(in.Display).IsZero() {
		builder.Display(in.Display)
	}
	if !reflect.ValueOf(in.Primary).IsZero() {
		builder.Primary(in.Primary)
	}
	if !reflect.ValueOf(in.Type).IsZero() {
		builder.Type(in.Type)
	}
	if !reflect.ValueOf(in.Value).IsZero() {
		builder.Value(in.Value)
	}
	return builder.Build()
}

func RoleEntFieldFromSCIM(s string) string {
	switch s {
	case resource.RoleDisplayKey:
		return role.FieldDisplay
	case resource.RolePrimaryKey:
		return role.FieldPrimary
	case resource.RoleTypeKey:
		return role.FieldType
	case resource.RoleValueKey:
		return role.FieldValue
	default:
		return s
	}
}

func roleStartsWithPredicate(q *ent.RoleQuery, scimField string, val interface{}) (predicate.Role, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.RoleDisplayKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleTypeKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleValueKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func roleEndsWithPredicate(q *ent.RoleQuery, scimField string, val interface{}) (predicate.Role, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.RoleDisplayKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleTypeKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleValueKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func roleContainsPredicate(q *ent.RoleQuery, scimField string, val interface{}) (predicate.Role, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.RoleDisplayKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleTypeKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleValueKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func roleEqualsPredicate(q *ent.RoleQuery, scimField string, val interface{}) (predicate.Role, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.RoleDisplayKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleTypeKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.RoleValueKey:
		entFieldName := RoleEntFieldFromSCIM(scimField)
		return predicate.Role(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func rolePresencePredicate(scimField string) predicate.Role {
	switch scimField {
	case resource.RoleDisplayKey:
		return role.And(role.DisplayNotNil(), role.DisplayNEQ(""))
	case resource.RoleTypeKey:
		return role.And(role.TypeNotNil(), role.TypeNEQ(""))
	default:
		return nil
	}
}
