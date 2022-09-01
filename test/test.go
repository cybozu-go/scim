package test

import (
	"context"
	"embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
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

// Global switches to control which features to test
// var supportsBulk bool  // current UNUSED
var supportsPatch bool

func init() {
	v, err := strconv.ParseBool(os.Getenv(`SCIM_TRACE`))
	if err == nil {
		if v {
			TraceWriter = os.Stdout
		}
	}
}

type testClient struct {
	httpcl *http.Client
	token  string
}

func (c *testClient) Do(r *http.Request) (*http.Response, error) {
	r.Header.Add(`Authorization`, fmt.Sprintf(`Bearer %s`, c.token))
	return c.httpcl.Do(r)
}

func RunConformanceTests(t *testing.T, name string, backend interface{}) {
	t.Run(name, func(t *testing.T) {
		hh, err := server.NewServer(backend)
		require.NoError(t, err, `server.NewServer should succeed`)

		tok := "123456"
		ohh := hh
		hh = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hdr := r.Header.Get(`Authorization`)
			if hdr == "" {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			got := strings.TrimPrefix(hdr, "Bearer ")
			if got != tok {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			ohh.ServeHTTP(w, r)
		})
		srv := httptest.NewServer(hh)

		httpcl := &testClient{
			httpcl: srv.Client(),
			token:  tok,
		}

		cl := client.New(srv.URL, client.WithClient(httpcl), client.WithTrace(TraceWriter))

		t.Run(`Meta`, func(t *testing.T) {
			t.Run(`ServiceProviderConfig`, ServiceProviderConfig(t, cl))
			t.Run("ResourceTypes", ResourceTypes(t, cl))
			t.Run("Schemas", Schemas(t, cl))
		})
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
		t.Run("Mixed Search", MixedSearch(t, cl))
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
					Do(context.TODO())
				if !assert.NoError(t, err, `user creation should succeed`) {
					t.Logf("invalid data: %s", e)
					return
				}
			}
		})
		t.Run("Load Group fixtures", func(t *testing.T) {
			var groups []*resource.Group

			// Unfortunately we can't just automatically create Groups because it
			// requires that we know the members' internal IDs
			defs := []struct {
				Name   string
				Filter string
			}{
				{
					Name:   `Zemeckis Crew`,
					Filter: `emails.value ew "zemeckis-crew.com"`,
				},
				{
					Name:   `Zemeckis Actors`,
					Filter: `emails.value ew "zemeckis-actors.com"`,
				},
			}
			for _, def := range defs {
				def := def
				t.Run(fmt.Sprintf("Create group %s", def.Name), func(t *testing.T) {
					list, err := cl.User().Search().
						Filter(def.Filter).
						Do(context.TODO())
					require.NoError(t, err, `user search should succeed`)

					createGroupCall := cl.Group().Create().
						DisplayName(def.Name)
					for _, r := range list.Resources() {
						createGroupCall.MemberFrom(r)
					}
					g, err := createGroupCall.Do(context.TODO())
					require.NoError(t, err, `group creation should succeed`)
					t.Logf(`Created group %q with %d members`, g.DisplayName(), len(list.Resources()))
					groups = append(groups, g)
				})
			}

			t.Run("Create group Everybody", func(t *testing.T) {
				// Now create a group that contains everybody
				createGroupCall := cl.Group().Create().
					DisplayName("Everybody")
				for _, g := range groups {
					createGroupCall.MemberFrom(g)
				}
				_, err := createGroupCall.Do(context.TODO())
				require.NoError(t, err, `group creation should succeed`)
			})
		})
	}
}

func stockUserCreateCall(cl *client.Client) *client.CreateUserCall {
	f, err := datasrc.Open(`data/stick.png`)
	if err != nil {
		panic("could not open figure.png: " + err.Error())
	}

	stickFigure, err := io.ReadAll(f)
	if err != nil {
		panic("could not read figure.png: " + err.Error())
	}

	stickFigureHex := hex.EncodeToString(stickFigure)

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
			MustBuild()).
		Roles(resource.NewRoleBuilder().
			Value("Director of Human Resources").
			MustBuild()).
		PhoneNumbers(resource.NewPhoneNumberBuilder().
			Value("tel:+1-999-9999-9999").
			MustBuild()).
		Photos(
			resource.NewPhotoBuilder().
				Value(`data:image/png;base64,` + stickFigureHex).
				MustBuild(),
		)
}

