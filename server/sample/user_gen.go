package sample

import (
	"fmt"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/server/sample/ent"
)

func UserResourceFromEnt(in *ent.User) (*resource.User, error) {
	var b resource.Builder

	meta, err := b.Meta().
		ResourceType("User").
		Location("https://foobar.com/scim/v2/User/" + in.ID.String()).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build meta information for User")
	}
	return b.User().
		Active(in.Active).
		DisplayName(in.DisplayName).
		ExternalID(in.ExternalID).
		ID(in.ID.String()).
		Locale(in.Locale).
		Password(in.Password).
		PreferredLanguage(in.PreferredLanguage).
		Timezone(in.Timezone).
		UserName(in.UserName).
		UserType(in.UserType).
		Meta(meta).
		Build()
}
