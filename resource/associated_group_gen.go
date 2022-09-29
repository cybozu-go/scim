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
	Register("AssociatedGroup", "", AssociatedGroup{})
	RegisterBuilder("AssociatedGroup", "", AssociatedGroupBuilder{})
}

type AssociatedGroup struct {
	mu      sync.RWMutex
	display *string
	ref     *string
	typ     *string
	value   *string
	extra   map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	AssociatedGroupDisplayKey   = "display"
	AssociatedGroupReferenceKey = "$ref"
	AssociatedGroupTypeKey      = "type"
	AssociatedGroupValueKey     = "value"
)

// Get retrieves the value associated with a key
func (v *AssociatedGroup) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case AssociatedGroupDisplayKey:
		if val := v.display; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case AssociatedGroupReferenceKey:
		if val := v.ref; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case AssociatedGroupTypeKey:
		if val := v.typ; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case AssociatedGroupValueKey:
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
func (v *AssociatedGroup) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case AssociatedGroupDisplayKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field display, got %T`, value)
		}
		v.display = &converted
	case AssociatedGroupReferenceKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field $ref, got %T`, value)
		}
		v.ref = &converted
	case AssociatedGroupTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field type, got %T`, value)
		}
		v.typ = &converted
	case AssociatedGroupValueKey:
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
func (v *AssociatedGroup) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *AssociatedGroup) HasReference() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ref != nil
}

func (v *AssociatedGroup) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *AssociatedGroup) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *AssociatedGroup) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.display; val != nil {
		return *val
	}
	return ""
}

func (v *AssociatedGroup) Reference() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.ref; val != nil {
		return *val
	}
	return ""
}

func (v *AssociatedGroup) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return ""
}

func (v *AssociatedGroup) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.value; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *AssociatedGroup) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case AssociatedGroupDisplayKey:
		v.display = nil
	case AssociatedGroupReferenceKey:
		v.ref = nil
	case AssociatedGroupTypeKey:
		v.typ = nil
	case AssociatedGroupValueKey:
		v.value = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *AssociatedGroup) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 4)
	if val := v.display; val != nil {
		pairs = append(pairs, &fieldPair{Name: AssociatedGroupDisplayKey, Value: *val})
	}
	if val := v.ref; val != nil {
		pairs = append(pairs, &fieldPair{Name: AssociatedGroupReferenceKey, Value: *val})
	}
	if val := v.typ; val != nil {
		pairs = append(pairs, &fieldPair{Name: AssociatedGroupTypeKey, Value: *val})
	}
	if val := v.value; val != nil {
		pairs = append(pairs, &fieldPair{Name: AssociatedGroupValueKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

func (v *AssociatedGroup) Clone() *AssociatedGroup {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &AssociatedGroup{
		display: v.display,
		ref:     v.ref,
		typ:     v.typ,
		value:   v.value,
	}
}

// MarshalJSON serializes AssociatedGroup into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *AssociatedGroup) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into AssociatedGroup.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *AssociatedGroup) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.display = nil
	v.ref = nil
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
			case AssociatedGroupDisplayKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AssociatedGroupDisplayKey, err)
				}
				v.display = &val
			case AssociatedGroupReferenceKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AssociatedGroupReferenceKey, err)
				}
				v.ref = &val
			case AssociatedGroupTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AssociatedGroupTypeKey, err)
				}
				v.typ = &val
			case AssociatedGroupValueKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AssociatedGroupValueKey, err)
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

type AssociatedGroupBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *AssociatedGroup
}

// NewAssociatedGroupBuilder creates a new AssociatedGroupBuilder instance.
// AssociatedGroupBuilder is safe to be used uninitialized as well.
func NewAssociatedGroupBuilder() *AssociatedGroupBuilder {
	return &AssociatedGroupBuilder{}
}

func (b *AssociatedGroupBuilder) initialize() {
	b.err = nil
	b.object = &AssociatedGroup{}
}
func (b *AssociatedGroupBuilder) Display(in string) *AssociatedGroupBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(AssociatedGroupDisplayKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *AssociatedGroupBuilder) Reference(in string) *AssociatedGroupBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(AssociatedGroupReferenceKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *AssociatedGroupBuilder) Type(in string) *AssociatedGroupBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(AssociatedGroupTypeKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *AssociatedGroupBuilder) Value(in string) *AssociatedGroupBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(AssociatedGroupValueKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *AssociatedGroupBuilder) Build() (*AssociatedGroup, error) {
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

func (b *AssociatedGroupBuilder) MustBuild() *AssociatedGroup {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *AssociatedGroupBuilder) From(in *AssociatedGroup) *AssociatedGroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *AssociatedGroup) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *AssociatedGroup) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) AssociatedGroup() *AssociatedGroupBuilder {
	return &AssociatedGroupBuilder{}
}
