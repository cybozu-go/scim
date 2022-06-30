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

// UserService the logical grouping of SCIM user related API calls
type UserService struct {
	client *Client
}

func (client *Client) User() *UserService {
	return &UserService{
		client: client,
	}
}

type GetUserCall struct {
	builder *resource.PartialResourceRepresentationRequestBuilder
	object  *resource.PartialResourceRepresentationRequest
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *GetUserCall) payload() (*resource.PartialResourceRepresentationRequest, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *GetUserCall) FromJSON(data []byte) *GetUserCall {
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

// Get creates an instance of GetUserCall that sends an HTTP GET request to
// /Users to retrieve the user associated with the specified ID.
func (svc *UserService) Get(id string) *GetUserCall {
	return &GetUserCall{
		builder: resource.NewPartialResourceRepresentationRequestBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *GetUserCall) Attributes(v ...string) *GetUserCall {
	call.builder.Attributes(v...)
	return call
}

func (call *GetUserCall) ExcludedAttributes(v ...string) *GetUserCall {
	call.builder.ExcludedAttributes(v...)
	return call
}

func (call *GetUserCall) Trace(w io.Writer) *GetUserCall {
	call.trace = w
	return call
}

func (call GetUserCall) makeURL() string {
	return call.client.baseURL + "/Users/" + call.id
}

func (call *GetUserCall) Do(ctx context.Context) (*resource.User, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for GetUserCall: %w`, err)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type CreateUserCall struct {
	builder *resource.UserBuilder
	object  *resource.User
	err     error
	client  *Client
	trace   io.Writer
}

func (call *CreateUserCall) payload() (*resource.User, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *CreateUserCall) FromJSON(data []byte) *CreateUserCall {
	if call.err != nil {
		return call
	}
	var in resource.User
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

// Create creates an insance of CreateUserCall that sends an HTTP POST request to
// /Users to create a new user.
func (svc *UserService) Create() *CreateUserCall {
	return &CreateUserCall{
		builder: resource.NewUserBuilder(),
		client:  svc.client,
	}
}

func (call *CreateUserCall) Active(v bool) *CreateUserCall {
	call.builder.Active(v)
	return call
}

func (call *CreateUserCall) Addresses(v ...*resource.Address) *CreateUserCall {
	call.builder.Addresses(v...)
	return call
}

func (call *CreateUserCall) DisplayName(v string) *CreateUserCall {
	call.builder.DisplayName(v)
	return call
}

func (call *CreateUserCall) Emails(v ...*resource.Email) *CreateUserCall {
	call.builder.Emails(v...)
	return call
}

func (call *CreateUserCall) Entitlements(v ...*resource.Entitlement) *CreateUserCall {
	call.builder.Entitlements(v...)
	return call
}

func (call *CreateUserCall) ExternalID(v string) *CreateUserCall {
	call.builder.ExternalID(v)
	return call
}

func (call *CreateUserCall) IMS(v ...*resource.IMS) *CreateUserCall {
	call.builder.IMS(v...)
	return call
}

func (call *CreateUserCall) Locale(v string) *CreateUserCall {
	call.builder.Locale(v)
	return call
}

func (call *CreateUserCall) Name(v *resource.Names) *CreateUserCall {
	call.builder.Name(v)
	return call
}

func (call *CreateUserCall) NickName(v string) *CreateUserCall {
	call.builder.NickName(v)
	return call
}

func (call *CreateUserCall) Password(v string) *CreateUserCall {
	call.builder.Password(v)
	return call
}

func (call *CreateUserCall) PhoneNumbers(v ...*resource.PhoneNumber) *CreateUserCall {
	call.builder.PhoneNumbers(v...)
	return call
}

func (call *CreateUserCall) Photos(v ...*resource.Photo) *CreateUserCall {
	call.builder.Photos(v...)
	return call
}

func (call *CreateUserCall) PreferredLanguage(v string) *CreateUserCall {
	call.builder.PreferredLanguage(v)
	return call
}

func (call *CreateUserCall) ProfileURL(v string) *CreateUserCall {
	call.builder.ProfileURL(v)
	return call
}

func (call *CreateUserCall) Roles(v ...*resource.Role) *CreateUserCall {
	call.builder.Roles(v...)
	return call
}

func (call *CreateUserCall) Timezone(v string) *CreateUserCall {
	call.builder.Timezone(v)
	return call
}

func (call *CreateUserCall) Title(v string) *CreateUserCall {
	call.builder.Title(v)
	return call
}

func (call *CreateUserCall) UserName(v string) *CreateUserCall {
	call.builder.UserName(v)
	return call
}

func (call *CreateUserCall) UserType(v string) *CreateUserCall {
	call.builder.UserType(v)
	return call
}

func (call *CreateUserCall) X509Certificates(v ...*resource.X509Certificate) *CreateUserCall {
	call.builder.X509Certificates(v...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *CreateUserCall) Extension(uri string, value interface{}) *CreateUserCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *CreateUserCall) Validator(v resource.UserValidator) *CreateUserCall {
	call.builder.Validator(v)
	return call
}

func (call *CreateUserCall) Trace(w io.Writer) *CreateUserCall {
	call.trace = w
	return call
}

func (call *CreateUserCall) makeURL() string {
	return call.client.baseURL + "/Users"
}

func (call *CreateUserCall) Do(ctx context.Context) (*resource.User, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for CreateUserCall: %w`, err)
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

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(`expected call response %d, got (%d)`, http.StatusCreated, res.StatusCode)
	}

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type ReplaceUserCall struct {
	builder *resource.UserBuilder
	object  *resource.User
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *ReplaceUserCall) payload() (*resource.User, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *ReplaceUserCall) FromJSON(data []byte) *ReplaceUserCall {
	if call.err != nil {
		return call
	}
	var in resource.User
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

// Replace creates an insance of ReplaceUserCall that sends an HTTP PUT request to
// /Users to replace an existing new user.
func (svc *UserService) Replace(id string) *ReplaceUserCall {
	return &ReplaceUserCall{
		builder: resource.NewUserBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *ReplaceUserCall) Active(v bool) *ReplaceUserCall {
	call.builder.Active(v)
	return call
}

func (call *ReplaceUserCall) Addresses(v ...*resource.Address) *ReplaceUserCall {
	call.builder.Addresses(v...)
	return call
}

func (call *ReplaceUserCall) DisplayName(v string) *ReplaceUserCall {
	call.builder.DisplayName(v)
	return call
}

func (call *ReplaceUserCall) Emails(v ...*resource.Email) *ReplaceUserCall {
	call.builder.Emails(v...)
	return call
}

func (call *ReplaceUserCall) Entitlements(v ...*resource.Entitlement) *ReplaceUserCall {
	call.builder.Entitlements(v...)
	return call
}

func (call *ReplaceUserCall) ExternalID(v string) *ReplaceUserCall {
	call.builder.ExternalID(v)
	return call
}

func (call *ReplaceUserCall) IMS(v ...*resource.IMS) *ReplaceUserCall {
	call.builder.IMS(v...)
	return call
}

func (call *ReplaceUserCall) Locale(v string) *ReplaceUserCall {
	call.builder.Locale(v)
	return call
}

func (call *ReplaceUserCall) Name(v *resource.Names) *ReplaceUserCall {
	call.builder.Name(v)
	return call
}

func (call *ReplaceUserCall) NickName(v string) *ReplaceUserCall {
	call.builder.NickName(v)
	return call
}

func (call *ReplaceUserCall) Password(v string) *ReplaceUserCall {
	call.builder.Password(v)
	return call
}

func (call *ReplaceUserCall) PhoneNumbers(v ...*resource.PhoneNumber) *ReplaceUserCall {
	call.builder.PhoneNumbers(v...)
	return call
}

func (call *ReplaceUserCall) Photos(v ...*resource.Photo) *ReplaceUserCall {
	call.builder.Photos(v...)
	return call
}

func (call *ReplaceUserCall) PreferredLanguage(v string) *ReplaceUserCall {
	call.builder.PreferredLanguage(v)
	return call
}

func (call *ReplaceUserCall) ProfileURL(v string) *ReplaceUserCall {
	call.builder.ProfileURL(v)
	return call
}

func (call *ReplaceUserCall) Roles(v ...*resource.Role) *ReplaceUserCall {
	call.builder.Roles(v...)
	return call
}

func (call *ReplaceUserCall) Timezone(v string) *ReplaceUserCall {
	call.builder.Timezone(v)
	return call
}

func (call *ReplaceUserCall) Title(v string) *ReplaceUserCall {
	call.builder.Title(v)
	return call
}

func (call *ReplaceUserCall) UserName(v string) *ReplaceUserCall {
	call.builder.UserName(v)
	return call
}

func (call *ReplaceUserCall) UserType(v string) *ReplaceUserCall {
	call.builder.UserType(v)
	return call
}

func (call *ReplaceUserCall) X509Certificates(v ...*resource.X509Certificate) *ReplaceUserCall {
	call.builder.X509Certificates(v...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *ReplaceUserCall) Extension(uri string, value interface{}) *ReplaceUserCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *ReplaceUserCall) Validator(v resource.UserValidator) *ReplaceUserCall {
	call.builder.Validator(v)
	return call
}

func (call *ReplaceUserCall) Trace(w io.Writer) *ReplaceUserCall {
	call.trace = w
	return call
}

func (call ReplaceUserCall) makeURL() string {
	return call.client.baseURL + "/Users/" + call.id
}

func (call *ReplaceUserCall) Do(ctx context.Context) (*resource.User, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for ReplaceUserCall: %w`, err)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type DeleteUserCall struct {
	builder *resource.UserBuilder
	object  *resource.User
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *DeleteUserCall) payload() (*resource.User, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *DeleteUserCall) FromJSON(data []byte) *DeleteUserCall {
	if call.err != nil {
		return call
	}
	var in resource.User
	if err := json.Unmarshal(data, &in); err != nil {
		call.err = fmt.Errorf("failed to decode data: %w", err)
		return call
	}
	call.object = &in
	return call
}

func (svc *UserService) Delete(id string) *DeleteUserCall {
	return &DeleteUserCall{
		builder: resource.NewUserBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *DeleteUserCall) Active(v bool) *DeleteUserCall {
	call.builder.Active(v)
	return call
}

func (call *DeleteUserCall) Addresses(v ...*resource.Address) *DeleteUserCall {
	call.builder.Addresses(v...)
	return call
}

func (call *DeleteUserCall) DisplayName(v string) *DeleteUserCall {
	call.builder.DisplayName(v)
	return call
}

func (call *DeleteUserCall) Emails(v ...*resource.Email) *DeleteUserCall {
	call.builder.Emails(v...)
	return call
}

func (call *DeleteUserCall) Entitlements(v ...*resource.Entitlement) *DeleteUserCall {
	call.builder.Entitlements(v...)
	return call
}

func (call *DeleteUserCall) ExternalID(v string) *DeleteUserCall {
	call.builder.ExternalID(v)
	return call
}

func (call *DeleteUserCall) Groups(v ...*resource.GroupMember) *DeleteUserCall {
	call.builder.Groups(v...)
	return call
}

func (call *DeleteUserCall) ID(v string) *DeleteUserCall {
	call.builder.ID(v)
	return call
}

func (call *DeleteUserCall) IMS(v ...*resource.IMS) *DeleteUserCall {
	call.builder.IMS(v...)
	return call
}

func (call *DeleteUserCall) Locale(v string) *DeleteUserCall {
	call.builder.Locale(v)
	return call
}

func (call *DeleteUserCall) Meta(v *resource.Meta) *DeleteUserCall {
	call.builder.Meta(v)
	return call
}

func (call *DeleteUserCall) Name(v *resource.Names) *DeleteUserCall {
	call.builder.Name(v)
	return call
}

func (call *DeleteUserCall) NickName(v string) *DeleteUserCall {
	call.builder.NickName(v)
	return call
}

func (call *DeleteUserCall) Password(v string) *DeleteUserCall {
	call.builder.Password(v)
	return call
}

func (call *DeleteUserCall) PhoneNumbers(v ...*resource.PhoneNumber) *DeleteUserCall {
	call.builder.PhoneNumbers(v...)
	return call
}

func (call *DeleteUserCall) Photos(v ...*resource.Photo) *DeleteUserCall {
	call.builder.Photos(v...)
	return call
}

func (call *DeleteUserCall) PreferredLanguage(v string) *DeleteUserCall {
	call.builder.PreferredLanguage(v)
	return call
}

func (call *DeleteUserCall) ProfileURL(v string) *DeleteUserCall {
	call.builder.ProfileURL(v)
	return call
}

func (call *DeleteUserCall) Roles(v ...*resource.Role) *DeleteUserCall {
	call.builder.Roles(v...)
	return call
}

func (call *DeleteUserCall) Timezone(v string) *DeleteUserCall {
	call.builder.Timezone(v)
	return call
}

func (call *DeleteUserCall) Title(v string) *DeleteUserCall {
	call.builder.Title(v)
	return call
}

func (call *DeleteUserCall) UserName(v string) *DeleteUserCall {
	call.builder.UserName(v)
	return call
}

func (call *DeleteUserCall) UserType(v string) *DeleteUserCall {
	call.builder.UserType(v)
	return call
}

func (call *DeleteUserCall) X509Certificates(v ...*resource.X509Certificate) *DeleteUserCall {
	call.builder.X509Certificates(v...)
	return call
}

func (call *DeleteUserCall) Trace(w io.Writer) *DeleteUserCall {
	call.trace = w
	return call
}

func (call DeleteUserCall) makeURL() string {
	return call.client.baseURL + "/Users/" + call.id
}

func (call *DeleteUserCall) Do(ctx context.Context) error {
	if err := call.err; err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return fmt.Errorf(`failed to generate request payload for DeleteUserCall: %w`, err)
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

type SearchUserCall struct {
	builder *resource.SearchRequestBuilder
	object  *resource.SearchRequest
	err     error
	client  *Client
	trace   io.Writer
}

func (call *SearchUserCall) payload() (*resource.SearchRequest, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *SearchUserCall) FromJSON(data []byte) *SearchUserCall {
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

func (svc *UserService) Search() *SearchUserCall {
	return &SearchUserCall{
		builder: resource.NewSearchRequestBuilder(),
		client:  svc.client,
	}
}

func (call *SearchUserCall) Attributes(v ...string) *SearchUserCall {
	call.builder.Attributes(v...)
	return call
}

func (call *SearchUserCall) Count(v int) *SearchUserCall {
	call.builder.Count(v)
	return call
}

func (call *SearchUserCall) ExludedAttributes(v ...string) *SearchUserCall {
	call.builder.ExludedAttributes(v...)
	return call
}

func (call *SearchUserCall) Filter(v string) *SearchUserCall {
	call.builder.Filter(v)
	return call
}

func (call *SearchUserCall) SortBy(v string) *SearchUserCall {
	call.builder.SortBy(v)
	return call
}

func (call *SearchUserCall) SortOrder(v string) *SearchUserCall {
	call.builder.SortOrder(v)
	return call
}

func (call *SearchUserCall) StartIndex(v int) *SearchUserCall {
	call.builder.StartIndex(v)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *SearchUserCall) Extension(uri string, value interface{}) *SearchUserCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *SearchUserCall) Validator(v resource.SearchRequestValidator) *SearchUserCall {
	call.builder.Validator(v)
	return call
}

func (call *SearchUserCall) Trace(w io.Writer) *SearchUserCall {
	call.trace = w
	return call
}

func (call *SearchUserCall) makeURL() string {
	return call.client.baseURL + "/Users/.search"
}

func (call *SearchUserCall) Do(ctx context.Context) (*resource.ListResponse, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for SearchUserCall: %w`, err)
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
