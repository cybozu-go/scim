package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const SearchRequestSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:SearchRequest"

func init() {
	RegisterExtension(SearchRequestSchemaURI, SearchRequest{})
}

type SearchRequest struct {
	mu                sync.RWMutex
	attributes        []string
	count             *int
	excludeAttributes []string
	filter            *string
	schema            *string
	schemas           *schemas
	sortBy            *string
	sortOrder         *string
	startIndex        *int
	extra             map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	SearchRequestAttributesKey        = "attributes"
	SearchRequestCountKey             = "count"
	SearchRequestExcludeAttributesKey = "excludeAttributes"
	SearchRequestFilterKey            = "filter"
	SearchRequestSchemaKey            = "schema"
	SearchRequestSchemasKey           = "schemas"
	SearchRequestSortByKey            = "sortBy"
	SearchRequestSortOrderKey         = "sortOrder"
	SearchRequestStartIndexKey        = "startIndex"
)

// Get retrieves the value associated with a key
func (v *SearchRequest) Get(key string, dst interface{}) error {
	switch key {
	case SearchRequestAttributesKey:
		if val := v.attributes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case SearchRequestCountKey:
		if val := v.count; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SearchRequestExcludeAttributesKey:
		if val := v.excludeAttributes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case SearchRequestFilterKey:
		if val := v.filter; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SearchRequestSchemaKey:
		if val := v.schema; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SearchRequestSchemasKey:
		if val := v.schemas; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SearchRequestSortByKey:
		if val := v.sortBy; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SearchRequestSortOrderKey:
		if val := v.sortOrder; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SearchRequestStartIndexKey:
		if val := v.startIndex; val != nil {
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
func (v *SearchRequest) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case SearchRequestAttributesKey:
		converted, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected value of type []string for field attributes, got %T`, value)
		}
		v.attributes = converted
	case SearchRequestCountKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field count, got %T`, value)
		}
		v.count = &converted
	case SearchRequestExcludeAttributesKey:
		converted, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected value of type []string for field excludeAttributes, got %T`, value)
		}
		v.excludeAttributes = converted
	case SearchRequestFilterKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field filter, got %T`, value)
		}
		v.filter = &converted
	case SearchRequestSchemaKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field schema, got %T`, value)
		}
		v.schema = &converted
	case SearchRequestSchemasKey:
		converted, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected value of type schemas for field schemas, got %T`, value)
		}
		v.schemas = &converted
	case SearchRequestSortByKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field sortBy, got %T`, value)
		}
		v.sortBy = &converted
	case SearchRequestSortOrderKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field sortOrder, got %T`, value)
		}
		v.sortOrder = &converted
	case SearchRequestStartIndexKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field startIndex, got %T`, value)
		}
		v.startIndex = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *SearchRequest) HasAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes != nil
}

func (v *SearchRequest) HasCount() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.count != nil
}

func (v *SearchRequest) HasExcludeAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.excludeAttributes != nil
}

func (v *SearchRequest) HasFilter() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.filter != nil
}

func (v *SearchRequest) HasSchema() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schema != nil
}

func (v *SearchRequest) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *SearchRequest) HasSortBy() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sortBy != nil
}

func (v *SearchRequest) HasSortOrder() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sortOrder != nil
}

func (v *SearchRequest) HasStartIndex() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.startIndex != nil
}

func (v *SearchRequest) Attributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.attributes; val != nil {
		return val
	}
	return []string(nil)
}

func (v *SearchRequest) Count() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.count; val != nil {
		return *val
	}
	return 0
}

func (v *SearchRequest) ExcludeAttributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.excludeAttributes; val != nil {
		return val
	}
	return []string(nil)
}

func (v *SearchRequest) Filter() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.filter; val != nil {
		return *val
	}
	return ""
}

func (v *SearchRequest) Schema() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schema; val != nil {
		return *val
	}
	return ""
}

func (v *SearchRequest) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

func (v *SearchRequest) SortBy() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.sortBy; val != nil {
		return *val
	}
	return ""
}

func (v *SearchRequest) SortOrder() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.sortOrder; val != nil {
		return *val
	}
	return ""
}

func (v *SearchRequest) StartIndex() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.startIndex; val != nil {
		return *val
	}
	return 0
}

// Remove removes the value associated with a key
func (v *SearchRequest) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case SearchRequestAttributesKey:
		v.attributes = nil
	case SearchRequestCountKey:
		v.count = nil
	case SearchRequestExcludeAttributesKey:
		v.excludeAttributes = nil
	case SearchRequestFilterKey:
		v.filter = nil
	case SearchRequestSchemaKey:
		v.schema = nil
	case SearchRequestSchemasKey:
		v.schemas = nil
	case SearchRequestSortByKey:
		v.sortBy = nil
	case SearchRequestSortOrderKey:
		v.sortOrder = nil
	case SearchRequestStartIndexKey:
		v.startIndex = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *SearchRequest) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 9)
	if val := v.attributes; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: SearchRequestAttributesKey, Value: val})
	}
	if val := v.count; val != nil {
		pairs = append(pairs, &fieldPair{Name: SearchRequestCountKey, Value: *val})
	}
	if val := v.excludeAttributes; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: SearchRequestExcludeAttributesKey, Value: val})
	}
	if val := v.filter; val != nil {
		pairs = append(pairs, &fieldPair{Name: SearchRequestFilterKey, Value: *val})
	}
	if val := v.schema; val != nil {
		pairs = append(pairs, &fieldPair{Name: SearchRequestSchemaKey, Value: *val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: SearchRequestSchemasKey, Value: *val})
	}
	if val := v.sortBy; val != nil {
		pairs = append(pairs, &fieldPair{Name: SearchRequestSortByKey, Value: *val})
	}
	if val := v.sortOrder; val != nil {
		pairs = append(pairs, &fieldPair{Name: SearchRequestSortOrderKey, Value: *val})
	}
	if val := v.startIndex; val != nil {
		pairs = append(pairs, &fieldPair{Name: SearchRequestStartIndexKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes SearchRequest into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *SearchRequest) MarshalJSON() ([]byte, error) {
	pairs := v.makePairs()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, pair := range pairs {
		if i > 0 {
			buf.WriteByte(',')
		}
		enc.Encode(pair.Name)
		buf.WriteByte(':')
		enc.Encode(pair.Value)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON deserializes a piece of JSON data into SearchRequest.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *SearchRequest) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.attributes = nil
	v.count = nil
	v.excludeAttributes = nil
	v.filter = nil
	v.schema = nil
	v.schemas = nil
	v.sortBy = nil
	v.sortOrder = nil
	v.startIndex = nil

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
			case SearchRequestAttributesKey:
				var val []string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestAttributesKey, err)
				}
				v.attributes = val
			case SearchRequestCountKey:
				var val int
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestCountKey, err)
				}
				v.count = &val
			case SearchRequestExcludeAttributesKey:
				var val []string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestExcludeAttributesKey, err)
				}
				v.excludeAttributes = val
			case SearchRequestFilterKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestFilterKey, err)
				}
				v.filter = &val
			case SearchRequestSchemaKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestSchemaKey, err)
				}
				v.schema = &val
			case SearchRequestSchemasKey:
				var val schemas
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestSchemasKey, err)
				}
				v.schemas = &val
			case SearchRequestSortByKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestSortByKey, err)
				}
				v.sortBy = &val
			case SearchRequestSortOrderKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestSortOrderKey, err)
				}
				v.sortOrder = &val
			case SearchRequestStartIndexKey:
				var val int
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestStartIndexKey, err)
				}
				v.startIndex = &val
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

type SearchRequestBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *SearchRequest
}

// NewSearchRequestBuilder creates a new SearchRequestBuilder instance.
// SearchRequestBuilder is safe to be used uninitialized as well.
func NewSearchRequestBuilder() *SearchRequestBuilder {
	return &SearchRequestBuilder{}
}

func (b *SearchRequestBuilder) initialize() {
	b.err = nil
	b.object = &SearchRequest{}
}
func (b *SearchRequestBuilder) Attributes(in []string) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestAttributesKey, in)
	return b
}
func (b *SearchRequestBuilder) Count(in int) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestCountKey, in)
	return b
}
func (b *SearchRequestBuilder) ExcludeAttributes(in []string) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestExcludeAttributesKey, in)
	return b
}
func (b *SearchRequestBuilder) Filter(in string) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestFilterKey, in)
	return b
}
func (b *SearchRequestBuilder) Schema(in string) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestSchemaKey, in)
	return b
}
func (b *SearchRequestBuilder) Schemas(in ...string) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestSchemasKey, in)
	return b
}
func (b *SearchRequestBuilder) SortBy(in string) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestSortByKey, in)
	return b
}
func (b *SearchRequestBuilder) SortOrder(in string) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestSortOrderKey, in)
	return b
}
func (b *SearchRequestBuilder) StartIndex(in int) *SearchRequestBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(SearchRequestStartIndexKey, in)
	return b
}

func (b *SearchRequestBuilder) Build() (*SearchRequest, error) {
	err := b.err
	if err != nil {
		return nil, err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *SearchRequestBuilder) MustBuild() *SearchRequest {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (v *SearchRequest) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *SearchRequest) GetExtension(name, uri string, dst interface{}) error {
	if uri == "" {
		return v.Get(name, dst)
	}
	var ext interface{}
	if err := v.Get(uri, ext); err != nil {
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

func (b *Builder) SearchRequest() *SearchRequestBuilder {
	return &SearchRequestBuilder{}
}
