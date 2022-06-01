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
// error should indicate that required field are missing
user, err := resource.NewBuilder().
  User().
    Title("Best Employee").
    Build()
```

The builder is sometimes seen as a convoluted way to construct
objects, but it has one major advantage, which is that it significantly
decreases the cognitive load for the code writer / reader to
recognize if a field is required or optional, as well as detecting
when an object is not properly initialized (as described in the previous
section). This is especially true for a resource like "User",
which has over 20 possible fields that can be populated. If this were
an open struct, you have to read up on it and manually check which
fields are required (and with what value must they be initialized with),
but with a builder all you have to is to call `Build()`, and any
mistakes you have made will be available in the error message.

Each resource type has its own builder type that specifically handles
initialization of the resource, but thes builders for individual resource
objects are accessible from a centralized `resource.Builder` object for convenience. 

You can access builders for other types by calling method names that match
their type names (e.g. `User()`, `Group()`, `Meta()`, etc)

```go
var b resource.Builder

user, err := b.User().
  // methods to initialize resource.User
  Build()

group, err := b.Group().
  // ditto
  Build()
```

The bulk of the code that implements the resources are generated,
in order to promote consistency and also to avoid any bugs caused
by fat-fingering.

The resource objects are generated using a custom YAML based DSL
for the time being. Direct use of the SCIM schema to derive the code
was briefly considered, but the schema lacks crucial Go-specific
hints that are required when generating code, and thus a custom
format was chosen. There is nothing barring this module from using
other formats if any, but the current implementation was chosen
because it was already available.


## Client

The client code follows a [Google Cloud Client style API](https://cloud.google.com/go).

The client is first initialized with the common configuration parameters
such as the base URL for the SCIM server, then you will need to obtain
the "service" objects that logically group API calls.

Each service object can provide "call" objects. The call objects are effectively
request builders, which also does validation of the request -- therefore we can
can invalid and/or incomplete requests before sending them to the client.

```go
// The client object
scimClient := client.New(baseURL)

// The service object
service := scimClient.User()

// The call object (notice `id` is required, so it must be passed to the constructor)
call := service.GetUser(id)

// Populate optional fields
call.Attributes(...)

// Make the HTTP request, and parse the response
user, err := call.Do(ctx)
```

All of these can be chained into a single "line":

```go
user, err := client.New(baseURL).
  User().
    GetUser(id).
    Attributes(...).
    Do(ctx)
```

This architecture is slightly more complicated than, say, implementing a client
and that can directly handle all API calls, for example:

```go
user, err := client.GetUser(id, attrs, ....)
```

But much like the builders for the resources, this quickly becomes problematic when
you have 20+ fields that you must be able to set values to:

```go
user, err := client.CreateUser( lots, and, lots, and, lots, of, fields... )
```

Using this Google Cloud style convention makes it much easier to give users
full control over all possible fields that are sent to the server, while
reducing the complexity of the method signature.

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
