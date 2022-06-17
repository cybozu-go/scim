package test

import (
	"context"
	"embed"
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

//go:embed data/*
var datasrc embed.FS

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

		t.Run("Prepare Fixtures", PrepareFixtures(t, cl))
		t.Run("Users", func(t *testing.T) {
			t.Run("Basic CRUD", UsersBasicCRUD(t, cl))
			t.Run("Fetch", UsersFetch(t, cl))
			t.Run("Search", UsersSearch(t, cl))
		})
		t.Run("Groups", func(t *testing.T) {
			t.Run("Basic CRUD", GroupsBasicCRUD(t, cl))
			t.Run("Search", GroupsSearch(t, cl))
		})
		t.Run("Meta", func(t *testing.T) {
			t.Run("ServiceProviderConfig", ServiceProviderConfig(t, cl))
			t.Run("ResourceTypes", ResourceTypes(t, cl))
			t.Run("Schemas", Schemas(t, cl))
		})
	})
}

func PrepareFixtures(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("Load User fixtures", func(t *testing.T) {
			f, err := datasrc.Open(`data/users.json`)
			require.NoError(t, err, `opening file should succeed`)

			var data []json.RawMessage
			require.NoError(t, json.NewDecoder(f).Decode(&data), `unmarshaling form file should succeed`)

			for _, e := range data {
				_, err := cl.User().Create().
					FromJSON(e).
					Trace(TraceWriter).
					Do(context.TODO())
				if !assert.NoError(t, err, `user creation should succeed`) {
					t.Logf("invalid data: %s", e)
					return
				}
			}
		})
		t.Run("Load Group fixtures", func(t *testing.T) {
			// Unfortunately we can't just automatically create Groups because it
			// requires that we know the members' internal IDs
			list, err := cl.User().Search().
				Trace(TraceWriter).
				Filter(`email.value ew "zemeckis-crew.com"`).
				Do(context.TODO())
			require.NoError(t, err, `user search should succeed`)

			createGroupCall := cl.Group().Create().Trace(TraceWriter).
				DisplayName(`zemeckis-crew`)
			for _, r := range list.Resources() {
				createGroupCall.MemberFrom(r)
			}
			createGroupCall.Do(context.TODO())
		})
	}
}

func stockUserCreateCall(cl *client.Client) *client.CreateUserCall {
	return cl.User().Create().
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

func stockGroupCreateCall(cl *client.Client) *client.CreateGroupCall {
	u1, err := cl.User().Create().
		Trace(TraceWriter).
		UserName("jsmith").
		ExternalID("jsmith").
		DisplayName("John Smith").
		Title("Manager").
		Do(context.TODO())
	if err != nil {
		panic(err)
	}
	u2, err := cl.User().Create().
		Trace(TraceWriter).
		UserName("acooper").
		ExternalID("acooper").
		DisplayName("Alice Cooper").
		Title("Manager").
		Do(context.TODO())
	if err != nil {
		panic(err)
	}
	u3, err := cl.User().Create().
		Trace(TraceWriter).
		UserName("wjohnson").
		ExternalID("wjohnson").
		DisplayName("William Johnson").
		Title("Lead Engineer").
		Do(context.TODO())
	if err != nil {
		panic(err)
	}
	members := []*resource.GroupMember{
		resource.NewGroupMemberBuilder().
			Value(u3.ID()).
			Ref(u3.Meta().Location()).
			MustBuild(),
	}
	g1, err := cl.Group().Create().
		Trace(TraceWriter).
		DisplayName("Product").
		Members(members...).
		Do(context.TODO())
	if err != nil {
		panic(err)
	}

	members = []*resource.GroupMember{
		resource.NewGroupMemberBuilder().
			Value(u1.ID()).
			Ref(u1.Meta().Location()).
			MustBuild(),
		resource.NewGroupMemberBuilder().
			Value(u2.ID()).
			Ref(u2.Meta().Location()).
			MustBuild(),
		resource.NewGroupMemberBuilder().
			Value(g1.ID()).
			Ref(g1.Meta().Location()).
			MustBuild(),
	}

	return cl.Group().Create().
		Trace(TraceWriter).
		DisplayName("Engineering").
		Members(members...)
}

func UsersSearch(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("search via /Users/.search", func(t *testing.T) {
			// TODO: need to create a group that matches the query
			// and make sure that the result doesn't contain any groups
			createdUser := createStockUser(t, cl)
			//nolint:errcheck
			defer cl.User().Delete(createdUser.ID()).
				Do(context.TODO())

			res, err := cl.User().Search().
				Trace(TraceWriter).
				Attributes(`displayName`, `userName`).
				Filter(`displayName sw "Barbara"`).
				StartIndex(1).
				Count(10).
				Do(context.TODO())
			require.NoError(t, err, `cl.Search should succeed`)
			require.Equal(t, 1, res.TotalResults(), `total results should be 1`)
		})
	}
}

