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

type MetaService struct {
	client *Client
}

// Meta creates a new Service object to perform an operation
func (client *Client) Meta() *MetaService {
	return &MetaService{
		client: client,
	}
}

// GetResourceTypesCall is an encapsulation of a SCIM operation.
type GetResourceTypesCall struct {
	err    error
	client *Client
	trace  io.Writer
}

func (svc *MetaService) GetResourceTypes() *GetResourceTypesCall {
	return &GetResourceTypesCall{
		client: svc.client,
	}
}

func (call *GetResourceTypesCall) Trace(w io.Writer) *GetResourceTypesCall {
	call.trace = w
	return call
}

func (call *GetResourceTypesCall) makeURL() string {
	return call.client.baseURL + "/ResourceTypes"
}

func (call *GetResourceTypesCall) Do(ctx context.Context) (*[]resource.ResourceType, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	trace := call.trace
	if trace == nil {
		trace = call.client.trace
	}
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, "trace: client sending call request to %q\n", u)
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

	var respayload []resource.ResourceType
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// GetServiceProviderConfigCall is an encapsulation of a SCIM operation.
type GetServiceProviderConfigCall struct {
	err    error
	client *Client
	trace  io.Writer
}

func (svc *MetaService) GetServiceProviderConfig() *GetServiceProviderConfigCall {
	return &GetServiceProviderConfigCall{
		client: svc.client,
	}
}

func (call *GetServiceProviderConfigCall) Trace(w io.Writer) *GetServiceProviderConfigCall {
	call.trace = w
	return call
}

func (call *GetServiceProviderConfigCall) makeURL() string {
	return call.client.baseURL + "/ServiceProviderConfig"
}

func (call *GetServiceProviderConfigCall) Do(ctx context.Context) (*resource.ServiceProviderConfig, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	trace := call.trace
	if trace == nil {
		trace = call.client.trace
	}
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, "trace: client sending call request to %q\n", u)
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

	var respayload resource.ServiceProviderConfig
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// GetSchemas is an encapsulation of a SCIM operation.
type GetSchemas struct {
	err    error
	client *Client
	trace  io.Writer
}

func (svc *MetaService) GetSchemas() *GetSchemas {
	return &GetSchemas{
		client: svc.client,
	}
}

func (call *GetSchemas) Trace(w io.Writer) *GetSchemas {
	call.trace = w
	return call
}

func (call *GetSchemas) makeURL() string {
	return call.client.baseURL + "/Schemas"
}

func (call *GetSchemas) Do(ctx context.Context) (*resource.ListResponse, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	trace := call.trace
	if trace == nil {
		trace = call.client.trace
	}
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, "trace: client sending call request to %q\n", u)
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

	var respayload resource.ListResponse
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}

// GetSchema is an encapsulation of a SCIM operation.
type GetSchema struct {
	err    error
	client *Client
	trace  io.Writer
	id     string
}

func (svc *MetaService) GetSchema(id string) *GetSchema {
	return &GetSchema{
		client: svc.client,
		id:     id,
	}
}

func (call *GetSchema) Trace(w io.Writer) *GetSchema {
	call.trace = w
	return call
}

func (call GetSchema) makeURL() string {
	return call.client.baseURL + "/Schemas/" + call.id
}

func (call *GetSchema) Do(ctx context.Context) (*resource.Schema, error) {
	if err := call.err; err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	trace := call.trace
	if trace == nil {
		trace = call.client.trace
	}
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, "trace: client sending call request to %q\n", u)
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

	var respayload resource.Schema
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}
