package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	searchRequestAttributesJSONKey        = "attributes"
	searchRequestCountJSONKey             = "count"
	searchRequestExludedAttributesJSONKey = "exludedAttributes"
	searchRequestFilterJSONKey            = "filter"
	searchRequestSchemasJSONKey           = "schemas"
	searchRequestSortByJSONKey            = "sortBy"
	searchRequestSortOrderJSONKey         = "sortOrder"
	searchRequestStartIndexJSONKey        = "startIndex"
)

const SearchRequestSchemaURI = "urn:ietf:params:scim:api:messages:2.0:SearchRequest"

func init() {
	RegisterExtension(SearchRequestSchemaURI, SearchRequest{})
}

type SearchRequest struct {
	attributes        []string
	count             *int
	exludedAttributes []string
	filter            *string
	schemas           schemas
	sortBy            *string
	sortOrder         *string
	startIndex        *int
	privateParams     map[string]interface{}
	mu                sync.RWMutex
}

type SearchRequestValidator interface {
	Validate(*SearchRequest) error
}

type SearchRequestValidateFunc func(v *SearchRequest) error

func (f SearchRequestValidateFunc) Validate(v *SearchRequest) error {
	return f(v)
}

var DefaultSearchRequestValidator SearchRequestValidator = SearchRequestValidateFunc(func(v *SearchRequest) error {
	return nil
})

func (v *SearchRequest) HasAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes != nil
}

func (v *SearchRequest) Attributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes
}

func (v *SearchRequest) HasCount() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.count != nil
}

func (v *SearchRequest) Count() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.count == nil {
		return 0
	}
	return *(v.count)
}

func (v *SearchRequest) HasExludedAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.exludedAttributes != nil
}

func (v *SearchRequest) ExludedAttributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.exludedAttributes
}

func (v *SearchRequest) HasFilter() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.filter != nil
}

func (v *SearchRequest) Filter() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.filter == nil {
		return ""
	}
	return *(v.filter)
}

func (v *SearchRequest) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *SearchRequest) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *SearchRequest) HasSortBy() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sortBy != nil
}

func (v *SearchRequest) SortBy() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.sortBy == nil {
		return ""
	}
	return *(v.sortBy)
}

func (v *SearchRequest) HasSortOrder() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sortOrder != nil
}

func (v *SearchRequest) SortOrder() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.sortOrder == nil {
		return ""
	}
	return *(v.sortOrder)
}

func (v *SearchRequest) HasStartIndex() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.startIndex != nil
}

func (v *SearchRequest) StartIndex() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.startIndex == nil {
		return 0
	}
	return *(v.startIndex)
}

func (v *SearchRequest) makePairs() []pair {
	pairs := make([]pair, 0, 8)
	if v.attributes != nil {
		pairs = append(pairs, pair{Key: "attributes", Value: v.attributes})
	}
	if v.count != nil {
		pairs = append(pairs, pair{Key: "count", Value: *(v.count)})
	}
	if v.exludedAttributes != nil {
		pairs = append(pairs, pair{Key: "exludedAttributes", Value: v.exludedAttributes})
	}
	if v.filter != nil {
		pairs = append(pairs, pair{Key: "filter", Value: *(v.filter)})
	}
	if v.schemas != nil {
		pairs = append(pairs, pair{Key: "schemas", Value: v.schemas})
	}
	if v.sortBy != nil {
		pairs = append(pairs, pair{Key: "sortBy", Value: *(v.sortBy)})
	}
	if v.sortOrder != nil {
		pairs = append(pairs, pair{Key: "sortOrder", Value: *(v.sortOrder)})
	}
	if v.startIndex != nil {
		pairs = append(pairs, pair{Key: "startIndex", Value: *(v.startIndex)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *SearchRequest) MarshalJSON() ([]byte, error) {
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

func (v *SearchRequest) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case searchRequestAttributesJSONKey:
		if v.attributes == nil {
			return nil, false
		}
		return v.attributes, true
	case searchRequestCountJSONKey:
		if v.count == nil {
			return nil, false
		}
		return *(v.count), true
	case searchRequestExludedAttributesJSONKey:
		if v.exludedAttributes == nil {
			return nil, false
		}
		return v.exludedAttributes, true
	case searchRequestFilterJSONKey:
		if v.filter == nil {
			return nil, false
		}
		return *(v.filter), true
	case searchRequestSchemasJSONKey:
		if v.schemas == nil {
			return nil, false
		}
		return v.schemas, true
	case searchRequestSortByJSONKey:
		if v.sortBy == nil {
			return nil, false
		}
		return *(v.sortBy), true
	case searchRequestSortOrderJSONKey:
		if v.sortOrder == nil {
			return nil, false
		}
		return *(v.sortOrder), true
	case searchRequestStartIndexJSONKey:
		if v.startIndex == nil {
			return nil, false
		}
		return *(v.startIndex), true
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

func (v *SearchRequest) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case searchRequestAttributesJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "attributes", but got %T`, value)
		}
		v.attributes = tmp
		return nil
	case searchRequestCountJSONKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "count", but got %T`, value)
		}
		v.count = &tmp
		return nil
	case searchRequestExludedAttributesJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "exludedAttributes", but got %T`, value)
		}
		v.exludedAttributes = tmp
		return nil
	case searchRequestFilterJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "filter", but got %T`, value)
		}
		v.filter = &tmp
		return nil
	case searchRequestSchemasJSONKey:
		var tmp schemas
		tmp, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected schemas for field "schemas", but got %T`, value)
		}
		v.schemas = tmp
		return nil
	case searchRequestSortByJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "sortBy", but got %T`, value)
		}
		v.sortBy = &tmp
		return nil
	case searchRequestSortOrderJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "sortOrder", but got %T`, value)
		}
		v.sortOrder = &tmp
		return nil
	case searchRequestStartIndexJSONKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "startIndex", but got %T`, value)
		}
		v.startIndex = &tmp
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

