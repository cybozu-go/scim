//go:generate ../tools/cmd/genclient.sh

package client

import (
	"net/http"
	"strings"
)

// Client represents an object that acts on behalf of the user to
// interact with the SCIM server
type Client struct {
	baseURL string
	httpcl  *http.Client
}

// New creates a Client instance. The `baseURL` parameter is required,
// and must point to the root of the SCIM server.
//
// If you need to perform authentication, create an HTTP client that
// handles the authentication, and pass it as an option
func New(baseURL string, options ...NewOption) *Client {
	httpcl := http.DefaultClient
	//nolint:forcetypeassert
	for _, option := range options {
		switch option.Ident() {
		case identClient{}:
			httpcl = option.Value().(*http.Client)
		}
	}

	return &Client{
		// Strip trailing slash
		baseURL: strings.TrimSuffix(baseURL, "/"),
		httpcl:  httpcl,
	}
}
