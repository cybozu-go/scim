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
	Register("GroupMember", "", GroupMember{})
	RegisterBuilder("GroupMember", "", GroupMemberBuilder{})
}

type GroupMember struct {
	mu    sync.RWMutex
	value *string
	ref   *string
	typ   *string
	extra map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	GroupMemberValueKey     = "value"
	GroupMemberReferenceKey = "$ref"
	GroupMemberTypeKey      = "type"
)

// Get retrieves the value associated with a key
func (v *GroupMember) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *GroupMember) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case GroupMemberValueKey:
		if val := v.value; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GroupMemberReferenceKey:
		if val := v.ref; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GroupMemberTypeKey:
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
func (v *GroupMember) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case GroupMemberValueKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field value, got %T`, value)
		}
		v.value = &converted
	case GroupMemberReferenceKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field $ref, got %T`, value)
		}
		v.ref = &converted
	case GroupMemberTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field type, got %T`, value)
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
func (v *GroupMember) Has(name string) bool {
	switch name {
	case GroupMemberValueKey:
		return v.value != nil
	case GroupMemberReferenceKey:
		return v.ref != nil
	case GroupMemberTypeKey:
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
func (v *GroupMember) Keys() []string {
	keys := make([]string, 0, 3)
	if v.value != nil {
		keys = append(keys, GroupMemberValueKey)
	}
	if v.ref != nil {
		keys = append(keys, GroupMemberReferenceKey)
	}
	if v.typ != nil {
		keys = append(keys, GroupMemberTypeKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasValue returns true if the field `value` has been populated
func (v *GroupMember) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

// HasReference returns true if the field `$ref` has been populated
func (v *GroupMember) HasReference() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ref != nil
}

// HasType returns true if the field `type` has been populated
func (v *GroupMember) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *GroupMember) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.value; val != nil {
		return *val
	}
	return ""
}

func (v *GroupMember) Reference() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.ref; val != nil {
		return *val
	}
	return ""
}

func (v *GroupMember) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *GroupMember) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case GroupMemberValueKey:
		v.value = nil
	case GroupMemberReferenceKey:
		v.ref = nil
	case GroupMemberTypeKey:
		v.typ = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *GroupMember) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &GroupMember{
		value: v.value,
		ref:   v.ref,
		typ:   v.typ,
		extra: extra,
	})
}

// MarshalJSON serializes GroupMember into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *GroupMember) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into GroupMember.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *GroupMember) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.value = nil
	v.ref = nil
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
			case GroupMemberValueKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupMemberValueKey, err)
				}
				v.value = &val
			case GroupMemberReferenceKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupMemberReferenceKey, err)
				}
				v.ref = &val
			case GroupMemberTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupMemberTypeKey, err)
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

type GroupMemberBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *GroupMember
}

// NewGroupMemberBuilder creates a new GroupMemberBuilder instance.
// GroupMemberBuilder is safe to be used uninitialized as well.
func NewGroupMemberBuilder() *GroupMemberBuilder {
	return &GroupMemberBuilder{}
}

func (b *GroupMemberBuilder) initialize() {
	b.err = nil
	b.object = &GroupMember{}
}
func (b *GroupMemberBuilder) Value(in string) *GroupMemberBuilder {
	return b.SetField(GroupMemberValueKey, in)
}
func (b *GroupMemberBuilder) Reference(in string) *GroupMemberBuilder {
	return b.SetField(GroupMemberReferenceKey, in)
}
func (b *GroupMemberBuilder) Type(in string) *GroupMemberBuilder {
	return b.SetField(GroupMemberTypeKey, in)
}

// SetField sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *GroupMemberBuilder) SetField(name string, value interface{}) *GroupMemberBuilder {
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
func (b *GroupMemberBuilder) Build() (*GroupMember, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	if b.object.value == nil {
		return nil, fmt.Errorf("required field 'Value' not initialized")
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}
func (b *GroupMemberBuilder) MustBuild() *GroupMember {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *GroupMemberBuilder) From(in *GroupMember) *GroupMemberBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	var cloned GroupMember
	if err := in.Clone(&cloned); err != nil {
		b.err = err
		return b
	}

	b.object = &cloned
	return b
}

// AsMap returns the resource as a Go map
func (v *GroupMember) AsMap(m map[string]interface{}) error {
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
func (v *GroupMember) GetExtension(name, uri string, dst interface{}) error {
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

func (*GroupMember) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
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

func (b *Builder) GroupMember() *GroupMemberBuilder {
	return &GroupMemberBuilder{}
}
