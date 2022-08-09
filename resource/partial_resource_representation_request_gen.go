package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	PartialResourceRepresentationRequestAttributesKey         = "attributes"
	PartialResourceRepresentationRequestExcludedAttributesKey = "excludedAttributes"
)

type PartialResourceRepresentationRequest struct {
	attributes         []string
	excludedAttributes []string
	privateParams      map[string]interface{}
	mu                 sync.RWMutex
}

type PartialResourceRepresentationRequestValidator interface {
	Validate(*PartialResourceRepresentationRequest) error
}

type PartialResourceRepresentationRequestValidateFunc func(v *PartialResourceRepresentationRequest) error

func (f PartialResourceRepresentationRequestValidateFunc) Validate(v *PartialResourceRepresentationRequest) error {
	return f(v)
}

var DefaultPartialResourceRepresentationRequestValidator PartialResourceRepresentationRequestValidator = PartialResourceRepresentationRequestValidateFunc(func(v *PartialResourceRepresentationRequest) error {
	return nil
})

func (v *PartialResourceRepresentationRequest) HasAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes != nil
}

func (v *PartialResourceRepresentationRequest) Attributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.attributes
}

func (v *PartialResourceRepresentationRequest) HasExcludedAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.excludedAttributes != nil
}

func (v *PartialResourceRepresentationRequest) ExcludedAttributes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.excludedAttributes
}

func (v *PartialResourceRepresentationRequest) makePairs() []pair {
	pairs := make([]pair, 0, 2)
	if v.attributes != nil {
		pairs = append(pairs, pair{Key: "attributes", Value: v.attributes})
	}
	if v.excludedAttributes != nil {
		pairs = append(pairs, pair{Key: "excludedAttributes", Value: v.excludedAttributes})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *PartialResourceRepresentationRequest) MarshalJSON() ([]byte, error) {
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

func (v *PartialResourceRepresentationRequest) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case PartialResourceRepresentationRequestAttributesKey:
		if v.attributes == nil {
			return nil, false
		}
		return v.attributes, true
	case PartialResourceRepresentationRequestExcludedAttributesKey:
		if v.excludedAttributes == nil {
			return nil, false
		}
		return v.excludedAttributes, true
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

func (v *PartialResourceRepresentationRequest) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case PartialResourceRepresentationRequestAttributesKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "attributes", but got %T`, value)
		}
		v.attributes = tmp
		return nil
	case PartialResourceRepresentationRequestExcludedAttributesKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "excludedAttributes", but got %T`, value)
		}
		v.excludedAttributes = tmp
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

func (v *PartialResourceRepresentationRequest) Clone() *PartialResourceRepresentationRequest {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &PartialResourceRepresentationRequest{
		attributes:         v.attributes,
		excludedAttributes: v.excludedAttributes,
	}
}

func (v *PartialResourceRepresentationRequest) UnmarshalJSON(data []byte) error {
	v.attributes = nil
	v.excludedAttributes = nil
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
			case PartialResourceRepresentationRequestAttributesKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "attributes": %w`, err)
				}
				v.attributes = x
			case PartialResourceRepresentationRequestExcludedAttributesKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "excludedAttributes": %w`, err)
				}
				v.excludedAttributes = x
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

func (v *PartialResourceRepresentationRequest) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// PartialResourceRepresentationRequestBuilder creates a PartialResourceRepresentationRequest resource
type PartialResourceRepresentationRequestBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator PartialResourceRepresentationRequestValidator
	object    *PartialResourceRepresentationRequest
}

func (b *Builder) PartialResourceRepresentationRequest() *PartialResourceRepresentationRequestBuilder {
	return NewPartialResourceRepresentationRequestBuilder()
}

func NewPartialResourceRepresentationRequestBuilder() *PartialResourceRepresentationRequestBuilder {
	var b PartialResourceRepresentationRequestBuilder
	b.init()
	return &b
}

func (b *PartialResourceRepresentationRequestBuilder) From(in *PartialResourceRepresentationRequest) *PartialResourceRepresentationRequestBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *PartialResourceRepresentationRequestBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &PartialResourceRepresentationRequest{}
}

func (b *PartialResourceRepresentationRequestBuilder) Attributes(v ...string) *PartialResourceRepresentationRequestBuilder {
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

func (b *PartialResourceRepresentationRequestBuilder) ExcludedAttributes(v ...string) *PartialResourceRepresentationRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("excludedAttributes", v); err != nil {
		b.err = err
	}
	return b
}

func (b *PartialResourceRepresentationRequestBuilder) Validator(v PartialResourceRepresentationRequestValidator) *PartialResourceRepresentationRequestBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *PartialResourceRepresentationRequestBuilder) Build() (*PartialResourceRepresentationRequest, error) {
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
		return nil, fmt.Errorf("resource.PartialResourceRepresentationRequestBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultPartialResourceRepresentationRequestValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *PartialResourceRepresentationRequestBuilder) MustBuild() *PartialResourceRepresentationRequest {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
