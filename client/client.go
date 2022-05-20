//go:generate ../tools/cmd/genclient.sh

package client

import (
	"net/http"
	"strings"
)

type Client struct {
	baseURL string
	httpcl  *http.Client
}

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
