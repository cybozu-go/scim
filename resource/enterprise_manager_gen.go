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

type EnterpriseManager struct {
	mu          sync.RWMutex
	displayName *string
	id          *string
	ref         *string
	extra       map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	EnterpriseManagerDisplayNameKey = "displayName"
	EnterpriseManagerIDKey          = "id"
	EnterpriseManagerReferenceKey   = "$ref"
)

// Get retrieves the value associated with a key
func (v *EnterpriseManager) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case EnterpriseManagerDisplayNameKey:
		if val := v.displayName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EnterpriseManagerIDKey:
		if val := v.id; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EnterpriseManagerReferenceKey:
		if val := v.ref; val != nil {
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
func (v *EnterpriseManager) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case EnterpriseManagerDisplayNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field displayName, got %T`, value)
		}
		v.displayName = &converted
	case EnterpriseManagerIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field id, got %T`, value)
		}
		v.id = &converted
	case EnterpriseManagerReferenceKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field $ref, got %T`, value)
		}
		v.ref = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *EnterpriseManager) HasDisplayName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.displayName != nil
}

func (v *EnterpriseManager) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *EnterpriseManager) HasReference() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ref != nil
}

func (v *EnterpriseManager) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.displayName; val != nil {
		return *val
	}
	return ""
}

func (v *EnterpriseManager) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.id; val != nil {
		return *val
	}
	return ""
}

func (v *EnterpriseManager) Reference() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.ref; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *EnterpriseManager) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case EnterpriseManagerDisplayNameKey:
		v.displayName = nil
	case EnterpriseManagerIDKey:
		v.id = nil
	case EnterpriseManagerReferenceKey:
		v.ref = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *EnterpriseManager) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 3)
	if val := v.displayName; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseManagerDisplayNameKey, Value: *val})
	}
	if val := v.id; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseManagerIDKey, Value: *val})
	}
	if val := v.ref; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseManagerReferenceKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes EnterpriseManager into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *EnterpriseManager) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into EnterpriseManager.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *EnterpriseManager) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.displayName = nil
	v.id = nil
	v.ref = nil

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
			case EnterpriseManagerDisplayNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseManagerDisplayNameKey, err)
				}
				v.displayName = &val
			case EnterpriseManagerIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseManagerIDKey, err)
				}
				v.id = &val
			case EnterpriseManagerReferenceKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseManagerReferenceKey, err)
				}
				v.ref = &val
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

type EnterpriseManagerBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *EnterpriseManager
}

// NewEnterpriseManagerBuilder creates a new EnterpriseManagerBuilder instance.
// EnterpriseManagerBuilder is safe to be used uninitialized as well.
func NewEnterpriseManagerBuilder() *EnterpriseManagerBuilder {
	return &EnterpriseManagerBuilder{}
}

func (b *EnterpriseManagerBuilder) initialize() {
	b.err = nil
	b.object = &EnterpriseManager{}
}
func (b *EnterpriseManagerBuilder) DisplayName(in string) *EnterpriseManagerBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(EnterpriseManagerDisplayNameKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *EnterpriseManagerBuilder) ID(in string) *EnterpriseManagerBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(EnterpriseManagerIDKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *EnterpriseManagerBuilder) Reference(in string) *EnterpriseManagerBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(EnterpriseManagerReferenceKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseManagerBuilder) Build() (*EnterpriseManager, error) {
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

func (b *EnterpriseManagerBuilder) MustBuild() *EnterpriseManager {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *EnterpriseManagerBuilder) From(in *EnterpriseManager) *EnterpriseManagerBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.object = in.Clone()
	return b
}

func (v *EnterpriseManager) Clone() *EnterpriseManager {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &EnterpriseManager{
		displayName: v.displayName,
		id:          v.id,
		ref:         v.ref,
	}
}

func (v *EnterpriseManager) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *EnterpriseManager) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) EnterpriseManager() *EnterpriseManagerBuilder {
	return &EnterpriseManagerBuilder{}
}
