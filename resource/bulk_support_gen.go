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
func (v *BulkSupport) HasMaxOperations() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxOperations != nil
}

func (v *BulkSupport) HasMaxPayloadSize() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxPayloadSize != nil
}

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

func (v *BulkSupport) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 3)
	if val := v.maxOperations; val != nil {
		pairs = append(pairs, &fieldPair{Name: BulkSupportMaxOperationsKey, Value: *val})
	}
	if val := v.maxPayloadSize; val != nil {
		pairs = append(pairs, &fieldPair{Name: BulkSupportMaxPayloadSizeKey, Value: *val})
	}
	if val := v.supported; val != nil {
		pairs = append(pairs, &fieldPair{Name: BulkSupportSupportedKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes BulkSupport into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *BulkSupport) MarshalJSON() ([]byte, error) {
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
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(BulkSupportMaxOperationsKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *BulkSupportBuilder) MaxPayloadSize(in int) *BulkSupportBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(BulkSupportMaxPayloadSizeKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *BulkSupportBuilder) Supported(in bool) *BulkSupportBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(BulkSupportSupportedKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *BulkSupportBuilder) Build() (*BulkSupport, error) {
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
	b.object = in.Clone()
	return b
}

func (v *BulkSupport) Clone() *BulkSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &BulkSupport{
		maxOperations:  v.maxOperations,
		maxPayloadSize: v.maxPayloadSize,
		supported:      v.supported,
	}
}

func (v *BulkSupport) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
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

func (b *Builder) BulkSupport() *BulkSupportBuilder {
	return &BulkSupportBuilder{}
}
