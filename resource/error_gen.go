package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	ErrorDetailKey   = "detail"
	ErrorScimTypeKey = "scimType"
	ErrorStatusKey   = "status"
)

type Error struct {
	detail        *string
	scimType      *ErrorType
	status        *int
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type ErrorValidator interface {
	Validate(*Error) error
}

type ErrorValidateFunc func(v *Error) error

func (f ErrorValidateFunc) Validate(v *Error) error {
	return f(v)
}

var DefaultErrorValidator ErrorValidator = ErrorValidateFunc(func(v *Error) error {
	return nil
})

func (v *Error) HasDetail() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.detail != nil
}

func (v *Error) Detail() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.detail == nil {
		return ""
	}
	return *(v.detail)
}

func (v *Error) HasScimType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.scimType != nil
}

func (v *Error) ScimType() ErrorType {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.scimType == nil {
		return ErrUnknown
	}
	return *(v.scimType)
}

func (v *Error) HasStatus() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.status != nil
}

func (v *Error) Status() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.status == nil {
		return 0
	}
	return *(v.status)
}

func (v *Error) makePairs() []pair {
	pairs := make([]pair, 0, 3)
	if v.detail != nil {
		pairs = append(pairs, pair{Key: "detail", Value: *(v.detail)})
	}
	if v.scimType != nil {
		pairs = append(pairs, pair{Key: "scimType", Value: *(v.scimType)})
	}
	if v.status != nil {
		pairs = append(pairs, pair{Key: "status", Value: *(v.status)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *Error) MarshalJSON() ([]byte, error) {
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

func (v *Error) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case ErrorDetailKey:
		if v.detail == nil {
			return nil, false
		}
		return *(v.detail), true
	case ErrorScimTypeKey:
		if v.scimType == nil {
			return nil, false
		}
		return *(v.scimType), true
	case ErrorStatusKey:
		if v.status == nil {
			return nil, false
		}
		return *(v.status), true
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

func (v *Error) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case ErrorDetailKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "detail", but got %T`, value)
		}
		v.detail = &tmp
		return nil
	case ErrorScimTypeKey:
		var tmp ErrorType
		tmp, ok := value.(ErrorType)
		if !ok {
			return fmt.Errorf(`expected ErrorType for field "scimType", but got %T`, value)
		}
		v.scimType = &tmp
		return nil
	case ErrorStatusKey:
		var tmp int
		tmp, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected int for field "status", but got %T`, value)
		}
		v.status = &tmp
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

func (v *Error) Clone() *Error {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &Error{
		detail:   v.detail,
		scimType: v.scimType,
		status:   v.status,
	}
}

func (v *Error) UnmarshalJSON(data []byte) error {
	v.detail = nil
	v.scimType = nil
	v.status = nil
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
			case ErrorDetailKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "detail": %w`, err)
				}
				v.detail = &x
			case ErrorScimTypeKey:
				var x ErrorType
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "scimType": %w`, err)
				}
				v.scimType = &x
			case ErrorStatusKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "status": %w`, err)
				}
				v.status = &x
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

func (v *Error) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// ErrorBuilder creates a Error resource
type ErrorBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator ErrorValidator
	object    *Error
}

func (b *Builder) Error() *ErrorBuilder {
	return NewErrorBuilder()
}

func NewErrorBuilder() *ErrorBuilder {
	var b ErrorBuilder
	b.init()
	return &b
}

func (b *ErrorBuilder) From(in *Error) *ErrorBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *ErrorBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &Error{}
}

func (b *ErrorBuilder) Detail(v string) *ErrorBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("detail", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ErrorBuilder) ScimType(v ErrorType) *ErrorBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("scimType", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ErrorBuilder) Status(v int) *ErrorBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("status", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ErrorBuilder) Validator(v ErrorValidator) *ErrorBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *ErrorBuilder) Build() (*Error, error) {
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
		return nil, fmt.Errorf("resource.ErrorBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultErrorValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *ErrorBuilder) MustBuild() *Error {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