func stockGroupCreateCall(cl *client.Client) *client.CreateGroupCall {
	u1, err := cl.User().Create().
		UserName("jsmith").
		ExternalID("jsmith").
		DisplayName("John Smith").
		Title("Manager").
		Do(context.TODO())
	if err != nil {
		panic(err)
	}
	u2, err := cl.User().Create().
		UserName("acooper").
		ExternalID("acooper").
		DisplayName("Alice Cooper").
		Title("Manager").
		Do(context.TODO())
	if err != nil {
		panic(err)
	}
	u3, err := cl.User().Create().
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
		DisplayName("Engineering").
		Members(members...)
}

func UsersSearch(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("search via /Users/.search", func(t *testing.T) {
			testcases := []struct {
				Query        string
				TotalResults int
			}{
				{
					Query:        `roles.value eq "director"`,
					TotalResults: 1,
				},
				{
					Query:        `roles.value eq "actor"`,
					TotalResults: 4,
				},
				{
					Query:        `roles.value eq "actor" OR roles.value eq "director"`,
					TotalResults: 5,
				},
				{
					Query:        `roles.value eq "actor" AND roles.value eq "director"`,
					TotalResults: 0,
				},
				{
					Query:        `phoneNumbers.value co "123-456-7890"`,
					TotalResults: 1,
				},
				{
					Query:        `userName sw "fg"`,
					TotalResults: 1,
				},
				{
					Query:        `roles.value sw "act"`,
					TotalResults: 4,
				},
			}

			for _, tc := range testcases {
				tc := tc
				t.Run(tc.Query, func(t *testing.T) {
					res, err := cl.User().Search().
						Filter(tc.Query).
						Do(context.TODO())
					require.NoError(t, err, `search should succeed`)
					require.Equal(t, tc.TotalResults, res.TotalResults(), `total results should be %d`, tc.TotalResults)
				})
			}
		})
	}
}

