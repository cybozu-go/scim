package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	GroupMemberDisplayKey = "display"
	GroupMemberRefKey     = "$ref"
	GroupMemberTypeKey    = "type"
	GroupMemberValueKey   = "value"
)

type GroupMember struct {
	display       *string
	ref           *string
	typ           *string
	value         *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type GroupMemberValidator interface {
	Validate(*GroupMember) error
}

type GroupMemberValidateFunc func(v *GroupMember) error

func (f GroupMemberValidateFunc) Validate(v *GroupMember) error {
	return f(v)
}

var DefaultGroupMemberValidator GroupMemberValidator = GroupMemberValidateFunc(func(v *GroupMember) error {
	return nil
})

func (v *GroupMember) HasDisplay() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.display != nil
}

func (v *GroupMember) Display() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.display == nil {
		return ""
	}
	return *(v.display)
}

func (v *GroupMember) HasRef() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ref != nil
}

func (v *GroupMember) Ref() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.ref == nil {
		return ""
	}
	return *(v.ref)
}

func (v *GroupMember) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *GroupMember) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.typ == nil {
		return ""
	}
	return *(v.typ)
}

func (v *GroupMember) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *GroupMember) Value() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.value == nil {
		return ""
	}
	return *(v.value)
}

func (v *GroupMember) makePairs() []pair {
	pairs := make([]pair, 0, 4)
	if v.display != nil {
		pairs = append(pairs, pair{Key: "display", Value: *(v.display)})
	}
	if v.ref != nil {
		pairs = append(pairs, pair{Key: "$ref", Value: *(v.ref)})
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

func (v *GroupMember) MarshalJSON() ([]byte, error) {
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

func (v *GroupMember) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case GroupMemberDisplayKey:
		if v.display == nil {
			return nil, false
		}
		return *(v.display), true
	case GroupMemberRefKey:
		if v.ref == nil {
			return nil, false
		}
		return *(v.ref), true
	case GroupMemberTypeKey:
		if v.typ == nil {
			return nil, false
		}
		return *(v.typ), true
	case GroupMemberValueKey:
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

func (v *GroupMember) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case GroupMemberDisplayKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "display", but got %T`, value)
		}
		v.display = &tmp
		return nil
	case GroupMemberRefKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "$ref", but got %T`, value)
		}
		v.ref = &tmp
		return nil
	case GroupMemberTypeKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "type", but got %T`, value)
		}
		v.typ = &tmp
		return nil
	case GroupMemberValueKey:
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

func (v *GroupMember) Clone() *GroupMember {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &GroupMember{
		display: v.display,
		ref:     v.ref,
		typ:     v.typ,
		value:   v.value,
	}
}

func (v *GroupMember) UnmarshalJSON(data []byte) error {
	v.display = nil
	v.ref = nil
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
			case GroupMemberDisplayKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "display": %w`, err)
				}
				v.display = &x
			case GroupMemberRefKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "$ref": %w`, err)
				}
				v.ref = &x
			case GroupMemberTypeKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "type": %w`, err)
				}
				v.typ = &x
			case GroupMemberValueKey:
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

func (v *GroupMember) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// GroupMemberBuilder creates a GroupMember resource
type GroupMemberBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator GroupMemberValidator
	object    *GroupMember
}

func (b *Builder) GroupMember() *GroupMemberBuilder {
	return NewGroupMemberBuilder()
}

func NewGroupMemberBuilder() *GroupMemberBuilder {
	var b GroupMemberBuilder
	b.init()
	return &b
}

func (b *GroupMemberBuilder) From(in *GroupMember) *GroupMemberBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *GroupMemberBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &GroupMember{}
}

func (b *GroupMemberBuilder) Display(v string) *GroupMemberBuilder {
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

func (b *GroupMemberBuilder) Ref(v string) *GroupMemberBuilder {
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

func (b *GroupMemberBuilder) Type(v string) *GroupMemberBuilder {
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

func (b *GroupMemberBuilder) Value(v string) *GroupMemberBuilder {
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

func (b *GroupMemberBuilder) Validator(v GroupMemberValidator) *GroupMemberBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *GroupMemberBuilder) Build() (*GroupMember, error) {
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
		return nil, fmt.Errorf("resource.GroupMemberBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultGroupMemberValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *GroupMemberBuilder) MustBuild() *GroupMember {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
