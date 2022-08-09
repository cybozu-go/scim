package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

// JSON key names for ResourceType resource
const (
	ResourceTypeDescriptionKey      = "description"
	ResourceTypeEndpointKey         = "endpoint"
	ResourceTypeIDKey               = "id"
	ResourceTypeNameKey             = "name"
	ResourceTypeSchemaKey           = "schema"
	ResourceTypeSchemaExtensionsKey = "schemaExtensions"
	ResourceTypeSchemasKey          = "schemas"
)

const ResourceTypeSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:ResourceType"

func init() {
	RegisterExtension(ResourceTypeSchemaURI, ResourceType{})
}

type ResourceType struct {
	description      *string
	endpoint         *string
	id               *string
	name             *string
	schema           *string
	schemaExtensions []*SchemaExtension
	schemas          schemas
	privateParams    map[string]interface{}
	mu               sync.RWMutex
}

type ResourceTypeValidator interface {
	Validate(*ResourceType) error
}

type ResourceTypeValidateFunc func(v *ResourceType) error

func (f ResourceTypeValidateFunc) Validate(v *ResourceType) error {
	return f(v)
}

var DefaultResourceTypeValidator ResourceTypeValidator = ResourceTypeValidateFunc(func(v *ResourceType) error {
	if v.endpoint == nil {
		return fmt.Errorf(`required field "endpoint" is missing in "ResourceType"`)
	}
	if v.schema == nil {
		return fmt.Errorf(`required field "schema" is missing in "ResourceType"`)
	}
	return nil
})

func (v *ResourceType) HasDescription() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.description != nil
}

func (v *ResourceType) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.description == nil {
		return ""
	}
	return *(v.description)
}

func (v *ResourceType) HasEndpoint() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.endpoint != nil
}

func (v *ResourceType) Endpoint() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.endpoint == nil {
		return ""
	}
	return *(v.endpoint)
}

func (v *ResourceType) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *ResourceType) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *ResourceType) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

func (v *ResourceType) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.name == nil {
		return ""
	}
	return *(v.name)
}

func (v *ResourceType) HasSchema() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schema != nil
}

func (v *ResourceType) Schema() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.schema == nil {
		return ""
	}
	return *(v.schema)
}

func (v *ResourceType) HasSchemaExtensions() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemaExtensions != nil
}

func (v *ResourceType) SchemaExtensions() []*SchemaExtension {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemaExtensions
}

