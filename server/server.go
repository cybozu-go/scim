package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	stdlibpath "path"
	"sync"

	"github.com/cybozu-go/scim/resource"
	"github.com/lestrrat-go/mux"
)

var ctKey = `Content-Type`
var mimeSCIM = `application/scim+json`

func NewServer(backend interface{}) (http.Handler, error) {
	var b Builder

	if v, ok := backend.(CreateGroupBackend); ok {
		b.CreateGroup(CreateGroupEndpoint(v))
	}
	if v, ok := backend.(DeleteGroupBackend); ok {
		b.DeleteGroup(DeleteGroupEndpoint(v))
	}

	if v, ok := backend.(ReplaceGroupBackend); ok {
		b.ReplaceGroup(ReplaceGroupEndpoint(v))
	}

	if v, ok := backend.(RetrieveGroupBackend); ok {
		b.RetrieveGroup(RetrieveGroupEndpoint(v))
	}

	if v, ok := backend.(CreateUserBackend); ok {
		b.CreateUser(CreateUserEndpoint(v))
	}

	if v, ok := backend.(DeleteUserBackend); ok {
		b.DeleteUser(DeleteUserEndpoint(v))
	}

	if v, ok := backend.(ReplaceUserBackend); ok {
		b.ReplaceUser(ReplaceUserEndpoint(v))
	}

	if v, ok := backend.(RetrieveUserBackend); ok {
		b.RetrieveUser(RetrieveUserEndpoint(v))
	}

	if v, ok := backend.(SearchGroupBackend); ok {
		b.SearchGroup(SearchGroupEndpoint(v))
	}

	if v, ok := backend.(SearchUserBackend); ok {
		b.SearchUser(SearchUserEndpoint(v))
	}

	if v, ok := backend.(SearchBackend); ok {
		b.Search(SearchEndpoint(v))
	}

	if v, ok := backend.(RetrieveServiceProviderConfigBackend); ok {
		b.ServiceProviderConfig(RetrieveServiceProviderConfigEndpoint(v))
	}

	if v, ok := backend.(RetrieveResourceTypesBackend); ok {
		b.ResourceTypes(RetrieveResourceTypesEndpoint(v))
	}
	return b.Build()
}

type Middleware interface {
	Wrap(http.Handler) http.Handler
}

type MiddlewareFunc func(http.Handler) http.Handler

func (f MiddlewareFunc) Wrap(h http.Handler) http.Handler {
	return f(h)
}

type Handler struct {
	method      string
	path        string // .search, User, Group, etc
	handler     http.Handler
	middlewares []Middleware
}

type Builder struct {
	mu       sync.RWMutex
	err      error
	basePath string // default "/"
	handlers []*Handler
}

// must lock before using
func (b *Builder) init() {
	b.err = nil
	b.basePath = "/"
	b.handlers = nil
}

func (b *Builder) Handler(method, path string, hh http.Handler, options ...HandlerOption) *Builder {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.err != nil {
		return b
	}

	if path == "" {
		b.err = fmt.Errorf(`handler path must be specified`)
		return b
	}

	if hh == nil {
		b.err = fmt.Errorf(`handler object must be specified`)
		return b
	}

	b.handlers = append(b.handlers, &Handler{
		method:  method,
		path:    path,
		handler: hh,
	})
	return b
}

func (b *Builder) Build() (http.Handler, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := b.err
	handlers := b.handlers
	basePath := b.basePath
	b.init()
	if err != nil {
		return nil, err
	}

	if basePath == "" {
		basePath = "/"
	}

	var r mux.Router
	for _, h := range handlers {
		hh := h.handler
		for _, m := range h.middlewares {
			hh = m.Wrap(hh)
		}
		path := stdlibpath.Clean(basePath + h.path)
		if err := r.Handler(h.method, path, hh); err != nil {
			return nil, fmt.Errorf(`failed to register handler (method = %q, path =%q)`, h.method, path)
		}
	}
	return &r, nil
}

func ServiceProviderConfig(config *resource.ServiceProviderConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(config)
	})
}

func (b *Builder) CreateGroup(hh http.Handler) *Builder {
	b.Handler(http.MethodPost, `/Groups`, hh)
	return b
}

func (b *Builder) DeleteGroup(hh http.Handler) *Builder {
	b.Handler(http.MethodDelete, `/Groups/{id}`, hh)
	return b
}

func (b *Builder) ReplaceGroup(hh http.Handler) *Builder {
	b.Handler(http.MethodPut, `/Groups/{id}`, hh)
	return b
}

func (b *Builder) RetrieveGroup(hh http.Handler) *Builder {
	b.Handler(http.MethodGet, `/Groups/{id}`, hh)
	return b
}

func (b *Builder) CreateUser(hh http.Handler) *Builder {
	b.Handler(http.MethodPost, `/Users`, hh)
	return b
}

func (b *Builder) DeleteUser(hh http.Handler) *Builder {
	b.Handler(http.MethodDelete, `/Users/{id}`, hh)
	return b
}

func (b *Builder) ReplaceUser(hh http.Handler) *Builder {
	b.Handler(http.MethodPut, `/Users/{id}`, hh)
	return b
}

func (b *Builder) RetrieveUser(hh http.Handler) *Builder {
	b.Handler(http.MethodGet, `/Users/{id}`, hh)
	return b
}

func (b *Builder) SearchGroup(hh http.Handler) *Builder {
	b.Handler(http.MethodPost, `/Groups/.search`, hh)
	return b
}

func (b *Builder) SearchUser(hh http.Handler) *Builder {
	b.Handler(http.MethodPost, `/Users/.search`, hh)
	return b
}

func (b *Builder) Search(hh http.Handler) *Builder {
	b.Handler(http.MethodPost, `/.search`, hh)
	return b
}

func (b *Builder) ServiceProviderConfig(hh http.Handler) *Builder {
	b.Handler(http.MethodGet, `/ServiceProviderConfig`, hh)
	return b
}

func (b *Builder) ResourceTypes(hh http.Handler) *Builder {
	b.Handler(http.MethodGet, `/ResourceTypes`, hh)
	return b
}
