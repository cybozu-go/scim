scim
====
[![Go Reference](https://pkg.go.dev/badge/github.com/cybozu-go/scim.svg)](https://pkg.go.dev/github.com/cybozu-go/scim)

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
  "net/http/httptest"

  "github.com/cybozu-go/scim/server"
)

func ExampleClient_UserCreate() {
  srv := httptest.NewServer(server.MustNewServer(NewMockBackend()))
  defer srv.Close()

  // Create a SCIM client: this is wrapped in another function call
  // to accomodate for debug tracing, but you can replace it with
  // a client.New() call
  cl := NewClient(srv.URL)

  user, err := cl.User().Create().
    DisplayName(`Daisuke Maki`).
    ExternalID(`lestrrat`).
    Do(context.TODO())
  if err != nil {
    fmt.Printf("failed to create user: %s", err)
  }

  _ = user

  // OUTPUT:
  //
}
```
source: [./examples/client_user_create_example_test.go](https://github.com/cybozu-go/scim/blob/main/./examples/client_user_create_example_test.go)
<!-- END INCLUDE -->

# TODO

* Finish implementing sample server
  * Probably move it to github.com/cybozu-go/scim-ent
* Implement check in CI to diff against generated code
* Streamline code generation

# Code Generation

## go generate ./resource

Generates code in `resource`, `sample/ent/schema`, and `sample`. If you made any changes that result in
changes to code under `sample/ent/schema`, you need to run `go generate ./sample` as well.

Source file is [`tools/cmd/genresources/objects.yml`](./tools/cmd/genresources/object.yml)

## go generate ./sample

Generates code for [`ent`](https://entgo.io), based on files under `sample/ent/schema`.
This generates code from ent code, and thus have not source file.

## go generate ./client

Generates client code. Requires some synchronization between `resource`.

Source file is [`tools/cmd/genclient/calls.yml`](./tools/cmd/genclient/calls.yml)

## go generate ./filter

Generates filter parser code from `filter/parser.go.y` support objects.

Source file is [`tools/cmd/genfilter/objects.yml`](./tools/cmd/genfilter/object.yml)

## go generate .

Generates options, as well as run all of `go generate` for `resource`, `sample`, `client`, and `filer`

Source files are located in various places, under the name `options.yml`
