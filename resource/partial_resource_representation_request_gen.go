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
	Register("PartialResourceRepresentationRequest", "", PartialResourceRepresentationRequest{})
	RegisterBuilder("PartialResourceRepresentationRequest", "", PartialResourceRepresentationRequestBuilder{})
}

type PartialResourceRepresentationRequest struct {
	mu                 sync.RWMutex
	attributes         []string
	excludedAttributes []string
	extra              map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	PartialResourceRepresentationRequestAttributesKey         = "attributes"
	PartialResourceRepresentationRequestExcludedAttributesKey = "excludedAttributes"
)

// Get retrieves the value associated with a key
func (v *PartialResourceRepresentationRequest) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *PartialResourceRepresentationRequest) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case PartialResourceRepresentationRequestAttributesKey:
		if val := v.attributes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case PartialResourceRepresentationRequestExcludedAttributesKey:
		if val := v.excludedAttributes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
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
func (v *PartialResourceRepresentationRequest) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case PartialResourceRepresentationRequestAttributesKey:
		converted, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected value of type []string for field attributes, got %T`, value)
		}
		v.attributes = converted
	case PartialResourceRepresentationRequestExcludedAttributesKey:
		converted, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected value of type []string for field excludedAttributes, got %T`, value)
		}
		v.excludedAttributes = converted
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
func (v *PartialResourceRepresentationRequest) Has(name string) bool {
	switch name {
	case PartialResourceRepresentationRequestAttributesKey:
		return v.attributes != nil
	case PartialResourceRepresentationRequestExcludedAttributesKey:
		return v.excludedAttributes != nil
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
func (v *PartialResourceRepresentationRequest) Keys() []string {
	keys := make([]string, 0, 2)
	if v.attributes != nil {
		keys = append(keys, PartialResourceRepresentationRequestAttributesKey)
	}
	if v.excludedAttributes != nil {
		keys = append(keys, PartialResourceRepresentationRequestExcludedAttributesKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasAttributes returns true if the field `attributes` has been populated
func (v *PartialResourceRepresentationRequest) HasAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes != nil
}

// HasExcludedAttributes returns true if the field `excludedAttributes` has been populated
func (v *PartialResourceRepresentationRequest) HasExcludedAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.excludedAttributes != nil
}

func (v *PartialResourceRepresentationRequest) Attributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.attributes; val != nil {
		return val
	}
	return nil
}

func (v *PartialResourceRepresentationRequest) ExcludedAttributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.excludedAttributes; val != nil {
		return val
	}
	return nil
}

// Remove removes the value associated with a key
func (v *PartialResourceRepresentationRequest) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case PartialResourceRepresentationRequestAttributesKey:
		v.attributes = nil
	case PartialResourceRepresentationRequestExcludedAttributesKey:
		v.excludedAttributes = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *PartialResourceRepresentationRequest) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &PartialResourceRepresentationRequest{
		attributes:         v.attributes,
		excludedAttributes: v.excludedAttributes,
		extra:              extra,
	})
}

// MarshalJSON serializes PartialResourceRepresentationRequest into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *PartialResourceRepresentationRequest) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into PartialResourceRepresentationRequest.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *PartialResourceRepresentationRequest) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.attributes = nil
	v.excludedAttributes = nil

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
			case PartialResourceRepresentationRequestAttributesKey:
				var val []string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PartialResourceRepresentationRequestAttributesKey, err)
				}
				v.attributes = val
			case PartialResourceRepresentationRequestExcludedAttributesKey:
				var val []string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PartialResourceRepresentationRequestExcludedAttributesKey, err)
				}
				v.excludedAttributes = val
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

type PartialResourceRepresentationRequestBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *PartialResourceRepresentationRequest
}

// NewPartialResourceRepresentationRequestBuilder creates a new PartialResourceRepresentationRequestBuilder instance.
// PartialResourceRepresentationRequestBuilder is safe to be used uninitialized as well.
func NewPartialResourceRepresentationRequestBuilder() *PartialResourceRepresentationRequestBuilder {
	return &PartialResourceRepresentationRequestBuilder{}
}

func (b *PartialResourceRepresentationRequestBuilder) initialize() {
	b.err = nil
	b.object = &PartialResourceRepresentationRequest{}
}
func (b *PartialResourceRepresentationRequestBuilder) Attributes(in ...string) *PartialResourceRepresentationRequestBuilder {
	return b.SetField(PartialResourceRepresentationRequestAttributesKey, in)
}
func (b *PartialResourceRepresentationRequestBuilder) ExcludedAttributes(in ...string) *PartialResourceRepresentationRequestBuilder {
	return b.SetField(PartialResourceRepresentationRequestExcludedAttributesKey, in)
}

// SetField sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *PartialResourceRepresentationRequestBuilder) SetField(name string, value interface{}) *PartialResourceRepresentationRequestBuilder {
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
func (b *PartialResourceRepresentationRequestBuilder) Build() (*PartialResourceRepresentationRequest, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}
func (b *PartialResourceRepresentationRequestBuilder) MustBuild() *PartialResourceRepresentationRequest {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *PartialResourceRepresentationRequestBuilder) From(in *PartialResourceRepresentationRequest) *PartialResourceRepresentationRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	var cloned PartialResourceRepresentationRequest
	if err := in.Clone(&cloned); err != nil {
		b.err = err
		return b
	}

	b.object = &cloned
	return b
}

// AsMap returns the resource as a Go map
func (v *PartialResourceRepresentationRequest) AsMap(m map[string]interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	for _, key := range v.Keys() {
		var val interface{}
		if err := v.getNoLock(key, &val, false); err != nil {
			return fmt.Errorf(`failed to retrieve value for key %q: %w`, key, err)
		}
		m[key] = val
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *PartialResourceRepresentationRequest) GetExtension(name, uri string, dst interface{}) error {
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

func (*PartialResourceRepresentationRequest) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
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

func (b *Builder) PartialResourceRepresentationRequest() *PartialResourceRepresentationRequestBuilder {
	return &PartialResourceRepresentationRequestBuilder{}
}
