package client

import (
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

func (client *Client) Meta() *MetaService {
	return &MetaService{
		client: client,
	}
}

type GetServiceProviderConfigCall struct {
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

	trace := call.trace
	u := call.makeURL()
	if trace != nil {
		fmt.Fprintf(trace, `trace: client sending call request to %q`, u)
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

	var respayload resource.ServiceProviderConfig
	if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {
		return nil, fmt.Errorf(`failed to decode call response: %w`, err)
	}

	return &respayload, nil
}
