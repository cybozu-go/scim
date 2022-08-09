package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	SchemaExtensionRequiredKey = "required"
	SchemaExtensionSchemaKey   = "schema"
)

type SchemaExtension struct {
	required      *bool
	schema        *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type SchemaExtensionValidator interface {
	Validate(*SchemaExtension) error
}

type SchemaExtensionValidateFunc func(v *SchemaExtension) error

func (f SchemaExtensionValidateFunc) Validate(v *SchemaExtension) error {
	return f(v)
}

var DefaultSchemaExtensionValidator SchemaExtensionValidator = SchemaExtensionValidateFunc(func(v *SchemaExtension) error {
	return nil
})

func (v *SchemaExtension) HasRequired() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.required != nil
}

func (v *SchemaExtension) Required() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.required == nil {
		return false
	}
	return *(v.required)
}

func (v *SchemaExtension) HasSchema() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schema != nil
}

func (v *SchemaExtension) Schema() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.schema == nil {
		return ""
	}
	return *(v.schema)
}

func (v *SchemaExtension) makePairs() []pair {
	pairs := make([]pair, 0, 2)
	if v.required != nil {
		pairs = append(pairs, pair{Key: "required", Value: *(v.required)})
	}
	if v.schema != nil {
		pairs = append(pairs, pair{Key: "schema", Value: *(v.schema)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *SchemaExtension) MarshalJSON() ([]byte, error) {
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

func (v *SchemaExtension) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case SchemaExtensionRequiredKey:
		if v.required == nil {
			return nil, false
		}
		return *(v.required), true
	case SchemaExtensionSchemaKey:
		if v.schema == nil {
			return nil, false
		}
		return *(v.schema), true
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

func (v *SchemaExtension) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case SchemaExtensionRequiredKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "required", but got %T`, value)
		}
		v.required = &tmp
		return nil
	case SchemaExtensionSchemaKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "schema", but got %T`, value)
		}
		v.schema = &tmp
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

func (v *SchemaExtension) Clone() *SchemaExtension {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &SchemaExtension{
		required: v.required,
		schema:   v.schema,
	}
}

func (v *SchemaExtension) UnmarshalJSON(data []byte) error {
	v.required = nil
	v.schema = nil
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
			case SchemaExtensionRequiredKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "required": %w`, err)
				}
				v.required = &x
			case SchemaExtensionSchemaKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schema": %w`, err)
				}
				v.schema = &x
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

func (v *SchemaExtension) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// SchemaExtensionBuilder creates a SchemaExtension resource
type SchemaExtensionBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator SchemaExtensionValidator
	object    *SchemaExtension
}

func (b *Builder) SchemaExtension() *SchemaExtensionBuilder {
	return NewSchemaExtensionBuilder()
}

func NewSchemaExtensionBuilder() *SchemaExtensionBuilder {
	var b SchemaExtensionBuilder
	b.init()
	return &b
}

func (b *SchemaExtensionBuilder) From(in *SchemaExtension) *SchemaExtensionBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *SchemaExtensionBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &SchemaExtension{}
}

func (b *SchemaExtensionBuilder) Required(v bool) *SchemaExtensionBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("required", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaExtensionBuilder) Schema(v string) *SchemaExtensionBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("schema", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaExtensionBuilder) Validator(v SchemaExtensionValidator) *SchemaExtensionBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *SchemaExtensionBuilder) Build() (*SchemaExtension, error) {
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
		return nil, fmt.Errorf("resource.SchemaExtensionBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultSchemaExtensionValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *SchemaExtensionBuilder) MustBuild() *SchemaExtension {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