func (v *ResourceType) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *ResourceType) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *ResourceType) makePairs() []pair {
	pairs := make([]pair, 0, 7)
	if v.description != nil {
		pairs = append(pairs, pair{Key: "description", Value: *(v.description)})
	}
	if v.endpoint != nil {
		pairs = append(pairs, pair{Key: "endpoint", Value: *(v.endpoint)})
	}
	if v.id != nil {
		pairs = append(pairs, pair{Key: "id", Value: *(v.id)})
	}
	if v.name != nil {
		pairs = append(pairs, pair{Key: "name", Value: *(v.name)})
	}
	if v.schema != nil {
		pairs = append(pairs, pair{Key: "schema", Value: *(v.schema)})
	}
	if v.schemaExtensions != nil {
		pairs = append(pairs, pair{Key: "schemaExtensions", Value: v.schemaExtensions})
	}
	if v.schemas != nil {
		pairs = append(pairs, pair{Key: "schemas", Value: v.schemas})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *ResourceType) MarshalJSON() ([]byte, error) {
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

func (v *ResourceType) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case ResourceTypeDescriptionKey:
		if v.description == nil {
			return nil, false
		}
		return *(v.description), true
	case ResourceTypeEndpointKey:
		if v.endpoint == nil {
			return nil, false
		}
		return *(v.endpoint), true
	case ResourceTypeIDKey:
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case ResourceTypeNameKey:
		if v.name == nil {
			return nil, false
		}
		return *(v.name), true
	case ResourceTypeSchemaKey:
		if v.schema == nil {
			return nil, false
		}
		return *(v.schema), true
	case ResourceTypeSchemaExtensionsKey:
		if v.schemaExtensions == nil {
			return nil, false
		}
		return v.schemaExtensions, true
	case ResourceTypeSchemasKey:
		if v.schemas == nil {
			return nil, false
		}
		return v.schemas, true
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

func (v *ResourceType) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case ResourceTypeDescriptionKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "description", but got %T`, value)
		}
		v.description = &tmp
		return nil
	case ResourceTypeEndpointKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "endpoint", but got %T`, value)
		}
		v.endpoint = &tmp
		return nil
	case ResourceTypeIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case ResourceTypeNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "name", but got %T`, value)
		}
		v.name = &tmp
		return nil
	case ResourceTypeSchemaKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "schema", but got %T`, value)
		}
		v.schema = &tmp
		return nil
	case ResourceTypeSchemaExtensionsKey:
		var tmp []*SchemaExtension
		tmp, ok := value.([]*SchemaExtension)
		if !ok {
			return fmt.Errorf(`expected []*SchemaExtension for field "schemaExtensions", but got %T`, value)
		}
		v.schemaExtensions = tmp
		return nil
	case ResourceTypeSchemasKey:
		var tmp schemas
		tmp, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected schemas for field "schemas", but got %T`, value)
		}
		v.schemas = tmp
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

func (v *ResourceType) Clone() *ResourceType {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &ResourceType{
		description:      v.description,
		endpoint:         v.endpoint,
		id:               v.id,
		name:             v.name,
		schema:           v.schema,
		schemaExtensions: v.schemaExtensions,
		schemas:          v.schemas,
	}
}

func (v *ResourceType) UnmarshalJSON(data []byte) error {
	v.description = nil
	v.endpoint = nil
	v.id = nil
	v.name = nil
	v.schema = nil
	v.schemaExtensions = nil
	v.schemas = nil
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
			case ResourceTypeDescriptionKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "description": %w`, err)
				}
				v.description = &x
			case ResourceTypeEndpointKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "endpoint": %w`, err)
				}
				v.endpoint = &x
			case ResourceTypeIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case ResourceTypeNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "name": %w`, err)
				}
				v.name = &x
			case ResourceTypeSchemaKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schema": %w`, err)
				}
				v.schema = &x
			case ResourceTypeSchemaExtensionsKey:
				var x []*SchemaExtension
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemaExtensions": %w`, err)
				}
				v.schemaExtensions = x
			case ResourceTypeSchemasKey:
				var x schemas
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = x
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

func (v *ResourceType) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// ResourceTypeBuilder creates a ResourceType resource
type ResourceTypeBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator ResourceTypeValidator
	object    *ResourceType
}

func (b *Builder) ResourceType() *ResourceTypeBuilder {
	return NewResourceTypeBuilder()
}

func NewResourceTypeBuilder() *ResourceTypeBuilder {
	var b ResourceTypeBuilder
	b.init()
	return &b
}

func (b *ResourceTypeBuilder) From(in *ResourceType) *ResourceTypeBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *ResourceTypeBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &ResourceType{}

	b.object.schemas = make(schemas)
	b.object.schemas.Add(ResourceTypeSchemaURI)
}

func (b *ResourceTypeBuilder) Description(v string) *ResourceTypeBuilder {
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

func (b *ResourceTypeBuilder) Endpoint(v string) *ResourceTypeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("endpoint", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ResourceTypeBuilder) ID(v string) *ResourceTypeBuilder {
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

func (b *ResourceTypeBuilder) Name(v string) *ResourceTypeBuilder {
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

func (b *ResourceTypeBuilder) Schema(v string) *ResourceTypeBuilder {
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

func (b *ResourceTypeBuilder) SchemaExtensions(v ...*SchemaExtension) *ResourceTypeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("schemaExtensions", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ResourceTypeBuilder) Schemas(v ...string) *ResourceTypeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	for _, schema := range v {
		b.object.schemas.Add(schema)
	}
	return b
}

func (b *ResourceTypeBuilder) Extension(uri string, value interface{}) *ResourceTypeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (b *ResourceTypeBuilder) Validator(v ResourceTypeValidator) *ResourceTypeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *ResourceTypeBuilder) Build() (*ResourceType, error) {
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
		return nil, fmt.Errorf("resource.ResourceTypeBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultResourceTypeValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *ResourceTypeBuilder) MustBuild() *ResourceType {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
