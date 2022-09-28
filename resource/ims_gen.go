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
	Register("IMS", "", IMS{})
	RegisterBuilder("IMS", "", IMSBuilder{})
}

type IMS struct {
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
	IMSDisplayKey = "display"
	IMSPrimaryKey = "primary"
	IMSTypeKey    = "type"
	IMSValueKey   = "value"
)

// Get retrieves the value associated with a key
func (v *IMS) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case IMSDisplayKey:
		if val := v.display; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case IMSPrimaryKey:
		if val := v.primary; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case IMSTypeKey:
		if val := v.typ; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case IMSValueKey:
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
func (v *IMS) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case IMSDisplayKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field display, got %T`, value)
		}
		v.display = &converted
	case IMSPrimaryKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field primary, got %T`, value)
		}
		v.primary = &converted
	case IMSTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field type, got %T`, value)
		}
		v.typ = &converted
	case IMSValueKey:
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
func (v *IMS) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *IMS) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *IMS) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *IMS) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *IMS) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.display; val != nil {
		return *val
	}
	return ""
}

func (v *IMS) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.primary; val != nil {
		return *val
	}
	return false
}

func (v *IMS) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return ""
}

func (v *IMS) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.value; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *IMS) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case IMSDisplayKey:
		v.display = nil
	case IMSPrimaryKey:
		v.primary = nil
	case IMSTypeKey:
		v.typ = nil
	case IMSValueKey:
		v.value = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *IMS) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 4)
	if val := v.display; val != nil {
		pairs = append(pairs, &fieldPair{Name: IMSDisplayKey, Value: *val})
	}
	if val := v.primary; val != nil {
		pairs = append(pairs, &fieldPair{Name: IMSPrimaryKey, Value: *val})
	}
	if val := v.typ; val != nil {
		pairs = append(pairs, &fieldPair{Name: IMSTypeKey, Value: *val})
	}
	if val := v.value; val != nil {
		pairs = append(pairs, &fieldPair{Name: IMSValueKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes IMS into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *IMS) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into IMS.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *IMS) UnmarshalJSON(data []byte) error {
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
			case IMSDisplayKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IMSDisplayKey, err)
				}
				v.display = &val
			case IMSPrimaryKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IMSPrimaryKey, err)
				}
				v.primary = &val
			case IMSTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IMSTypeKey, err)
				}
				v.typ = &val
			case IMSValueKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IMSValueKey, err)
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

type IMSBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *IMS
}

// NewIMSBuilder creates a new IMSBuilder instance.
// IMSBuilder is safe to be used uninitialized as well.
func NewIMSBuilder() *IMSBuilder {
	return &IMSBuilder{}
}

func (b *IMSBuilder) initialize() {
	b.err = nil
	b.object = &IMS{}
}
func (b *IMSBuilder) Display(in string) *IMSBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(IMSDisplayKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *IMSBuilder) Primary(in bool) *IMSBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(IMSPrimaryKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *IMSBuilder) Type(in string) *IMSBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(IMSTypeKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *IMSBuilder) Value(in string) *IMSBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(IMSValueKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *IMSBuilder) Build() (*IMS, error) {
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

func (b *IMSBuilder) MustBuild() *IMS {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *IMSBuilder) From(in *IMS) *IMSBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *IMS) Clone() *IMS {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &IMS{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *IMS) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *IMS) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) IMS() *IMSBuilder {
	return &IMSBuilder{}
}
