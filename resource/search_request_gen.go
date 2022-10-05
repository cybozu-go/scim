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

const SearchRequestSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:SearchRequest"

func init() {
	Register("SearchRequest", SearchRequestSchemaURI, SearchRequest{})
	RegisterBuilder("SearchRequest", SearchRequestSchemaURI, SearchRequestBuilder{})
}

type SearchRequest struct {
	mu                 sync.RWMutex
	attributes         []string
	count              *int
	excludedAttributes []string
	filter             *string
	schema             *string
	schemas            *schemas
	sortBy             *string
	sortOrder          *string
	startIndex         *int
	extra              map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	SearchRequestAttributesKey         = "attributes"
	SearchRequestCountKey              = "count"
	SearchRequestExcludedAttributesKey = "excludedAttributes"
	SearchRequestFilterKey             = "filter"
	SearchRequestSchemaKey             = "schema"
	SearchRequestSchemasKey            = "schemas"
	SearchRequestSortByKey             = "sortBy"
	SearchRequestSortOrderKey          = "sortOrder"
	SearchRequestStartIndexKey         = "startIndex"
)

// Get retrieves the value associated with a key
func (v *SearchRequest) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *SearchRequest) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case SearchRequestAttributesKey:
		if val := v.attributes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case SearchRequestCountKey:
		if val := v.count; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SearchRequestExcludedAttributesKey:
		if val := v.excludedAttributes; val != nil {
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
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
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
	case SearchRequestExcludedAttributesKey:
		converted, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected value of type []string for field excludedAttributes, got %T`, value)
		}
		v.excludedAttributes = converted
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
		var object schemas
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.schemas = &object
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

// Has returns true if the field specified by the argument has been populated.
// The field name must be the JSON field name, not the Go-structure's field name.
func (v *SearchRequest) Has(name string) bool {
	switch name {
	case SearchRequestAttributesKey:
		return v.attributes != nil
	case SearchRequestCountKey:
		return v.count != nil
	case SearchRequestExcludedAttributesKey:
		return v.excludedAttributes != nil
	case SearchRequestFilterKey:
		return v.filter != nil
	case SearchRequestSchemaKey:
		return v.schema != nil
	case SearchRequestSchemasKey:
		return v.schemas != nil
	case SearchRequestSortByKey:
		return v.sortBy != nil
	case SearchRequestSortOrderKey:
		return v.sortOrder != nil
	case SearchRequestStartIndexKey:
		return v.startIndex != nil
	default:
		if v.extra != nil {
			if _, ok := v.extra[name]; ok {
				return true
			}
		}
		return false
	}
}

// Keys returns a slice of string comprising of JSON field names whose values
// are present in the object.
func (v *SearchRequest) Keys() []string {
	keys := make([]string, 0, 9)
	if v.attributes != nil {
		keys = append(keys, SearchRequestAttributesKey)
	}
	if v.count != nil {
		keys = append(keys, SearchRequestCountKey)
	}
	if v.excludedAttributes != nil {
		keys = append(keys, SearchRequestExcludedAttributesKey)
	}
	if v.filter != nil {
		keys = append(keys, SearchRequestFilterKey)
	}
	if v.schema != nil {
		keys = append(keys, SearchRequestSchemaKey)
	}
	if v.schemas != nil {
		keys = append(keys, SearchRequestSchemasKey)
	}
	if v.sortBy != nil {
		keys = append(keys, SearchRequestSortByKey)
	}
	if v.sortOrder != nil {
		keys = append(keys, SearchRequestSortOrderKey)
	}
	if v.startIndex != nil {
		keys = append(keys, SearchRequestStartIndexKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasAttributes returns true if the field `attributes` has been populated
func (v *SearchRequest) HasAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes != nil
}

// HasCount returns true if the field `count` has been populated
func (v *SearchRequest) HasCount() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.count != nil
}

// HasExcludedAttributes returns true if the field `excludedAttributes` has been populated
func (v *SearchRequest) HasExcludedAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.excludedAttributes != nil
}

// HasFilter returns true if the field `filter` has been populated
func (v *SearchRequest) HasFilter() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.filter != nil
}

// HasSchema returns true if the field `schema` has been populated
func (v *SearchRequest) HasSchema() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schema != nil
}

// HasSchemas returns true if the field `schemas` has been populated
func (v *SearchRequest) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

// HasSortBy returns true if the field `sortBy` has been populated
func (v *SearchRequest) HasSortBy() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sortBy != nil
}

// HasSortOrder returns true if the field `sortOrder` has been populated
func (v *SearchRequest) HasSortOrder() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sortOrder != nil
}

// HasStartIndex returns true if the field `startIndex` has been populated
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

func (v *SearchRequest) ExcludedAttributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.excludedAttributes; val != nil {
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
		return val.GetValue()
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
	case SearchRequestExcludedAttributesKey:
		v.excludedAttributes = nil
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

func (v *SearchRequest) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &SearchRequest{
		attributes:         v.attributes,
		count:              v.count,
		excludedAttributes: v.excludedAttributes,
		filter:             v.filter,
		schema:             v.schema,
		schemas:            v.schemas,
		sortBy:             v.sortBy,
		sortOrder:          v.sortOrder,
		startIndex:         v.startIndex,
		extra:              extra,
	})
}

// MarshalJSON serializes SearchRequest into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *SearchRequest) MarshalJSON() ([]byte, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, k := range v.Keys() {
		var val interface{}
		if err := v.getNoLock(k, &val, true); err != nil {
			return nil, fmt.Errorf(`failed to retrieve value for field %q: %w`, k, err)
		}

		if i > 0 {
			buf.WriteByte(',')
		}
		if err := enc.Encode(k); err != nil {
			return nil, fmt.Errorf(`failed to encode map key name: %w`, err)
		}
		buf.WriteByte(':')
		if err := enc.Encode(val); err != nil {
			return nil, fmt.Errorf(`failed to encode map value for %q: %w`, k, err)
		}
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
	v.excludedAttributes = nil
	v.filter = nil
	v.schema = nil
	v.schemas = nil
	v.sortBy = nil
	v.sortOrder = nil
	v.startIndex = nil

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
			case SearchRequestExcludedAttributesKey:
				var val []string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SearchRequestExcludedAttributesKey, err)
				}
				v.excludedAttributes = val
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
				if err := v.decodeExtraField(tok, dec, &val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, tok, err)
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
	b.object.schemas = &schemas{}
	b.object.schemas.Add(SearchRequestSchemaURI)
}
func (b *SearchRequestBuilder) Attributes(in ...string) *SearchRequestBuilder {
	return b.SetField(SearchRequestAttributesKey, in)
}
func (b *SearchRequestBuilder) Count(in int) *SearchRequestBuilder {
	return b.SetField(SearchRequestCountKey, in)
}
func (b *SearchRequestBuilder) ExcludedAttributes(in ...string) *SearchRequestBuilder {
	return b.SetField(SearchRequestExcludedAttributesKey, in)
}
func (b *SearchRequestBuilder) Filter(in string) *SearchRequestBuilder {
	return b.SetField(SearchRequestFilterKey, in)
}
func (b *SearchRequestBuilder) Schema(in string) *SearchRequestBuilder {
	return b.SetField(SearchRequestSchemaKey, in)
}
func (b *SearchRequestBuilder) Schemas(in ...string) *SearchRequestBuilder {
	return b.SetField(SearchRequestSchemasKey, in)
}
func (b *SearchRequestBuilder) SortBy(in string) *SearchRequestBuilder {
	return b.SetField(SearchRequestSortByKey, in)
}
func (b *SearchRequestBuilder) SortOrder(in string) *SearchRequestBuilder {
	return b.SetField(SearchRequestSortOrderKey, in)
}
func (b *SearchRequestBuilder) StartIndex(in int) *SearchRequestBuilder {
	return b.SetField(SearchRequestStartIndexKey, in)
}

// SetField sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *SearchRequestBuilder) SetField(name string, value interface{}) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(name, value); err != nil {
		b.err = err
	}
	return b
}
func (b *SearchRequestBuilder) Build() (*SearchRequest, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
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

func (b *SearchRequestBuilder) From(in *SearchRequest) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	var cloned SearchRequest
	if err := in.Clone(&cloned); err != nil {
		b.err = err
		return b
	}

	b.object = &cloned
	return b
}

func (b *SearchRequestBuilder) Extension(uri string, value interface{}) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}
	if b.object.schemas == nil {
		b.object.schemas = &schemas{}
		b.object.schemas.Add(SearchRequestSchemaURI)
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

// AsMap returns the resource as a Go map
func (v *SearchRequest) AsMap(m map[string]interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	for _, key := range v.Keys() {
		var val interface{}
		if err := v.getNoLock(key, &val, false); err != nil {
			return fmt.Errorf(`failed to retrieve value for key %q: %w`, key, err)
		}
		m[key] = val
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

func (*SearchRequest) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
	// we can get an instance of the resource object
	if rx, ok := registry.LookupByURI(name); ok {
		if err := dec.Decode(&rx); err != nil {
			return fmt.Errorf(`failed to decode value for key %q: %w`, name, err)
		}
		if err := blackmagic.AssignIfCompatible(dst, rx); err != nil {
			return err
		}
	} else {
		if err := dec.Decode(dst); err != nil {
			return fmt.Errorf(`failed to decode value for key %q: %w`, name, err)
		}
	}
	return nil
}

func (b *Builder) SearchRequest() *SearchRequestBuilder {
	return &SearchRequestBuilder{}
}
