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
	Register("BulkSupport", "", BulkSupport{})
	RegisterBuilder("BulkSupport", "", BulkSupportBuilder{})
}

type BulkSupport struct {
	mu             sync.RWMutex
	maxOperations  *int
	maxPayloadSize *int
	supported      *bool
	extra          map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	BulkSupportMaxOperationsKey  = "maxOperations"
	BulkSupportMaxPayloadSizeKey = "maxPayloadSize"
	BulkSupportSupportedKey      = "supported"
)

// Get retrieves the value associated with a key
func (v *BulkSupport) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *BulkSupport) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case BulkSupportMaxOperationsKey:
		if val := v.maxOperations; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case BulkSupportMaxPayloadSizeKey:
		if val := v.maxPayloadSize; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case BulkSupportSupportedKey:
		if val := v.supported; val != nil {
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
func (v *BulkSupport) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case BulkSupportMaxOperationsKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field maxOperations, got %T`, value)
		}
		v.maxOperations = &converted
	case BulkSupportMaxPayloadSizeKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field maxPayloadSize, got %T`, value)
		}
		v.maxPayloadSize = &converted
	case BulkSupportSupportedKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field supported, got %T`, value)
		}
		v.supported = &converted
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
func (v *BulkSupport) Has(name string) bool {
	switch name {
	case BulkSupportMaxOperationsKey:
		return v.maxOperations != nil
	case BulkSupportMaxPayloadSizeKey:
		return v.maxPayloadSize != nil
	case BulkSupportSupportedKey:
		return v.supported != nil
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
func (v *BulkSupport) Keys() []string {
	keys := make([]string, 0, 3)
	if v.maxOperations != nil {
		keys = append(keys, BulkSupportMaxOperationsKey)
	}
	if v.maxPayloadSize != nil {
		keys = append(keys, BulkSupportMaxPayloadSizeKey)
	}
	if v.supported != nil {
		keys = append(keys, BulkSupportSupportedKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasMaxOperations returns true if the field `maxOperations` has been populated
func (v *BulkSupport) HasMaxOperations() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxOperations != nil
}

// HasMaxPayloadSize returns true if the field `maxPayloadSize` has been populated
func (v *BulkSupport) HasMaxPayloadSize() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxPayloadSize != nil
}

// HasSupported returns true if the field `supported` has been populated
func (v *BulkSupport) HasSupported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.supported != nil
}

func (v *BulkSupport) MaxOperations() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.maxOperations; val != nil {
		return *val
	}
	return 0
}

func (v *BulkSupport) MaxPayloadSize() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.maxPayloadSize; val != nil {
		return *val
	}
	return 0
}

func (v *BulkSupport) Supported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.supported; val != nil {
		return *val
	}
	return false
}

// Remove removes the value associated with a key
func (v *BulkSupport) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case BulkSupportMaxOperationsKey:
		v.maxOperations = nil
	case BulkSupportMaxPayloadSizeKey:
		v.maxPayloadSize = nil
	case BulkSupportSupportedKey:
		v.supported = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *BulkSupport) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &BulkSupport{
		maxOperations:  v.maxOperations,
		maxPayloadSize: v.maxPayloadSize,
		supported:      v.supported,
		extra:          extra,
	})
}

// MarshalJSON serializes BulkSupport into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *BulkSupport) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into BulkSupport.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *BulkSupport) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.maxOperations = nil
	v.maxPayloadSize = nil
	v.supported = nil

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
			case BulkSupportMaxOperationsKey:
				var val int
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, BulkSupportMaxOperationsKey, err)
				}
				v.maxOperations = &val
			case BulkSupportMaxPayloadSizeKey:
				var val int
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, BulkSupportMaxPayloadSizeKey, err)
				}
				v.maxPayloadSize = &val
			case BulkSupportSupportedKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, BulkSupportSupportedKey, err)
				}
				v.supported = &val
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

type BulkSupportBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *BulkSupport
}

// NewBulkSupportBuilder creates a new BulkSupportBuilder instance.
// BulkSupportBuilder is safe to be used uninitialized as well.
func NewBulkSupportBuilder() *BulkSupportBuilder {
	return &BulkSupportBuilder{}
}

func (b *BulkSupportBuilder) initialize() {
	b.err = nil
	b.object = &BulkSupport{}
}
func (b *BulkSupportBuilder) MaxOperations(in int) *BulkSupportBuilder {
	return b.SetField(BulkSupportMaxOperationsKey, in)
}
func (b *BulkSupportBuilder) MaxPayloadSize(in int) *BulkSupportBuilder {
	return b.SetField(BulkSupportMaxPayloadSizeKey, in)
}
func (b *BulkSupportBuilder) Supported(in bool) *BulkSupportBuilder {
	return b.SetField(BulkSupportSupportedKey, in)
}

// SetField sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *BulkSupportBuilder) SetField(name string, value interface{}) *BulkSupportBuilder {
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
func (b *BulkSupportBuilder) Build() (*BulkSupport, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	if b.object.maxOperations == nil {
		return nil, fmt.Errorf("required field 'MaxOperations' not initialized")
	}
	if b.object.maxPayloadSize == nil {
		return nil, fmt.Errorf("required field 'MaxPayloadSize' not initialized")
	}
	if b.object.supported == nil {
		return nil, fmt.Errorf("required field 'Supported' not initialized")
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}
func (b *BulkSupportBuilder) MustBuild() *BulkSupport {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *BulkSupportBuilder) From(in *BulkSupport) *BulkSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	var cloned BulkSupport
	if err := in.Clone(&cloned); err != nil {
		b.err = err
		return b
	}

	b.object = &cloned
	return b
}

// AsMap returns the resource as a Go map
func (v *BulkSupport) AsMap(m map[string]interface{}) error {
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
func (v *BulkSupport) GetExtension(name, uri string, dst interface{}) error {
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

func (*BulkSupport) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
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

func (b *Builder) BulkSupport() *BulkSupportBuilder {
	return &BulkSupportBuilder{}
}
