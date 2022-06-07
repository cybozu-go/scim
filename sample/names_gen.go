package sample

import (
	"reflect"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
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
