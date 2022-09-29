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
	Register("GenericSupport", "", GenericSupport{})
	RegisterBuilder("GenericSupport", "", GenericSupportBuilder{})
}

type GenericSupport struct {
	mu        sync.RWMutex
	supported *bool
	extra     map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	GenericSupportSupportedKey = "supported"
)

// Get retrieves the value associated with a key
func (v *GenericSupport) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case GenericSupportSupportedKey:
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
func (v *GenericSupport) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case GenericSupportSupportedKey:
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

func (v *GenericSupport) HasSupported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.supported != nil
}

func (v *GenericSupport) Supported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.supported; val != nil {
		return *val
	}
	return false
}

// Remove removes the value associated with a key
func (v *GenericSupport) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case GenericSupportSupportedKey:
		v.supported = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *GenericSupport) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 1)
	if val := v.supported; val != nil {
		pairs = append(pairs, &fieldPair{Name: GenericSupportSupportedKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

func (v *GenericSupport) Clone() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &GenericSupport{
		supported: v.supported,
	}
}

// MarshalJSON serializes GenericSupport into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *GenericSupport) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into GenericSupport.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *GenericSupport) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
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
			case GenericSupportSupportedKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GenericSupportSupportedKey, err)
				}
				v.supported = &val
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

type GenericSupportBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *GenericSupport
}

// NewGenericSupportBuilder creates a new GenericSupportBuilder instance.
// GenericSupportBuilder is safe to be used uninitialized as well.
func NewGenericSupportBuilder() *GenericSupportBuilder {
	return &GenericSupportBuilder{}
}

func (b *GenericSupportBuilder) initialize() {
	b.err = nil
	b.object = &GenericSupport{}
}
func (b *GenericSupportBuilder) Supported(in bool) *GenericSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(GenericSupportSupportedKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *GenericSupportBuilder) Build() (*GenericSupport, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	if b.object.supported == nil {
		return nil, fmt.Errorf("required field 'Supported' not initialized")
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *GenericSupportBuilder) MustBuild() *GenericSupport {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *GenericSupportBuilder) From(in *GenericSupport) *GenericSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *GenericSupport) AsMap(dst map[string]interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *GenericSupport) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) GenericSupport() *GenericSupportBuilder {
	return &GenericSupportBuilder{}
}
