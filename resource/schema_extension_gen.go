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
	Register("SchemaExtension", "", SchemaExtension{})
	RegisterBuilder("SchemaExtension", "", SchemaExtensionBuilder{})
}

type SchemaExtension struct {
	mu       sync.RWMutex
	schema   *string
	required *bool
	extra    map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	SchemaExtensionSchemaKey   = "schema"
	SchemaExtensionRequiredKey = "required"
)

// Get retrieves the value associated with a key
func (v *SchemaExtension) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case SchemaExtensionSchemaKey:
		if val := v.schema; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaExtensionRequiredKey:
		if val := v.required; val != nil {
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
func (v *SchemaExtension) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case SchemaExtensionSchemaKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field schema, got %T`, value)
		}
		v.schema = &converted
	case SchemaExtensionRequiredKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field required, got %T`, value)
		}
		v.required = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *SchemaExtension) HasSchema() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schema != nil
}

func (v *SchemaExtension) HasRequired() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.required != nil
}

func (v *SchemaExtension) Schema() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schema; val != nil {
		return *val
	}
	return ""
}

func (v *SchemaExtension) Required() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.required; val != nil {
		return *val
	}
	return false
}

// Remove removes the value associated with a key
func (v *SchemaExtension) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case SchemaExtensionSchemaKey:
		v.schema = nil
	case SchemaExtensionRequiredKey:
		v.required = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *SchemaExtension) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 2)
	if val := v.schema; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaExtensionSchemaKey, Value: *val})
	}
	if val := v.required; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaExtensionRequiredKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

func (v *SchemaExtension) Clone() *SchemaExtension {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &SchemaExtension{
		schema:   v.schema,
		required: v.required,
	}
}

// MarshalJSON serializes SchemaExtension into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *SchemaExtension) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into SchemaExtension.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *SchemaExtension) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.schema = nil
	v.required = nil

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
			case SchemaExtensionSchemaKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaExtensionSchemaKey, err)
				}
				v.schema = &val
			case SchemaExtensionRequiredKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaExtensionRequiredKey, err)
				}
				v.required = &val
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

type SchemaExtensionBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *SchemaExtension
}

// NewSchemaExtensionBuilder creates a new SchemaExtensionBuilder instance.
// SchemaExtensionBuilder is safe to be used uninitialized as well.
func NewSchemaExtensionBuilder() *SchemaExtensionBuilder {
	return &SchemaExtensionBuilder{}
}

func (b *SchemaExtensionBuilder) initialize() {
	b.err = nil
	b.object = &SchemaExtension{}
}
func (b *SchemaExtensionBuilder) Schema(in string) *SchemaExtensionBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaExtensionSchemaKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaExtensionBuilder) Required(in bool) *SchemaExtensionBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaExtensionRequiredKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaExtensionBuilder) Build() (*SchemaExtension, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err := b.err; err != nil {
		return nil, err
	}
	if b.object.schema == nil {
		return nil, fmt.Errorf("required field 'Schema' not initialized")
	}
	if b.object.required == nil {
		return nil, fmt.Errorf("required field 'Required' not initialized")
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *SchemaExtensionBuilder) MustBuild() *SchemaExtension {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *SchemaExtensionBuilder) From(in *SchemaExtension) *SchemaExtensionBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *SchemaExtension) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *SchemaExtension) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) SchemaExtension() *SchemaExtensionBuilder {
	return &SchemaExtensionBuilder{}
}
