package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	emailDisplayJSONKey = "display"
	emailPrimaryJSONKey = "primary"
	emailTypeJSONKey    = "type"
	emailValueJSONKey   = "value"
)

type Email struct {
	display       *string
	primary       *bool
	typ           *string
	value         *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type EmailValidator interface {
	Validate(*Email) error
}

type EmailValidateFunc func(v *Email) error

func (f EmailValidateFunc) Validate(v *Email) error {
	return f(v)
}

var DefaultEmailValidator EmailValidator = EmailValidateFunc(func(v *Email) error {
	return nil
})

func (v *Email) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *Email) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.display == nil {
		return ""
	}
	return *(v.display)
}

func (v *Email) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *Email) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.primary == nil {
		return false
	}
	return *(v.primary)
}

func (v *Email) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *Email) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.typ == nil {
		return ""
	}
	return *(v.typ)
}

func (v *Email) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *Email) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.value == nil {
		return ""
	}
	return *(v.value)
}

func (v *Email) makePairs() []pair {
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

func (v *Email) MarshalJSON() ([]byte, error) {
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

func (v *Email) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case emailDisplayJSONKey:
		if v.display == nil {
			return nil, false
		}
		return *(v.display), true
	case emailPrimaryJSONKey:
		if v.primary == nil {
			return nil, false
		}
		return *(v.primary), true
	case emailTypeJSONKey:
		if v.typ == nil {
			return nil, false
		}
		return *(v.typ), true
	case emailValueJSONKey:
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

func (v *Email) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case emailDisplayJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "display", but got %T`, value)
		}
		v.display = &tmp
		return nil
	case emailPrimaryJSONKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "primary", but got %T`, value)
		}
		v.primary = &tmp
		return nil
	case emailTypeJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "type", but got %T`, value)
		}
		v.typ = &tmp
		return nil
	case emailValueJSONKey:
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

func (v *Email) Clone() *Email {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &Email{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *Email) UnmarshalJSON(data []byte) error {
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
			case emailDisplayJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "display": %w`, err)
				}
				v.display = &x
			case emailPrimaryJSONKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "primary": %w`, err)
				}
				v.primary = &x
			case emailTypeJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "type": %w`, err)
				}
				v.typ = &x
			case emailValueJSONKey:
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

func (v *Email) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type EmailBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator EmailValidator
	object    *Email
}

func (b *Builder) Email() *EmailBuilder {
	return NewEmailBuilder()
}

func NewEmailBuilder() *EmailBuilder {
	var b EmailBuilder
	b.init()
	return &b
}

func (b *EmailBuilder) From(in *Email) *EmailBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *EmailBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &Email{}
}

func (b *EmailBuilder) Display(v string) *EmailBuilder {
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

func (b *EmailBuilder) Primary(v bool) *EmailBuilder {
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

func (b *EmailBuilder) Type(v string) *EmailBuilder {
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

func (b *EmailBuilder) Value(v string) *EmailBuilder {
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

func (b *EmailBuilder) Validator(v EmailValidator) *EmailBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *EmailBuilder) Build() (*Email, error) {
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
		return nil, fmt.Errorf("resource.EmailBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultEmailValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *EmailBuilder) MustBuild() *Email {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
