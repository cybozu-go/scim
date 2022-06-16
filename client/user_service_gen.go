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
	client  *Client
	trace   io.Writer
	id      string
}

func (svc *UserService) GetUser(id string) *GetUserCall {
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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for GetUserCall: %w`, err)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type CreateUserCall struct {
	builder *resource.UserBuilder
	client  *Client
	trace   io.Writer
}

func (svc *UserService) CreateUser() *CreateUserCall {
	return &CreateUserCall{
		builder: resource.NewUserBuilder(),
		client:  svc.client,
	}
}

func (call *CreateUserCall) Active(v bool) *CreateUserCall {
	call.builder.Active(v)
	return call
}

func (call *CreateUserCall) Addresses(v ...string) *CreateUserCall {
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

func (call *CreateUserCall) Entitlements(v ...string) *CreateUserCall {
	call.builder.Entitlements(v...)
	return call
}

func (call *CreateUserCall) ExternalID(v string) *CreateUserCall {
	call.builder.ExternalID(v)
	return call
}

func (call *CreateUserCall) IMS(v ...string) *CreateUserCall {
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

func (call *CreateUserCall) PhoneNumbers(v ...string) *CreateUserCall {
	call.builder.PhoneNumbers(v...)
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

func (call *CreateUserCall) Roles(v ...string) *CreateUserCall {
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

func (call *CreateUserCall) X509Certificates(v ...string) *CreateUserCall {
	call.builder.X509Certificates(v...)
	return call
}

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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for CreateUserCall: %w`, err)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type ReplaceUserCall struct {
	builder *resource.UserBuilder
	client  *Client
	trace   io.Writer
	id      string
}

func (svc *UserService) ReplaceUser(id string) *ReplaceUserCall {
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

func (call *ReplaceUserCall) Addresses(v ...string) *ReplaceUserCall {
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

func (call *ReplaceUserCall) Entitlements(v ...string) *ReplaceUserCall {
	call.builder.Entitlements(v...)
	return call
}

func (call *ReplaceUserCall) ExternalID(v string) *ReplaceUserCall {
	call.builder.ExternalID(v)
	return call
}

func (call *ReplaceUserCall) IMS(v ...string) *ReplaceUserCall {
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

func (call *ReplaceUserCall) PhoneNumbers(v ...string) *ReplaceUserCall {
	call.builder.PhoneNumbers(v...)
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

func (call *ReplaceUserCall) Roles(v ...string) *ReplaceUserCall {
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

func (call *ReplaceUserCall) X509Certificates(v ...string) *ReplaceUserCall {
	call.builder.X509Certificates(v...)
	return call
}

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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for ReplaceUserCall: %w`, err)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

type DeleteUserCall struct {
	builder *resource.UserBuilder
	client  *Client
	trace   io.Writer
	id      string
}

func (svc *UserService) DeleteUser(id string) *DeleteUserCall {
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

func (call *DeleteUserCall) Addresses(v ...string) *DeleteUserCall {
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

func (call *DeleteUserCall) Entitlements(v ...string) *DeleteUserCall {
	call.builder.Entitlements(v...)
	return call
}

func (call *DeleteUserCall) ExternalID(v string) *DeleteUserCall {
	call.builder.ExternalID(v)
	return call
}

func (call *DeleteUserCall) Groups(v ...string) *DeleteUserCall {
	call.builder.Groups(v...)
	return call
}

func (call *DeleteUserCall) ID(v string) *DeleteUserCall {
	call.builder.ID(v)
	return call
}

func (call *DeleteUserCall) IMS(v ...string) *DeleteUserCall {
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

func (call *DeleteUserCall) PhoneNumbers(v ...string) *DeleteUserCall {
	call.builder.PhoneNumbers(v...)
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

func (call *DeleteUserCall) Roles(v ...string) *DeleteUserCall {
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

func (call *DeleteUserCall) X509Certificates(v ...string) *DeleteUserCall {
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
	payload, err := call.builder.Build()
	if err != nil {
		return fmt.Errorf(`failed to generate request payload for DeleteUserCall: %w`, err)
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

type SearchUserCall struct {
	builder *resource.SearchRequestBuilder
	client  *Client
	trace   io.Writer
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
	payload, err := call.builder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for SearchUserCall: %w`, err)
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
