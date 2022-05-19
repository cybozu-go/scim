package resource_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/lestrrat-go/scim/resource"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	var b resource.Builder

	created, _ := time.Parse(time.RFC3339, "2010-01-23T04:56:22Z")
	lastModified, _ := time.Parse(time.RFC3339, "2011-05-13T04:42:34Z")

	u, err := b.User().
		Schemas("urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User").
		ID("2819c223-7f76-453a-413861904646").
		ExternalID("701984").
		UserName("bjensen@example.com").
		Name(
			b.Names().
				Formatted("Ms. Barbara J Jensen, III").
				FamilyName("Jensen").
				GivenName("Barbara").
				MiddleName("Jane").
				HonorificPrefix("Ms.").
				HonorificSuffix("III").
				MustBuild(),
		).
		Extension(
			resource.EnterpriseUserSchemaURI,
			b.EnterpriseUser().
				EmployeeNumber("701984").
				CostCenter("4130").
				MustBuild(),
		).
		Meta(b.Meta().
			ResourceType("User").
			Created(created).
			LastModified(lastModified).
			Version("W/\"3694e05e9dff591\"").
			Location("https://example.com/v2/Users/2819c223-7f76-453a-413861904646").
			MustBuild(),
		).
		Build()
	require.NoError(t, err, `UserBuilder.Build() should succeed`)
	require.Equal(t, []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"}, u.Schemas(), `values should match`)
	require.Equal(t, `bjensen@example.com`, u.UserName(), `values should match`)

	t.Run("EnterpriseUser", func(t *testing.T) {
		en, ok := u.Get(`employeeNumber`, resource.WithExtension(resource.EnterpriseUserSchemaURI))
		require.True(t, ok, `u.Get("employeeNumber") should succeed`)
		require.Equal(t, `701984`, en)
		cc, ok := u.Get(`costCenter`, resource.WithExtension(resource.EnterpriseUserSchemaURI))
		require.True(t, ok, `u.Get("costCenter") should succeed`)
		require.Equal(t, `4130`, cc)
	})
	t.Run("Name", func(t *testing.T) {
		names := u.Name()
		require.Equal(t, `Ms. Barbara J Jensen, III`, names.Formatted())
		require.Equal(t, `Jensen`, names.FamilyName())
		require.Equal(t, `Barbara`, names.GivenName())
		require.Equal(t, `Jane`, names.MiddleName())
		require.Equal(t, `Ms.`, names.HonorificPrefix())
		require.Equal(t, `III`, names.HonorificSuffix())
	})
	t.Run("Meta", func(t *testing.T) {
		meta := u.Meta()
		require.Equal(t, "User", meta.ResourceType())
		require.Equal(t, created, meta.Created())
		require.Equal(t, lastModified, meta.LastModified())
		require.Equal(t, "W/\"3694e05e9dff591\"", meta.Version())
		require.Equal(t, "https://example.com/v2/Users/2819c223-7f76-453a-413861904646", meta.Location())
	})

	t.Run("JSON", func(t *testing.T) {
		buf, err := json.MarshalIndent(u, "", "  ")
		require.NoError(t, err)
		t.Logf("%s", buf)
	})
}

func TestUnmarshal(t *testing.T) {
	t.Run(`sample`, func(t *testing.T) {
		const src = `{
     "schemas":
       ["urn:ietf:params:scim:schemas:core:2.0:User",
         "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"],
     "id": "2819c223-7f76-453a-413861904646",
     "externalId": "701984",
     "userName": "bjensen@example.com",
     "name": {
       "formatted": "Ms. Barbara J Jensen, III",
       "familyName": "Jensen",
       "givenName": "Barbara",
       "middleName": "Jane",
       "honorificPrefix": "Ms.",
       "honorificSuffix": "III"
     },
     "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User": {
       "employeeNumber": "701984",
       "costCenter": "4130"
     },
     "meta": {
       "resourceType": "User",
       "created": "2010-01-23T04:56:22Z",
       "lastModified": "2011-05-13T04:42:34Z",
       "version": "W\/\"3694e05e9dff591\"",
       "location":
         "https://example.com/v2/Users/2819c223-7f76-453a-413861904646"
     }
    }`
		var user resource.User
		require.NoError(t, json.Unmarshal([]byte(src), &user), `json.Marshal should succeed`)
		require.Equal(t, []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"}, user.Schemas(), `user.Schemas() should match`)
		require.Equal(t, `701984`, user.ExternalID(), `user.ExternalID() should match`)
		require.Equal(t, `bjensen@example.com`, user.UserName(), `user.UserName() should match`)
	})
}
