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
