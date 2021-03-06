package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	EnterpriseManagerDisplayNameKey = "displayName"
	EnterpriseManagerIDKey          = "id"
	EnterpriseManagerReferenceKey   = "$ref"
)

type EnterpriseManager struct {
	displayName   *string
	id            *string
	ref           *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type EnterpriseManagerValidator interface {
	Validate(*EnterpriseManager) error
}

type EnterpriseManagerValidateFunc func(v *EnterpriseManager) error

func (f EnterpriseManagerValidateFunc) Validate(v *EnterpriseManager) error {
	return f(v)
}

var DefaultEnterpriseManagerValidator EnterpriseManagerValidator = EnterpriseManagerValidateFunc(func(v *EnterpriseManager) error {
	return nil
})

func (v *EnterpriseManager) HasDisplayName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.displayName != nil
}

func (v *EnterpriseManager) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.displayName == nil {
		return ""
	}
	return *(v.displayName)
}

func (v *EnterpriseManager) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *EnterpriseManager) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *EnterpriseManager) HasReference() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ref != nil
}

func (v *EnterpriseManager) Reference() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.ref == nil {
		return ""
	}
	return *(v.ref)
}

func (v *EnterpriseManager) makePairs() []pair {
	pairs := make([]pair, 0, 3)
	if v.displayName != nil {
		pairs = append(pairs, pair{Key: "displayName", Value: *(v.displayName)})
	}
	if v.id != nil {
		pairs = append(pairs, pair{Key: "id", Value: *(v.id)})
	}
	if v.ref != nil {
		pairs = append(pairs, pair{Key: "$ref", Value: *(v.ref)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *EnterpriseManager) MarshalJSON() ([]byte, error) {
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

func (v *EnterpriseManager) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case EnterpriseManagerDisplayNameKey:
		if v.displayName == nil {
			return nil, false
		}
		return *(v.displayName), true
	case EnterpriseManagerIDKey:
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case EnterpriseManagerReferenceKey:
		if v.ref == nil {
			return nil, false
		}
		return *(v.ref), true
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

func (v *EnterpriseManager) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case EnterpriseManagerDisplayNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "displayName", but got %T`, value)
		}
		v.displayName = &tmp
		return nil
	case EnterpriseManagerIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case EnterpriseManagerReferenceKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "$ref", but got %T`, value)
		}
		v.ref = &tmp
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

func (v *EnterpriseManager) Clone() *EnterpriseManager {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &EnterpriseManager{
		displayName: v.displayName,
		id:          v.id,
		ref:         v.ref,
	}
}

func (v *EnterpriseManager) UnmarshalJSON(data []byte) error {
	v.displayName = nil
	v.id = nil
	v.ref = nil
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
			case EnterpriseManagerDisplayNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "displayName": %w`, err)
				}
				v.displayName = &x
			case EnterpriseManagerIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case EnterpriseManagerReferenceKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "$ref": %w`, err)
				}
				v.ref = &x
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

func (v *EnterpriseManager) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type EnterpriseManagerBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator EnterpriseManagerValidator
	object    *EnterpriseManager
}

func (b *Builder) EnterpriseManager() *EnterpriseManagerBuilder {
	return NewEnterpriseManagerBuilder()
}

func NewEnterpriseManagerBuilder() *EnterpriseManagerBuilder {
	var b EnterpriseManagerBuilder
	b.init()
	return &b
}

func (b *EnterpriseManagerBuilder) From(in *EnterpriseManager) *EnterpriseManagerBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *EnterpriseManagerBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &EnterpriseManager{}
}

func (b *EnterpriseManagerBuilder) DisplayName(v string) *EnterpriseManagerBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("displayName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseManagerBuilder) ID(v string) *EnterpriseManagerBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("id", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseManagerBuilder) Reference(v string) *EnterpriseManagerBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("$ref", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseManagerBuilder) Validator(v EnterpriseManagerValidator) *EnterpriseManagerBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *EnterpriseManagerBuilder) Build() (*EnterpriseManager, error) {
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
		return nil, fmt.Errorf("resource.EnterpriseManagerBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultEnterpriseManagerValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *EnterpriseManagerBuilder) MustBuild() *EnterpriseManager {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
