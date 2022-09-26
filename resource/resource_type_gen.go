package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const ResourceTypeSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:ResourceType"

func init() {
	RegisterExtension(ResourceTypeSchemaURI, ResourceType{})
}

type ResourceType struct {
	mu              sync.RWMutex
	description     *string
	endpoint        *string
	id              *string
	name            *string
	schema          *string
	schemaExtension []*SchemaExtension
	schemas         *schemas
	extra           map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	ResourceTypeDescriptionKey     = "description"
	ResourceTypeEndpointKey        = "endpoint"
	ResourceTypeIDKey              = "id"
	ResourceTypeNameKey            = "name"
	ResourceTypeSchemaKey          = "schema"
	ResourceTypeSchemaExtensionKey = "schemaExtension"
	ResourceTypeSchemasKey         = "schemas"
)

// Get retrieves the value associated with a key
func (v *ResourceType) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case ResourceTypeDescriptionKey:
		if val := v.description; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ResourceTypeEndpointKey:
		if val := v.endpoint; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ResourceTypeIDKey:
		if val := v.id; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ResourceTypeNameKey:
		if val := v.name; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ResourceTypeSchemaKey:
		if val := v.schema; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ResourceTypeSchemaExtensionKey:
		if val := v.schemaExtension; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ResourceTypeSchemasKey:
		if val := v.schemas; val != nil {
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
func (v *ResourceType) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case ResourceTypeDescriptionKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field description, got %T`, value)
		}
		v.description = &converted
	case ResourceTypeEndpointKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field endpoint, got %T`, value)
		}
		v.endpoint = &converted
	case ResourceTypeIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field id, got %T`, value)
		}
		v.id = &converted
	case ResourceTypeNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field name, got %T`, value)
		}
		v.name = &converted
	case ResourceTypeSchemaKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field schema, got %T`, value)
		}
		v.schema = &converted
	case ResourceTypeSchemaExtensionKey:
		converted, ok := value.([]*SchemaExtension)
		if !ok {
			return fmt.Errorf(`expected value of type []*SchemaExtension for field schemaExtension, got %T`, value)
		}
		v.schemaExtension = converted
	case ResourceTypeSchemasKey:
		converted, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected value of type schemas for field schemas, got %T`, value)
		}
		v.schemas = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *ResourceType) HasDescription() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.description != nil
}

func (v *ResourceType) HasEndpoint() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.endpoint != nil
}

func (v *ResourceType) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *ResourceType) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

func (v *ResourceType) HasSchema() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schema != nil
}

func (v *ResourceType) HasSchemaExtension() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemaExtension != nil
}

func (v *ResourceType) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *ResourceType) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.description; val != nil {
		return *val
	}
	return ""
}

func (v *ResourceType) Endpoint() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.endpoint; val != nil {
		return *val
	}
	return ""
}

func (v *ResourceType) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.id; val != nil {
		return *val
	}
	return ""
}

func (v *ResourceType) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.name; val != nil {
		return *val
	}
	return ""
}

func (v *ResourceType) Schema() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schema; val != nil {
		return *val
	}
	return ""
}

func (v *ResourceType) SchemaExtension() []*SchemaExtension {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemaExtension; val != nil {
		return val
	}
	return nil
}

func (v *ResourceType) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

