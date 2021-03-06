//go:generate ../tools/cmd/genclient.sh

package client

import (
	"io"
	"net/http"
	"strings"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client represents an object that acts on behalf of the user to
// interact with the SCIM server
type Client struct {
	baseURL string
	httpcl  HTTPClient
	trace   io.Writer
}

// New creates a Client instance. The `baseURL` parameter is required,
// and must point to the root of the SCIM server.
//
// If you need to perform authentication, create an HTTP client that
// handles the authentication, and pass it as an option
func New(baseURL string, options ...NewOption) *Client {
	var httpcl HTTPClient = http.DefaultClient
	trace := io.Discard
	//nolint:forcetypeassert
	for _, option := range options {
		switch option.Ident() {
		case identClient{}:
			httpcl = option.Value().(HTTPClient)
		case identTrace{}:
			trace = option.Value().(io.Writer)
		}
	}

	return &Client{
		// Strip trailing slash
		baseURL: strings.TrimSuffix(baseURL, "/"),
		httpcl:  httpcl,
		trace:   trace,
	}
}

func (client *Client) Search() *SearchCall {
	// This is a very rare case where the service name (Search) stutters
	// withe Call verb (Search)... so we break rule and just return SearchCall
	return (&SearchService{client: client}).Search()
}
