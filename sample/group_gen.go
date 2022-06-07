package sample

import (
	"fmt"
	"reflect"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/group"
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

func GroupResourceFromEnt(in *ent.Group) (*resource.Group, error) {
	var b resource.Builder

	builder := b.Group()

	meta, err := b.Meta().
		ResourceType("Group").
		Location("https://foobar.com/scim/v2/Group/" + in.ID.String()).
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
	return builder.Build()
}
