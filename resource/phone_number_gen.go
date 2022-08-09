package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"sync"
)

// JSON key names for PhoneNumber resource
const (
	PhoneNumberDisplayKey = "display"
	PhoneNumberPrimaryKey = "primary"
	PhoneNumberTypeKey    = "type"
	PhoneNumberValueKey   = "value"
)

type PhoneNumber struct {
	display       *string
	primary       *bool
	typ           *string
	value         *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type PhoneNumberValidator interface {
	Validate(*PhoneNumber) error
}

type PhoneNumberValidateFunc func(v *PhoneNumber) error

func (f PhoneNumberValidateFunc) Validate(v *PhoneNumber) error {
	return f(v)
}

var DefaultPhoneNumberValidator PhoneNumberValidator = PhoneNumberValidateFunc(func(v *PhoneNumber) error {
	// This is a very simplified version of validating RFC3966.
	// We just make sure that it's a valid URI, and that the schema
	// is "tel"
	if ptr := v.value; ptr != nil {
		u, err := url.Parse(*ptr)
		if err != nil {
			return fmt.Errorf(`failed to parse phone number: %w`, err)
		}
		if u.Scheme != `tel` {
			return fmt.Errorf(`phone number scheme must be "tel", got %s`, u.Scheme)
		}
	}

	return nil
})

func (v *PhoneNumber) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *PhoneNumber) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.display == nil {
		return ""
	}
	return *(v.display)
}

func (v *PhoneNumber) HasPrimary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.primary != nil
}

func (v *PhoneNumber) Primary() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.primary == nil {
		return false
	}
	return *(v.primary)
}

func (v *PhoneNumber) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *PhoneNumber) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.typ == nil {
		return ""
	}
	return *(v.typ)
}

func (v *PhoneNumber) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *PhoneNumber) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.value == nil {
		return ""
	}
	return *(v.value)
}

func (v *PhoneNumber) makePairs() []pair {
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

func (v *PhoneNumber) MarshalJSON() ([]byte, error) {
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

func (v *PhoneNumber) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case PhoneNumberDisplayKey:
		if v.display == nil {
			return nil, false
		}
		return *(v.display), true
	case PhoneNumberPrimaryKey:
		if v.primary == nil {
			return nil, false
		}
		return *(v.primary), true
	case PhoneNumberTypeKey:
		if v.typ == nil {
			return nil, false
		}
		return *(v.typ), true
	case PhoneNumberValueKey:
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

func (v *PhoneNumber) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case PhoneNumberDisplayKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "display", but got %T`, value)
		}
		v.display = &tmp
		return nil
	case PhoneNumberPrimaryKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "primary", but got %T`, value)
		}
		v.primary = &tmp
		return nil
	case PhoneNumberTypeKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "type", but got %T`, value)
		}
		v.typ = &tmp
		return nil
	case PhoneNumberValueKey:
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

func (v *PhoneNumber) Clone() *PhoneNumber {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &PhoneNumber{
		display: v.display,
		primary: v.primary,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *PhoneNumber) UnmarshalJSON(data []byte) error {
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
			case PhoneNumberDisplayKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "display": %w`, err)
				}
				v.display = &x
			case PhoneNumberPrimaryKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "primary": %w`, err)
				}
				v.primary = &x
			case PhoneNumberTypeKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "type": %w`, err)
				}
				v.typ = &x
			case PhoneNumberValueKey:
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

func (v *PhoneNumber) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// PhoneNumberBuilder creates a PhoneNumber resource
type PhoneNumberBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator PhoneNumberValidator
	object    *PhoneNumber
}

func (b *Builder) PhoneNumber() *PhoneNumberBuilder {
	return NewPhoneNumberBuilder()
}

func NewPhoneNumberBuilder() *PhoneNumberBuilder {
	var b PhoneNumberBuilder
	b.init()
	return &b
}

func (b *PhoneNumberBuilder) From(in *PhoneNumber) *PhoneNumberBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *PhoneNumberBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &PhoneNumber{}
}

func (b *PhoneNumberBuilder) Display(v string) *PhoneNumberBuilder {
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

func (b *PhoneNumberBuilder) Primary(v bool) *PhoneNumberBuilder {
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

func (b *PhoneNumberBuilder) Type(v string) *PhoneNumberBuilder {
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

func (b *PhoneNumberBuilder) Value(v string) *PhoneNumberBuilder {
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

func (b *PhoneNumberBuilder) Validator(v PhoneNumberValidator) *PhoneNumberBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *PhoneNumberBuilder) Build() (*PhoneNumber, error) {
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
		return nil, fmt.Errorf("resource.PhoneNumberBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultPhoneNumberValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *PhoneNumberBuilder) MustBuild() *PhoneNumber {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
