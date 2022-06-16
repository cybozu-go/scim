package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/cybozu-go/scim/resource"
	"github.com/lestrrat-go/mux"
)

type CreateGroupBackend interface {
	CreateGroup(*resource.Group) (*resource.Group, error)
}

type DeleteGroupBackend interface {
	DeleteGroup(id string) error
}

type ReplaceGroupBackend interface {
	ReplaceGroup(string, *resource.Group) (*resource.Group, error)
}

type RetrieveGroupBackend interface {
	RetrieveGroup(string, ...string) (*resource.Group, error)
}

type CreateUserBackend interface {
	CreateUser(*resource.User) (*resource.User, error)
}

type DeleteUserBackend interface {
	DeleteUser(id string) error
}

type ReplaceUserBackend interface {
	ReplaceUser(id string, user *resource.User) (*resource.User, error)
}

type RetrieveUserBackend interface {
	RetrieveUser(string, ...string) (*resource.User, error)
}

type SearchBackend interface {
	Search(*resource.SearchRequest) (*resource.ListResponse, error)
}

type SearchUserBackend interface {
	SearchUser(*resource.SearchRequest) (*resource.ListResponse, error)
}

type SearchGroupBackend interface {
	SearchGroup(*resource.SearchRequest) (*resource.ListResponse, error)
}

type RetrieveServiceProviderConfigBackend interface {
	RetrieveServiceProviderConfig() (*resource.ServiceProviderConfig, error)
}

type RetrieveResourceTypesBackend interface {
	RetrieveResourceTypes() ([]*resource.ResourceType, error)
}

func DeleteGroupEndpoint(b DeleteGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := b.DeleteGroup(id); err != nil {
			// TODO: log
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func ReplaceGroupEndpoint(b ReplaceGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var group resource.Group
		if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		replaced, err := b.ReplaceGroup(id, &group)
		if err != nil {
			// TODO: log
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(replaced)
	})
}

func RetrieveGroupEndpoint(b RetrieveGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// TODO: handle "attributes" and stuff?
		group, err := b.RetrieveGroup(id)
		if err != nil {
			// TODO: distinguish between error and not found error
			// TODO: log
			log.Printf("%s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(group)
	})
}

func CreateGroupEndpoint(b CreateGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var group resource.Group
		if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		created, err := b.CreateGroup(&group)
		if err != nil {
			// TODO: log
			log.Printf("%s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set(ctKey, mimeSCIM)
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(created)
	})
}

func DeleteUserEndpoint(b DeleteUserBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := b.DeleteUser(id); err != nil {
			// TODO: log
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func ReplaceUserEndpoint(b ReplaceUserBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var user resource.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newUser, err := b.ReplaceUser(id, &user)
		if err != nil {
			err = fmt.Errorf(`replace user operation failed: %w`, err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(newUser)
	})
}

func RetrieveUserEndpoint(b RetrieveUserBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var attrs []string
		if v := r.URL.Query().Get(`attributes`); v != "" {
			attrs = strings.Split(v, ",")
		}
		user, err := b.RetrieveUser(id, attrs...)
		if err != nil {
			// TODO: distinguish between error and not found error
			// TODO: log
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user)
	})
}

func CreateUserEndpoint(b CreateUserBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user resource.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			// TODO: log
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		created, err := b.CreateUser(&user)
		if err != nil {
			log.Printf("%s", err)
			// TODO: log
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set(ctKey, mimeSCIM)
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(created)
	})
}

// Creates an instance of reference implementation http.Handler that
// uses the specified Backend
func SearchEndpoint(b SearchBackend) http.Handler {
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
			fmt.Fprint(w, err.Error())
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

func SearchUserEndpoint(b SearchUserBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var q resource.SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// TODO: log
			return
		}

		lr, err := b.SearchUser(&q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: log
			fmt.Fprint(w, err.Error())
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

func SearchGroupEndpoint(b SearchGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var q resource.SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// TODO: log
			return
		}

		lr, err := b.SearchGroup(&q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: log
			fmt.Fprint(w, err.Error())
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

func RetrieveServiceProviderConfigEndpoint(b RetrieveServiceProviderConfigBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scp, err := b.RetrieveServiceProviderConfig()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(scp); err != nil {
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

func RetrieveResourceTypesEndpoint(b RetrieveResourceTypesBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rts, err := b.RetrieveResourceTypes()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rts); err != nil {
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
