package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	X509CertificateDisplayKey = "display"
	X509CertificatePrimaryKey = "primary"
	X509CertificateTypeKey    = "type"
	X509CertificateValueKey   = "value"
)

type X509Certificate struct {
	display       *string
	primary       *bool
	typ           *string
	value         *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type X509CertificateValidator interface {
	Validate(*X509Certificate) error
}

type X509CertificateValidateFunc func(v *X509Certificate) error

func (f X509CertificateValidateFunc) Validate(v *X509Certificate) error {
	return f(v)
}

var DefaultX509CertificateValidator X509CertificateValidator = X509CertificateValidateFunc(func(v *X509Certificate) error {
	return nil
})

func (v *X509Certificate) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *X509Certificate) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.display == nil {
		return ""
	}
	return *(v.display)
}

func (v *X509Certificate) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *X509Certificate) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.primary == nil {
		return false
	}
	return *(v.primary)
}

func (v *X509Certificate) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *X509Certificate) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.typ == nil {
		return ""
	}
	return *(v.typ)
}

func (v *X509Certificate) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *X509Certificate) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.value == nil {
		return ""
	}
	return *(v.value)
}

func (v *X509Certificate) makePairs() []pair {
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

func (v *X509Certificate) MarshalJSON() ([]byte, error) {
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

func (v *X509Certificate) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case X509CertificateDisplayKey:
		if v.display == nil {
			return nil, false
		}
		return *(v.display), true
	case X509CertificatePrimaryKey:
		if v.primary == nil {
			return nil, false
		}
		return *(v.primary), true
	case X509CertificateTypeKey:
		if v.typ == nil {
			return nil, false
		}
		return *(v.typ), true
	case X509CertificateValueKey:
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

func (v *X509Certificate) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case X509CertificateDisplayKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "display", but got %T`, value)
		}
		v.display = &tmp
		return nil
	case X509CertificatePrimaryKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "primary", but got %T`, value)
		}
		v.primary = &tmp
		return nil
	case X509CertificateTypeKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "type", but got %T`, value)
		}
		v.typ = &tmp
		return nil
	case X509CertificateValueKey:
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

func (v *X509Certificate) Clone() *X509Certificate {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &X509Certificate{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *X509Certificate) UnmarshalJSON(data []byte) error {
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
			case X509CertificateDisplayKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "display": %w`, err)
				}
				v.display = &x
			case X509CertificatePrimaryKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "primary": %w`, err)
				}
				v.primary = &x
			case X509CertificateTypeKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "type": %w`, err)
				}
				v.typ = &x
			case X509CertificateValueKey:
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

func (v *X509Certificate) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type X509CertificateBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator X509CertificateValidator
	object    *X509Certificate
}

func (b *Builder) X509Certificate() *X509CertificateBuilder {
	return NewX509CertificateBuilder()
}

func NewX509CertificateBuilder() *X509CertificateBuilder {
	var b X509CertificateBuilder
	b.init()
	return &b
}

func (b *X509CertificateBuilder) From(in *X509Certificate) *X509CertificateBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *X509CertificateBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &X509Certificate{}
}

func (b *X509CertificateBuilder) Display(v string) *X509CertificateBuilder {
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

func (b *X509CertificateBuilder) Primary(v bool) *X509CertificateBuilder {
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

func (b *X509CertificateBuilder) Type(v string) *X509CertificateBuilder {
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

func (b *X509CertificateBuilder) Value(v string) *X509CertificateBuilder {
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

func (b *X509CertificateBuilder) Validator(v X509CertificateValidator) *X509CertificateBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *X509CertificateBuilder) Build() (*X509Certificate, error) {
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
		return nil, fmt.Errorf("resource.X509CertificateBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultX509CertificateValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *X509CertificateBuilder) MustBuild() *X509Certificate {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
