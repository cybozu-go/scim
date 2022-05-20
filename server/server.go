package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/lestrrat-go/scim/resource"
)

var ctKey = `Content-Type`
var mimeSCIM = `application/scim+json`

func NewServer(backend Backend) (http.Handler, error) {
	var b Builder

	return b.
		SearchEndpoint(SearchEndpoint(backend)).
		UsersEndpoint(UsersEndpoint(backend)).
		Build()
}

type Middleware interface {
	Wrap(http.Handler) http.Handler
}

type MiddlewareFunc func(http.Handler) http.Handler

func (f MiddlewareFunc) Wrap(h http.Handler) http.Handler {
	return f(h)
}

type Backend interface {
	CreateUser(user *resource.User) (*resource.User, error)
	DeleteUser(id string) error
	ReplaceUser(id string, user *resource.User) error
	RetrieveUser(string) (*resource.User, error)
	Search(*resource.SearchRequest) (*resource.ListResponse, error)
}

type Handler struct {
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

func (b *Builder) WithBackend(backend Backend) *Builder {
	return b.
		//		Handler(`/Users`, UsersEndpoint(backend)).
		//		Handler(`/Groups`, GroupsEndpoint(backend)).
		//		Handler(`/Self`, SelfEndpoint(backend))
		Handler(`/.search`, SearchEndpoint(backend))
}

func (b *Builder) Handler(path string, hh http.Handler, options ...HandlerOption) *Builder {
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
		path:    path,
		handler: hh,
	})
	return b
}

func (b *Builder) SearchEndpoint(handler http.Handler, options ...HandlerOption) *Builder {
	path := `.search`
	var remaining []HandlerOption
	//nolint:forcetypeassert
	for _, option := range options {
		switch option.Ident() {
		case identPath{}:
			path = option.Value().(string)
		default:
			remaining = append(remaining, option)
		}
	}

	b.Handler(path, handler, remaining...)
	return b
}

func (b *Builder) UsersEndpoint(handler http.Handler, options ...HandlerOption) *Builder {
	path := `Users`
	var remaining []HandlerOption
	//nolint:forcetypeassert
	for _, option := range options {
		switch option.Ident() {
		case identPath{}:
			path = option.Value().(string)
		default:
			remaining = append(remaining, option)
		}
	}

	b.Handler(path, handler, remaining...)
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

	var mux http.ServeMux
	for _, h := range handlers {
		hh := h.handler
		for _, m := range h.middlewares {
			hh = m.Wrap(hh)
		}
		path := basePath + h.path
		mux.Handle(path, hh)
		// HACK: there must be a better way to do this, but I'm too hungry to think
		if !strings.HasSuffix(path, "/") {
			mux.Handle(path+"/", hh)
		}
	}
	return &mux, nil
}

func ServiceProviderConfig(config *resource.ServiceProviderConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(config)
	})
}

func UsersEndpoint(b Backend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet, http.MethodPut, http.MethodDelete:
			// We must match the path path/(id)
			i := strings.LastIndexByte(r.URL.Path, '/')
			if i < 0 || i == len(r.URL.Path) { // pedantic
				// TODO: log
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Everything after the last slash is the ID
			id := r.URL.Path[i+1:]

			if r.Method == http.MethodDelete {
				if err := b.DeleteUser(id); err != nil {
					// TODO: log
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusNoContent)
				return
			}

			if r.Method == http.MethodPut {
				var user resource.User
				if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
					// TODO: log
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if err := b.ReplaceUser(id, &user); err != nil {
					// TODO: log
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			// TODO: handle "attributes" and stuff?
			user, err := b.RetrieveUser(id)
			if err != nil {
				// TODO: distinguish between error and not found error
				// TODO: log
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(user)
		case http.MethodPost:
			var user resource.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				// TODO: log
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			created, err := b.CreateUser(&user)
			if err != nil {
				// TODO: log
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set(ctKey, mimeSCIM)
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(created)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}

// Creates an instance of reference implementation http.Handler that
// uses the specified Backend
func SearchEndpoint(b Backend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var q resource.SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// TODO: log
			return
		}

		lr, err := b.Search(&q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: log
			return
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(lr); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: log
			return
		}

		hdr := w.Header()
		hdr.Set(ctKey, mimeSCIM)
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, &buf) // not much you can do by this point
	})
}
