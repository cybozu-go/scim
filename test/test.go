package test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/cybozu-go/scim/client"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/server"
	"github.com/stretchr/testify/require"
)

func RunConformanceTests(t *testing.T, name string, backend interface{}) {
	t.Run(name, func(t *testing.T) {
		hh, err := server.NewServer(backend)
		require.NoError(t, err, `server.NewServer should succeed`)

		srv := httptest.NewServer(hh)

		cl := client.New(srv.URL, client.WithClient(srv.Client()))

		t.Run("search via /.search", func(t *testing.T) {
			lr, err := cl.Search().Search().
				Attributes(`displayName`, `userName`).
				Filter(`displayName sw "smith"`).
				StartIndex(1).
				Count(10).
				Do(context.TODO())
			require.NoError(t, err, `cl.Search should succeed`)
			t.Logf("%#v", lr)
		})
		t.Run("/Users", func(t *testing.T) {
			t.Run("POST /Users", func(t *testing.T) {
				createdUser, err := cl.User().CreateUser().
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
					Do(context.TODO())
				require.NoError(t, err, `CreateUser should succeed`)
				t.Run(fmt.Sprintf("GET /Users/%s", createdUser.ID()), func(t *testing.T) {
					_, err := cl.User().GetUser(createdUser.ID()).
						Do(context.TODO())
					require.NoError(t, err, `GetUser should succeed`)
				})
				t.Run(fmt.Sprintf("DELETE /Users/%s", createdUser.ID()), func(t *testing.T) {
					err := cl.User().DeleteUser(createdUser.ID()).
						Do(context.TODO())
					require.NoError(t, err, `DeleteUser should succeed`)
				})
				t.Run(fmt.Sprintf("GET /Users/%s (after delete)", createdUser.ID()), func(t *testing.T) {
					_, err := cl.User().GetUser(createdUser.ID()).
						Do(context.TODO())
					require.Error(t, err, `GetUser should fail`)
				})
			})
			t.Run("GET /Users/foobar", func(t *testing.T) {
				_, err := cl.User().GetUser("foobar").
					Do(context.TODO())
				require.Error(t, err, `GetUser should fail`)
			})
		})
	})
}