// Remove removes the value associated with a key
func (v *ResourceType) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case ResourceTypeDescriptionKey:
		v.description = nil
	case ResourceTypeEndpointKey:
		v.endpoint = nil
	case ResourceTypeIDKey:
		v.id = nil
	case ResourceTypeNameKey:
		v.name = nil
	case ResourceTypeSchemaKey:
		v.schema = nil
	case ResourceTypeSchemaExtensionKey:
		v.schemaExtension = nil
	case ResourceTypeSchemasKey:
		v.schemas = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *ResourceType) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 7)
	if val := v.description; val != nil {
		pairs = append(pairs, &fieldPair{Name: ResourceTypeDescriptionKey, Value: *val})
	}
	if val := v.endpoint; val != nil {
		pairs = append(pairs, &fieldPair{Name: ResourceTypeEndpointKey, Value: *val})
	}
	if val := v.id; val != nil {
		pairs = append(pairs, &fieldPair{Name: ResourceTypeIDKey, Value: *val})
	}
	if val := v.name; val != nil {
		pairs = append(pairs, &fieldPair{Name: ResourceTypeNameKey, Value: *val})
	}
	if val := v.schema; val != nil {
		pairs = append(pairs, &fieldPair{Name: ResourceTypeSchemaKey, Value: *val})
	}
	if val := v.schemaExtension; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: ResourceTypeSchemaExtensionKey, Value: val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: ResourceTypeSchemasKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes ResourceType into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *ResourceType) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into ResourceType.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *ResourceType) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.description = nil
	v.endpoint = nil
	v.id = nil
	v.name = nil
	v.schema = nil
	v.schemaExtension = nil
	v.schemas = nil

	dec := json.NewDecoder(bytes.NewReader(data))

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
			case ResourceTypeDescriptionKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ResourceTypeDescriptionKey, err)
				}
				v.description = &val
			case ResourceTypeEndpointKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ResourceTypeEndpointKey, err)
				}
				v.endpoint = &val
			case ResourceTypeIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ResourceTypeIDKey, err)
				}
				v.id = &val
			case ResourceTypeNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ResourceTypeNameKey, err)
				}
				v.name = &val
			case ResourceTypeSchemaKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ResourceTypeSchemaKey, err)
				}
				v.schema = &val
			case ResourceTypeSchemaExtensionKey:
				var val []*SchemaExtension
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ResourceTypeSchemaExtensionKey, err)
				}
				v.schemaExtension = val
			case ResourceTypeSchemasKey:
				var val schemas
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ResourceTypeSchemasKey, err)
				}
				v.schemas = &val
			default:
				var val interface{}
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, tok, err)
				}
				if v.extra == nil {
					v.extra = make(map[string]interface{})
				}
				v.extra[tok] = val
			}
		}
	}
	return nil
}

type ResourceTypeBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *ResourceType
}

// NewResourceTypeBuilder creates a new ResourceTypeBuilder instance.
// ResourceTypeBuilder is safe to be used uninitialized as well.
func NewResourceTypeBuilder() *ResourceTypeBuilder {
	return &ResourceTypeBuilder{}
}

func (b *ResourceTypeBuilder) initialize() {
	b.err = nil
	b.object = &ResourceType{}
}
func (b *ResourceTypeBuilder) Description(in string) *ResourceTypeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ResourceTypeDescriptionKey, in)
	return b
}
func (b *ResourceTypeBuilder) Endpoint(in string) *ResourceTypeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ResourceTypeEndpointKey, in)
	return b
}
func (b *ResourceTypeBuilder) ID(in string) *ResourceTypeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ResourceTypeIDKey, in)
	return b
}
func (b *ResourceTypeBuilder) Name(in string) *ResourceTypeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ResourceTypeNameKey, in)
	return b
}
func (b *ResourceTypeBuilder) Schema(in string) *ResourceTypeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ResourceTypeSchemaKey, in)
	return b
}
func (b *ResourceTypeBuilder) SchemaExtension(in ...*SchemaExtension) *ResourceTypeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ResourceTypeSchemaExtensionKey, in)
	return b
}
func (b *ResourceTypeBuilder) Schemas(in ...string) *ResourceTypeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ResourceTypeSchemasKey, in)
	return b
}

func (b *ResourceTypeBuilder) Build() (*ResourceType, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	err := b.err
	if err != nil {
		return nil, err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *ResourceTypeBuilder) MustBuild() *ResourceType {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *ResourceTypeBuilder) From(in *ResourceType) *ResourceTypeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.object = in.Clone()
	return b
}

func (b *ResourceTypeBuilder) Extension(uri string, value interface{}) *ResourceTypeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}
	if b.object.schemas == nil {
		b.object.schemas = &schemas{}
		b.object.schemas.Add(ResourceTypeSchemaURI)
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (v *ResourceType) Clone() *ResourceType {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &ResourceType{
		description:     v.description,
		endpoint:        v.endpoint,
		id:              v.id,
		name:            v.name,
		schema:          v.schema,
		schemaExtension: v.schemaExtension,
		schemas:         v.schemas,
	}
}

func (v *ResourceType) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *ResourceType) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) ResourceType() *ResourceTypeBuilder {
	return &ResourceTypeBuilder{}
}
