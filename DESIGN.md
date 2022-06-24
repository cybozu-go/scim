## Resources

The resources described in RFC7643 are covered. The resources are immutable opaque objects
instead of an open Go struct.

This allows us to reject any invalid or incomplete resource to be created
because the objects cannot be initialized directly.

For example, with an open struct the following would be allowed:

```go
// no required fields such as userName, but has other fields
user := &resource.User{
  Title: "Best Employee",
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

### Extensions

All resources will be able to mimic adding extensions.

SCIM extensions are represented as sort of a "fully qualified resource name".
For example, an "enterprise user" extension can be embedded in the
core User resource as follows:

```json
{
  "id": ...,
  "userName": ....,
  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User": {
    "costCenter": ...,
    "department": ...,
    "division": ...,
    "employeeNumber": ...,
    "manager": ...,
    "organization": ...
  }
}
```

To create a User resource with the enterprise user extension,
You simply need to use the `Extension()` method. 

<!-- INCLUDE(./examples/resource_extension_example_test.go) -->
```go
package examples_test

import (
  "encoding/json"
  "fmt"
  "os"

  "github.com/cybozu-go/scim/resource"
)

func Example_ResourceExtension() {
  var b resource.Builder

  user, err := b.User().
    ID("foo").
    UserName("foo").
    Extension(
      resource.EnterpriseUserSchemaURI,
      b.EnterpriseUser().
        CostCenter("foo").
        Department("foo").
        Division("foo").
        EmployeeNumber("foo").
        Organization("foo").
        MustBuild(),
    ).
    Build()
  if err != nil {
    fmt.Printf("%s", err)
    return
  }

  json.NewEncoder(os.Stdout).Encode(user)
  // OUTPUT:
  // {"id":"foo","schemas":["urn:ietf:params:scim:schemas:core:2.0:User","urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"],"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User":{"costCenter":"foo","department":"foo","division":"foo","employeeNumber":"foo","organization":"foo"},"userName":"foo"}
}
```
source: [./examples/resource_extension_example_test.go](https://github.com/cybozy-go/scim/blob/main/./examples/resource_extension_example_test.go)
<!-- END INCLUDE -->

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
user, err := client.User().Create( lots, and, lots, and, lots, of, fields... )
```

Using this Google Cloud style convention makes it much easier to give users
full control over all possible fields that are sent to the server, while
reducing the complexity of the method signature.

The Call object is most of the times a thin wrapper around the builders for
the resources they are creating, and therefore implements the same API as
their respective builders. For example, for Call objects that expect to
create a `User` object, it gives you everything available for `resource.UserBuilder`

## Server

The server code in this module is a stub implementation. 

The server code is not expected to be robust and reusable. You may be able to reuse
parts of the server, especially the routing parts and the request validation part,
but in general you should NOT expect the code in module to be able
to handle production loads or to simply be able to plugin to your architecture.

This is because the details of implementation will differ significantly depending
on your own SCIM architecture. For example, a SCIM server running on a single
VPC or a fleet of containers will most likely have different expectations and assumptions.
The same goes for a scenario where data is stored in a few sets of on-premise
databases vs something that is using an API-based microservices-ish architecture.

Instead what this module aims to provide in terms of a server is the
basic API structure and basic behavior that a simple SCIM server is
expected to implement, so that it can be used as a reference implementation of
sorts.

The server architecture will be as pluggable as possible, providing some
minimal building blocks for a SCIM server, but its main focus will be on
the ease of testing, not anything aimed for prduction use.

For any serious implementations you will have to implement your own server anyways.
