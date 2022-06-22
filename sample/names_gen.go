package sample

import (
	"fmt"
	"reflect"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/names"
	"github.com/cybozu-go/scim/sample/ent/predicate"
)

func NamesResourceFromEnt(in *ent.Names) (*resource.Names, error) {
	var b resource.Builder

	builder := b.Names()
	if !reflect.ValueOf(in.FamilyName).IsZero() {
		builder.FamilyName(in.FamilyName)
	}
	if !reflect.ValueOf(in.Formatted).IsZero() {
		builder.Formatted(in.Formatted)
	}
	if !reflect.ValueOf(in.GivenName).IsZero() {
		builder.GivenName(in.GivenName)
	}
	if !reflect.ValueOf(in.HonorificPrefix).IsZero() {
		builder.HonorificPrefix(in.HonorificPrefix)
	}
	if !reflect.ValueOf(in.HonorificSuffix).IsZero() {
		builder.HonorificSuffix(in.HonorificSuffix)
	}
	if !reflect.ValueOf(in.MiddleName).IsZero() {
		builder.MiddleName(in.MiddleName)
	}
	return builder.Build()
}

func NamesEntFieldFromSCIM(s string) string {
	switch s {
	case resource.NamesFamilyNameKey:
		return names.FieldFamilyName
	case resource.NamesFormattedKey:
		return names.FieldFormatted
	case resource.NamesGivenNameKey:
		return names.FieldGivenName
	case resource.NamesHonorificPrefixKey:
		return names.FieldHonorificPrefix
	case resource.NamesHonorificSuffixKey:
		return names.FieldHonorificSuffix
	case resource.NamesMiddleNameKey:
		return names.FieldMiddleName
	default:
		return s
	}
}

func namesStartsWithPredicate(q *ent.NamesQuery, scimField string, val interface{}) (predicate.Names, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.NamesFamilyNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesFormattedKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesGivenNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificPrefixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificSuffixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesMiddleNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func namesEndsWithPredicate(q *ent.NamesQuery, scimField string, val interface{}) (predicate.Names, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.NamesFamilyNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesFormattedKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesGivenNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificPrefixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificSuffixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesMiddleNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func namesContainsPredicate(q *ent.NamesQuery, scimField string, val interface{}) (predicate.Names, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.NamesFamilyNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesFormattedKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesGivenNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificPrefixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificSuffixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesMiddleNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func namesEqualsPredicate(q *ent.NamesQuery, scimField string, val interface{}) (predicate.Names, error) {
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.NamesFamilyNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesFormattedKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesGivenNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificPrefixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesHonorificSuffixKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.NamesMiddleNameKey:
		entFieldName := NamesEntFieldFromSCIM(scimField)
		return predicate.Names(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func namesPresencePredicate(scimField string) predicate.Names {
	switch scimField {
	case resource.NamesFamilyNameKey:
		return names.And(names.FamilyNameNotNil(), names.FamilyNameNEQ(""))
	case resource.NamesFormattedKey:
		return names.And(names.FormattedNotNil(), names.FormattedNEQ(""))
	case resource.NamesGivenNameKey:
		return names.And(names.GivenNameNotNil(), names.GivenNameNEQ(""))
	case resource.NamesHonorificPrefixKey:
		return names.And(names.HonorificPrefixNotNil(), names.HonorificPrefixNEQ(""))
	case resource.NamesHonorificSuffixKey:
		return names.And(names.HonorificSuffixNotNil(), names.HonorificSuffixNEQ(""))
	case resource.NamesMiddleNameKey:
		return names.And(names.MiddleNameNotNil(), names.MiddleNameNEQ(""))
	default:
		return nil
	}
}
