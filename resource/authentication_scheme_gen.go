package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	authenticationSchemeDescriptionJSONKey      = "description"
	authenticationSchemeDocumentationURIJSONKey = "documentationUri"
	authenticationSchemeNameJSONKey             = "name"
	authenticationSchemeSpecURIJSONKey          = "specUri"
	authenticationSchemeTypJSONKey              = "type"
)

type AuthenticationScheme struct {
	description      *string
	documentationURI *string
	name             *string
	specURI          *string
	typ              *AuthenticationSchemeType
	privateParams    map[string]interface{}
	mu               sync.RWMutex
}

type AuthenticationSchemeValidator interface {
	Validate(*AuthenticationScheme) error
}

type AuthenticationSchemeValidateFunc func(v *AuthenticationScheme) error

func (f AuthenticationSchemeValidateFunc) Validate(v *AuthenticationScheme) error {
	return f(v)
}

var DefaultAuthenticationSchemeValidator AuthenticationSchemeValidator = AuthenticationSchemeValidateFunc(func(v *AuthenticationScheme) error {
	if v.description == nil {
		return fmt.Errorf(`required field "description" is missing`)
	}
	if v.name == nil {
		return fmt.Errorf(`required field "name" is missing`)
	}
	if v.typ == nil {
		return fmt.Errorf(`required field "type" is missing`)
	}
	return nil
})

func (v *AuthenticationScheme) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.description == nil {
		return ""
	}
	return *(v.description)
}

func (v *AuthenticationScheme) DocumentationURI() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.documentationURI == nil {
		return ""
	}
	return *(v.documentationURI)
}

func (v *AuthenticationScheme) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.name == nil {
		return ""
	}
	return *(v.name)
}

func (v *AuthenticationScheme) SpecURI() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.specURI == nil {
		return ""
	}
	return *(v.specURI)
}

func (v *AuthenticationScheme) Typ() AuthenticationSchemeType {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.typ == nil {
		return InvalidAuthenticationScheme
	}
	return *(v.typ)
}

func (v *AuthenticationScheme) makePairs() []pair {
	pairs := make([]pair, 0, 5)
	if v.description != nil {
		pairs = append(pairs, pair{Key: "description", Value: *(v.description)})
	}
	if v.documentationURI != nil {
		pairs = append(pairs, pair{Key: "documentationUri", Value: *(v.documentationURI)})
	}
	if v.name != nil {
		pairs = append(pairs, pair{Key: "name", Value: *(v.name)})
	}
	if v.specURI != nil {
		pairs = append(pairs, pair{Key: "specUri", Value: *(v.specURI)})
	}
	if v.typ != nil {
		pairs = append(pairs, pair{Key: "type", Value: *(v.typ)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *AuthenticationScheme) MarshalJSON() ([]byte, error) {
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

func (v *AuthenticationScheme) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case authenticationSchemeDescriptionJSONKey:
		if v.description == nil {
			return nil, false
		}
		return *(v.description), true
	case authenticationSchemeDocumentationURIJSONKey:
		if v.documentationURI == nil {
			return nil, false
		}
		return *(v.documentationURI), true
	case authenticationSchemeNameJSONKey:
		if v.name == nil {
			return nil, false
		}
		return *(v.name), true
	case authenticationSchemeSpecURIJSONKey:
		if v.specURI == nil {
			return nil, false
		}
		return *(v.specURI), true
	case authenticationSchemeTypJSONKey:
		if v.typ == nil {
			return nil, false
		}
		return *(v.typ), true
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

func (v *AuthenticationScheme) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case authenticationSchemeDescriptionJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "description", but got %T`, value)
		}
		v.description = &tmp
		return nil
	case authenticationSchemeDocumentationURIJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "documentationUri", but got %T`, value)
		}
		v.documentationURI = &tmp
		return nil
	case authenticationSchemeNameJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "name", but got %T`, value)
		}
		v.name = &tmp
		return nil
	case authenticationSchemeSpecURIJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "specUri", but got %T`, value)
		}
		v.specURI = &tmp
		return nil
	case authenticationSchemeTypJSONKey:
		var tmp AuthenticationSchemeType
		tmp, ok := value.(AuthenticationSchemeType)
		if !ok {
			return fmt.Errorf(`expected AuthenticationSchemeType for field "type", but got %T`, value)
		}
		v.typ = &tmp
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

func (v *AuthenticationScheme) UnmarshalJSON(data []byte) error {
	v.description = nil
	v.documentationURI = nil
	v.name = nil
	v.specURI = nil
	v.typ = nil
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
			case authenticationSchemeDescriptionJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "description": %w`, err)
				}
				v.description = &x
			case authenticationSchemeDocumentationURIJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "documentationUri": %w`, err)
				}
				v.documentationURI = &x
			case authenticationSchemeNameJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "name": %w`, err)
				}
				v.name = &x
			case authenticationSchemeSpecURIJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "specUri": %w`, err)
				}
				v.specURI = &x
			case authenticationSchemeTypJSONKey:
				var x AuthenticationSchemeType
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "type": %w`, err)
				}
				v.typ = &x
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

func (v *AuthenticationScheme) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type AuthenticationSchemeBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator AuthenticationSchemeValidator
	object    *AuthenticationScheme
}

func (b *Builder) AuthenticationScheme() *AuthenticationSchemeBuilder {
	return NewAuthenticationSchemeBuilder()
}

func NewAuthenticationSchemeBuilder() *AuthenticationSchemeBuilder {
	var b AuthenticationSchemeBuilder
	b.init()
	return &b
}

func (b *AuthenticationSchemeBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &AuthenticationScheme{}
}

func (b *AuthenticationSchemeBuilder) Description(v string) *AuthenticationSchemeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("description", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AuthenticationSchemeBuilder) DocumentationURI(v string) *AuthenticationSchemeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("documentationUri", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AuthenticationSchemeBuilder) Name(v string) *AuthenticationSchemeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("name", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AuthenticationSchemeBuilder) SpecURI(v string) *AuthenticationSchemeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("specUri", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AuthenticationSchemeBuilder) Typ(v AuthenticationSchemeType) *AuthenticationSchemeBuilder {
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

func (b *AuthenticationSchemeBuilder) Validator(v AuthenticationSchemeValidator) *AuthenticationSchemeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *AuthenticationSchemeBuilder) Build() (*AuthenticationScheme, error) {
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
		return nil, fmt.Errorf("resource.AuthenticationSchemeBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultAuthenticationSchemeValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *AuthenticationSchemeBuilder) MustBuild() *AuthenticationScheme {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
