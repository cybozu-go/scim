package server_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/test"
)

type mockBackend struct {
	mu    sync.RWMutex
	users map[string]*resource.User
}

func (m *mockBackend) DeleteUser(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.users[id]; !ok {
		return fmt.Errorf(`id not found`)
	}
	delete(m.users, id)
	return nil
}

func (m *mockBackend) ReplaceUser(id string, in *resource.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	// TODO: no check is beig done here
	_, err := m.RetrieveUser(id)
	if err != nil {
		return err
	}

	// HACK: attributes may need to be merged, etc
	m.users[id] = in
	return nil
}

func (m *mockBackend) CreateUser(in *resource.User) (*resource.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := fmt.Sprintf(`%d`, rand.Int63()) // TODO: check for clashes

	var b resource.UserBuilder
	user, err := b.ID(id).
		UserName(in.UserName()).
		ExternalID(in.ExternalID()).
		Name(in.Name()).
		Build()
	if err != nil {
		return nil, err
	}
	m.users[id] = user
	return user, nil
}

func (m *mockBackend) RetrieveUser(id string) (*resource.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	user, ok := m.users[id]
	if !ok {
		return nil, fmt.Errorf(`user not found`)
	}
	return user, nil
}

func (m *mockBackend) Search(*resource.SearchRequest) (*resource.ListResponse, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var b resource.Builder

	return b.ListResponse().
		TotalResults(2).
		StartIndex(1).
		ItemsPerPage(10).
		Resources(
			b.User().
				ID("2819c223-7f76-413861904646").
				UserName("jsmith").
				DisplayName("Smith, James").
				MustBuild(),
			b.Group().
				ID("c8596b90-7539-4f20968d1908").
				DisplayName("Smith Family").
				MustBuild(),
		).
		Build()
}

func TestServer(t *testing.T) {
	test.RunConformanceTests("Mock backend", &mockBackend{
		users: make(map[string]*resource.User),
	})
}
