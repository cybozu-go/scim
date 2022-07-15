package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	PatchRequestOperationsKey = "operations"
	PatchRequestSchemasKey    = "schemas"
)

const PatchRequestSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:PatchOp"

func init() {
	RegisterExtension(PatchRequestSchemaURI, PatchRequest{})
}

type PatchRequest struct {
	operations    []*PatchOperation
	schemas       schemas
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type PatchRequestValidator interface {
	Validate(*PatchRequest) error
}

type PatchRequestValidateFunc func(v *PatchRequest) error

func (f PatchRequestValidateFunc) Validate(v *PatchRequest) error {
	return f(v)
}

var DefaultPatchRequestValidator PatchRequestValidator = PatchRequestValidateFunc(func(v *PatchRequest) error {
	return nil
})

func (v *PatchRequest) HasOperations() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.operations != nil
}

func (v *PatchRequest) Operations() []*PatchOperation {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.operations
}

func (v *PatchRequest) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *PatchRequest) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *PatchRequest) makePairs() []pair {
	pairs := make([]pair, 0, 2)
	if v.operations != nil {
		pairs = append(pairs, pair{Key: "operations", Value: v.operations})
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

func (v *PatchRequest) MarshalJSON() ([]byte, error) {
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

func (v *PatchRequest) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case PatchRequestOperationsKey:
		if v.operations == nil {
			return nil, false
		}
		return v.operations, true
	case PatchRequestSchemasKey:
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

func (v *PatchRequest) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case PatchRequestOperationsKey:
		var tmp []*PatchOperation
		tmp, ok := value.([]*PatchOperation)
		if !ok {
			return fmt.Errorf(`expected []*PatchOperation for field "operations", but got %T`, value)
		}
		v.operations = tmp
		return nil
	case PatchRequestSchemasKey:
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

func (v *PatchRequest) Clone() *PatchRequest {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &PatchRequest{
		operations: v.operations,
		schemas:    v.schemas,
	}
}

func (v *PatchRequest) UnmarshalJSON(data []byte) error {
	v.operations = nil
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
			case PatchRequestOperationsKey:
				var x []*PatchOperation
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "operations": %w`, err)
				}
				v.operations = x
			case PatchRequestSchemasKey:
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

func (v *PatchRequest) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type PatchRequestBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator PatchRequestValidator
	object    *PatchRequest
}

func (b *Builder) PatchRequest() *PatchRequestBuilder {
	return NewPatchRequestBuilder()
}

func NewPatchRequestBuilder() *PatchRequestBuilder {
	var b PatchRequestBuilder
	b.init()
	return &b
}

func (b *PatchRequestBuilder) From(in *PatchRequest) *PatchRequestBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *PatchRequestBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &PatchRequest{}

	b.object.schemas = make(schemas)
	b.object.schemas.Add(PatchRequestSchemaURI)
}

func (b *PatchRequestBuilder) Operations(v ...*PatchOperation) *PatchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("operations", v); err != nil {
		b.err = err
	}
	return b
}

func (b *PatchRequestBuilder) Schemas(v ...string) *PatchRequestBuilder {
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

func (b *PatchRequestBuilder) Extension(uri string, value interface{}) *PatchRequestBuilder {
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

func (b *PatchRequestBuilder) Validator(v PatchRequestValidator) *PatchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *PatchRequestBuilder) Build() (*PatchRequest, error) {
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
		return nil, fmt.Errorf("resource.PatchRequestBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultPatchRequestValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *PatchRequestBuilder) MustBuild() *PatchRequest {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
