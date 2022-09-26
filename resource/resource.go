//go:generate ../tools/cmd/genresources.sh

package resource

import (
	"encoding/json"
	"sort"
	"sync"
)

// schemas is a container for schemas. it dedupes schema URIs,
// and marshals to / unmarshals from []string
type schemas struct {
	once    sync.Once
	storage map[string]struct{}
}

func (s *schemas) initialize() {
	s.storage = make(map[string]struct{})
}

func (s *schemas) Add(v string) {
	s.once.Do(s.initialize)
	s.storage[v] = struct{}{}
}

func (s *schemas) UnmarshalJSON(data []byte) error {
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

func (s schemas) Get() []string {
	return s.List()
}

func (s schemas) List() []string {
	list := make([]string, 0, len(s.storage))
	for u := range s.storage {
		list = append(list, u)
	}

	sort.Strings(list)
	return list
}

func (s schemas) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.List())
}

type pair struct {
	Key   string
	Value interface{}
}

func defaultUserValidate(v *User) error {
	return nil
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
