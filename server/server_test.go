package server_test

import (
	"context"
	"fmt"
	"math/rand"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/lestrrat-go/scim/client"
	"github.com/lestrrat-go/scim/resource"
	"github.com/lestrrat-go/scim/server"
	"github.com/stretchr/testify/require"
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
	t.Run("mock server", func(t *testing.T) {
		hh, err := server.NewServer(&mockBackend{
			users: make(map[string]*resource.User),
		})
		require.NoError(t, err, `server.NewServer should succeed`)

		srv := httptest.NewServer(hh)

		cl := client.New(srv.URL, client.WithClient(srv.Client()))

		t.Run("search via /.search", func(t *testing.T) {
			lr, err := cl.Search().Search().
				Attributes(`displayName`, `userName`).
				Filter(`displayName sw "smith"`).
				StartIndex(1).
				Count(10).
				Do(context.TODO())
			require.NoError(t, err, `cl.Search should succeed`)
			t.Logf("%#v", lr)
		})
		t.Run("/Users", func(t *testing.T) {
			t.Run("POST /Users", func(t *testing.T) {
				createdUser, err := cl.User().CreateUser().
					UserName("bjensen").
					ExternalID("bjensen").
					Emails(resource.NewEmailBuilder().
						Value("babs@jensen.org").
						Primary(true).
						MustBuild()).
					Name(resource.NewNamesBuilder().
						Formatted("Ms. Barbara J Jensen III").
						FamilyName("Jensen").
						GivenName("Barbara").
						MustBuild()).
					Do(context.TODO())
				require.NoError(t, err, `CreateUser should succeed`)
				t.Run(fmt.Sprintf("GET /Users/%s", createdUser.ID()), func(t *testing.T) {
					_, err := cl.User().GetUser(createdUser.ID()).
						Do(context.TODO())
					require.NoError(t, err, `GetUser should succeed`)
				})
				t.Run(fmt.Sprintf("DELETE /Users/%s", createdUser.ID()), func(t *testing.T) {
					err := cl.User().DeleteUser(createdUser.ID()).
						Do(context.TODO())
					require.NoError(t, err, `DeleteUser should succeed`)
				})
				t.Run(fmt.Sprintf("GET /Users/%s (after delete)", createdUser.ID()), func(t *testing.T) {
					_, err := cl.User().GetUser(createdUser.ID()).
						Do(context.TODO())
					require.Error(t, err, `GetUser should fail`)
				})
			})
			t.Run("GET /Users/foobar", func(t *testing.T) {
				_, err := cl.User().GetUser("foobar").
					Do(context.TODO())
				require.Error(t, err, `GetUser should fail`)
			})
		})
	})
}
