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

func (client *Client) Group() *GroupService {
	return &GroupService{
		client: client,
	}
}

type GetGroupCall struct {
	builder *resource.PartialResourceRepresentationRequestBuilder
	client  *Client
	trace   io.Writer
	id      string
}

func (svc *GroupService) GetGroup(id string) *GetGroupCall {
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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for GetGroupCall: %w`, err)
	}

	trace := call.trace
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

	var respayload resource.Group
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type CreateGroupCall struct {
	builder *resource.GroupBuilder
	client  *Client
	trace   io.Writer
}

func (svc *GroupService) CreateGroup() *CreateGroupCall {
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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for CreateGroupCall: %w`, err)
	}

	trace := call.trace
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

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(`expected call response %d, got (%d)`, http.StatusCreated, res.StatusCode)
	}

	var respayload resource.Group
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type ReplaceGroupCall struct {
	builder *resource.GroupBuilder
	client  *Client
	trace   io.Writer
	id      string
}

func (svc *GroupService) ReplaceGroup(id string) *ReplaceGroupCall {
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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for ReplaceGroupCall: %w`, err)
	}

	trace := call.trace
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

	var respayload resource.Group
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type DeleteGroupCall struct {
	builder *resource.GroupBuilder
	client  *Client
	trace   io.Writer
	id      string
}

func (svc *GroupService) DeleteGroup(id string) *DeleteGroupCall {
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
	payload, err := call.builder.Build()
	if err != nil {
		return fmt.Errorf(`failed to generate request payload for DeleteGroupCall: %w`, err)
	}

	trace := call.trace
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
	if trace != nil {
		buf, _ := httputil.DumpResponse(res, true)
		fmt.Fprintf(trace, "%s\n", buf)
	}
	if err != nil {
		return fmt.Errorf(`failed to send request to %q: %w`, u, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf(`expected call response %d, got (%d)`, http.StatusNoContent, res.StatusCode)
	}

	return nil
}

type SearchGroupCall struct {
	builder *resource.SearchRequestBuilder
	client  *Client
	trace   io.Writer
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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for SearchGroupCall: %w`, err)
	}

	trace := call.trace
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
