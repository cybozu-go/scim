package sample

import (
	"reflect"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/email"
)

func emailLoadEntFields(q *ent.EmailQuery, fields []string) {
	selectNames := make([]string, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "display":
			selectNames = append(selectNames, email.FieldDisplay)
		case "primary":
			selectNames = append(selectNames, email.FieldPrimary)
		case "typ":
			selectNames = append(selectNames, email.FieldType)
		case "value":
			selectNames = append(selectNames, email.FieldValue)
		}
	}
	q.Select(selectNames...)
}

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
