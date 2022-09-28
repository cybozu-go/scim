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

// User creates a new Service object to perform an operation
func (client *Client) User() *UserService {
	return &UserService{
		client: client,
	}
}

// GetUserCall is an encapsulation of a SCIM operation.
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

func (call *GetUserCall) Attributes(in ...string) *GetUserCall {
	call.builder.Attributes(in...)
	return call
}

func (call *GetUserCall) ExcludedAttributes(in ...string) *GetUserCall {
	call.builder.ExcludedAttributes(in...)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// CreateUserCall is an encapsulation of a SCIM operation.
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

func (call *CreateUserCall) Active(in bool) *CreateUserCall {
	call.builder.Active(in)
	return call
}

func (call *CreateUserCall) Addresses(in ...*resource.Address) *CreateUserCall {
	call.builder.Addresses(in...)
	return call
}

func (call *CreateUserCall) DisplayName(in string) *CreateUserCall {
	call.builder.DisplayName(in)
	return call
}

func (call *CreateUserCall) Emails(in ...*resource.Email) *CreateUserCall {
	call.builder.Emails(in...)
	return call
}

func (call *CreateUserCall) Entitlements(in ...*resource.Entitlement) *CreateUserCall {
	call.builder.Entitlements(in...)
	return call
}

func (call *CreateUserCall) ExternalID(in string) *CreateUserCall {
	call.builder.ExternalID(in)
	return call
}

func (call *CreateUserCall) IMS(in ...*resource.IMS) *CreateUserCall {
	call.builder.IMS(in...)
	return call
}

func (call *CreateUserCall) Locale(in string) *CreateUserCall {
	call.builder.Locale(in)
	return call
}

func (call *CreateUserCall) Name(in *resource.Names) *CreateUserCall {
	call.builder.Name(in)
	return call
}

func (call *CreateUserCall) NickName(in string) *CreateUserCall {
	call.builder.NickName(in)
	return call
}

func (call *CreateUserCall) Password(in string) *CreateUserCall {
	call.builder.Password(in)
	return call
}

func (call *CreateUserCall) PhoneNumbers(in ...*resource.PhoneNumber) *CreateUserCall {
	call.builder.PhoneNumbers(in...)
	return call
}

func (call *CreateUserCall) Photos(in ...*resource.Photo) *CreateUserCall {
	call.builder.Photos(in...)
	return call
}

func (call *CreateUserCall) PreferredLanguage(in string) *CreateUserCall {
	call.builder.PreferredLanguage(in)
	return call
}

func (call *CreateUserCall) ProfileURL(in string) *CreateUserCall {
	call.builder.ProfileURL(in)
	return call
}

func (call *CreateUserCall) Roles(in ...*resource.Role) *CreateUserCall {
	call.builder.Roles(in...)
	return call
}

func (call *CreateUserCall) Timezone(in string) *CreateUserCall {
	call.builder.Timezone(in)
	return call
}

func (call *CreateUserCall) Title(in string) *CreateUserCall {
	call.builder.Title(in)
	return call
}

func (call *CreateUserCall) UserName(in string) *CreateUserCall {
	call.builder.UserName(in)
	return call
}

func (call *CreateUserCall) UserType(in string) *CreateUserCall {
	call.builder.UserType(in)
	return call
}

func (call *CreateUserCall) X509Certificates(in ...*resource.X509Certificate) *CreateUserCall {
	call.builder.X509Certificates(in...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *CreateUserCall) Extension(uri string, value interface{}) *CreateUserCall {
	call.builder.Extension(uri, value)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// ReplaceUserCall is an encapsulation of a SCIM operation.
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

func (call *ReplaceUserCall) Active(in bool) *ReplaceUserCall {
	call.builder.Active(in)
	return call
}

func (call *ReplaceUserCall) Addresses(in ...*resource.Address) *ReplaceUserCall {
	call.builder.Addresses(in...)
	return call
}

func (call *ReplaceUserCall) DisplayName(in string) *ReplaceUserCall {
	call.builder.DisplayName(in)
	return call
}

func (call *ReplaceUserCall) Emails(in ...*resource.Email) *ReplaceUserCall {
	call.builder.Emails(in...)
	return call
}

func (call *ReplaceUserCall) Entitlements(in ...*resource.Entitlement) *ReplaceUserCall {
	call.builder.Entitlements(in...)
	return call
}

func (call *ReplaceUserCall) ExternalID(in string) *ReplaceUserCall {
	call.builder.ExternalID(in)
	return call
}

func (call *ReplaceUserCall) IMS(in ...*resource.IMS) *ReplaceUserCall {
	call.builder.IMS(in...)
	return call
}

func (call *ReplaceUserCall) Locale(in string) *ReplaceUserCall {
	call.builder.Locale(in)
	return call
}

func (call *ReplaceUserCall) Name(in *resource.Names) *ReplaceUserCall {
	call.builder.Name(in)
	return call
}

func (call *ReplaceUserCall) NickName(in string) *ReplaceUserCall {
	call.builder.NickName(in)
	return call
}

func (call *ReplaceUserCall) Password(in string) *ReplaceUserCall {
	call.builder.Password(in)
	return call
}

func (call *ReplaceUserCall) PhoneNumbers(in ...*resource.PhoneNumber) *ReplaceUserCall {
	call.builder.PhoneNumbers(in...)
	return call
}

func (call *ReplaceUserCall) Photos(in ...*resource.Photo) *ReplaceUserCall {
	call.builder.Photos(in...)
	return call
}

func (call *ReplaceUserCall) PreferredLanguage(in string) *ReplaceUserCall {
	call.builder.PreferredLanguage(in)
	return call
}

func (call *ReplaceUserCall) ProfileURL(in string) *ReplaceUserCall {
	call.builder.ProfileURL(in)
	return call
}

func (call *ReplaceUserCall) Roles(in ...*resource.Role) *ReplaceUserCall {
	call.builder.Roles(in...)
	return call
}

func (call *ReplaceUserCall) Timezone(in string) *ReplaceUserCall {
	call.builder.Timezone(in)
	return call
}

func (call *ReplaceUserCall) Title(in string) *ReplaceUserCall {
	call.builder.Title(in)
	return call
}

func (call *ReplaceUserCall) UserName(in string) *ReplaceUserCall {
	call.builder.UserName(in)
	return call
}

func (call *ReplaceUserCall) UserType(in string) *ReplaceUserCall {
	call.builder.UserType(in)
	return call
}

func (call *ReplaceUserCall) X509Certificates(in ...*resource.X509Certificate) *ReplaceUserCall {
	call.builder.X509Certificates(in...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *ReplaceUserCall) Extension(uri string, value interface{}) *ReplaceUserCall {
	call.builder.Extension(uri, value)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// PatchUserCall is an encapsulation of a SCIM operation.
type PatchUserCall struct {
	builder *resource.PatchRequestBuilder
	object  *resource.PatchRequest
	err     error
	client  *Client
	trace   io.Writer
	id      string
}

func (call *PatchUserCall) payload() (*resource.PatchRequest, error) {
	if object := call.object; object != nil {
		return object, nil
	}
	return call.builder.Build()
}

func (call *PatchUserCall) FromJSON(data []byte) *PatchUserCall {
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

// Patch allows the user to patch parts of the user object
func (svc *UserService) Patch(id string) *PatchUserCall {
	return &PatchUserCall{
		builder: resource.NewPatchRequestBuilder(),
		client:  svc.client,
		id:      id,
	}
}

func (call *PatchUserCall) Operations(in ...*resource.PatchOperation) *PatchUserCall {
	call.builder.Operations(in...)
	return call
}

func (call *PatchUserCall) Schemas(in ...string) *PatchUserCall {
	call.builder.Schemas(in...)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *PatchUserCall) Extension(uri string, value interface{}) *PatchUserCall {
	call.builder.Extension(uri, value)
	return call
}

func (call *PatchUserCall) Trace(w io.Writer) *PatchUserCall {
	call.trace = w
	return call
}

func (call PatchUserCall) makeURL() string {
	return call.client.baseURL + "/Users/" + call.id
}

func (call *PatchUserCall) Do(ctx context.Context) (*resource.User, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	payload, err := call.payload()
	if err != nil {
		return nil, fmt.Errorf(`failed to generate request payload for PatchUserCall: %w`, err)
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

	var respayload resource.User
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// DeleteUserCall is an encapsulation of a SCIM operation.
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

func (call *DeleteUserCall) Active(in bool) *DeleteUserCall {
	call.builder.Active(in)
	return call
}

func (call *DeleteUserCall) Addresses(in ...*resource.Address) *DeleteUserCall {
	call.builder.Addresses(in...)
	return call
}

func (call *DeleteUserCall) DisplayName(in string) *DeleteUserCall {
	call.builder.DisplayName(in)
	return call
}

func (call *DeleteUserCall) Emails(in ...*resource.Email) *DeleteUserCall {
	call.builder.Emails(in...)
	return call
}

func (call *DeleteUserCall) Entitlements(in ...*resource.Entitlement) *DeleteUserCall {
	call.builder.Entitlements(in...)
	return call
}

func (call *DeleteUserCall) ExternalID(in string) *DeleteUserCall {
	call.builder.ExternalID(in)
	return call
}

func (call *DeleteUserCall) Groups(in ...*resource.AssociatedGroup) *DeleteUserCall {
	call.builder.Groups(in...)
	return call
}

func (call *DeleteUserCall) ID(in string) *DeleteUserCall {
	call.builder.ID(in)
	return call
}

func (call *DeleteUserCall) IMS(in ...*resource.IMS) *DeleteUserCall {
	call.builder.IMS(in...)
	return call
}

func (call *DeleteUserCall) Locale(in string) *DeleteUserCall {
	call.builder.Locale(in)
	return call
}

func (call *DeleteUserCall) Meta(in *resource.Meta) *DeleteUserCall {
	call.builder.Meta(in)
	return call
}

func (call *DeleteUserCall) Name(in *resource.Names) *DeleteUserCall {
	call.builder.Name(in)
	return call
}

func (call *DeleteUserCall) NickName(in string) *DeleteUserCall {
	call.builder.NickName(in)
	return call
}

func (call *DeleteUserCall) Password(in string) *DeleteUserCall {
	call.builder.Password(in)
	return call
}

func (call *DeleteUserCall) PhoneNumbers(in ...*resource.PhoneNumber) *DeleteUserCall {
	call.builder.PhoneNumbers(in...)
	return call
}

func (call *DeleteUserCall) Photos(in ...*resource.Photo) *DeleteUserCall {
	call.builder.Photos(in...)
	return call
}

func (call *DeleteUserCall) PreferredLanguage(in string) *DeleteUserCall {
	call.builder.PreferredLanguage(in)
	return call
}

func (call *DeleteUserCall) ProfileURL(in string) *DeleteUserCall {
	call.builder.ProfileURL(in)
	return call
}

func (call *DeleteUserCall) Roles(in ...*resource.Role) *DeleteUserCall {
	call.builder.Roles(in...)
	return call
}

func (call *DeleteUserCall) Schemas(in ...string) *DeleteUserCall {
	call.builder.Schemas(in...)
	return call
}

func (call *DeleteUserCall) Timezone(in string) *DeleteUserCall {
	call.builder.Timezone(in)
	return call
}

func (call *DeleteUserCall) Title(in string) *DeleteUserCall {
	call.builder.Title(in)
	return call
}

func (call *DeleteUserCall) UserName(in string) *DeleteUserCall {
	call.builder.UserName(in)
	return call
}

func (call *DeleteUserCall) UserType(in string) *DeleteUserCall {
	call.builder.UserType(in)
	return call
}

func (call *DeleteUserCall) X509Certificates(in ...*resource.X509Certificate) *DeleteUserCall {
	call.builder.X509Certificates(in...)
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

// SearchUserCall is an encapsulation of a SCIM operation.
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

func (call *SearchUserCall) Attributes(in ...string) *SearchUserCall {
	call.builder.Attributes(in...)
	return call
}

func (call *SearchUserCall) Count(in int) *SearchUserCall {
	call.builder.Count(in)
	return call
}

func (call *SearchUserCall) ExcludedAttributes(in ...string) *SearchUserCall {
	call.builder.ExcludedAttributes(in...)
	return call
}

func (call *SearchUserCall) Filter(in string) *SearchUserCall {
	call.builder.Filter(in)
	return call
}

func (call *SearchUserCall) Schema(in string) *SearchUserCall {
	call.builder.Schema(in)
	return call
}

func (call *SearchUserCall) Schemas(in ...string) *SearchUserCall {
	call.builder.Schemas(in...)
	return call
}

func (call *SearchUserCall) SortBy(in string) *SearchUserCall {
	call.builder.SortBy(in)
	return call
}

func (call *SearchUserCall) SortOrder(in string) *SearchUserCall {
	call.builder.SortOrder(in)
	return call
}

func (call *SearchUserCall) StartIndex(in int) *SearchUserCall {
	call.builder.StartIndex(in)
	return call
}

// Extension allows users to register an extension using the fully qualified URI
func (call *SearchUserCall) Extension(uri string, value interface{}) *SearchUserCall {
	call.builder.Extension(uri, value)
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
