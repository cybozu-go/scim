package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/cybozu-go/scim/resource"
	"github.com/lestrrat-go/mux"
)

type CreateGroupBackend interface {
	CreateGroup(context.Context, *resource.Group) (*resource.Group, error)
}

type DeleteGroupBackend interface {
	DeleteGroup(context.Context, string) error
}

type ReplaceGroupBackend interface {
	ReplaceGroup(context.Context, string, *resource.Group) (*resource.Group, error)
}

type RetrieveGroupBackend interface {
	RetrieveGroup(context.Context, string, []string, []string) (*resource.Group, error)
}

type PatchGroupBackend interface {
	PatchGroup(context.Context, string, *resource.PatchRequest) (*resource.Group, error)
}

type CreateUserBackend interface {
	CreateUser(context.Context, *resource.User) (*resource.User, error)
}

type DeleteUserBackend interface {
	DeleteUser(context.Context, string) error
}

type ReplaceUserBackend interface {
	ReplaceUser(context.Context, string, *resource.User) (*resource.User, error)
}

type RetrieveUserBackend interface {
	RetrieveUser(context.Context, string, []string, []string) (*resource.User, error)
}

type PatchUserBackend interface {
	PatchUser(context.Context, string, *resource.PatchRequest) (*resource.User, error)
}

type SearchBackend interface {
	Search(context.Context, *resource.SearchRequest) (*resource.ListResponse, error)
}

type SearchUserBackend interface {
	SearchUser(context.Context, *resource.SearchRequest) (*resource.ListResponse, error)
}

type SearchGroupBackend interface {
	SearchGroup(context.Context, *resource.SearchRequest) (*resource.ListResponse, error)
}

type RetrieveServiceProviderConfigBackend interface {
	RetrieveServiceProviderConfig(context.Context) (*resource.ServiceProviderConfig, error)
}

type RetrieveResourceTypesBackend interface {
	RetrieveResourceTypes(context.Context) ([]*resource.ResourceType, error)
}

type ListSchemasBackend interface {
	ListSchemas(context.Context) (*resource.ListResponse, error)
}

type RetrieveSchemaBackend interface {
	RetrieveSchema(context.Context, string) (*resource.Schema, error)
}

// WriteSCIMError creates a resource.Error from the given input and
// writes to the response writer
func WriteSCIMError(w http.ResponseWriter, st int, msg string) {
	// log.Printf("%d: %s", st, msg)
	w.WriteHeader(st)
	serr := resource.NewErrorBuilder().
		Status(st).
		Detail(msg).
		SCIMType(resource.ErrUnknown).
		MustBuild()

	// Try one more time in vein to write the error of Encode?
	// Look, I've explicitly stated to ignore errors, you linters
	// should just let me be, OK?
	//nolint:errchkjson
	_ = json.NewEncoder(w).Encode(serr)
}

func WriteError(w http.ResponseWriter, err error) {
	var serr *resource.Error
	if errors.As(err, &serr) {
		w.WriteHeader(serr.Status())
		// Look, I've explicitly stated to ignore errors, you linters
		// should just let me be, OK?
		//nolint:errchkjson
		_ = json.NewEncoder(w).Encode(serr)
		return
	}

	WriteSCIMError(w, http.StatusInternalServerError, err.Error())
}