func UsersFetch(t *testing.T, cl *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("Fetch unknown user ID", func(t *testing.T) {
			u, err := cl.User().Get("foobar").
				Do(context.TODO())
			require.Nil(t, u, `Get return value should be nil`)
			require.Error(t, err, `Get should fail`)
		})
		t.Run("Fetch user", func(t *testing.T) {
			createdUser, err := stockUserCreateCall(cl).
				Do(context.TODO())
			require.NoError(t, err, `Create should succeed`)

			//nolint:errcheck
			defer cl.User().Delete(createdUser.ID()).
				Do(context.TODO())

			u, err := cl.User().Get(createdUser.ID()).
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
		t.Run("Fetch user with excludedAttributes", func(t *testing.T) {
			t.Run("exclude attributes with returned=always", func(t *testing.T) {
				createdUser := createStockUser(t, cl)
				u, err := cl.User().Get(createdUser.ID()).
					ExcludedAttributes("id", "name").
					Do(context.TODO())
				require.NoError(t, err, `Get should succeed`)

				//nolint:errcheck
				defer cl.User().Delete(createdUser.ID()).
					Do(context.TODO())

					// ID is returned=always
				require.Equal(t, createdUser.ID(), u.ID(), `ID should match`)
				require.Equal(t, createdUser.UserName(), u.UserName(), `UserName should match`)
				require.Equal(t, createdUser.Emails(), u.Emails(), `Emails should match`)
				require.NotNil(t, u.Meta(), `Meta should not be nil`)
				require.Equal(t, `User`, u.Meta().ResourceType(), `meta.resource_type should match`)
				require.NotEmpty(t, u.Meta().Version(), `meta.version should be populated`)
				require.Nil(t, u.Name(), `Name should be nil`)
			})
		})
		t.Run("Fetch (search) user with indirect group memberships", func(t *testing.T) {
			ctx := context.TODO()
			res, err := cl.User().Search().
				Filter(`displayName eq "Robert Zemeckis"`).
				Do(context.TODO())
			require.NoError(t, err, `cl.Search should succeed`)
			require.Equal(t, 1, res.TotalResults(), `total results should be 1`)

			for _, r := range res.Resources() {
				u, ok := r.(*resource.User)
				require.True(t, ok, `resources should be an instance of *resource.User`)
				if !assert.Len(t, u.Groups(), 2, `user should be a member of two groups`) {
					for i, gm := range u.Groups() {
						t.Logf("%0d: %s", i, gm.Value())
					}
					return
				}

				// we're not going to count the number of direct/indirect
				// memberships, but we're just going to make sure that
				// both exist in the list
				var directs int
				var indirects int
				for _, gm := range u.Groups() {
					g, err := cl.Group().Get(gm.Value()).Do(ctx)
					require.NoError(t, err, `fetching group should succeed`)
					t.Logf("user belongs to group %q (%s)", g.DisplayName(), gm.Type())
					switch gm.Type() {
					case `direct`:
						directs++
					case `indirect`:
						indirects++
					default:
						require.Fail(t, `expected either "direct" or "indirect", got %q`, gm.Type())
					}
				}

				require.True(t, directs > 0, `there should be more than 1 direct memberships`)
				require.True(t, indirects > 0, `there should be more than 1 indirect memberships`)
			}
		})
	}
}

func createStockUser(t *testing.T, cl *client.Client) *resource.User {
	var created *resource.User
	t.Run("Create stock user object", func(t *testing.T) {
		u, err := stockUserCreateCall(cl).
			Do(context.TODO())
		require.NoError(t, err, `Create should succeed`)

		// Make sure that it's really there
		fetched, err := cl.User().Get(u.ID()).
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
		require.True(t, len(createdUser.Photos()) > 0, `user should have some photos`)
		for _, photo := range createdUser.Photos() {
			require.Regexp(t, `^https://.+\.png$`, photo.Value(), `value should be a URL: got %s`, photo.Value())
		}

		etag := createdUser.Meta().Version()

		t.Run("Fetch user", func(t *testing.T) {
			u, err := cl.User().Get(createdUser.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)
			require.NotNil(t, u.Meta(), `meta should be non-nil`)
			require.Equal(t, etag, u.Meta().Version(), `versions should match`)
		})
		t.Run("Patch user", func(t *testing.T) {
			if !supportsPatch {
				t.Skip("Skipping test because Patch operations are not supported on this server.")
			}

			// Copied from RFC7644:
			// The result of the add operation depends upon what the target location
			// indicated by "path" references:
			//
			// >>> NOT IMPLEMENTED <<<
			// o  If omitted, the target location is assumed to be the resource
			//    itself.  The "value" parameter contains a set of attributes to be
			//    added to the resource.
			//
			// o  If the target location does not exist, the attribute and value are
			//    added.
			//
			// o  If the target location specifies a complex attribute, a set of
			//    sub-attributes SHALL be specified in the "value" parameter.
			//
			// o  If the target location specifies a multi-valued attribute, a new
			//    value is added to the attribute.
			//
			// o  If the target location specifies a single-valued attribute, the
			//    existing value is replaced.
			//
			// o  If the target location specifies an attribute that does not exist
			//    (has no value), the attribute is added with the new value.
			//
			// o  If the target location exists, the value is replaced.
			//
			// o  If the target location already contains the value specified, no
			//    changes SHOULD be made to the resource, and a success response
			//    SHOULD be returned.  Unless other operations change the resource,
			//    this operation SHALL NOT change the modify timestamp of the
			//    resource.
			t.Run("Add", func(t *testing.T) {
				t.Run("Empty single-valued attribute", func(t *testing.T) {
					require.Empty(t, createdUser.Title())
					patched, err := cl.User().Patch(createdUser.ID()).
						Operations(
							resource.NewPatchOperationBuilder().
								Op(resource.PatchAdd).
								Path(`title`).
								Value(`added title`).
								MustBuild(),
						).
						Do(context.TODO())
					require.NoError(t, err, `patch should succeed`)
					require.Equal(t, `added title`, patched.Title())
				})
				t.Run("Existing single-valued attribute", func(t *testing.T) {
					// This is confusing, but SCIM PATCH `add` operation _REPLACES_
					// a value when the target value is a single valued element
					patched, err := cl.User().Patch(createdUser.ID()).
						Operations(
							resource.NewPatchOperationBuilder().
								Op(resource.PatchAdd).
								Path(`title`).
								Value(`patched title`).
								MustBuild(),
						).
						Do(context.TODO())
					require.NoError(t, err, `patch should succeed`)
					require.Equal(t, `patched title`, patched.Title())
				})
			})
			t.Run("Add complex", func(t *testing.T) {
				addDirectorOfFinance := func(expected int) func(*testing.T) {
					return func(t *testing.T) {
						// This adds to the initially empty field
						require.Empty(t, createdUser.IMS())
						patched, err := cl.User().Patch(createdUser.ID()).
							Operations(
								resource.NewPatchOperationBuilder().
									Op(resource.PatchAdd).
									Path(`ims`).
									Value(
										resource.NewIMSBuilder().
											Value("babs919587").
											Type("ICQ").
											MustBuild(),
									).
									MustBuild(),
							).
							Do(context.TODO())
						require.NoError(t, err, `patch should succeed`)
						require.Len(t, patched.IMS(), expected)
					}
				}
				t.Run("Empty multi-valued field", addDirectorOfFinance(1))
				t.Run("Non-empty multi-valued field", func(t *testing.T) {
					// This should result in one more item in the field
					patched, err := cl.User().Patch(createdUser.ID()).
						Operations(
							resource.NewPatchOperationBuilder().
								Op(resource.PatchAdd).
								Path(`ims`).
								Value(
									resource.NewIMSBuilder().
										Value("bjensen@hotmail.com").
										Type("Skype").
										MustBuild(),
								).
								MustBuild(),
						).
						Do(context.TODO())
					require.NoError(t, err, `patch should succeed`)
					require.Len(t, patched.IMS(), 2)
				})
				// Duplicate items should not be added
				t.Run("Duplicate item on multi-valued field", addDirectorOfFinance(2))
			})
			// The "remove" operation removes the value at the target location
			// specified by the required attribute "path".  The operation performs
			// the following functions, depending on the target location specified
			// by "path":
			//
			// o  If "path" is unspecified, the operation fails with HTTP status
			//    code 400 and a "scimType" error code of "noTarget".
			//
			// o  If the target location is a single-value attribute, the attribute
			//    and its associated value is removed, and the attribute SHALL be
			//    considered unassigned.
			//
			// o  If the target location is a multi-valued attribute and no filter
			//    is specified, the attribute and all values are removed, and the
			//    attribute SHALL be considered unassigned.
			//
			// o  If the target location is a multi-valued attribute and a complex
			//    filter is specified comparing a "value", the values matched by the
			//    filter are removed.  If no other values remain after removal of
			//    the selected values, the multi-valued attribute SHALL be
			//    considered unassigned.
			//
			// o  If the target location is a complex multi-valued attribute and a
			//    complex filter is specified based on the attribute's
			//    sub-attributes, the matching records are removed.  Sub-attributes
			//    whose values have been removed SHALL be considered unassigned.  If
			//    the complex multi-valued attribute has no remaining records, the
			//    attribute SHALL be considered unassigned.
			//
			// If an attribute is removed or becomes unassigned and is defined as a
			// required attribute or a read-only attribute, the server SHALL return
			// an HTTP response status code and a JSON detail error response as
			// defined in Section 3.12, with a "scimType" error code of "mutability".
			t.Run("Remove", func(t *testing.T) {
				t.Run("Empty path", func(t *testing.T) {
					_, err := cl.User().Patch(createdUser.ID()).
						Operations(
							resource.NewPatchOperationBuilder().
								Op(resource.PatchRemove).
								Path(``).
								MustBuild(),
						).
						Do(context.TODO())
					var serr *resource.Error
					require.True(t, errors.As(err, &serr), `error should be a resource.Error type`)
					require.Equal(t, resource.ErrNoTarget, serr.ScimType())
				})
				t.Run("Single-value attribute", func(t *testing.T) {
					user, err := cl.User().Patch(createdUser.ID()).
						Operations(
							resource.NewPatchOperationBuilder().
								Op(resource.PatchRemove).
								Path(`title`).
								MustBuild(),
						).
						Do(context.TODO())
					require.NoError(t, err, `patch should succeed`)
					require.NotNil(t, user, `user should be sent back`)
					require.Equal(t, ``, user.Title())
				})
			})
		})
		t.Run("Replace user", func(t *testing.T) {
			replaced, err := cl.User().Replace(createdUser.ID()).
				ExternalID(createdUser.ExternalID()).
				Emails(resource.NewEmailBuilder().
					Value("babs-new@jensen.org").
					Primary(true).
					MustBuild()).
				PhoneNumbers(resource.NewPhoneNumberBuilder().
					Value("tel:+91-111-1111-1111").
					MustBuild()).
				Do(context.TODO())
			require.NoError(t, err, `ReplaceUser should succeed`)

			// Note: this implementation may differ from server to server.
			// In this version we assume that while the `userName` is read-write,
			// it cannot be empty. Thus, we assume that even when you replace an existing
			// user object with another version that does not have a valid `userName`
			// field, the old value should still persist
			require.Equal(t, replaced.UserName(), createdUser.UserName(), `UserName should match`)

			// We explicitly asked to persist these fields
			require.Equal(t, replaced.ExternalID(), createdUser.ExternalID(), `ExternalID should match`)

			// We cleared these fields
			require.Empty(t, replaced.DisplayName(), `DisplayName should be empty`)
			require.Empty(t, replaced.Name(), `Name should be empty`)
			require.Empty(t, replaced.Roles(), `Roles should be empty`)
			require.Empty(t, replaced.Photos(), `Photos should be empty`)

			// These fields were explicitly set to a different value
			require.NotEqual(t, replaced.Emails(), createdUser.Emails())
			require.NotEqual(t, replaced.PhoneNumbers(), createdUser.PhoneNumbers())

			// we need to validate the result from PUT and GET
			fetched, err := cl.User().Get(createdUser.ID()).
				Do(context.TODO())
			require.NoError(t, err, `Get should succeed`)

			// First off, u and fetched should be the same object
			require.Equal(t, replaced, fetched)

			// Just making sure that the versions are updated
			require.NotEqual(t, createdUser.Meta().Version(), replaced.Meta().Version(), "meta.version should NOT match")

			testcases := []struct {
				Name string
				User *resource.User
			}{
				{
					Name: "Result from issuing replace",
					User: replaced,
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
					require.NotEqual(t, etag, u.Meta().Version(), `versions should NOT match`)

					emails := u.Emails()
					require.Len(t, emails, 1)
					for _, email := range emails {
						// hardcoded for loop testing is weird, I know.
						// am just expecting it be expanded in the future
						require.Equal(t, `babs-new@jensen.org`, email.Value())
						require.True(t, email.Primary())
					}

					phoneNumbers := u.PhoneNumbers()
					require.Len(t, phoneNumbers, 1, `the number of phone numbers should match`)
					for _, pn := range phoneNumbers {
						require.Equal(t, `tel:+91-111-1111-1111`, pn.Value())
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
		ctx := context.TODO()

		// Create a single Group, apply Fetch, Replace, Delete
		createdUser, err := stockUserCreateCall(cl).Do(ctx)
		require.NoError(t, err, `Create should succeed`)
		createdGroup, err := stockGroupCreateCall(cl).Do(ctx)
		require.NoError(t, err, `Create should succeed`)

		etag := createdGroup.Meta().Version()

		t.Run("Fetch group", func(t *testing.T) {
			g, err := cl.Group().Get(createdGroup.ID()).Do(ctx)
			require.NoError(t, err, `Get should succeed`)
			require.Len(t, g.Members(), len(createdGroup.Members()), `there should be %d members`, len(createdGroup.Members()))
			require.NotNil(t, g.Meta(), `meta should be non-nil`)
			require.Equal(t, etag, g.Meta().Version(), `versions should match`)
		})
		t.Run("Patch group", func(t *testing.T) {
			if !supportsPatch {
				t.Skip("Skipping because Patch operations are not supported on this server")
			}
			g, err := cl.Group().Patch(createdGroup.ID()).
				Operations(
					resource.NewPatchOperationBuilder().
						Op(resource.PatchAdd).
						Path(`members`).
						Value(resource.NewGroupMemberBuilder().
							FromResource(createdUser).
							MustBuild(),
						).
						MustBuild(),
				).
				Do(ctx)
			require.NoError(t, err, `cl.Group().Patch() (add) should succeed`)
			require.Len(t, g.Members(), len(createdGroup.Members())+1)
			g2, err := cl.Group().Patch(createdGroup.ID()).
				Operations(
					resource.NewPatchOperationBuilder().
						Op(resource.PatchRemove).
						Path(fmt.Sprintf(`members[value eq %q]`, createdUser.ID())).
						MustBuild(),
				).
				Do(ctx)
			require.NoError(t, err, `cl.Group().Patch() (remove) should succeed`)
			require.Len(t, g2.Members(), len(createdGroup.Members()))
		})
		t.Run("Replace group", func(t *testing.T) {
			require.True(t, len(createdGroup.Members()) > 1, `more than 1 members required for this test to proceed`) // and avoid a panic

			u, err := cl.Group().Replace(createdGroup.ID()).
				DisplayName(createdGroup.DisplayName()).
				Members((createdGroup.Members())[1:]...).
				Do(ctx)
			require.NoError(t, err, `Replace should succeed`)

			// we need to validate the result from PUT and GET
			fetched, err := cl.Group().Get(createdGroup.ID()).Do(ctx)
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
					Attributes(`displayName`).
					Filter(`displayName ew "test1"`).
					StartIndex(1).
					Count(10).
					Do(context.TODO())
				require.NoError(t, err, `cl.Search should succeed`)
				require.Equal(t, 1, res.TotalResults(), `total results should be 2`)
				for _, r := range res.Resources() {
					g, ok := r.(*resource.Group)
					require.True(t, ok, `resource should be a Group`)
					require.True(t, strings.HasSuffix(g.DisplayName(), `test1`), `display name should match filter query`)
					require.Empty(t, g.Members())
				}
			})
			t.Run("Use `eq` predicate", func(t *testing.T) {
				res, err := cl.Group().Search().
					Attributes(`displayName`).
					Filter(`displayName eq "search-test1"`).
					StartIndex(1).
					Count(10).
					Do(context.TODO())
				require.NoError(t, err, `cl.Search should succeed`)
				require.Equal(t, 1, res.TotalResults(), `total results should be 1`)
				for _, r := range res.Resources() {
					g, ok := r.(*resource.Group)
					require.True(t, ok, `resource should be a Group`)
					require.Equal(t, g.DisplayName(), `search-test1`, `display name should match filter query`)
					require.Empty(t, g.Members())
				}
			})
		})
	}
}

func MixedSearch(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		res, err := cl.Search().
			Filter(`displayName co "Zemeckis"`).
			Do(context.TODO())
		require.NoError(t, err, `cl.Search should succeed`)

		// user Robert Zemckis, groups Zemeckis crew and Zemeckis actors
		const expectedCount = 3
		require.Equal(t, expectedCount, res.TotalResults(), `total results should match`)

		var groups int
		var users int
		var others int
		for _, r := range res.Resources() {
			switch r.(type) {
			case *resource.Group:
				groups++
			case *resource.User:
				users++
			default:
				others++
			}
		}

		require.Equal(t, 0, others, `there should be zero resources other than users/groups`)
		require.Equal(t, 1, users, `there should 1 user`)
		require.Equal(t, expectedCount-1, groups, `there should 1 group`)
	}
}

func ResourceTypes(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		spc, err := cl.Meta().GetResourceTypes().
			Do(context.TODO())
		require.NoError(t, err, `cl.ResourceTypes should succeed`)
		_ = spc // TODO: perform more checks
	}
}

func Schemas(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("Fetch all schemas", func(t *testing.T) {
			list, err := cl.Meta().GetSchemas().
				Do(context.TODO())
			require.NoError(t, err, `cl.GetSchema should succeed`)
			_ = list
		})
		t.Run("Fetch individual schemas", func(t *testing.T) {
			for _, u := range []string{resource.UserSchemaURI, resource.GroupSchemaURI} {
				s, err := cl.Meta().GetSchema(u).
					Do(context.TODO())
				require.NoError(t, err, `cl.GetSchemas should succeed`)
				_ = s
			}
		})
	}
}

func ServiceProviderConfig(t *testing.T, cl *client.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run(`Get /ServiceProviderConfig`, func(t *testing.T) {
			ctx := context.TODO()

			c, err := cl.Meta().GetServiceProviderConfig().Do(ctx)
			require.NoError(t, err, `GetServiceProviderConfig should succeed`)

			boolToSupported := func(b bool) string {
				if b {
					return "SUPPORTED"
				}
				return "N/A"
			}
			t.Logf("Server capabilities:")
			t.Logf("  Bulk operations : %s", boolToSupported(c.Bulk().Supported()))
			t.Logf("  Change password : %s", boolToSupported(c.ChangePassword().Supported()))
			t.Logf("  ETags           : %s", boolToSupported(c.ETag().Supported()))
			t.Logf("  Filters         : %s", boolToSupported(c.Filter().Supported()))
			t.Logf("  Patch operations: %s", boolToSupported(c.Patch().Supported()))
			t.Logf("  Sorting         : %s", boolToSupported(c.Sort().Supported()))

			supportsPatch = c.Patch().Supported()
		})
	}
}