func UsersFetch(t *testing.T, cl *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("Fetch unknown user ID", func(t *testing.T) {
			u, err := cl.User().Get("foobar").
				Trace(TraceWriter).
				Do(context.TODO())
			require.Nil(t, u, `Get return value should be nil`)
			require.Error(t, err, `Get should fail`)
		})
		t.Run("Fetch user", func(t *testing.T) {
			createdUser, err := stockUserCreateCall(cl).
				Trace(TraceWriter).
				Do(context.TODO())
			require.NoError(t, err, `Create should succeed`)

			//nolint:errcheck
			defer cl.User().Delete(createdUser.ID()).
				Do(context.TODO())

			u, err := cl.User().Get(createdUser.ID()).
				Trace(TraceWriter).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)

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
			u, err := cl.User().Get(createdUser.ID()).
				Attributes("userName", "emails").
				Trace(TraceWriter).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)

			//nolint:errcheck
			defer cl.User().Delete(createdUser.ID()).
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
		require.NoError(t, err, `Create should succeed`)

		// Make sure that it's really there
		fetched, err := cl.User().Get(u.ID()).
			Trace(TraceWriter).
			Do(context.TODO())
		require.NoError(t, err, `Get should succeed`)

		_ = json.NewEncoder(TraceWriter).Encode(fetched)
		created = u
	})

	if created == nil {
		t.Fatal()
	}
	return created
}

