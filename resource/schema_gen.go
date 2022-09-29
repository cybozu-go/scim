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
	Register("Schema", "", Schema{})
	RegisterBuilder("Schema", "", SchemaBuilder{})
}

// Schema represents a Schema resource as defined in the SCIM RFC
type Schema struct {
	mu          sync.RWMutex
	attributes  []*SchemaAttribute
	description *string
	id          *string
	name        *string
	extra       map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	SchemaAttributesKey  = "attributes"
	SchemaDescriptionKey = "description"
	SchemaIDKey          = "id"
	SchemaNameKey        = "name"
)

// Get retrieves the value associated with a key
func (v *Schema) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case SchemaAttributesKey:
		if val := v.attributes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case SchemaDescriptionKey:
		if val := v.description; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaIDKey:
		if val := v.id; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaNameKey:
		if val := v.name; val != nil {
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
func (v *Schema) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case SchemaAttributesKey:
		converted, ok := value.([]*SchemaAttribute)
		if !ok {
			return fmt.Errorf(`expected value of type []*SchemaAttribute for field attributes, got %T`, value)
		}
		v.attributes = converted
	case SchemaDescriptionKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field description, got %T`, value)
		}
		v.description = &converted
	case SchemaIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field id, got %T`, value)
		}
		v.id = &converted
	case SchemaNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field name, got %T`, value)
		}
		v.name = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}

func (v *Schema) HasAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes != nil
}

func (v *Schema) HasDescription() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.description != nil
}

func (v *Schema) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *Schema) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

func (v *Schema) Attributes() []*SchemaAttribute {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.attributes; val != nil {
		return val
	}
	return nil
}

func (v *Schema) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.description; val != nil {
		return *val
	}
	return ""
}

func (v *Schema) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.id; val != nil {
		return *val
	}
	return ""
}

func (v *Schema) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.name; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *Schema) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case SchemaAttributesKey:
		v.attributes = nil
	case SchemaDescriptionKey:
		v.description = nil
	case SchemaIDKey:
		v.id = nil
	case SchemaNameKey:
		v.name = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *Schema) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 4)
	if val := v.attributes; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributesKey, Value: val})
	}
	if val := v.description; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaDescriptionKey, Value: *val})
	}
	if val := v.id; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaIDKey, Value: *val})
	}
	if val := v.name; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaNameKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

func (v *Schema) Clone() *Schema {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &Schema{
		attributes:  v.attributes,
		description: v.description,
		id:          v.id,
		name:        v.name,
	}
}

// MarshalJSON serializes Schema into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *Schema) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into Schema.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *Schema) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.attributes = nil
	v.description = nil
	v.id = nil
	v.name = nil

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
			case SchemaAttributesKey:
				var val []*SchemaAttribute
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributesKey, err)
				}
				v.attributes = val
			case SchemaDescriptionKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaDescriptionKey, err)
				}
				v.description = &val
			case SchemaIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaIDKey, err)
				}
				v.id = &val
			case SchemaNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaNameKey, err)
				}
				v.name = &val
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

type SchemaBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *Schema
}

// NewSchemaBuilder creates a new SchemaBuilder instance.
// SchemaBuilder is safe to be used uninitialized as well.
func NewSchemaBuilder() *SchemaBuilder {
	return &SchemaBuilder{}
}

func (b *SchemaBuilder) initialize() {
	b.err = nil
	b.object = &Schema{}
}
func (b *SchemaBuilder) Attributes(in ...*SchemaAttribute) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributesKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaBuilder) Description(in string) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaDescriptionKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaBuilder) ID(in string) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaIDKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaBuilder) Name(in string) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaNameKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaBuilder) Build() (*Schema, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	if b.object.attributes == nil {
		return nil, fmt.Errorf("required field 'Attributes' not initialized")
	}
	if b.object.id == nil {
		return nil, fmt.Errorf("required field 'ID' not initialized")
	}
	if b.object.name == nil {
		return nil, fmt.Errorf("required field 'Name' not initialized")
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *SchemaBuilder) MustBuild() *Schema {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (v *Schema) AsMap(dst map[string]interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *Schema) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) Schema() *SchemaBuilder {
	return &SchemaBuilder{}
}
