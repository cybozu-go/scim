package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/cybozu-go/scim/resource"
)

type GroupService struct {
	client *Client
}

// Group creates a new Service object to perform an operation
func (client *Client) Group() *GroupService {
	return &GroupService{
		client: client,
	}
}

// GetGroupCall is an encapsulation of a SCIM operation.
type GetGroupCall struct {
	builder *resource.PartialResourceRepresentationRequestBuilder
	object  *resource.PartialResourceRepresentationRequest
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *GetGroupCall) payload() (*resource.PartialResourceRepresentationRequest, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *GetGroupCall) FromJSON(data []byte) *GetGroupCall {
	if call.err != nil {
		return call
	}
	var in resource.PartialResourceRepresentationRequest
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

func (svc *GroupService) Get(id string) *GetGroupCall {
	return &GetGroupCall{
		builder: resource.NewPartialResourceRepresentationRequestBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *GetGroupCall) Attributes(v ...string) *GetGroupCall {
	call.builder.Attributes(v...)
	return call
}

func (call *GetGroupCall) ExcludedAttributes(v ...string) *GetGroupCall {
	call.builder.ExcludedAttributes(v...)
	return call
}

func (call *GetGroupCall) Trace(w io.Writer) *GetGroupCall {
	call.trace = w
	return call
}

func (call GetGroupCall) makeURL() string {
	return call.client.baseURL + "/Groups/" + call.id
}

func (call *GetGroupCall) Do(ctx context.Context) (*resource.Group, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for GetGroupCall: %w`, err)
	}

	trace := call.trace
	if trace == nil {
		trace = call.client.trace
	}
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, "trace: client sending call request to %q\n", u)
	}

	var vals url.Values
	m := make(map[string]interface{})
	if err := payload.AsMap(m); err != nil {
		return nil, fmt.Errorf(`failed to convert resource into map: %w`, err)
	}
	if len(m) > 0 {
		vals = make(url.Values)
		for key, value := range m {
			switch value := value.(type) {
			case []string:
				vals.Add(key, strings.Join(value, ","))
			default:
				vals.Add(key, fmt.Sprintf(`%s`, value))
			}
		}
	}
	if enc := vals.Encode(); len(enc) > 0 {
		u = u + "?" + vals.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf(`failed to create new HTTP request: %w`, err)
	}
	req.Header.Set(`Accept`, `application/scim+json`)

	if trace != nil {
		buf, _ := httputil.DumpRequestOut(req, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}

	res, err := call.client.httpcl.Do(req)
	if err != nil {
		return nil, fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var serr resource.Error
		var resBody bytes.Buffer
		if err := json.NewDecoder(io.TeeReader(res.Body, &resBody)).Decode(&serr); err != nil {
			return nil, fmt.Errorf("expected %d (got %d): %s", http.StatusOK, res.StatusCode, resBody.String())
		}
		return nil, &serr
	}

	var respayload resource.Group
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// CreateGroupCall is an encapsulation of a SCIM operation.
type CreateGroupCall struct {
	builder *resource.GroupBuilder
	object  *resource.Group
	err     error
	client  *Client
	trace   io.Writer
}

func (call *CreateGroupCall) payload() (*resource.Group, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *CreateGroupCall) FromJSON(data []byte) *CreateGroupCall {
	if call.err != nil {
		return call
	}
	var in resource.Group
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

func (svc *GroupService) Create() *CreateGroupCall {
	return &CreateGroupCall{
		builder: resource.NewGroupBuilder(),
		client:  svc.client,
	}
}

func (call *CreateGroupCall) DisplayName(v string) *CreateGroupCall {
	call.builder.DisplayName(v)
	return call
}

func (call *CreateGroupCall) ExternalID(v string) *CreateGroupCall {
	call.builder.ExternalID(v)
	return call
}

func (call *CreateGroupCall) Members(v ...*resource.GroupMember) *CreateGroupCall {
	call.builder.Members(v...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *CreateGroupCall) Extension(uri string, value interface{}) *CreateGroupCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *CreateGroupCall) Validator(v resource.GroupValidator) *CreateGroupCall {
	call.builder.Validator(v)
	return call
}

func (call *CreateGroupCall) Trace(w io.Writer) *CreateGroupCall {
	call.trace = w
	return call
}

func (call *CreateGroupCall) makeURL() string {
	return call.client.baseURL + "/Groups"
}

func (call *CreateGroupCall) Do(ctx context.Context) (*resource.Group, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for CreateGroupCall: %w`, err)
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
	if err != nil {
		return nil, fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		var serr resource.Error
		var resBody bytes.Buffer
		if err := json.NewDecoder(io.TeeReader(res.Body, &resBody)).Decode(&serr); err != nil {
			return nil, fmt.Errorf("expected %d (got %d): %s", http.StatusCreated, res.StatusCode, resBody.String())
		}
		return nil, &serr
	}

	var respayload resource.Group
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// ReplaceGroupCall is an encapsulation of a SCIM operation.
type ReplaceGroupCall struct {
	builder *resource.GroupBuilder
	object  *resource.Group
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *ReplaceGroupCall) payload() (*resource.Group, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *ReplaceGroupCall) FromJSON(data []byte) *ReplaceGroupCall {
	if call.err != nil {
		return call
	}
	var in resource.Group
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

func (svc *GroupService) Replace(id string) *ReplaceGroupCall {
	return &ReplaceGroupCall{
		builder: resource.NewGroupBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *ReplaceGroupCall) DisplayName(v string) *ReplaceGroupCall {
	call.builder.DisplayName(v)
	return call
}

func (call *ReplaceGroupCall) ExternalID(v string) *ReplaceGroupCall {
	call.builder.ExternalID(v)
	return call
}

func (call *ReplaceGroupCall) Members(v ...*resource.GroupMember) *ReplaceGroupCall {
	call.builder.Members(v...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *ReplaceGroupCall) Extension(uri string, value interface{}) *ReplaceGroupCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *ReplaceGroupCall) Validator(v resource.GroupValidator) *ReplaceGroupCall {
	call.builder.Validator(v)
	return call
}

func (call *ReplaceGroupCall) Trace(w io.Writer) *ReplaceGroupCall {
	call.trace = w
	return call
}

func (call ReplaceGroupCall) makeURL() string {
	return call.client.baseURL + "/Groups/" + call.id
}

func (call *ReplaceGroupCall) Do(ctx context.Context) (*resource.Group, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for ReplaceGroupCall: %w`, err)
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u, &body)
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
	if err != nil {
		return nil, fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var serr resource.Error
		var resBody bytes.Buffer
		if err := json.NewDecoder(io.TeeReader(res.Body, &resBody)).Decode(&serr); err != nil {
			return nil, fmt.Errorf("expected %d (got %d): %s", http.StatusOK, res.StatusCode, resBody.String())
		}
		return nil, &serr
	}

	var respayload resource.Group
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// PatchGroupCall is an encapsulation of a SCIM operation.
type PatchGroupCall struct {
	builder *resource.PatchRequestBuilder
	object  *resource.PatchRequest
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *PatchGroupCall) payload() (*resource.PatchRequest, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *PatchGroupCall) FromJSON(data []byte) *PatchGroupCall {
	if call.err != nil {
		return call
	}
	var in resource.PatchRequest
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

// Patch allows the user to patch parts of the group object
func (svc *GroupService) Patch(id string) *PatchGroupCall {
	return &PatchGroupCall{
		builder: resource.NewPatchRequestBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *PatchGroupCall) Operations(v ...*resource.PatchOperation) *PatchGroupCall {
	call.builder.Operations(v...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *PatchGroupCall) Extension(uri string, value interface{}) *PatchGroupCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *PatchGroupCall) Validator(v resource.PatchRequestValidator) *PatchGroupCall {
	call.builder.Validator(v)
	return call
}

func (call *PatchGroupCall) Trace(w io.Writer) *PatchGroupCall {
	call.trace = w
	return call
}

func (call PatchGroupCall) makeURL() string {
	return call.client.baseURL + "/Groups/" + call.id
}

func (call *PatchGroupCall) Do(ctx context.Context) (*resource.Group, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for PatchGroupCall: %w`, err)
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, u, &body)
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
	if err != nil {
		return nil, fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent {
		//nolint:nilnil
		return nil, nil
	}

	if res.StatusCode != http.StatusOK {
		var serr resource.Error
		var resBody bytes.Buffer
		if err := json.NewDecoder(io.TeeReader(res.Body, &resBody)).Decode(&serr); err != nil {
			return nil, fmt.Errorf("expected %d (got %d): %s", http.StatusOK, res.StatusCode, resBody.String())
		}
		return nil, &serr
	}

	var respayload resource.Group
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// DeleteGroupCall is an encapsulation of a SCIM operation.
type DeleteGroupCall struct {
	builder *resource.GroupBuilder
	object  *resource.Group
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *DeleteGroupCall) payload() (*resource.Group, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *DeleteGroupCall) FromJSON(data []byte) *DeleteGroupCall {
	if call.err != nil {
		return call
	}
	var in resource.Group
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

func (svc *GroupService) Delete(id string) *DeleteGroupCall {
	return &DeleteGroupCall{
		builder: resource.NewGroupBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *DeleteGroupCall) DisplayName(v string) *DeleteGroupCall {
	call.builder.DisplayName(v)
	return call
}

func (call *DeleteGroupCall) ExternalID(v string) *DeleteGroupCall {
	call.builder.ExternalID(v)
	return call
}

func (call *DeleteGroupCall) ID(v string) *DeleteGroupCall {
	call.builder.ID(v)
	return call
}

func (call *DeleteGroupCall) Members(v ...*resource.GroupMember) *DeleteGroupCall {
	call.builder.Members(v...)
	return call
}

func (call *DeleteGroupCall) Meta(v *resource.Meta) *DeleteGroupCall {
	call.builder.Meta(v)
	return call
}

func (call *DeleteGroupCall) Trace(w io.Writer) *DeleteGroupCall {
	call.trace = w
	return call
}

func (call DeleteGroupCall) makeURL() string {
	return call.client.baseURL + "/Groups/" + call.id
}

func (call *DeleteGroupCall) Do(ctx context.Context) error {
	if err := call.err; err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return fmt.Errorf(`failed to generate request payload for DeleteGroupCall: %w`, err)
	}

	trace := call.trace
	if trace == nil {
		trace = call.client.trace
	}
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, "trace: client sending call request to %q\n", u)
	}

	var vals url.Values
	m := make(map[string]interface{})
	if err := payload.AsMap(m); err != nil {
		return fmt.Errorf(`failed to convert resource into map: %w`, err)
	}
	if len(m) > 0 {
		vals = make(url.Values)
		for key, value := range m {
			switch value := value.(type) {
			case []string:
				vals.Add(key, strings.Join(value, ","))
			default:
				vals.Add(key, fmt.Sprintf(`%s`, value))
			}
		}
	}
	if enc := vals.Encode(); len(enc) > 0 {
		u = u + "?" + vals.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return fmt.Errorf(`failed to create new HTTP request: %w`, err)
	}
	req.Header.Set(`Accept`, `application/scim+json`)

	if trace != nil {
		buf, _ := httputil.DumpRequestOut(req, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}

	res, err := call.client.httpcl.Do(req)
	if err != nil {
		return fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		var serr resource.Error
		var resBody bytes.Buffer
		if err := json.NewDecoder(io.TeeReader(res.Body, &resBody)).Decode(&serr); err != nil {
			return fmt.Errorf("expected %d (got %d): %s", http.StatusNoContent, res.StatusCode, resBody.String())
		}
		return &serr
	}

	return nil
}

// SearchGroupCall is an encapsulation of a SCIM operation.
type SearchGroupCall struct {
	builder *resource.SearchRequestBuilder
	object  *resource.SearchRequest
	err     error
	client  *Client
	trace   io.Writer
}

func (call *SearchGroupCall) payload() (*resource.SearchRequest, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *SearchGroupCall) FromJSON(data []byte) *SearchGroupCall {
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

func (svc *GroupService) Search() *SearchGroupCall {
	return &SearchGroupCall{
		builder: resource.NewSearchRequestBuilder(),
		client:  svc.client,
	}
}

func (call *SearchGroupCall) Attributes(v ...string) *SearchGroupCall {
	call.builder.Attributes(v...)
	return call
}

func (call *SearchGroupCall) Count(v int) *SearchGroupCall {
	call.builder.Count(v)
	return call
}

func (call *SearchGroupCall) ExludedAttributes(v ...string) *SearchGroupCall {
	call.builder.ExludedAttributes(v...)
	return call
}

func (call *SearchGroupCall) Filter(v string) *SearchGroupCall {
	call.builder.Filter(v)
	return call
}

func (call *SearchGroupCall) SortBy(v string) *SearchGroupCall {
	call.builder.SortBy(v)
	return call
}

func (call *SearchGroupCall) SortOrder(v string) *SearchGroupCall {
	call.builder.SortOrder(v)
	return call
}

func (call *SearchGroupCall) StartIndex(v int) *SearchGroupCall {
	call.builder.StartIndex(v)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *SearchGroupCall) Extension(uri string, value interface{}) *SearchGroupCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *SearchGroupCall) Validator(v resource.SearchRequestValidator) *SearchGroupCall {
	call.builder.Validator(v)
	return call
}

func (call *SearchGroupCall) Trace(w io.Writer) *SearchGroupCall {
	call.trace = w
	return call
}

func (call *SearchGroupCall) makeURL() string {
	return call.client.baseURL + "/Groups/.search"
}

func (call *SearchGroupCall) Do(ctx context.Context) (*resource.ListResponse, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for SearchGroupCall: %w`, err)
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
	if err != nil {
		return nil, fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var serr resource.Error
		var resBody bytes.Buffer
		if err := json.NewDecoder(io.TeeReader(res.Body, &resBody)).Decode(&serr); err != nil {
			return nil, fmt.Errorf("expected %d (got %d): %s", http.StatusOK, res.StatusCode, resBody.String())
		}
		return nil, &serr
	}

	var respayload resource.ListResponse
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}
