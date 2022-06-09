package server_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/test"
)

type mockBackend struct {
	mu     sync.RWMutex
	users  map[string]*resource.User
	groups map[string]*resource.Group
}

func (m *mockBackend) DeleteGroup(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.groups[id]; !ok {
		return fmt.Errorf(`id not found`)
	}
	delete(m.groups, id)
	return nil
}

func (m *mockBackend) RetrieveGroup(id string, fields ...string) (*resource.Group, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, err := m.retrieveGroupNoLock(id)
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 {
		return v, nil
	}

	// HACK (we want to delete certain fields)
	buf, _ := json.Marshal(v)
	var store map[string]interface{}
	_ = json.Unmarshal(buf, &store)
	want := make(map[string]struct{})
	want["id"] = struct{}{}
	for _, f := range fields {
		want[f] = struct{}{}
	}

	for k := range store {
		if _, ok := want[k]; !ok {
			delete(store, k)
		}
	}

	buf, _ = json.Marshal(store)
	var u resource.Group
	_ = json.Unmarshal(buf, &u)
	return &u, nil
}

func (m *mockBackend) ReplaceGroup(id string, in *resource.Group) (*resource.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// TODO: no check is beig done here
	g, err := m.retrieveGroupNoLock(id)
	if err != nil {
		return nil, err
	}

	var builder resource.Builder

	groupBuilder := builder.Group().From(g)

	if in.HasDisplayName() {
		groupBuilder.DisplayName(in.DisplayName())
	}

	groupBuilder.Members(in.Members()...)

	// TODO: handle other fields
	gNew, err := groupBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to create group: %w`, err)
	}

	m.groups[id] = gNew
	return gNew, nil
}

func (m *mockBackend) CreateGroup(in *resource.Group) (*resource.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := fmt.Sprintf(`%d`, rand.Int63()) // TODO: check for clashes

	var b resource.Builder

	meta, err := b.Meta().
		ResourceType("Group").
		Location(`https://foobar.com/scim/v2/Groups/` + id).
		Build()
	if err != nil {
		return nil, err
	}

	group, err := b.Group().ID(id).
		DisplayName(in.DisplayName()).
		Members(in.Members()...).
		Meta(meta).
		Build()
	if err != nil {
		return nil, err
	}
	m.groups[id] = group
	return group, nil
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

func (m *mockBackend) ReplaceUser(id string, in *resource.User) (*resource.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// TODO: no check is beig done here
	u, err := m.retrieveUserNoLock(id)
	if err != nil {
		return nil, err
	}

	var builder resource.Builder

	userBuilder := builder.User().From(u)

	if in.HasUserName() {
		userBuilder.UserName(in.UserName())
	}

	if in.HasEmails() {
		userBuilder.Emails(in.Emails()...)
	}

	// TODO: handle other fields
	uNew, err := userBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf(`failed to create user: %w`, err)
	}

	m.users[id] = uNew
	return uNew, nil
}

func (m *mockBackend) CreateUser(in *resource.User) (*resource.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := fmt.Sprintf(`%d`, rand.Int63()) // TODO: check for clashes

	var b resource.Builder

	meta, err := b.Meta().
		ResourceType("User").
		Location(`https://foobar.com/scim/v2/Users/` + id).
		Build()
	if err != nil {
		return nil, err
	}

	user, err := b.User().ID(id).
		UserName(in.UserName()).
		ExternalID(in.ExternalID()).
		Name(in.Name()).
		Meta(meta).
		Build()
	if err != nil {
		return nil, err
	}
	m.users[id] = user
	return user, nil
}

func (m *mockBackend) RetrieveUser(id string, fields ...string) (*resource.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, err := m.retrieveUserNoLock(id)
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 {
		return v, nil
	}

	// HACK (we want to delete certain fields)
	buf, _ := json.Marshal(v)
	var store map[string]interface{}
	_ = json.Unmarshal(buf, &store)
	want := make(map[string]struct{})
	want["id"] = struct{}{}
	for _, f := range fields {
		want[f] = struct{}{}
	}

	for k := range store {
		if _, ok := want[k]; !ok {
			delete(store, k)
		}
	}

	buf, _ = json.Marshal(store)
	var u resource.User
	_ = json.Unmarshal(buf, &u)
	return &u, nil
}

func (m *mockBackend) retrieveUserNoLock(id string) (*resource.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, fmt.Errorf(`user not found`)
	}
	return user, nil
}

func (m *mockBackend) retrieveGroupNoLock(id string) (*resource.Group, error) {
	group, ok := m.groups[id]
	if !ok {
		return nil, fmt.Errorf(`group not found`)
	}
	return group, nil
}

func (m *mockBackend) SearchUser(*resource.SearchRequest) (*resource.ListResponse, error) {
	return m.search(nil, true, false)
}

func (m *mockBackend) SearchGroup(*resource.SearchRequest) (*resource.ListResponse, error) {
	return m.search(nil, false, true)
}

func (m *mockBackend) search(_ *resource.SearchRequest, searchUser, searchGroup bool) (*resource.ListResponse, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var b resource.Builder

	var resources []interface{}
	if searchUser {
		resources = append(resources, b.User().
			ID("2819c223-7f76-413861904646").
			UserName("jsmith").
			DisplayName("Smith, James").
			MustBuild(),
		)
	}

	if searchGroup {
		// TODO: This is totally wrong as far as the conformance tests go
		// and WILL fail once we write proper tests
		resources = append(resources, b.Group().
			ID("c8596b90-7539-4f20968d1908").
			DisplayName("Smith Family").
			MustBuild(),
		)
		resources = append(resources, b.Group().
			ID("c8596b90-7539-5f30461a1b08").
			DisplayName("Johson Family").
			MustBuild(),
		)
	}

	return b.ListResponse().
		TotalResults(len(resources)).
		StartIndex(1).
		ItemsPerPage(10).
		Resources(resources...).
		Build()
}

func TestServer(t *testing.T) {
	test.RunConformanceTests(t, "Mock backend", &mockBackend{
		users:  make(map[string]*resource.User),
		groups: make(map[string]*resource.Group),
	})
}
