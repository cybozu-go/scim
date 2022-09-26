package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const ListResponseSchemaURI = "urn:ietf:params:scim:api:messages:2.0:ListResponse"

func init() {
	RegisterExtension(ListResponseSchemaURI, ListResponse{})
}

type ListResponse struct {
	mu           sync.RWMutex
	itemsPerPage *int
	resources    []interface{}
	startIndex   *int
	totalResults *int
	schemas      *schemas
	extra        map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	ListResponseItemsPerPageKey = "itemsPerPage"
	ListResponseResourcesKey    = "resources"
	ListResponseStartIndexKey   = "startIndex"
	ListResponseTotalResultsKey = "totalResults"
	ListResponseSchemasKey      = "schemas"
)

// Get retrieves the value associated with a key
func (v *ListResponse) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case ListResponseItemsPerPageKey:
		if val := v.itemsPerPage; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ListResponseResourcesKey:
		if val := v.resources; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ListResponseStartIndexKey:
		if val := v.startIndex; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ListResponseTotalResultsKey:
		if val := v.totalResults; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ListResponseSchemasKey:
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
func (v *ListResponse) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case ListResponseItemsPerPageKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field itemsPerPage, got %T`, value)
		}
		v.itemsPerPage = &converted
	case ListResponseResourcesKey:
		converted, ok := value.([]interface{})
		if !ok {
			return fmt.Errorf(`expected value of type []interface {} for field resources, got %T`, value)
		}
		v.resources = converted
	case ListResponseStartIndexKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field startIndex, got %T`, value)
		}
		v.startIndex = &converted
	case ListResponseTotalResultsKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field totalResults, got %T`, value)
		}
		v.totalResults = &converted
	case ListResponseSchemasKey:
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
func (v *ListResponse) HasItemsPerPage() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.itemsPerPage != nil
}

func (v *ListResponse) HasResources() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.resources != nil
}

func (v *ListResponse) HasStartIndex() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.startIndex != nil
}

func (v *ListResponse) HasTotalResults() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.totalResults != nil
}

func (v *ListResponse) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *ListResponse) ItemsPerPage() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.itemsPerPage; val != nil {
		return *val
	}
	return 0
}

func (v *ListResponse) Resources() []interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.resources; val != nil {
		return val
	}
	return []interface{}(nil)
}

func (v *ListResponse) StartIndex() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.startIndex; val != nil {
		return *val
	}
	return 0
}

func (v *ListResponse) TotalResults() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.totalResults; val != nil {
		return *val
	}
	return 0
}

func (v *ListResponse) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

// Remove removes the value associated with a key
func (v *ListResponse) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case ListResponseItemsPerPageKey:
		v.itemsPerPage = nil
	case ListResponseResourcesKey:
		v.resources = nil
	case ListResponseStartIndexKey:
		v.startIndex = nil
	case ListResponseTotalResultsKey:
		v.totalResults = nil
	case ListResponseSchemasKey:
		v.schemas = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *ListResponse) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 5)
	if val := v.itemsPerPage; val != nil {
		pairs = append(pairs, &fieldPair{Name: ListResponseItemsPerPageKey, Value: *val})
	}
	if val := v.resources; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: ListResponseResourcesKey, Value: val})
	}
	if val := v.startIndex; val != nil {
		pairs = append(pairs, &fieldPair{Name: ListResponseStartIndexKey, Value: *val})
	}
	if val := v.totalResults; val != nil {
		pairs = append(pairs, &fieldPair{Name: ListResponseTotalResultsKey, Value: *val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: ListResponseSchemasKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes ListResponse into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *ListResponse) MarshalJSON() ([]byte, error) {
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

type ListResponseBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *ListResponse
}

// NewListResponseBuilder creates a new ListResponseBuilder instance.
// ListResponseBuilder is safe to be used uninitialized as well.
func NewListResponseBuilder() *ListResponseBuilder {
	return &ListResponseBuilder{}
}

func (b *ListResponseBuilder) initialize() {
	b.err = nil
	b.object = &ListResponse{}
}
func (b *ListResponseBuilder) ItemsPerPage(in int) *ListResponseBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ListResponseItemsPerPageKey, in)
	return b
}
func (b *ListResponseBuilder) Resources(in ...interface{}) *ListResponseBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ListResponseResourcesKey, in)
	return b
}
func (b *ListResponseBuilder) StartIndex(in int) *ListResponseBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ListResponseStartIndexKey, in)
	return b
}
func (b *ListResponseBuilder) TotalResults(in int) *ListResponseBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ListResponseTotalResultsKey, in)
	return b
}
func (b *ListResponseBuilder) Schemas(in ...string) *ListResponseBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(ListResponseSchemasKey, in)
	return b
}

func (b *ListResponseBuilder) Build() (*ListResponse, error) {
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

func (b *ListResponseBuilder) MustBuild() *ListResponse {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *ListResponseBuilder) From(in *ListResponse) *ListResponseBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.object = in.Clone()
	return b
}

func (b *ListResponseBuilder) Extension(uri string, value interface{}) *ListResponseBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}
	if b.object.schemas == nil {
		b.object.schemas = &schemas{}
		b.object.schemas.Add(ListResponseSchemaURI)
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (v *ListResponse) Clone() *ListResponse {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &ListResponse{
		itemsPerPage: v.itemsPerPage,
		resources:    v.resources,
		startIndex:   v.startIndex,
		totalResults: v.totalResults,
		schemas:      v.schemas,
	}
}

func (v *ListResponse) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *ListResponse) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) ListResponse() *ListResponseBuilder {
	return &ListResponseBuilder{}
}
