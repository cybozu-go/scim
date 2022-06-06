package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	filterSupportMaxResultsJSONKey = "maxResults"
	filterSupportSupportedJSONKey  = "supported"
)

type FilterSupport struct {
	maxResults    *int
	supported     *bool
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type FilterSupportValidator interface {
	Validate(*FilterSupport) error
}

type FilterSupportValidateFunc func(v *FilterSupport) error

func (f FilterSupportValidateFunc) Validate(v *FilterSupport) error {
	return f(v)
}

var DefaultFilterSupportValidator FilterSupportValidator = FilterSupportValidateFunc(func(v *FilterSupport) error {
	if v.maxResults == nil {
		return fmt.Errorf(`required field "maxResults" is missing in "FilterSupport"`)
	}
	if v.supported == nil {
		return fmt.Errorf(`required field "supported" is missing in "FilterSupport"`)
	}
	return nil
})

func (v *FilterSupport) HasMaxResults() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxResults != nil
}

func (v *FilterSupport) MaxResults() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.maxResults == nil {
		return 0
	}
	return *(v.maxResults)
}

func (v *FilterSupport) HasSupported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.supported != nil
}

func (v *FilterSupport) Supported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.supported == nil {
		return false
	}
	return *(v.supported)
}

func (v *FilterSupport) makePairs() []pair {
	pairs := make([]pair, 0, 2)
	if v.maxResults != nil {
		pairs = append(pairs, pair{Key: "maxResults", Value: *(v.maxResults)})
	}
	if v.supported != nil {
		pairs = append(pairs, pair{Key: "supported", Value: *(v.supported)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *FilterSupport) MarshalJSON() ([]byte, error) {
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

func (v *FilterSupport) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case filterSupportMaxResultsJSONKey:
		if v.maxResults == nil {
			return nil, false
		}
		return *(v.maxResults), true
	case filterSupportSupportedJSONKey:
		if v.supported == nil {
			return nil, false
		}
		return *(v.supported), true
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

func (v *FilterSupport) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case filterSupportMaxResultsJSONKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "maxResults", but got %T`, value)
		}
		v.maxResults = &tmp
		return nil
	case filterSupportSupportedJSONKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "supported", but got %T`, value)
		}
		v.supported = &tmp
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

func (v *FilterSupport) Clone() *FilterSupport {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &FilterSupport{
		maxResults: v.maxResults,
		supported:  v.supported,
	}
}

func (v *FilterSupport) UnmarshalJSON(data []byte) error {
	v.maxResults = nil
	v.supported = nil
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
			case filterSupportMaxResultsJSONKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "maxResults": %w`, err)
				}
				v.maxResults = &x
			case filterSupportSupportedJSONKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "supported": %w`, err)
				}
				v.supported = &x
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

func (v *FilterSupport) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type FilterSupportBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator FilterSupportValidator
	object    *FilterSupport
}

func (b *Builder) FilterSupport() *FilterSupportBuilder {
	return NewFilterSupportBuilder()
}

func NewFilterSupportBuilder() *FilterSupportBuilder {
	var b FilterSupportBuilder
	b.init()
	return &b
}

func (b *FilterSupportBuilder) From(in *FilterSupport) *FilterSupportBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *FilterSupportBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &FilterSupport{}
}

func (b *FilterSupportBuilder) MaxResults(v int) *FilterSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("maxResults", v); err != nil {
		b.err = err
	}
	return b
}

func (b *FilterSupportBuilder) Supported(v bool) *FilterSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("supported", v); err != nil {
		b.err = err
	}
	return b
}

func (b *FilterSupportBuilder) Validator(v FilterSupportValidator) *FilterSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *FilterSupportBuilder) Build() (*FilterSupport, error) {
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
		return nil, fmt.Errorf("resource.FilterSupportBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultFilterSupportValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *FilterSupportBuilder) MustBuild() *FilterSupport {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
