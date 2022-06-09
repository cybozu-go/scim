package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	NamesFamilyNameKey      = "familyName"
	NamesFormattedKey       = "formatted"
	NamesGivenNameKey       = "givenName"
	NamesHonorificPrefixKey = "honorificPrefix"
	NamesHonorificSuffixKey = "honorificSuffix"
	NamesMiddleNameKey      = "middleName"
)

type Names struct {
	familyName      *string
	formatted       *string
	givenName       *string
	honorificPrefix *string
	honorificSuffix *string
	middleName      *string
	privateParams   map[string]interface{}
	mu              sync.RWMutex
}

type NamesValidator interface {
	Validate(*Names) error
}

type NamesValidateFunc func(v *Names) error

func (f NamesValidateFunc) Validate(v *Names) error {
	return f(v)
}

var DefaultNamesValidator NamesValidator = NamesValidateFunc(func(v *Names) error {
	return nil
})

func (v *Names) HasFamilyName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.familyName != nil
}

func (v *Names) FamilyName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.familyName == nil {
		return ""
	}
	return *(v.familyName)
}

func (v *Names) HasFormatted() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.formatted != nil
}

func (v *Names) Formatted() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.formatted == nil {
		return ""
	}
	return *(v.formatted)
}

func (v *Names) HasGivenName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.givenName != nil
}

func (v *Names) GivenName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.givenName == nil {
		return ""
	}
	return *(v.givenName)
}

func (v *Names) HasHonorificPrefix() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.honorificPrefix != nil
}

func (v *Names) HonorificPrefix() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.honorificPrefix == nil {
		return ""
	}
	return *(v.honorificPrefix)
}

func (v *Names) HasHonorificSuffix() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.honorificSuffix != nil
}

func (v *Names) HonorificSuffix() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.honorificSuffix == nil {
		return ""
	}
	return *(v.honorificSuffix)
}

func (v *Names) HasMiddleName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.middleName != nil
}

func (v *Names) MiddleName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.middleName == nil {
		return ""
	}
	return *(v.middleName)
}

func (v *Names) makePairs() []pair {
	pairs := make([]pair, 0, 6)
	if v.familyName != nil {
		pairs = append(pairs, pair{Key: "familyName", Value: *(v.familyName)})
	}
	if v.formatted != nil {
		pairs = append(pairs, pair{Key: "formatted", Value: *(v.formatted)})
	}
	if v.givenName != nil {
		pairs = append(pairs, pair{Key: "givenName", Value: *(v.givenName)})
	}
	if v.honorificPrefix != nil {
		pairs = append(pairs, pair{Key: "honorificPrefix", Value: *(v.honorificPrefix)})
	}
	if v.honorificSuffix != nil {
		pairs = append(pairs, pair{Key: "honorificSuffix", Value: *(v.honorificSuffix)})
	}
	if v.middleName != nil {
		pairs = append(pairs, pair{Key: "middleName", Value: *(v.middleName)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *Names) MarshalJSON() ([]byte, error) {
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

func (v *Names) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case NamesFamilyNameKey:
		if v.familyName == nil {
			return nil, false
		}
		return *(v.familyName), true
	case NamesFormattedKey:
		if v.formatted == nil {
			return nil, false
		}
		return *(v.formatted), true
	case NamesGivenNameKey:
		if v.givenName == nil {
			return nil, false
		}
		return *(v.givenName), true
	case NamesHonorificPrefixKey:
		if v.honorificPrefix == nil {
			return nil, false
		}
		return *(v.honorificPrefix), true
	case NamesHonorificSuffixKey:
		if v.honorificSuffix == nil {
			return nil, false
		}
		return *(v.honorificSuffix), true
	case NamesMiddleNameKey:
		if v.middleName == nil {
			return nil, false
		}
		return *(v.middleName), true
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

func (v *Names) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case NamesFamilyNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "familyName", but got %T`, value)
		}
		v.familyName = &tmp
		return nil
	case NamesFormattedKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "formatted", but got %T`, value)
		}
		v.formatted = &tmp
		return nil
	case NamesGivenNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "givenName", but got %T`, value)
		}
		v.givenName = &tmp
		return nil
	case NamesHonorificPrefixKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "honorificPrefix", but got %T`, value)
		}
		v.honorificPrefix = &tmp
		return nil
	case NamesHonorificSuffixKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "honorificSuffix", but got %T`, value)
		}
		v.honorificSuffix = &tmp
		return nil
	case NamesMiddleNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "middleName", but got %T`, value)
		}
		v.middleName = &tmp
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

func (v *Names) Clone() *Names {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &Names{
		familyName:      v.familyName,
		formatted:       v.formatted,
		givenName:       v.givenName,
		honorificPrefix: v.honorificPrefix,
		honorificSuffix: v.honorificSuffix,
		middleName:      v.middleName,
	}
}

func (v *Names) UnmarshalJSON(data []byte) error {
	v.familyName = nil
	v.formatted = nil
	v.givenName = nil
	v.honorificPrefix = nil
	v.honorificSuffix = nil
	v.middleName = nil
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
			case NamesFamilyNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "familyName": %w`, err)
				}
				v.familyName = &x
			case NamesFormattedKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "formatted": %w`, err)
				}
				v.formatted = &x
			case NamesGivenNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "givenName": %w`, err)
				}
				v.givenName = &x
			case NamesHonorificPrefixKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "honorificPrefix": %w`, err)
				}
				v.honorificPrefix = &x
			case NamesHonorificSuffixKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "honorificSuffix": %w`, err)
				}
				v.honorificSuffix = &x
			case NamesMiddleNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "middleName": %w`, err)
				}
				v.middleName = &x
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

func (v *Names) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type NamesBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator NamesValidator
	object    *Names
}

func (b *Builder) Names() *NamesBuilder {
	return NewNamesBuilder()
}

func NewNamesBuilder() *NamesBuilder {
	var b NamesBuilder
	b.init()
	return &b
}

func (b *NamesBuilder) From(in *Names) *NamesBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *NamesBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &Names{}
}

func (b *NamesBuilder) FamilyName(v string) *NamesBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("familyName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *NamesBuilder) Formatted(v string) *NamesBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("formatted", v); err != nil {
		b.err = err
	}
	return b
}

func (b *NamesBuilder) GivenName(v string) *NamesBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("givenName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *NamesBuilder) HonorificPrefix(v string) *NamesBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("honorificPrefix", v); err != nil {
		b.err = err
	}
	return b
}

func (b *NamesBuilder) HonorificSuffix(v string) *NamesBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("honorificSuffix", v); err != nil {
		b.err = err
	}
	return b
}

func (b *NamesBuilder) MiddleName(v string) *NamesBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("middleName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *NamesBuilder) Validator(v NamesValidator) *NamesBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *NamesBuilder) Build() (*Names, error) {
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
		return nil, fmt.Errorf("resource.NamesBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultNamesValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *NamesBuilder) MustBuild() *Names {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
