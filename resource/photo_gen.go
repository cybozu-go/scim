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
	Register("Photo", "", Photo{})
	RegisterBuilder("Photo", "", PhotoBuilder{})
}

type Photo struct {
	mu      sync.RWMutex
	display *string
	primary *bool
	typ     *string
	value   *string
	extra   map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	PhotoDisplayKey = "display"
	PhotoPrimaryKey = "primary"
	PhotoTypeKey    = "type"
	PhotoValueKey   = "value"
)

// Get retrieves the value associated with a key
func (v *Photo) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *Photo) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case PhotoDisplayKey:
		if val := v.display; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PhotoPrimaryKey:
		if val := v.primary; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PhotoTypeKey:
		if val := v.typ; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PhotoValueKey:
		if val := v.value; val != nil {
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
func (v *Photo) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case PhotoDisplayKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field display, got %T`, value)
		}
		v.display = &converted
	case PhotoPrimaryKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field primary, got %T`, value)
		}
		v.primary = &converted
	case PhotoTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field type, got %T`, value)
		}
		v.typ = &converted
	case PhotoValueKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field value, got %T`, value)
		}
		v.value = &converted
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
func (v *Photo) Has(name string) bool {
	switch name {
	case PhotoDisplayKey:
		return v.display != nil
	case PhotoPrimaryKey:
		return v.primary != nil
	case PhotoTypeKey:
		return v.typ != nil
	case PhotoValueKey:
		return v.value != nil
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
func (v *Photo) Keys() []string {
	keys := make([]string, 0, 4)
	if v.display != nil {
		keys = append(keys, PhotoDisplayKey)
	}
	if v.primary != nil {
		keys = append(keys, PhotoPrimaryKey)
	}
	if v.typ != nil {
		keys = append(keys, PhotoTypeKey)
	}
	if v.value != nil {
		keys = append(keys, PhotoValueKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasDisplay returns true if the field `display` has been populated
func (v *Photo) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

// HasPrimary returns true if the field `primary` has been populated
func (v *Photo) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

// HasType returns true if the field `type` has been populated
func (v *Photo) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

// HasValue returns true if the field `value` has been populated
func (v *Photo) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *Photo) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.display; val != nil {
		return *val
	}
	return ""
}

func (v *Photo) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.primary; val != nil {
		return *val
	}
	return false
}

func (v *Photo) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return ""
}

func (v *Photo) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.value; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *Photo) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case PhotoDisplayKey:
		v.display = nil
	case PhotoPrimaryKey:
		v.primary = nil
	case PhotoTypeKey:
		v.typ = nil
	case PhotoValueKey:
		v.value = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *Photo) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &Photo{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
		extra:   extra,
	})
}

// MarshalJSON serializes Photo into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *Photo) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into Photo.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *Photo) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.display = nil
	v.primary = nil
	v.typ = nil
	v.value = nil

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
			case PhotoDisplayKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhotoDisplayKey, err)
				}
				v.display = &val
			case PhotoPrimaryKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhotoPrimaryKey, err)
				}
				v.primary = &val
			case PhotoTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhotoTypeKey, err)
				}
				v.typ = &val
			case PhotoValueKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhotoValueKey, err)
				}
				v.value = &val
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

type PhotoBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *Photo
}

// NewPhotoBuilder creates a new PhotoBuilder instance.
// PhotoBuilder is safe to be used uninitialized as well.
func NewPhotoBuilder() *PhotoBuilder {
	return &PhotoBuilder{}
}

func (b *PhotoBuilder) initialize() {
	b.err = nil
	b.object = &Photo{}
}
func (b *PhotoBuilder) Display(in string) *PhotoBuilder {
	return b.SetField(PhotoDisplayKey, in)
}
func (b *PhotoBuilder) Primary(in bool) *PhotoBuilder {
	return b.SetField(PhotoPrimaryKey, in)
}
func (b *PhotoBuilder) Type(in string) *PhotoBuilder {
	return b.SetField(PhotoTypeKey, in)
}
func (b *PhotoBuilder) Value(in string) *PhotoBuilder {
	return b.SetField(PhotoValueKey, in)
}

// SetField sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *PhotoBuilder) SetField(name string, value interface{}) *PhotoBuilder {
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
func (b *PhotoBuilder) Build() (*Photo, error) {
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
func (b *PhotoBuilder) MustBuild() *Photo {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *PhotoBuilder) From(in *Photo) *PhotoBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	var cloned Photo
	if err := in.Clone(&cloned); err != nil {
		b.err = err
		return b
	}

	b.object = &cloned
	return b
}

// AsMap returns the resource as a Go map
func (v *Photo) AsMap(m map[string]interface{}) error {
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
func (v *Photo) GetExtension(name, uri string, dst interface{}) error {
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

func (*Photo) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
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

func (b *Builder) Photo() *PhotoBuilder {
	return &PhotoBuilder{}
}
