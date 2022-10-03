// Generated by "sketch" utility. DO NOT EDIT
package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

func init() {
	Register("AuthenticationScheme", "", AuthenticationScheme{})
	RegisterBuilder("AuthenticationScheme", "", AuthenticationSchemeBuilder{})
}

type AuthenticationScheme struct {
	mu               sync.RWMutex
	description      *string
	documentationURI *string
	name             *string
	specURI          *string
	typ              *AuthenticationSchemeType
	extra            map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	AuthenticationSchemeDescriptionKey      = "description"
	AuthenticationSchemeDocumentationURIKey = "documentationUri"
	AuthenticationSchemeNameKey             = "name"
	AuthenticationSchemeSpecURIKey          = "specUri"
	AuthenticationSchemeTypeKey             = "type"
)

// Get retrieves the value associated with a key
func (v *AuthenticationScheme) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *AuthenticationScheme) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case AuthenticationSchemeDescriptionKey:
		if val := v.description; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case AuthenticationSchemeDocumentationURIKey:
		if val := v.documentationURI; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case AuthenticationSchemeNameKey:
		if val := v.name; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case AuthenticationSchemeSpecURIKey:
		if val := v.specURI; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case AuthenticationSchemeTypeKey:
		if val := v.typ; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	default:
		if v.extra != nil {
			val, ok := v.extra[key]
			if ok {
				return blackmagic.AssignIfCompatible(dst, val)
			}
		}
	}
	return fmt.Errorf(`no such key %q`, key)
}

// Set sets the value of the specified field. The name must be a JSON
// field name, not the Go name
func (v *AuthenticationScheme) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case AuthenticationSchemeDescriptionKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field description, got %T`, value)
		}
		v.description = &converted
	case AuthenticationSchemeDocumentationURIKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field documentationUri, got %T`, value)
		}
		v.documentationURI = &converted
	case AuthenticationSchemeNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field name, got %T`, value)
		}
		v.name = &converted
	case AuthenticationSchemeSpecURIKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field specUri, got %T`, value)
		}
		v.specURI = &converted
	case AuthenticationSchemeTypeKey:
		converted, ok := value.(AuthenticationSchemeType)
		if !ok {
			return fmt.Errorf(`expected value of type AuthenticationSchemeType for field type, got %T`, value)
		}
		v.typ = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}

// Has returns true if the field specified by the argument has been populated.
// The field name must be the JSON field name, not the Go-structure's field name.
func (v *AuthenticationScheme) Has(name string) bool {
	switch name {
	case AuthenticationSchemeDescriptionKey:
		return v.description != nil
	case AuthenticationSchemeDocumentationURIKey:
		return v.documentationURI != nil
	case AuthenticationSchemeNameKey:
		return v.name != nil
	case AuthenticationSchemeSpecURIKey:
		return v.specURI != nil
	case AuthenticationSchemeTypeKey:
		return v.typ != nil
	default:
		if v.extra != nil {
			if _, ok := v.extra[name]; ok {
				return true
			}
		}
		return false
	}
}

// Keys returns a slice of string comprising of JSON field names whose values
// are present in the object.
func (v *AuthenticationScheme) Keys() []string {
	keys := make([]string, 0, 5)
	if v.description != nil {
		keys = append(keys, AuthenticationSchemeDescriptionKey)
	}
	if v.documentationURI != nil {
		keys = append(keys, AuthenticationSchemeDocumentationURIKey)
	}
	if v.name != nil {
		keys = append(keys, AuthenticationSchemeNameKey)
	}
	if v.specURI != nil {
		keys = append(keys, AuthenticationSchemeSpecURIKey)
	}
	if v.typ != nil {
		keys = append(keys, AuthenticationSchemeTypeKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasDescription returns true if the field `description` has been populated
func (v *AuthenticationScheme) HasDescription() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.description != nil
}

// HasDocumentationURI returns true if the field `documentationUri` has been populated
func (v *AuthenticationScheme) HasDocumentationURI() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.documentationURI != nil
}

// HasName returns true if the field `name` has been populated
func (v *AuthenticationScheme) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

// HasSpecURI returns true if the field `specUri` has been populated
func (v *AuthenticationScheme) HasSpecURI() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.specURI != nil
}

// HasType returns true if the field `type` has been populated
func (v *AuthenticationScheme) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *AuthenticationScheme) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.description; val != nil {
		return *val
	}
	return ""
}

func (v *AuthenticationScheme) DocumentationURI() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.documentationURI; val != nil {
		return *val
	}
	return ""
}

func (v *AuthenticationScheme) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.name; val != nil {
		return *val
	}
	return ""
}

func (v *AuthenticationScheme) SpecURI() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.specURI; val != nil {
		return *val
	}
	return ""
}

func (v *AuthenticationScheme) Type() AuthenticationSchemeType {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return InvalidAuthenticationScheme
}

// Remove removes the value associated with a key
func (v *AuthenticationScheme) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case AuthenticationSchemeDescriptionKey:
		v.description = nil
	case AuthenticationSchemeDocumentationURIKey:
		v.documentationURI = nil
	case AuthenticationSchemeNameKey:
		v.name = nil
	case AuthenticationSchemeSpecURIKey:
		v.specURI = nil
	case AuthenticationSchemeTypeKey:
		v.typ = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *AuthenticationScheme) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &AuthenticationScheme{
		description:      v.description,
		documentationURI: v.documentationURI,
		name:             v.name,
		specURI:          v.specURI,
		typ:              v.typ,
		extra:            extra,
	})
}

