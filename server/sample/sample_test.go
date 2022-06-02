package sample_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/server/sample"
	"github.com/stretchr/testify/require"
)

func TestSample(t *testing.T) {
	s, err := sample.New("file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err, `sample.New should succeed`)

	var b resource.Builder

	u, err := s.CreateUser(
		b.User().
			UserName("bjensen").
			ExternalID("bjensen").
			Emails(resource.NewEmailBuilder().
				Value("babs@jensen.org").
				Primary(true).
				MustBuild()).
			Name(resource.NewNamesBuilder().
				Formatted("Ms. Barbara J Jensen III").
				FamilyName("Jensen").
				GivenName("Barbara").
				MustBuild()).
			Password("foobar").
			MustBuild(),
	)
	require.NoError(t, err, `s.CreateUser should succeed`)

	g, err := s.CreateGroup(
		b.Group().
			DisplayName(`Engineering`).
			Members(
				b.GroupMember().
					Value(u.ID()).
					Ref(u.Meta().Location()).
					MustBuild(),
			).
			MustBuild(),
	)
	require.NoError(t, err, `s.CreateGroup should succeed`)

	json.NewEncoder(os.Stdout).Encode(u)
	json.NewEncoder(os.Stdout).Encode(g)

}