func DeleteGroupEndpoint(b DeleteGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		if err := b.DeleteGroup(r.Context(), id); err != nil {
			WriteError(w, err)
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
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		var group resource.Group
		if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		replaced, err := b.ReplaceGroup(r.Context(), id, &group)
		if err != nil {
			WriteError(w, err)
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
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		var attrs []string
		if v := r.URL.Query().Get(`attributes`); v != "" {
			attrs = strings.Split(v, ",")
		}

		var excluded []string
		if v := r.URL.Query().Get(`excludedAttributes`); v != "" {
			excluded = strings.Split(v, ",")
		}
		group, err := b.RetrieveGroup(r.Context(), id, attrs, excluded)
		if err != nil {
			WriteError(w, err)
			return
		}

		if meta := group.Meta(); meta != nil {
			if v := meta.Version(); v != "" {
				w.Header().Set(`ETag`, v)
			}
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(group)
	})
}

func CreateGroupEndpoint(b CreateGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var group resource.Group
		if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		created, err := b.CreateGroup(r.Context(), &group)
		if err != nil {
			WriteError(w, err)
			return
		}

		if meta := created.Meta(); meta != nil {
			if v := meta.Version(); v != "" {
				w.Header().Set(`ETag`, v)
			}
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
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		if err := b.DeleteUser(r.Context(), id); err != nil {
			WriteError(w, err)
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
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		var user resource.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		newUser, err := b.ReplaceUser(r.Context(), id, &user)
		if err != nil {
			WriteError(w, err)
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
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		var attrs []string
		if v := r.URL.Query().Get(`attributes`); v != "" {
			attrs = strings.Split(v, ",")
		}

		var excluded []string
		if v := r.URL.Query().Get(`excludedAttributes`); v != "" {
			excluded = strings.Split(v, ",")
		}
		user, err := b.RetrieveUser(r.Context(), id, attrs, excluded)
		if err != nil {
			WriteError(w, err)
			return
		}

		if meta := user.Meta(); meta != nil {
			if v := meta.Version(); v != "" {
				w.Header().Set(`ETag`, v)
			}
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user)
	})
}

func PatchUserEndpoint(b PatchUserBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		defer r.Body.Close()
		var preq resource.PatchRequest
		if err := json.NewDecoder(r.Body).Decode(&preq); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		user, err := b.PatchUser(r.Context(), id, &preq)
		if err != nil {
			WriteError(w, err)
			return
		}

		if user == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if meta := user.Meta(); meta != nil {
			if v := meta.Version(); v != "" {
				w.Header().Set(`ETag`, v)
			}
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user)
	})
}

func PatchGroupEndpoint(b PatchGroupBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		defer r.Body.Close()
		var preq resource.PatchRequest
		if err := json.NewDecoder(r.Body).Decode(&preq); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		group, err := b.PatchGroup(r.Context(), id, &preq)
		if err != nil {
			WriteError(w, err)
			return
		}

		if group == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if meta := group.Meta(); meta != nil {
			if v := meta.Version(); v != "" {
				w.Header().Set(`ETag`, v)
			}
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(group)
	})
}

func CreateUserEndpoint(b CreateUserBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user resource.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		created, err := b.CreateUser(r.Context(), &user)
		if err != nil {
			WriteError(w, err)
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
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		lr, err := b.Search(r.Context(), &q)
		if err != nil {
			WriteError(w, err)
			return
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(lr); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to encode response`)
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
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		lr, err := b.SearchUser(r.Context(), &q)
		if err != nil {
			WriteError(w, err)
			return
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(lr); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to encode response`)
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
			WriteSCIMError(w, http.StatusBadRequest, `failed to parse payload`)
			return
		}

		lr, err := b.SearchGroup(r.Context(), &q)
		if err != nil {
			WriteError(w, err)
			return
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(lr); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to encode response`)
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
		scp, err := b.RetrieveServiceProviderConfig(r.Context())
		if err != nil {
			WriteError(w, err)
			return
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(scp); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to encode response`)
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
		rts, err := b.RetrieveResourceTypes(r.Context())
		if err != nil {
			WriteError(w, err)
			return
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rts); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to encode response`)
			return
		}

		hdr := w.Header()
		hdr.Set(ctKey, mimeSCIM)
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, &buf) // not much you can do by this point
	})
}

func ListSchemasEndpoint(b ListSchemasBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		schemas, err := b.ListSchemas(r.Context())
		if err != nil {
			WriteError(w, err)
			return
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(schemas); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to encode response`)
			return
		}

		hdr := w.Header()
		hdr.Set(ctKey, mimeSCIM)
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, &buf) // not much you can do by this point
	})
}

func RetrieveSchemaEndpoint(b RetrieveSchemaBackend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars.Get(`id`)
		if id == "" {
			WriteSCIMError(w, http.StatusBadRequest, `missing ID`)
			return
		}

		schema, err := b.RetrieveSchema(r.Context(), id)
		if err != nil {
			WriteError(w, err)
			return
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(schema); err != nil {
			WriteSCIMError(w, http.StatusBadRequest, `failed to encode response`)
			return
		}

		hdr := w.Header()
		hdr.Set(ctKey, mimeSCIM)
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, &buf) // not much you can do by this point
	})
}
