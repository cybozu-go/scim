package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	IMSDisplayKey = "display"
	IMSPrimaryKey = "primary"
	IMSTypeKey    = "type"
	IMSValueKey   = "value"
)

type IMS struct {
	display       *string
	primary       *bool
	typ           *string
	value         *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type IMSValidator interface {
	Validate(*IMS) error
}

type IMSValidateFunc func(v *IMS) error

func (f IMSValidateFunc) Validate(v *IMS) error {
	return f(v)
}

var DefaultIMSValidator IMSValidator = IMSValidateFunc(func(v *IMS) error {
	return nil
})

func (v *IMS) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *IMS) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.display == nil {
		return ""
	}
	return *(v.display)
}

func (v *IMS) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *IMS) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.primary == nil {
		return false
	}
	return *(v.primary)
}

func (v *IMS) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *IMS) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.typ == nil {
		return ""
	}
	return *(v.typ)
}

func (v *IMS) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *IMS) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.value == nil {
		return ""
	}
	return *(v.value)
}

func (v *IMS) makePairs() []pair {
	pairs := make([]pair, 0, 4)
	if v.display != nil {
		pairs = append(pairs, pair{Key: "display", Value: *(v.display)})
	}
	if v.primary != nil {
		pairs = append(pairs, pair{Key: "primary", Value: *(v.primary)})
	}
	if v.typ != nil {
		pairs = append(pairs, pair{Key: "type", Value: *(v.typ)})
	}
	if v.value != nil {
		pairs = append(pairs, pair{Key: "value", Value: *(v.value)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *IMS) MarshalJSON() ([]byte, error) {
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

func (v *IMS) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case IMSDisplayKey:
		if v.display == nil {
			return nil, false
		}
		return *(v.display), true
	case IMSPrimaryKey:
		if v.primary == nil {
			return nil, false
		}
		return *(v.primary), true
	case IMSTypeKey:
		if v.typ == nil {
			return nil, false
		}
		return *(v.typ), true
	case IMSValueKey:
		if v.value == nil {
			return nil, false
		}
		return *(v.value), true
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

func (v *IMS) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case IMSDisplayKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "display", but got %T`, value)
		}
		v.display = &tmp
		return nil
	case IMSPrimaryKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "primary", but got %T`, value)
		}
		v.primary = &tmp
		return nil
	case IMSTypeKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "type", but got %T`, value)
		}
		v.typ = &tmp
		return nil
	case IMSValueKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "value", but got %T`, value)
		}
		v.value = &tmp
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

func (v *IMS) Clone() *IMS {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &IMS{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *IMS) UnmarshalJSON(data []byte) error {
	v.display = nil
	v.primary = nil
	v.typ = nil
	v.value = nil
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
			case IMSDisplayKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "display": %w`, err)
				}
				v.display = &x
			case IMSPrimaryKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "primary": %w`, err)
				}
				v.primary = &x
			case IMSTypeKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "type": %w`, err)
				}
				v.typ = &x
			case IMSValueKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "value": %w`, err)
				}
				v.value = &x
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

func (v *IMS) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// IMSBuilder creates a IMS resource
type IMSBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator IMSValidator
	object    *IMS
}

func (b *Builder) IMS() *IMSBuilder {
	return NewIMSBuilder()
}

func NewIMSBuilder() *IMSBuilder {
	var b IMSBuilder
	b.init()
	return &b
}

func (b *IMSBuilder) From(in *IMS) *IMSBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *IMSBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &IMS{}
}

func (b *IMSBuilder) Display(v string) *IMSBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("display", v); err != nil {
		b.err = err
	}
	return b
}

func (b *IMSBuilder) Primary(v bool) *IMSBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("primary", v); err != nil {
		b.err = err
	}
	return b
}

func (b *IMSBuilder) Type(v string) *IMSBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("type", v); err != nil {
		b.err = err
	}
	return b
}

func (b *IMSBuilder) Value(v string) *IMSBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("value", v); err != nil {
		b.err = err
	}
	return b
}

func (b *IMSBuilder) Validator(v IMSValidator) *IMSBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *IMSBuilder) Build() (*IMS, error) {
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
		return nil, fmt.Errorf("resource.IMSBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultIMSValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *IMSBuilder) MustBuild() *IMS {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
