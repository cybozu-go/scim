package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/cybozu-go/scim/resource"
)

type SearchService struct {
	client *Client
}

func (client *Client) Search() *SearchService {
	return &SearchService{
		client: client,
	}
}

type SearchCall struct {
	builder *resource.SearchRequestBuilder
	object  *resource.SearchRequest
	err     error
	client  *Client
	trace   io.Writer
}

func (call *SearchCall) payload() (*resource.SearchRequest, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *SearchCall) FromJSON(data []byte) *SearchCall {
	if call.err != nil {
		return call
	}
	var in resource.SearchRequest
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

func (svc *SearchService) Search() *SearchCall {
	return &SearchCall{
		builder: resource.NewSearchRequestBuilder(),
		client:  svc.client,
	}
}

func (call *SearchCall) Attributes(v ...string) *SearchCall {
	call.builder.Attributes(v...)
	return call
}

func (call *SearchCall) Count(v int) *SearchCall {
	call.builder.Count(v)
	return call
}

func (call *SearchCall) ExludedAttributes(v ...string) *SearchCall {
	call.builder.ExludedAttributes(v...)
	return call
}

func (call *SearchCall) Filter(v string) *SearchCall {
	call.builder.Filter(v)
	return call
}

func (call *SearchCall) SortBy(v string) *SearchCall {
	call.builder.SortBy(v)
	return call
}

func (call *SearchCall) SortOrder(v string) *SearchCall {
	call.builder.SortOrder(v)
	return call
}

func (call *SearchCall) StartIndex(v int) *SearchCall {
	call.builder.StartIndex(v)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *SearchCall) Extension(uri string, value interface{}) *SearchCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *SearchCall) Validator(v resource.SearchRequestValidator) *SearchCall {
	call.builder.Validator(v)
	return call
}

func (call *SearchCall) Trace(w io.Writer) *SearchCall {
	call.trace = w
	return call
}

func (call *SearchCall) makeURL() string {
	return call.client.baseURL + "/.search"
}

func (call *SearchCall) Do(ctx context.Context) (*resource.ListResponse, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for SearchCall: %w`, err)
	}

	trace := call.trace
	if trace == nil {
		trace = call.client.trace
	}
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, "trace: client sending call request to %q\n", u)
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(payload); err != nil {
		return nil, fmt.Errorf(`failed to encode call request: %w`, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, &body)
	if err != nil {
		return nil, fmt.Errorf(`failed to create new HTTP request: %w`, err)
	}

	req.Header.Set(`Content-Type`, `application/scim+json`)
	req.Header.Set(`Accept`, `application/scim+json`)

	if trace != nil {
		buf, _ := httputil.DumpRequestOut(req, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}

	res, err := call.client.httpcl.Do(req)
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	if err != nil {
		return nil, fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`expected call response %d, got (%d)`, http.StatusOK, res.StatusCode)
	}

	var respayload resource.ListResponse
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}
