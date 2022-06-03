package sample

import (
	"fmt"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/server/sample/ent"
)

func GroupResourceFromEnt(in *ent.Group) (*resource.Group, error) {
	var b resource.Builder

	meta, err := b.Meta().
		ResourceType("Group").
		Location("https://foobar.com/scim/v2/Group/" + in.ID.String()).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build meta information for Group")
	}
	return b.Group().
		DisplayName(in.DisplayName).
		ExternalID(in.ExternalID).
		ID(in.ID.String()).
		Meta(meta).
		Build()
}
