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
		t.Run("Users", func(t *testing.T) {
			t.Run("Basic CRUD", func(t *testing.T) {
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
				t.Run(fmt.Sprintf("Fetch user %q", createdUser.ID()), func(t *testing.T) {
					_, err := cl.User().GetUser(createdUser.ID()).
						Do(context.TODO())
					require.NoError(t, err, `GetUser should succeed`)
				})
				t.Run(fmt.Sprintf("Replace user %q", createdUser.ID()), func(t *testing.T) {
					u, err := cl.User().ReplaceUser(createdUser.ID()).
						Emails(resource.NewEmailBuilder().
							Value("babs-new@jensen.org").
							Primary(true).
							MustBuild()).
						Do(context.TODO())
					require.NoError(t, err, `ReplaceUser should succeed`)

					require.Equal(t, "bjensen", u.ExternalID())

					// Sanity
					require.Equal(t, createdUser.ID(), u.ID())

					emails := u.Emails()
					require.Len(t, emails, 1)
					for _, email := range emails {
						// hardcoded for loop testing is weird, I know.
						// am just expecting it be expanded in the future
						require.Equal(t, `babs-new@jensen.org`, email.Value())
						require.True(t, email.Primary())
					}
				})
				t.Run(fmt.Sprintf("Delete user %q", createdUser.ID()), func(t *testing.T) {
					err := cl.User().DeleteUser(createdUser.ID()).
						Do(context.TODO())
					require.NoError(t, err, `DeleteUser should succeed`)
				})
				t.Run(fmt.Sprintf("Fetch users %q (after delete)", createdUser.ID()), func(t *testing.T) {
					_, err := cl.User().GetUser(createdUser.ID()).
						Do(context.TODO())
					require.Error(t, err, `GetUser should fail`)
				})
			})
			t.Run("Fetch unknown user ID", func(t *testing.T) {
				_, err := cl.User().GetUser("foobar").
					Do(context.TODO())
				require.Error(t, err, `GetUser should fail`)
			})
		})
	})
}