func UsersBasicCRUD(t *testing.T, cl *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		// Create a single User, apply Fetch, Replace, Delete
		createdUser, err := stockUserCreateCall(cl).
			Do(context.TODO())
		require.NoError(t, err, `Create should succeed`)
		require.Empty(t, createdUser.Password(), `user should not return password`)

		t.Run("Fetch user", func(t *testing.T) {
			_, err = cl.User().Get(createdUser.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)
		})
		t.Run("Replace user", func(t *testing.T) {
			u, err := cl.User().Replace(createdUser.ID()).
				Emails(resource.NewEmailBuilder().
					Value("babs-new@jensen.org").
					Primary(true).
					MustBuild()).
				Do(context.TODO())
			require.NoError(t, err, `ReplaceUser should succeed`)

			// we need to validate the result from PUT and GET
			fetched, err := cl.User().Get(createdUser.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)

			testcases := []struct {
				Name string
				User *resource.User
			}{
				{
					Name: "Result from issuing replace",
					User: u,
				},
				{
					Name: "Result from fetch after replace",
					User: fetched,
				},
			}

			for _, tc := range testcases {
				tc := tc
				t.Run(tc.Name, func(t *testing.T) {
					u := tc.User
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
			}
		})
		t.Run("Delete user", func(t *testing.T) {
			err := cl.User().Delete(createdUser.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Delete should succeed`)
			t.Run("Fetch users (after delete)", func(t *testing.T) {
				_, err := cl.User().Get(createdUser.ID()).
					Do(context.TODO())
				require.Error(t, err, `Get should fail`)
			})
		})
	}
}

func GroupsBasicCRUD(t *testing.T, cl *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		// Create a single Group, apply Fetch, Replace, Delete
		createdGroup, err := stockGroupCreateCall(cl).
			Do(context.TODO())
		require.NoError(t, err, `Create should succeed`)
		_ = createdGroup

		t.Run("Fetch group", func(t *testing.T) {
			group, err := cl.Group().Get(createdGroup.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)
			require.Len(t, group.Members(), len(createdGroup.Members()), `there should be %d members`, len(createdGroup.Members()))
		})
		t.Run("Replace group", func(t *testing.T) {
			u, err := cl.Group().Replace(createdGroup.ID()).
				DisplayName(createdGroup.DisplayName()).
				Members((createdGroup.Members())[1:]...).
				Do(context.TODO())
			require.NoError(t, err, `Replace should succeed`)

			// we need to validate the result from PUT and GET
			fetched, err := cl.Group().Get(createdGroup.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)

			testcases := []struct {
				Name  string
				Group *resource.Group
			}{
				{
					Name:  "Result from issuing replace",
					Group: u,
				},
				{
					Name:  "Result from fetch after replace",
					Group: fetched,
				},
			}

			for _, tc := range testcases {
				tc := tc
				t.Run(tc.Name, func(t *testing.T) {
					g := tc.Group

					// Sanity
					require.Equal(t, createdGroup.ID(), g.ID())
					require.Len(t, g.Members(), len(createdGroup.Members())-1)
				})
			}
		})
		t.Run("Delete group", func(t *testing.T) {
			err := cl.Group().Delete(createdGroup.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Delete should succeed`)
			t.Run("Fetch group (after delete)", func(t *testing.T) {
				_, err := cl.Group().Get(createdGroup.ID()).
					Do(context.TODO())
				require.Error(t, err, `Get should fail`)
			})
		})
	}
}

func GroupsSearch(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("search via /Groups/.search", func(t *testing.T) {
			g1, err := cl.Group().Create().
				DisplayName("search-test1").
				Do(context.TODO())
			require.NoError(t, err, `Create should succeed`)

			g2, err := cl.Group().Create().
				DisplayName("search-test2").
				Do(context.TODO())
			require.NoError(t, err, `Create should succeed`)

			_ = g1
			_ = g2

			//nolint:errcheck
			defer cl.Group().Delete(g1.ID()).
				Do(context.TODO())
			//nolint:errcheck
			defer cl.Group().Delete(g2.ID()).
				Do(context.TODO())

			t.Run("Use `sw` predicate", func(t *testing.T) {
				res, err := cl.Group().Search().
					Trace(TraceWriter).
					Attributes(`displayName`).
					Filter(`displayName sw "search-test"`).
					StartIndex(1).
					Count(10).
					Do(context.TODO())
				require.NoError(t, err, `cl.Search should succeed`)
				require.Equal(t, 2, res.TotalResults(), `total results should be 2`)
			})
			t.Run("Use `co` predicate", func(t *testing.T) {
				res, err := cl.Group().Search().
					Trace(TraceWriter).
					Attributes(`displayName`).
					Filter(`displayName co "arch-test"`).
					StartIndex(1).
					Count(10).
					Do(context.TODO())
				require.NoError(t, err, `cl.Search should succeed`)
				require.Equal(t, 2, res.TotalResults(), `total results should be 2`)
			})
			t.Run("Use `ew` predicate", func(t *testing.T) {
				res, err := cl.Group().Search().
					Trace(TraceWriter).
					Attributes(`displayName`).
					Filter(`displayName ew "test1"`).
					StartIndex(1).
					Count(10).
					Do(context.TODO())
				require.NoError(t, err, `cl.Search should succeed`)
				require.Equal(t, 1, res.TotalResults(), `total results should be 2`)
			})
			t.Run("Use `eq` predicate", func(t *testing.T) {
				res, err := cl.Group().Search().
					Trace(TraceWriter).
					Attributes(`displayName`).
					Filter(`displayName eq "search-test1"`).
					StartIndex(1).
					Count(10).
					Do(context.TODO())
				require.NoError(t, err, `cl.Search should succeed`)
				require.Equal(t, 1, res.TotalResults(), `total results should be 2`)
			})
		})
	}
}

func ServiceProviderConfig(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		spc, err := cl.Meta().GetServiceProviderConfig().
			Trace(TraceWriter).
			Do(context.TODO())
		require.NoError(t, err, `cl.GetServiceProviderConfig should succeed`)
		_ = spc // TODO: perform more checks
	}
}

func ResourceTypes(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		spc, err := cl.Meta().GetResourceTypes().
			Trace(TraceWriter).
			Do(context.TODO())
		require.NoError(t, err, `cl.ResourceTypes should succeed`)
		_ = spc // TODO: perform more checks
	}
}

func Schemas(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("Fetch all schemas", func(t *testing.T) {
			list, err := cl.Meta().GetSchemas().
				Trace(TraceWriter).
				Do(context.TODO())
			require.NoError(t, err, `cl.GetSchema should succeed`)
			_ = list
		})
		t.Run("Fetch individual schemas", func(t *testing.T) {
			for _, u := range []string{resource.UserSchemaURI, resource.GroupSchemaURI} {
				s, err := cl.Meta().GetSchema(u).
					Trace(TraceWriter).
					Do(context.TODO())
				require.NoError(t, err, `cl.GetSchemas should succeed`)
				_ = s
			}
		})
	}
}
