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

  "github.com/lestrrat-go/scim/client"
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
source: [./examples/client_user_create_example_test.go](https://github.com/lestrrat-go/scim/blob/main/./examples/client_user_create_example_test.go)
<!-- END INCLUDE -->

# TODO

* Implement a sample server
* Implement check in CI to diff against generated code


