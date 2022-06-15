package sample_test

import (
	"testing"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample"
	"github.com/cybozu-go/scim/test"
	"github.com/stretchr/testify/require"
)

func TestSample(t *testing.T) {
	var b resource.Builder
	spc, err := b.ServiceProviderConfig().
		AuthenticationSchemes(
			b.AuthenticationScheme().
				Name("OAuth Bearer Token").
				Description("Authentication scheme using the OAuth Bearer Token Standard").
				SpecURI("http://www.rfc-editor.org/info/rfc6750").
				DocumentationURI("http://example.com/help/oauth.html").
				Type(resource.OAuthBearerToken).
				MustBuild(),
		).
		Bulk(b.BulkSupport().
			Supported(false).
			MaxOperations(0).
			MaxPayloadSize(0).
			MustBuild(),
		).
		Filter(b.FilterSupport().
			Supported(true).
			MaxResults(200). // TODO: arbitrary value used
			MustBuild(),
		).
		Sort(b.GenericSupport().
			Supported(false).
			MustBuild(),
		).
		Etag(b.GenericSupport().
			Supported(false).
			MustBuild(),
		).
		Patch(b.GenericSupport().
			Supported(false).
			MustBuild(),
		).
		ChangePassword(b.GenericSupport().
			Supported(false).
			MustBuild(),
		).
		Build()
	require.NoError(t, err, `ServiceProviderConfig should succeed`)

	s, err := sample.New("file:ent?mode=memory&cache=shared&_fk=1", spc)
	require.NoError(t, err, `sample.New should succeed`)

	test.RunConformanceTests(t, "Sample backend", s)
}
