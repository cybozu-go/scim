package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/cybozu-go/scim/client"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var TraceWriter = io.Discard

func init() {
	v, err := strconv.ParseBool(os.Getenv(`SCIM_TRACE`))
	if err == nil {
		if v {
			TraceWriter = os.Stdout
		}
	}
}

func RunConformanceTests(t *testing.T, name string, backend interface{}) {
	t.Run(name, func(t *testing.T) {
		hh, err := server.NewServer(backend)
		require.NoError(t, err, `server.NewServer should succeed`)

		srv := httptest.NewServer(hh)

		cl := client.New(srv.URL, client.WithClient(srv.Client()))

		t.Run("search via /.search", func(t *testing.T) {
			createdUser := createStockUser(t, cl)
			//nolint:errcheck
			defer cl.User().DeleteUser(createdUser.ID()).
				Do(context.TODO())

			res, err := cl.Search().Search().
				Trace(TraceWriter).
				Attributes(`displayName`, `userName`).
				// Filter(`displayName sw "smith"`).
				Filter(`displayName pr`).
				StartIndex(1).
				Count(10).
				Do(context.TODO())
			require.NoError(t, err, `cl.Search should succeed`)
			require.Equal(t, 1, res.TotalResults(), `total results should be 1`)
		})
		t.Run("Users", func(t *testing.T) {
			t.Run("Basic CRUD", UsersBasicCRUD(t, cl))
			t.Run("Fetch", UsersFetch(t, cl))
		})
	})
}

func UsersFetch(t *testing.T, cl *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("Fetch unknown user ID", func(t *testing.T) {
			u, err := cl.User().GetUser("foobar").
				Trace(TraceWriter).
				Do(context.TODO())
			require.Nil(t, u, `GetUser return value should be nil`)
			require.Error(t, err, `GetUser should fail`)
		})
		t.Run("Fetch user", func(t *testing.T) {
			createdUser, err := stockUserCreateCall(cl).
				Trace(TraceWriter).
				Do(context.TODO())
			require.NoError(t, err, `CreateUser should succeed`)

			//nolint:errcheck
			defer cl.User().DeleteUser(createdUser.ID()).
				Do(context.TODO())

			u, err := cl.User().GetUser(createdUser.ID()).
				Trace(TraceWriter).
				Do(context.TODO())
			require.NoError(t, err, `GetUser should succeed`)

			// fetched user should have non-empty userName and emails, as
			// well as the attributes with the "returned" value of always
			require.Equal(t, createdUser.ID(), u.ID(), `ID should match`)
			require.Equal(t, createdUser.UserName(), u.UserName(), `UserName should match`)
			require.Equal(t, createdUser.Emails(), u.Emails(), `Emails should match`)
			if n := u.Name(); assert.NotNil(t, n, `Name should not be nil`) {
				require.Equal(t, `Ms. Barbara J Jensen III`, n.Formatted(), `Formatted should match`)
				require.Equal(t, `Jensen`, n.FamilyName(), `FamilyName should match`)
				require.Equal(t, `Barbara`, n.GivenName(), `GivenName should match`)
			}
		})
		t.Run("Fetch user with attributes", func(t *testing.T) {
			createdUser := createStockUser(t, cl)
			u, err := cl.User().GetUser(createdUser.ID()).
				Attributes("userName", "emails").
				Trace(TraceWriter).
				Do(context.TODO())
			require.NoError(t, err, `GetUser should succeed`)

			//nolint:errcheck
			defer cl.User().DeleteUser(createdUser.ID()).
				Do(context.TODO())

			// fetched user should have non-empty userName and emails, as
			// well as the attributes with the "returned" value of always
			require.Equal(t, createdUser.ID(), u.ID(), `ID should match`)
			require.Equal(t, createdUser.UserName(), u.UserName(), `UserName should match`)
			require.Equal(t, createdUser.Emails(), u.Emails(), `Emails should match`)
			require.Nil(t, u.Name(), `Name should be nil`)
		})
	}
}

func createStockUser(t *testing.T, cl *client.Client) *resource.User {
	var created *resource.User
	t.Run("Create stock user object", func(t *testing.T) {
		u, err := stockUserCreateCall(cl).
			Trace(TraceWriter).
			Do(context.TODO())
		require.NoError(t, err, `CreateUser should succeed`)

		// Make sure that it's really there
		fetched, err := cl.User().GetUser(u.ID()).
			Trace(TraceWriter).
			Do(context.TODO())
		require.NoError(t, err, `GetUser should succeed`)

		_ = json.NewEncoder(TraceWriter).Encode(fetched)
		created = u
	})

	if created == nil {
		t.Fatal()
	}
	return created
}
func stockUserCreateCall(cl *client.Client) *client.CreateUserCall {
	return cl.User().CreateUser().
		UserName("bjensen").
		ExternalID("bjensen").
		DisplayName("Barbara Jensen").
		Emails(resource.NewEmailBuilder().
			Value("babs@jensen.org").
			Primary(true).
			MustBuild()).
		Name(resource.NewNamesBuilder().
			Formatted("Ms. Barbara J Jensen III").
			FamilyName("Jensen").
			GivenName("Barbara").
			MustBuild())
}

func UsersBasicCRUD(t *testing.T, cl *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("Basic CRUD", func(t *testing.T) {
			// Create a single User, apply Fetch, Replace, Delete
			createdUser, err := stockUserCreateCall(cl).
				Do(context.TODO())
			require.NoError(t, err, `CreateUser should succeed`)

			t.Run("Fetch user", func(t *testing.T) {
				_, err = cl.User().GetUser(createdUser.ID()).
					Do(context.TODO())
				require.NoError(t, err, `GetUser should succeed`)
			})
			t.Run("Replace user", func(t *testing.T) {
				u, err := cl.User().ReplaceUser(createdUser.ID()).
					Emails(resource.NewEmailBuilder().
						Value("babs-new@jensen.org").
						Primary(true).
						MustBuild()).
					Do(context.TODO())
				require.NoError(t, err, `ReplaceUser should succeed`)

				require.Equal(t, "bjensen", u.ExternalID(), `externalID should match`)

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
			t.Run("Delete user", func(t *testing.T) {
				err := cl.User().DeleteUser(createdUser.ID()).
					Do(context.TODO())
				require.NoError(t, err, `DeleteUser should succeed`)
			})
			t.Run("Fetch users (after delete)", func(t *testing.T) {
				_, err := cl.User().GetUser(createdUser.ID()).
					Do(context.TODO())
				require.Error(t, err, `GetUser should fail`)
			})
		})
	}
}
