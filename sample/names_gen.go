package sample

import (
	"reflect"

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

func NamesEntFileFromSCIM(s string) string {
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
