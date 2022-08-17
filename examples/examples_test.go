package examples_test

import (
	"context"
	"os"
	"strconv"

	"github.com/cybozu-go/scim/client"
	"github.com/cybozu-go/scim/resource"
	"github.com/google/uuid"
)

func NewClient(u string, options ...client.NewOption) *client.Client {
	v, err := strconv.ParseBool(os.Getenv(`SCIM_TRACE`))
	if err == nil {
		if v {
			options = append(options, client.WithTrace(os.Stderr))
		}
	}
	return client.New(u, options...)
}

// MockBackend is used for testing, and can only perform minimal tasks
type MockBackend struct {
	store map[string]*resource.User
}

func NewMockBackend() *MockBackend {
	return &MockBackend{
		store: make(map[string]*resource.User),
	}
}

func (s *MockBackend) CreateUser(_ context.Context, in *resource.User) (*resource.User, error) {
	id := uuid.New().String()

	var b resource.Builder
	u, err := b.User().From(in).ID(id).Build()
	if err != nil {
		return nil, err
	}

	s.store[id] = u
	return u, nil
}