func (v *SearchRequest) Clone() *SearchRequest {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &SearchRequest{
		attributes:        v.attributes,
		count:             v.count,
		exludedAttributes: v.exludedAttributes,
		filter:            v.filter,
		schemas:           v.schemas,
		sortBy:            v.sortBy,
		sortOrder:         v.sortOrder,
		startIndex:        v.startIndex,
	}
}

func (v *SearchRequest) UnmarshalJSON(data []byte) error {
	v.attributes = nil
	v.count = nil
	v.exludedAttributes = nil
	v.filter = nil
	v.schemas = nil
	v.sortBy = nil
	v.sortOrder = nil
	v.startIndex = nil
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
			case searchRequestAttributesJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "attributes": %w`, err)
				}
				v.attributes = x
			case searchRequestCountJSONKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "count": %w`, err)
				}
				v.count = &x
			case searchRequestExludedAttributesJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "exludedAttributes": %w`, err)
				}
				v.exludedAttributes = x
			case searchRequestFilterJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "filter": %w`, err)
				}
				v.filter = &x
			case searchRequestSchemasJSONKey:
				var x schemas
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = x
			case searchRequestSortByJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "sortBy": %w`, err)
				}
				v.sortBy = &x
			case searchRequestSortOrderJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "sortOrder": %w`, err)
				}
				v.sortOrder = &x
			case searchRequestStartIndexJSONKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "startIndex": %w`, err)
				}
				v.startIndex = &x
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

func (v *SearchRequest) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type SearchRequestBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator SearchRequestValidator
	object    *SearchRequest
}

func (b *Builder) SearchRequest() *SearchRequestBuilder {
	return NewSearchRequestBuilder()
}

func NewSearchRequestBuilder() *SearchRequestBuilder {
	var b SearchRequestBuilder
	b.init()
	return &b
}

func (b *SearchRequestBuilder) From(in *SearchRequest) *SearchRequestBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *SearchRequestBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &SearchRequest{}
}

func (b *SearchRequestBuilder) Attributes(v ...string) *SearchRequestBuilder {
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

func (b *SearchRequestBuilder) Count(v int) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("count", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SearchRequestBuilder) ExludedAttributes(v ...string) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("exludedAttributes", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SearchRequestBuilder) Filter(v string) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("filter", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SearchRequestBuilder) Schemas(v ...string) *SearchRequestBuilder {
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

func (b *SearchRequestBuilder) SortBy(v string) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("sortBy", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SearchRequestBuilder) SortOrder(v string) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("sortOrder", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SearchRequestBuilder) StartIndex(v int) *SearchRequestBuilder {
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

func (b *SearchRequestBuilder) Extension(uri string, value interface{}) *SearchRequestBuilder {
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

func (b *SearchRequestBuilder) Validator(v SearchRequestValidator) *SearchRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *SearchRequestBuilder) Build() (*SearchRequest, error) {
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
		return nil, fmt.Errorf("resource.SearchRequestBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultSearchRequestValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *SearchRequestBuilder) MustBuild() *SearchRequest {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
