package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	ListResponseItemsPerPageKey = "itemsPerPage"
	ListResponseResourcesKey    = "resources"
	ListResponseSchemasKey      = "schemas"
	ListResponseStartIndexKey   = "startIndex"
	ListResponseTotalResultsKey = "totalResults"
)

const ListResponseSchemaURI = "urn:ietf:params:scim:api:messages:2.0:ListResponse"

func init() {
	RegisterExtension(ListResponseSchemaURI, ListResponse{})
}

type ListResponse struct {
	itemsPerPage  *int
	resources     []interface{}
	schemas       schemas
	startIndex    *int
	totalResults  *int
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type ListResponseValidator interface {
	Validate(*ListResponse) error
}

type ListResponseValidateFunc func(v *ListResponse) error

func (f ListResponseValidateFunc) Validate(v *ListResponse) error {
	return f(v)
}

var DefaultListResponseValidator ListResponseValidator = ListResponseValidateFunc(func(v *ListResponse) error {
	return nil
})

func (v *ListResponse) HasItemsPerPage() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.itemsPerPage != nil
}

func (v *ListResponse) ItemsPerPage() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.itemsPerPage == nil {
		return 0
	}
	return *(v.itemsPerPage)
}

func (v *ListResponse) HasResources() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.resources != nil
}

func (v *ListResponse) Resources() []interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.resources
}

func (v *ListResponse) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *ListResponse) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *ListResponse) HasStartIndex() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.startIndex != nil
}

func (v *ListResponse) StartIndex() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.startIndex == nil {
		return 0
	}
	return *(v.startIndex)
}

func (v *ListResponse) HasTotalResults() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.totalResults != nil
}

func (v *ListResponse) TotalResults() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.totalResults == nil {
		return 0
	}
	return *(v.totalResults)
}

func (v *ListResponse) makePairs() []pair {
	pairs := make([]pair, 0, 5)
	if v.itemsPerPage != nil {
		pairs = append(pairs, pair{Key: "itemsPerPage", Value: *(v.itemsPerPage)})
	}
	if v.resources != nil {
		pairs = append(pairs, pair{Key: "resources", Value: v.resources})
	}
	if v.schemas != nil {
		pairs = append(pairs, pair{Key: "schemas", Value: v.schemas})
	}
	if v.startIndex != nil {
		pairs = append(pairs, pair{Key: "startIndex", Value: *(v.startIndex)})
	}
	if v.totalResults != nil {
		pairs = append(pairs, pair{Key: "totalResults", Value: *(v.totalResults)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *ListResponse) MarshalJSON() ([]byte, error) {
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

func (v *ListResponse) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case ListResponseItemsPerPageKey:
		if v.itemsPerPage == nil {
			return nil, false
		}
		return *(v.itemsPerPage), true
	case ListResponseResourcesKey:
		if v.resources == nil {
			return nil, false
		}
		return v.resources, true
	case ListResponseSchemasKey:
		if v.schemas == nil {
			return nil, false
		}
		return v.schemas, true
	case ListResponseStartIndexKey:
		if v.startIndex == nil {
			return nil, false
		}
		return *(v.startIndex), true
	case ListResponseTotalResultsKey:
		if v.totalResults == nil {
			return nil, false
		}
		return *(v.totalResults), true
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

func (v *ListResponse) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case ListResponseItemsPerPageKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "itemsPerPage", but got %T`, value)
		}
		v.itemsPerPage = &tmp
		return nil
	case ListResponseResourcesKey:
		var tmp []interface{}
		tmp, ok := value.([]interface{})
		if !ok {
			return fmt.Errorf(`expected []interface{} for field "resources", but got %T`, value)
		}
		v.resources = tmp
		return nil
	case ListResponseSchemasKey:
		var tmp schemas
		tmp, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected schemas for field "schemas", but got %T`, value)
		}
		v.schemas = tmp
		return nil
	case ListResponseStartIndexKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "startIndex", but got %T`, value)
		}
		v.startIndex = &tmp
		return nil
	case ListResponseTotalResultsKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "totalResults", but got %T`, value)
		}
		v.totalResults = &tmp
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

func (v *ListResponse) Clone() *ListResponse {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &ListResponse{
		itemsPerPage: v.itemsPerPage,
		resources:    v.resources,
		schemas:      v.schemas,
		startIndex:   v.startIndex,
		totalResults: v.totalResults,
	}
}

func (v *ListResponse) UnmarshalJSON(data []byte) error {
	v.itemsPerPage = nil
	v.resources = nil
	v.schemas = nil
	v.startIndex = nil
	v.totalResults = nil
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
			case ListResponseItemsPerPageKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "itemsPerPage": %w`, err)
				}
				v.itemsPerPage = &x
			case ListResponseResourcesKey:
				var x []interface{}
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "resources": %w`, err)
				}
				v.resources = x
			case ListResponseSchemasKey:
				var x schemas
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = x
			case ListResponseStartIndexKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "startIndex": %w`, err)
				}
				v.startIndex = &x
			case ListResponseTotalResultsKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "totalResults": %w`, err)
				}
				v.totalResults = &x
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

func (v *ListResponse) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type ListResponseBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator ListResponseValidator
	object    *ListResponse
}

func (b *Builder) ListResponse() *ListResponseBuilder {
	return NewListResponseBuilder()
}

func NewListResponseBuilder() *ListResponseBuilder {
	var b ListResponseBuilder
	b.init()
	return &b
}

func (b *ListResponseBuilder) From(in *ListResponse) *ListResponseBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *ListResponseBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &ListResponse{}
}

func (b *ListResponseBuilder) ItemsPerPage(v int) *ListResponseBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("itemsPerPage", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ListResponseBuilder) Resources(v ...interface{}) *ListResponseBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("resources", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ListResponseBuilder) Schemas(v ...string) *ListResponseBuilder {
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

func (b *ListResponseBuilder) StartIndex(v int) *ListResponseBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("startIndex", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ListResponseBuilder) TotalResults(v int) *ListResponseBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("totalResults", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ListResponseBuilder) Extension(uri string, value interface{}) *ListResponseBuilder {
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

func (b *ListResponseBuilder) Validator(v ListResponseValidator) *ListResponseBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *ListResponseBuilder) Build() (*ListResponse, error) {
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
		return nil, fmt.Errorf("resource.ListResponseBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultListResponseValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *ListResponseBuilder) MustBuild() *ListResponse {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
