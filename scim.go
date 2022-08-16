//go:generate ./tools/cmd/genoptions.sh
//go:generate go generate ./filter
//go:generate go generate ./resource
//go:generate go generate ./schema
//go:generate go generate ./client

// Package scim contains SCIM resource definitions, SCIM client implementation, and
// server-side shim for a simplistic SCIM service.
//
// # `resource` Directory
//
// Directory `resource` contains resource definitions, such as `resource.User`,
// and `resource.Group`.
//
// The resource objects are considered immutable once
// initialized. Use the corresponding builder objects to create new instances.
//
// Each resource object is implemented such that when/if the payload from the
// wire contains extra fields, those fields can safely be unmarshaled as well,
// allowing you to define extra fields, extra sub-resource types, extensions
// to existing SCIM resources via schema-qualified fields, etc.
//
// # `client` Directory
//
// Directory `client` contains a SCIM client.
//
// The structure is loosely similar to Google Cloud Go SDK. At the top level
// there is a `Client` object. The client has `Service` objects that group
// individual `Call` objects.
//
// The `Service` object is only a logical grouping of `Call` objects.
//
// Each `Call` object encapsulates an operation to an SCIM operation endpoint.
// `Call` objects are created by executing methods on the parent `Service`
// object. These method may require parameters: these exist because these
// parameters are required fields in the SCIM operation.
//
// All other optional parameters are passed via method calls on the `Call`
// object, in a builder-pattern.
//
// When the `Call` object is ready, invoke the `Do` method, which will
// issue an HTTP request to the endpoint, do the necessary conversions and
// return the response or an error.
//
// # Customizing the HTTP request
//
// When/if you want to control the behavior of your HTTP client, you can pass
// your customized HTTP client to the `client.New` contstructor:
//
//    customClient := ... // *http.Client object, or anything that satisfies client.HTTPClient
//    cl := client.New(url, client.WithClient(customClient))
package scim
