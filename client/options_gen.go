// This file is auto-generated by tools/cmd/genoptions/main.go. DO NOT EDIT

package client

import (
	"io"
	"net/http"

	"github.com/lestrrat-go/option"
)

type Option = option.Interface

// NewOption is a type of option that can be passed to `client.New()`
type NewOption interface {
	Option
	newOption()
}

type newOption struct {
	Option
}

func (*newOption) newOption() {}

// RequestOption is a type of option that can be passed to any of the
// client request methods.
type RequestOption interface {
	Option
	requestOption()
}

type requestOption struct {
	Option
}

func (*requestOption) requestOption() {}

type identClient struct{}
type identTrace struct{}

func (identClient) String() string {
	return "WithClient"
}

func (identTrace) String() string {
	return "WithTrace"
}

// WithClient specifies the http.Client instance to use.
func WithClient(v *http.Client) NewOption {
	return &newOption{option.New(identClient{}, v)}
}

// WithTrace allows users to specify a destination to write out the
// details of a request
func WithTrace(v io.Writer) RequestOption {
	return &requestOption{option.New(identTrace{}, v)}
}