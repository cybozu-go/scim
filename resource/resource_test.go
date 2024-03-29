package resource_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/cybozu-go/scim/resource"
	"github.com/stretchr/testify/require"
)

func TestEnterpriseUser(t *testing.T) {
	u, err := resource.NewUserBuilder().
		UserName(`johndoe`).
		Extension(resource.EnterpriseUserSchemaURI, resource.NewEnterpriseUserBuilder().
			Organization(`IT`).MustBuild(),
		).
		Build()
	require.NoError(t, err, `building resource.User should succeed`)

	var eu resource.EnterpriseUser
	require.NoError(t, u.Get(resource.EnterpriseUserSchemaURI, &eu), `u.Get(EnterpriseUserSchemaURI) should succeed`)
	_ = &eu
}

func TestDateTime(t *testing.T) {
	// Load the timezone that is not local
	var tz *time.Location
	v, err := time.LoadLocation(`Asia/Tokyo`)
	require.NoError(t, err, `time.LoadLocation should succeed`)
	tz = v

	ref := time.Date(2022, 2, 22, 14, 22, 22, 987654321, time.UTC)
	testcases := []struct {
		Value    string
		Expected time.Time
		Error    bool
	}{
		{
			Value:    ref.Format(`2006-01-02T15:04:05Z`),
			Expected: ref.Truncate(time.Second),
		},
		{
			Value:    ref.Format(`2006-01-02T15:04:05.9999999999Z`),
			Expected: ref,
		},
		{
			Value:    ref.In(tz).Format(`2006-01-02T15:04:05Z0700`),
			Expected: ref.Truncate(time.Second).In(tz),
		},
		{
			Value:    ref.In(tz).Format(`2006-01-02T15:04:05Z07:00`),
			Expected: ref.Truncate(time.Second).In(tz),
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.Value, func(t *testing.T) {
			parsed, err := resource.ParseDateTime(tc.Value)

			require.NoError(t, err, `resource.ParseDateTime should succeed`)
			require.Equal(t, tc.Expected.UTC(), parsed.UTC())
		})
	}
}

func TestPhoneNumber(t *testing.T) {
	testcases := []struct {
		Value string
		Error bool
	}{
		{
			Value: `tel:+1-201-555-0123`,
		},
		{
			Value: `denwa:+1-201-555-0123`,
			Error: true,
		},
		{
			Value: `+1-201-555-0123`,
			Error: true,
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.Value, func(t *testing.T) {
			var b resource.Builder
			r, err := b.PhoneNumber().
				Value(tc.Value).
				Build()
			_ = r
			if tc.Error {
				require.Error(t, err, `building phone number %q should fail`, tc.Value)
			} else {
				require.NoError(t, err, `building phone number %q should succeed`, tc.Value)
			}
		})
	}
}

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
		var en string
		require.NoError(t, u.GetExtension(`employeeNumber`, resource.EnterpriseUserSchemaURI, &en), `u.GetExtension should succeed`)
		require.Equal(t, `701984`, en)

		var cc string
		require.NoError(t, u.GetExtension(`costCenter`, resource.EnterpriseUserSchemaURI, &cc), `u.GetExtension should succeed`)
		require.Equal(t, `4130`, cc)
	})
	t.Run("Names", func(t *testing.T) {
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