// MarshalJSON serializes AuthenticationScheme into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *AuthenticationScheme) MarshalJSON() ([]byte, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, k := range v.Keys() {
		var val interface{}
		if err := v.getNoLock(k, &val, true); err != nil {
			return nil, fmt.Errorf(`failed to retrieve value for field %q: %w`, k, err)
		}

		if i > 0 {
			buf.WriteByte(',')
		}
		if err := enc.Encode(k); err != nil {
			return nil, fmt.Errorf(`failed to encode map key name: %w`, err)
		}
		buf.WriteByte(':')
		if err := enc.Encode(val); err != nil {
			return nil, fmt.Errorf(`failed to encode map value for %q: %w`, k, err)
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON deserializes a piece of JSON data into AuthenticationScheme.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *AuthenticationScheme) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.description = nil
	v.documentationURI = nil
	v.name = nil
	v.specURI = nil
	v.typ = nil

	dec := json.NewDecoder(bytes.NewReader(data))
	var extra map[string]interface{}

LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading JSON token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			if tok == '}' { // end of object
				break LOOP
			}
			// we should only get into this clause at the very beginning, and just once
			if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string:
			switch tok {
			case AuthenticationSchemeDescriptionKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AuthenticationSchemeDescriptionKey, err)
				}
				v.description = &val
			case AuthenticationSchemeDocumentationURIKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AuthenticationSchemeDocumentationURIKey, err)
				}
				v.documentationURI = &val
			case AuthenticationSchemeNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AuthenticationSchemeNameKey, err)
				}
				v.name = &val
			case AuthenticationSchemeSpecURIKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AuthenticationSchemeSpecURIKey, err)
				}
				v.specURI = &val
			case AuthenticationSchemeTypeKey:
				var val AuthenticationSchemeType
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AuthenticationSchemeTypeKey, err)
				}
				v.typ = &val
			default:
				var val interface{}
				if err := v.decodeExtraField(tok, dec, &val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, tok, err)
				}
				if extra == nil {
					extra = make(map[string]interface{})
				}
				extra[tok] = val
			}
		}
	}

	if extra != nil {
		v.extra = extra
	}
	return nil
}

type AuthenticationSchemeBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *AuthenticationScheme
}

// NewAuthenticationSchemeBuilder creates a new AuthenticationSchemeBuilder instance.
// AuthenticationSchemeBuilder is safe to be used uninitialized as well.
func NewAuthenticationSchemeBuilder() *AuthenticationSchemeBuilder {
	return &AuthenticationSchemeBuilder{}
}

func (b *AuthenticationSchemeBuilder) initialize() {
	b.err = nil
	b.object = &AuthenticationScheme{}
}
func (b *AuthenticationSchemeBuilder) Description(in string) *AuthenticationSchemeBuilder {
	return b.SetField(AuthenticationSchemeDescriptionKey, in)
}
func (b *AuthenticationSchemeBuilder) DocumentationURI(in string) *AuthenticationSchemeBuilder {
	return b.SetField(AuthenticationSchemeDocumentationURIKey, in)
}
func (b *AuthenticationSchemeBuilder) Name(in string) *AuthenticationSchemeBuilder {
	return b.SetField(AuthenticationSchemeNameKey, in)
}
func (b *AuthenticationSchemeBuilder) SpecURI(in string) *AuthenticationSchemeBuilder {
	return b.SetField(AuthenticationSchemeSpecURIKey, in)
}
func (b *AuthenticationSchemeBuilder) Type(in AuthenticationSchemeType) *AuthenticationSchemeBuilder {
	return b.SetField(AuthenticationSchemeTypeKey, in)
}

// SetField sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *AuthenticationSchemeBuilder) SetField(name string, value interface{}) *AuthenticationSchemeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(name, value); err != nil {
		b.err = err
	}
	return b
}
func (b *AuthenticationSchemeBuilder) Build() (*AuthenticationScheme, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	if b.object.description == nil {
		return nil, fmt.Errorf("required field 'Description' not initialized")
	}
	if b.object.name == nil {
		return nil, fmt.Errorf("required field 'Name' not initialized")
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}
func (b *AuthenticationSchemeBuilder) MustBuild() *AuthenticationScheme {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *AuthenticationSchemeBuilder) From(in *AuthenticationScheme) *AuthenticationSchemeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	var cloned AuthenticationScheme
	if err := in.Clone(&cloned); err != nil {
		b.err = err
		return b
	}

	b.object = &cloned
	return b
}

// AsMap returns the resource as a Go map
func (v *AuthenticationScheme) AsMap(m map[string]interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	for _, key := range v.Keys() {
		var val interface{}
		if err := v.getNoLock(key, &val, false); err != nil {
			m[key] = val
		}
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *AuthenticationScheme) GetExtension(name, uri string, dst interface{}) error {
	if uri == "" {
		return v.Get(name, dst)
	}
	var ext interface{}
	if err := v.Get(uri, &ext); err != nil {
		return fmt.Errorf(`failed to fetch extension %q: %w`, uri, err)
	}

	getter, ok := ext.(interface {
		Get(string, interface{}) error
	})
	if !ok {
		return fmt.Errorf(`extension does not implement Get(string, interface{}) error`)
	}
	return getter.Get(name, dst)
}

func (*AuthenticationScheme) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
	// we can get an instance of the resource object
	if rx, ok := registry.LookupByURI(name); ok {
		if err := dec.Decode(&rx); err != nil {
			return fmt.Errorf(`failed to decode value for key %q: %w`, name, err)
		}
		if err := blackmagic.AssignIfCompatible(dst, rx); err != nil {
			return err
		}
	} else {
		if err := dec.Decode(dst); err != nil {
			return fmt.Errorf(`failed to decode value for key %q: %w`, name, err)
		}
	}
	return nil
}

func (b *Builder) AuthenticationScheme() *AuthenticationSchemeBuilder {
	return &AuthenticationSchemeBuilder{}
}
