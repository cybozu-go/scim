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
