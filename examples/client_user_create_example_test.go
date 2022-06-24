package examples_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cybozu-go/scim/client"
)

func ExampleClient_UserCreate() {
	// TODO: setup a toy SCIM server
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{}`)
	}))
	defer srv.Close()

	cl := client.New(srv.URL)

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
