package sample

import (
	"reflect"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/names"
)

func namesLoadEntFields(q *ent.NamesQuery, fields []string) {
	if len(fields) == 0 {
		fields = []string{"familyName", "formatted", "givenName", "honorificPrefix", "honorificSuffix", "middleName"}
	}
	selectNames := make([]string, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "familyName":
			selectNames = append(selectNames, names.FieldFamilyName)
		case "formatted":
			selectNames = append(selectNames, names.FieldFormatted)
		case "givenName":
			selectNames = append(selectNames, names.FieldGivenName)
		case "honorificPrefix":
			selectNames = append(selectNames, names.FieldHonorificPrefix)
		case "honorificSuffix":
			selectNames = append(selectNames, names.FieldHonorificSuffix)
		case "middleName":
			selectNames = append(selectNames, names.FieldMiddleName)
		}
	}
	q.Select(selectNames...)
}

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
