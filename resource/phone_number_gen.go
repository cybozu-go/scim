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

type PhoneNumber struct {
	mu      sync.RWMutex
	display *string
	primary *bool
	typ     *string
	value   *PhoneNumberValue
	extra   map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	PhoneNumberDisplayKey = "display"
	PhoneNumberPrimaryKey = "primary"
	PhoneNumberTypeKey    = "type"
	PhoneNumberValueKey   = "value"
)

// Get retrieves the value associated with a key
func (v *PhoneNumber) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case PhoneNumberDisplayKey:
		if val := v.display; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PhoneNumberPrimaryKey:
		if val := v.primary; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PhoneNumberTypeKey:
		if val := v.typ; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PhoneNumberValueKey:
		if val := v.value; val != nil {
			return blackmagic.AssignIfCompatible(dst, val.Get())
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
func (v *PhoneNumber) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case PhoneNumberDisplayKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field display, got %T`, value)
		}
		v.display = &converted
	case PhoneNumberPrimaryKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field primary, got %T`, value)
		}
		v.primary = &converted
	case PhoneNumberTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field type, got %T`, value)
		}
		v.typ = &converted
	case PhoneNumberValueKey:
		var object PhoneNumberValue
		if err := object.Accept(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.value = &object
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *PhoneNumber) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *PhoneNumber) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *PhoneNumber) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *PhoneNumber) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *PhoneNumber) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.display; val != nil {
		return *val
	}
	return ""
}

func (v *PhoneNumber) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.primary; val != nil {
		return *val
	}
	return false
}

func (v *PhoneNumber) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return ""
}

func (v *PhoneNumber) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.value; val != nil {
		return val.Get()
	}
	return ""
}

// Remove removes the value associated with a key
func (v *PhoneNumber) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case PhoneNumberDisplayKey:
		v.display = nil
	case PhoneNumberPrimaryKey:
		v.primary = nil
	case PhoneNumberTypeKey:
		v.typ = nil
	case PhoneNumberValueKey:
		v.value = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *PhoneNumber) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 4)
	if val := v.display; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhoneNumberDisplayKey, Value: *val})
	}
	if val := v.primary; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhoneNumberPrimaryKey, Value: *val})
	}
	if val := v.typ; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhoneNumberTypeKey, Value: *val})
	}
	if val := v.value; val != nil {
		pairs = append(pairs, &fieldPair{Name: PhoneNumberValueKey, Value: val.Get()})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes PhoneNumber into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *PhoneNumber) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into PhoneNumber.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *PhoneNumber) UnmarshalJSON(data []byte) error {
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
			case PhoneNumberDisplayKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhoneNumberDisplayKey, err)
				}
				v.display = &val
			case PhoneNumberPrimaryKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhoneNumberPrimaryKey, err)
				}
				v.primary = &val
			case PhoneNumberTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhoneNumberTypeKey, err)
				}
				v.typ = &val
			case PhoneNumberValueKey:
				var val PhoneNumberValue
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhoneNumberValueKey, err)
				}
				v.value = &val
			default:
				var val interface{}
				if err := dec.Decode(&val); err != nil {
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

type PhoneNumberBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *PhoneNumber
}

// NewPhoneNumberBuilder creates a new PhoneNumberBuilder instance.
// PhoneNumberBuilder is safe to be used uninitialized as well.
func NewPhoneNumberBuilder() *PhoneNumberBuilder {
	return &PhoneNumberBuilder{}
}

func (b *PhoneNumberBuilder) initialize() {
	b.err = nil
	b.object = &PhoneNumber{}
}
func (b *PhoneNumberBuilder) Display(in string) *PhoneNumberBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhoneNumberDisplayKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *PhoneNumberBuilder) Primary(in bool) *PhoneNumberBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhoneNumberPrimaryKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *PhoneNumberBuilder) Type(in string) *PhoneNumberBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhoneNumberTypeKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *PhoneNumberBuilder) Value(in string) *PhoneNumberBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(PhoneNumberValueKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *PhoneNumberBuilder) Build() (*PhoneNumber, error) {
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

func (b *PhoneNumberBuilder) MustBuild() *PhoneNumber {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *PhoneNumberBuilder) From(in *PhoneNumber) *PhoneNumberBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *PhoneNumber) Clone() *PhoneNumber {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &PhoneNumber{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *PhoneNumber) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *PhoneNumber) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) PhoneNumber() *PhoneNumberBuilder {
	return &PhoneNumberBuilder{}
}
