scim
====

SCIM tools for Go (Work In Progress)

* [server](./server) - SCIM server
* [client](./client) - SCIM client
* [resource](./resource) - Definition of SCIM resource types
* [filter](./filter) - Basic SCIM filter parsing, with a simple SQL generator

# SYNOPSIS

<!-- INCLUDE(./examples/client_user_create_example_test.go) -->
```go
package examples_test

import (
  "context"
  "fmt"
  "testing"

  "github.com/cybozu-go/scim/client"
)

func TestClient_CreateUser(t *testing.T) {
  const baseURL = `https://scim.example.com`
  cl := client.New(baseURL)

  user, err := cl.User().CreateUser().
    DisplayName(`Daisuke Maki`).
    ExternalID(`lestrrat`).
    Do(context.TODO())
  if err != nil {
    fmt.Printf("failed to create user: %s", err)
  }

  _ = user
}
```
source: [./examples/client_user_create_example_test.go](https://github.com/cybozu-go/scim/blob/main/./examples/client_user_create_example_test.go)
<!-- END INCLUDE -->

# TODO

* Implement a sample server
* Implement check in CI to diff against generated code

# DESCRIPTION

The goal of this library is to implement most boilerplate code that is required to write a SCIM client/server.
This includes:

* A set of generic resources (Users, Groups) and ways to create and access them
* A generic client to make requests to the server
* A skeleton framework for a generic server

## Resources

The resources described in RFC7643 are covered. The resources are immutable opaque objects
instead of an open Go struct.

This allows us to reject any invalid or incomplete resource to be created
because the objects cannot be initialized directly.

For example, with an open struct the following would be allowed:

```go
// no required fields such as userName, but has other fields
user := &resource.User{
  title: "Best Employee",
}
```

Instead, we force the user to create an object using builder.
The following example will properly report errors when required
fields are not available. This assures as that the when we get
a hold of a resource, it is in fact a legal one (at least in
terms of object structure).

```go
user, err := resource.Builder().
  User().
    Title("Best Employee").
    Build()
```

TODO: More on implementation descisions

## Client

The client code follows a [Google Cloud Client style API](https://cloud.google.com/go).

TODO: Why, etc

## Server

The server code in this module is a stub implementation. You will need to provide
either the "backend" object that the stub implementation expects, or you will have to
provide the actual `http.Handler` object with your complete implementation.

The server code is not expected to be robust and reusable, because the details of
implementation will differ significantly depending on your own SCIM architecture.
For example, a SCIM server running on a single VPC or a fleet of containers will
most likely have different expectations and assumptions.

This implementation attempts to provide with the minimal building blocks, but
be aware that for any serious implementations you will have to implement your
own server anyways.

A toy implementation for the server will be provided, but only for testing purposes. 

TODO: more details
