package sample

import (
	"fmt"
	"reflect"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/email"
	"github.com/cybozu-go/scim/sample/ent/predicate"
)

func EmailResourceFromEnt(in *ent.Email) (*resource.Email, error) {
	var b resource.Builder

	builder := b.Email()
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

func EmailEntFieldFromSCIM(s string) string {
	switch s {
	case resource.EmailDisplayKey:
		return email.FieldDisplay
	case resource.EmailPrimaryKey:
		return email.FieldPrimary
	case resource.EmailTypeKey:
		return email.FieldType
	case resource.EmailValueKey:
		return email.FieldValue
	default:
		return s
	}
}

func emailStartsWithPredicate(q *ent.EmailQuery, scimField string, val interface{}) (predicate.Email, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.EmailDisplayKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailTypeKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailValueKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func emailEndsWithPredicate(q *ent.EmailQuery, scimField string, val interface{}) (predicate.Email, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.EmailDisplayKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailTypeKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailValueKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func emailContainsPredicate(q *ent.EmailQuery, scimField string, val interface{}) (predicate.Email, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.EmailDisplayKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailTypeKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailValueKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func emailEqualsPredicate(q *ent.EmailQuery, scimField string, val interface{}) (predicate.Email, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.EmailDisplayKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailTypeKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.EmailValueKey:
		entFieldName := EmailEntFieldFromSCIM(scimField)
		return predicate.Email(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func emailPresencePredicate(scimField string) predicate.Email {
	switch scimField {
	case resource.EmailDisplayKey:
		return email.And(email.DisplayNotNil(), email.DisplayNEQ(""))
	case resource.EmailTypeKey:
		return email.And(email.TypeNotNil(), email.TypeNEQ(""))
	default:
		return nil
	}
}
