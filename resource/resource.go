//go:generate ../tools/gen-resource.sh

package resource

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

// schemas is a container for schemas. it dedupes schema URIs,
// and marshals to / unmarshals from []string
type schemas struct {
	mu      sync.RWMutex
	storage map[string]struct{}
}

func (s *schemas) Add(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.storage == nil {
		s.storage = make(map[string]struct{})
	}
	s.storage[v] = struct{}{}
}

func (s *schemas) UnmarshalJSON(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	var list []string
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	s.storage = make(map[string]struct{})
	for _, u := range list {
		s.storage[u] = struct{}{}
	}
	return nil
}

func (s *schemas) AcceptValue(v interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	storage := make(map[string]struct{})
	switch v := v.(type) {
	case []interface{}:
		for i, e := range v {
			es, ok := e.(string)
			if !ok {
				return fmt.Errorf(`failed to accept value for schemas: element %d of []interface{} was %T (expected string)`, i, e)
			}
			storage[es] = struct{}{}
		}
	case []string:
		for _, e := range v {
			storage[e] = struct{}{}
		}
	default:
		return fmt.Errorf(`schemas can only accept []string values (got %[1]T): %[1]v`, v)
	}
	s.storage = storage
	return nil
}

func (s *schemas) GetValue() []string {
	return s.List()
}

func (s *schemas) List() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]string, 0, len(s.storage))
	for u := range s.storage {
		list = append(list, u)
	}

	sort.Strings(list)
	return list
}

func (s *schemas) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.List())
}

// Builder is a centralized store for other type-specific builders,
// which exists mainly for convenience. Its zero-value can be used
// safely, but you could also use the `resource.NewXXXXBuilder()` function
// (where `XXXX` may be a resource type name) to start a method calling chain.
type Builder struct{}

// NewBuilder creates a new instance of the Builder object.
// This method exists solely as convenience, as the zero-value for
// the `resource.Builder` can safely be used without any initialization
func NewBuilder() *Builder {
	return &Builder{}
}

type AuthenticationSchemeType string

const (
	InvalidAuthenticationScheme AuthenticationSchemeType = ""
	OAuth                       AuthenticationSchemeType = "oauth"
	OAuth2                      AuthenticationSchemeType = "oauth2"
	OAuthBearerToken            AuthenticationSchemeType = "oauthbearertoken"
	HTTPBasic                   AuthenticationSchemeType = "httpbasic"
	HTTPDigest                  AuthenticationSchemeType = "httpdigest"
)

type DataType string

const (
	InvalidDataType DataType = "__INVALID__"
	String          DataType = "string"
	Boolean         DataType = "boolean"
	Decimal         DataType = "decimal"
	Integer         DataType = "integer"
	DateTime        DataType = "dateTime"
	Reference       DataType = "reference"
	Complex         DataType = "complex"
)

type Mutability string

const (
	MutReadOnly  Mutability = `readOnly`
	MutReadWrite Mutability = `readWrite`
	MutImmutable Mutability = `immutable`
	MutWriteOnly Mutability = `writeOnly`
)

type Uniqueness string

const (
	UniqNone   Uniqueness = `none`
	UniqServer Uniqueness = `server`
	UniqGlobal Uniqueness = `global`
)

type Returned string

const (
	ReturnedAlways  Returned = "always"
	ReturnedNever   Returned = "never"
	ReturnedDefault Returned = "default"
	ReturnedRequest Returned = "request"
)

type ErrorType string

const (
	ErrUnknown       ErrorType = `unknown` // not part of RFC7644, but used as our "zero value"
	ErrInvalidFilter ErrorType = `invalidFilter`
	ErrTooMany       ErrorType = `tooMany`
	ErrUniqueness    ErrorType = `uniqueness`
	ErrMutability    ErrorType = `mutability`
	ErrInvalidSyntax ErrorType = `invalidSyntax`
	ErrInvalidPath   ErrorType = `invalidPath`
	ErrNoTarget      ErrorType = `noTarget`
	ErrInvalidValue  ErrorType = `invalidValue`
	ErrInvalidVers   ErrorType = `invalidVers`
	ErrSensitive     ErrorType = `sensitive`
)

type PatchOperationType string

const (
	PatchInvalid PatchOperationType = ``
	PatchAdd     PatchOperationType = `add`
	PatchRemove  PatchOperationType = `remove`
	PatchReplace PatchOperationType = `replace`
)
