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

func (v *Photo) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *Photo) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *Photo) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

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

func (v *Photo) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 4)
	if val := v.display; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhotoDisplayKey, Value: *val})
	}
	if val := v.primary; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhotoPrimaryKey, Value: *val})
	}
	if val := v.typ; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhotoTypeKey, Value: *val})
	}
	if val := v.value; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhotoValueKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

func (v *Photo) Clone() *Photo {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &Photo{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

// MarshalJSON serializes Photo into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *Photo) MarshalJSON() ([]byte, error) {
	pairs := v.makePairs()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, pair := range pairs {
		if i > 0 {
			buf.WriteByte(',')
		}
		if err := enc.Encode(pair.Name); err != nil {
			return nil, fmt.Errorf(`failed to encode map key name: %w`, err)
		}
		buf.WriteByte(':')
		if err := enc.Encode(pair.Value); err != nil {
			return nil, fmt.Errorf(`failed to encode map value for %q: %w`, pair.Name, err)
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
				if err := extraFieldsDecoder(tok, dec, &val); err != nil {
					return err
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
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhotoDisplayKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *PhotoBuilder) Primary(in bool) *PhotoBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhotoPrimaryKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *PhotoBuilder) Type(in string) *PhotoBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhotoTypeKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *PhotoBuilder) Value(in string) *PhotoBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhotoValueKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *PhotoBuilder) Build() (*Photo, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err := b.err; err != nil {
		return nil, err
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
	b.object = in.Clone()
	return b
}

func (v *Photo) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
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

func (b *Builder) Photo() *PhotoBuilder {
	return &PhotoBuilder{}
}
