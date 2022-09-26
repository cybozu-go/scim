package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const PatchRequestSchemaURI = "urn:ietf:params:scim:api:messages:2.0:PatchOp"

func init() {
	RegisterExtension(PatchRequestSchemaURI, PatchRequest{})
}

type PatchRequest struct {
	mu         sync.RWMutex
	operations []*PatchOperation
	schemas    *schemas
	extra      map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	PatchRequestOperationsKey = "operations"
	PatchRequestSchemasKey    = "schemas"
)

// Get retrieves the value associated with a key
func (v *PatchRequest) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case PatchRequestOperationsKey:
		if val := v.operations; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case PatchRequestSchemasKey:
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
func (v *PatchRequest) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case PatchRequestOperationsKey:
		converted, ok := value.([]*PatchOperation)
		if !ok {
			return fmt.Errorf(`expected value of type []*PatchOperation for field operations, got %T`, value)
		}
		v.operations = converted
	case PatchRequestSchemasKey:
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
func (v *PatchRequest) HasOperations() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.operations != nil
}

func (v *PatchRequest) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *PatchRequest) Operations() []*PatchOperation {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.operations; val != nil {
		return val
	}
	return nil
}

func (v *PatchRequest) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

// Remove removes the value associated with a key
func (v *PatchRequest) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case PatchRequestOperationsKey:
		v.operations = nil
	case PatchRequestSchemasKey:
		v.schemas = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *PatchRequest) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 2)
	if val := v.operations; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: PatchRequestOperationsKey, Value: val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: PatchRequestSchemasKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes PatchRequest into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *PatchRequest) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into PatchRequest.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *PatchRequest) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.operations = nil
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
			case PatchRequestOperationsKey:
				var val []*PatchOperation
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PatchRequestOperationsKey, err)
				}
				v.operations = val
			case PatchRequestSchemasKey:
				var val schemas
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PatchRequestSchemasKey, err)
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

type PatchRequestBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *PatchRequest
}

// NewPatchRequestBuilder creates a new PatchRequestBuilder instance.
// PatchRequestBuilder is safe to be used uninitialized as well.
func NewPatchRequestBuilder() *PatchRequestBuilder {
	return &PatchRequestBuilder{}
}

func (b *PatchRequestBuilder) initialize() {
	b.err = nil
	b.object = &PatchRequest{}
}
func (b *PatchRequestBuilder) Operations(in ...*PatchOperation) *PatchRequestBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(PatchRequestOperationsKey, in)
	return b
}
func (b *PatchRequestBuilder) Schemas(in ...string) *PatchRequestBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(PatchRequestSchemasKey, in)
	return b
}

func (b *PatchRequestBuilder) Build() (*PatchRequest, error) {
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

func (b *PatchRequestBuilder) MustBuild() *PatchRequest {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *PatchRequestBuilder) From(in *PatchRequest) *PatchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.object = in.Clone()
	return b
}

func (b *PatchRequestBuilder) Extension(uri string, value interface{}) *PatchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}
	if b.object.schemas == nil {
		b.object.schemas = &schemas{}
		b.object.schemas.Add(PatchRequestSchemaURI)
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (v *PatchRequest) Clone() *PatchRequest {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &PatchRequest{
		operations: v.operations,
		schemas:    v.schemas,
	}
}

func (v *PatchRequest) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *PatchRequest) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) PatchRequest() *PatchRequestBuilder {
	return &PatchRequestBuilder{}
}
