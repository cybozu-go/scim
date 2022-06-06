package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	bulkSupportMaxOperationsJSONKey  = "maxOperations"
	bulkSupportMaxPayloadSizeJSONKey = "maxPayloadSize"
	bulkSupportSupportedJSONKey      = "supported"
)

type BulkSupport struct {
	maxOperations  *int
	maxPayloadSize *int
	supported      *bool
	privateParams  map[string]interface{}
	mu             sync.RWMutex
}

type BulkSupportValidator interface {
	Validate(*BulkSupport) error
}

type BulkSupportValidateFunc func(v *BulkSupport) error

func (f BulkSupportValidateFunc) Validate(v *BulkSupport) error {
	return f(v)
}

var DefaultBulkSupportValidator BulkSupportValidator = BulkSupportValidateFunc(func(v *BulkSupport) error {
	if v.maxOperations == nil {
		return fmt.Errorf(`required field "maxOperations" is missing in "BulkSupport"`)
	}
	if v.maxPayloadSize == nil {
		return fmt.Errorf(`required field "maxPayloadSize" is missing in "BulkSupport"`)
	}
	if v.supported == nil {
		return fmt.Errorf(`required field "supported" is missing in "BulkSupport"`)
	}
	return nil
})

func (v *BulkSupport) HasMaxOperations() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxOperations != nil
}

func (v *BulkSupport) MaxOperations() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.maxOperations == nil {
		return 0
	}
	return *(v.maxOperations)
}

func (v *BulkSupport) HasMaxPayloadSize() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxPayloadSize != nil
}

func (v *BulkSupport) MaxPayloadSize() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.maxPayloadSize == nil {
		return 0
	}
	return *(v.maxPayloadSize)
}

func (v *BulkSupport) HasSupported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.supported != nil
}

func (v *BulkSupport) Supported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.supported == nil {
		return false
	}
	return *(v.supported)
}

func (v *BulkSupport) makePairs() []pair {
	pairs := make([]pair, 0, 3)
	if v.maxOperations != nil {
		pairs = append(pairs, pair{Key: "maxOperations", Value: *(v.maxOperations)})
	}
	if v.maxPayloadSize != nil {
		pairs = append(pairs, pair{Key: "maxPayloadSize", Value: *(v.maxPayloadSize)})
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

func (v *BulkSupport) MarshalJSON() ([]byte, error) {
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

func (v *BulkSupport) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case bulkSupportMaxOperationsJSONKey:
		if v.maxOperations == nil {
			return nil, false
		}
		return *(v.maxOperations), true
	case bulkSupportMaxPayloadSizeJSONKey:
		if v.maxPayloadSize == nil {
			return nil, false
		}
		return *(v.maxPayloadSize), true
	case bulkSupportSupportedJSONKey:
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

func (v *BulkSupport) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case bulkSupportMaxOperationsJSONKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "maxOperations", but got %T`, value)
		}
		v.maxOperations = &tmp
		return nil
	case bulkSupportMaxPayloadSizeJSONKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "maxPayloadSize", but got %T`, value)
		}
		v.maxPayloadSize = &tmp
		return nil
	case bulkSupportSupportedJSONKey:
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

func (v *BulkSupport) Clone() *BulkSupport {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &BulkSupport{
		maxOperations:  v.maxOperations,
		maxPayloadSize: v.maxPayloadSize,
		supported:      v.supported,
	}
}

func (v *BulkSupport) UnmarshalJSON(data []byte) error {
	v.maxOperations = nil
	v.maxPayloadSize = nil
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
			case bulkSupportMaxOperationsJSONKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "maxOperations": %w`, err)
				}
				v.maxOperations = &x
			case bulkSupportMaxPayloadSizeJSONKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "maxPayloadSize": %w`, err)
				}
				v.maxPayloadSize = &x
			case bulkSupportSupportedJSONKey:
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

func (v *BulkSupport) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type BulkSupportBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator BulkSupportValidator
	object    *BulkSupport
}

func (b *Builder) BulkSupport() *BulkSupportBuilder {
	return NewBulkSupportBuilder()
}

func NewBulkSupportBuilder() *BulkSupportBuilder {
	var b BulkSupportBuilder
	b.init()
	return &b
}

func (b *BulkSupportBuilder) From(in *BulkSupport) *BulkSupportBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *BulkSupportBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &BulkSupport{}
}

func (b *BulkSupportBuilder) MaxOperations(v int) *BulkSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("maxOperations", v); err != nil {
		b.err = err
	}
	return b
}

func (b *BulkSupportBuilder) MaxPayloadSize(v int) *BulkSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("maxPayloadSize", v); err != nil {
		b.err = err
	}
	return b
}

func (b *BulkSupportBuilder) Supported(v bool) *BulkSupportBuilder {
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

func (b *BulkSupportBuilder) Validator(v BulkSupportValidator) *BulkSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *BulkSupportBuilder) Build() (*BulkSupport, error) {
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
		return nil, fmt.Errorf("resource.BulkSupportBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultBulkSupportValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *BulkSupportBuilder) MustBuild() *BulkSupport {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
