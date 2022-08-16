package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

// JSON key names for Schema resource
const (
	SchemaAttributesKey  = "attributes"
	SchemaDescriptionKey = "description"
	SchemaIDKey          = "id"
	SchemaNameKey        = "name"
)

type Schema struct {
	attributes         []*SchemaAttribute
	description        *string
	id                 *string
	name               *string
	attrByNameInitOnce sync.Once
	attrByName         map[string]*SchemaAttribute
	privateParams      map[string]interface{}
	mu                 sync.RWMutex
}

type SchemaValidator interface {
	Validate(*Schema) error
}

type SchemaValidateFunc func(v *Schema) error

func (f SchemaValidateFunc) Validate(v *Schema) error {
	return f(v)
}

var DefaultSchemaValidator SchemaValidator = SchemaValidateFunc(func(v *Schema) error {
	if v.id == nil {
		return fmt.Errorf(`required field "id" is missing in "Schema"`)
	}
	return nil
})

func (v *Schema) HasAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes != nil
}

func (v *Schema) Attributes() []*SchemaAttribute {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes
}

func (v *Schema) HasDescription() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.description != nil
}

func (v *Schema) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.description == nil {
		return ""
	}
	return *(v.description)
}

func (v *Schema) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *Schema) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *Schema) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

func (v *Schema) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.name == nil {
		return ""
	}
	return *(v.name)
}

func (v *Schema) makePairs() []pair {
	pairs := make([]pair, 0, 4)
	if v.attributes != nil {
		pairs = append(pairs, pair{Key: "attributes", Value: v.attributes})
	}
	if v.description != nil {
		pairs = append(pairs, pair{Key: "description", Value: *(v.description)})
	}
	if v.id != nil {
		pairs = append(pairs, pair{Key: "id", Value: *(v.id)})
	}
	if v.name != nil {
		pairs = append(pairs, pair{Key: "name", Value: *(v.name)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *Schema) MarshalJSON() ([]byte, error) {
	pairs := v.makePairs()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, pair := range pairs {
		if i > 0 {
			buf.WriteRune(',')
		}
		fmt.Fprintf(&buf, "%q:", pair.Key)
		if err := enc.Encode(pair.Value); err != nil {
			return nil, fmt.Errorf("failed to encode value for key %q: %w", pair.Key, err)
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (v *Schema) Get(name string, options ...GetOption) (interface{}, bool) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	var ext string
	//nolint:forcetypeassert
	for _, option := range options {
		switch option.Ident() {
		case identExtension{}:
			ext = option.Value().(string)
		}
	}
	switch name {
	case SchemaAttributesKey:
		if v.attributes == nil {
			return nil, false
		}
		return v.attributes, true
	case SchemaDescriptionKey:
		if v.description == nil {
			return nil, false
		}
		return *(v.description), true
	case SchemaIDKey:
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case SchemaNameKey:
		if v.name == nil {
			return nil, false
		}
		return *(v.name), true
	default:
		pp := v.privateParams
		if pp == nil {
			return nil, false
		}
		if ext == "" {
			ret, ok := pp[name]
			return ret, ok
		}
		obj, ok := pp[ext]
		if !ok {
			return nil, false
		}
		getter, ok := obj.(interface {
			Get(string, ...GetOption) (interface{}, bool)
		})
		if !ok {
			return nil, false
		}
		return getter.Get(name)
	}
}

func (v *Schema) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case SchemaAttributesKey:
		var tmp []*SchemaAttribute
		tmp, ok := value.([]*SchemaAttribute)
		if !ok {
			return fmt.Errorf(`expected []*SchemaAttribute for field "attributes", but got %T`, value)
		}
		v.attributes = tmp
		return nil
	case SchemaDescriptionKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "description", but got %T`, value)
		}
		v.description = &tmp
		return nil
	case SchemaIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case SchemaNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "name", but got %T`, value)
		}
		v.name = &tmp
		return nil
	default:
		pp := v.privateParams
		if pp == nil {
			pp = make(map[string]interface{})
			v.privateParams = pp
		}
		pp[name] = value
		return nil
	}
}

func (v *Schema) Clone() *Schema {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &Schema{
		attributes:  v.attributes,
		description: v.description,
		id:          v.id,
		name:        v.name,
	}
}

func (v *Schema) UnmarshalJSON(data []byte) error {
	v.attributes = nil
	v.description = nil
	v.id = nil
	v.name = nil
	v.privateParams = nil
	dec := json.NewDecoder(bytes.NewReader(data))
	{ // first token
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf("failed to read next token: %s", err)
		}
		tok, ok := tok.(json.Delim)
		if !ok {
			return fmt.Errorf("expected first token to be '{', got %c", tok)
		}
	}
	var privateParams map[string]interface{}

LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf("failed to read next token: %s", err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			if tok == '}' {
				break LOOP
			} else {
				return fmt.Errorf("unexpected token %c found", tok)
			}
		case string:
			switch tok {
			case SchemaAttributesKey:
				var x []*SchemaAttribute
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "attributes": %w`, err)
				}
				v.attributes = x
			case SchemaDescriptionKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "description": %w`, err)
				}
				v.description = &x
			case SchemaIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case SchemaNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "name": %w`, err)
				}
				v.name = &x
			default:
				var x interface{}
				if rx, ok := registry.Get(tok); ok {
					x = rx
					if err := dec.Decode(x); err != nil {
						return fmt.Errorf(`failed to decode value for key %q: %w`, tok, err)
					}
				} else {
					if err := dec.Decode(&x); err != nil {
						return fmt.Errorf(`failed to decode value for key %q: %w`, tok, err)
					}
				}
				if privateParams == nil {
					privateParams = make(map[string]interface{})
				}
				privateParams[tok] = x
			}
		}
	}
	if privateParams != nil {
		v.privateParams = privateParams
	}
	return nil
}

func (v *Schema) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// SchemaBuilder creates a Schema resource
type SchemaBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator SchemaValidator
	object    *Schema
}

func (b *Builder) Schema() *SchemaBuilder {
	return NewSchemaBuilder()
}

func NewSchemaBuilder() *SchemaBuilder {
	var b SchemaBuilder
	b.init()
	return &b
}

func (b *SchemaBuilder) From(in *Schema) *SchemaBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *SchemaBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &Schema{}
}

func (b *SchemaBuilder) Attributes(v ...*SchemaAttribute) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("attributes", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaBuilder) Description(v string) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("description", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaBuilder) ID(v string) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("id", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaBuilder) Name(v string) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("name", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaBuilder) Validator(v SchemaValidator) *SchemaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *SchemaBuilder) Build() (*Schema, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	object := b.object
	validator := b.validator
	err := b.err
	b.once = sync.Once{}
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("resource.SchemaBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultSchemaValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *SchemaBuilder) MustBuild() *Schema {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
