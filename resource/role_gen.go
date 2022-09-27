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
	Register("Role", "", Role{})
}

type Role struct {
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
	RoleDisplayKey = "display"
	RolePrimaryKey = "primary"
	RoleTypeKey    = "type"
	RoleValueKey   = "value"
)

// Get retrieves the value associated with a key
func (v *Role) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case RoleDisplayKey:
		if val := v.display; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case RolePrimaryKey:
		if val := v.primary; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case RoleTypeKey:
		if val := v.typ; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case RoleValueKey:
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
func (v *Role) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case RoleDisplayKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field display, got %T`, value)
		}
		v.display = &converted
	case RolePrimaryKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field primary, got %T`, value)
		}
		v.primary = &converted
	case RoleTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field type, got %T`, value)
		}
		v.typ = &converted
	case RoleValueKey:
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
func (v *Role) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *Role) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *Role) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *Role) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *Role) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.display; val != nil {
		return *val
	}
	return ""
}

func (v *Role) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.primary; val != nil {
		return *val
	}
	return false
}

func (v *Role) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return ""
}

func (v *Role) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.value; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *Role) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case RoleDisplayKey:
		v.display = nil
	case RolePrimaryKey:
		v.primary = nil
	case RoleTypeKey:
		v.typ = nil
	case RoleValueKey:
		v.value = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *Role) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 4)
	if val := v.display; val != nil {
		pairs = append(pairs, &fieldPair{Name: RoleDisplayKey, Value: *val})
	}
	if val := v.primary; val != nil {
		pairs = append(pairs, &fieldPair{Name: RolePrimaryKey, Value: *val})
	}
	if val := v.typ; val != nil {
		pairs = append(pairs, &fieldPair{Name: RoleTypeKey, Value: *val})
	}
	if val := v.value; val != nil {
		pairs = append(pairs, &fieldPair{Name: RoleValueKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes Role into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *Role) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into Role.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *Role) UnmarshalJSON(data []byte) error {
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
			case RoleDisplayKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, RoleDisplayKey, err)
				}
				v.display = &val
			case RolePrimaryKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, RolePrimaryKey, err)
				}
				v.primary = &val
			case RoleTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, RoleTypeKey, err)
				}
				v.typ = &val
			case RoleValueKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, RoleValueKey, err)
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

type RoleBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *Role
}

// NewRoleBuilder creates a new RoleBuilder instance.
// RoleBuilder is safe to be used uninitialized as well.
func NewRoleBuilder() *RoleBuilder {
	return &RoleBuilder{}
}

func (b *RoleBuilder) initialize() {
	b.err = nil
	b.object = &Role{}
}
func (b *RoleBuilder) Display(in string) *RoleBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(RoleDisplayKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *RoleBuilder) Primary(in bool) *RoleBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(RolePrimaryKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *RoleBuilder) Type(in string) *RoleBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(RoleTypeKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *RoleBuilder) Value(in string) *RoleBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(RoleValueKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *RoleBuilder) Build() (*Role, error) {
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

func (b *RoleBuilder) MustBuild() *Role {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *RoleBuilder) From(in *Role) *RoleBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *Role) Clone() *Role {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &Role{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *Role) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *Role) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) Role() *RoleBuilder {
	return &RoleBuilder{}
}
