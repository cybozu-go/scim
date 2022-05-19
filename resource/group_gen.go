package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

type Group struct {
	displayName   *string
	externalID    *string
	id            *string
	members       []*User
	meta          *Meta
	schemas       []string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type GroupValidator interface {
	Validate(*Group) error
}

type GroupValidateFunc func(v *Group) error

func (f GroupValidateFunc) Validate(v *Group) error {
	return f(v)
}

var DefaultGroupValidator GroupValidator

func (v *Group) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.displayName == nil {
		return ""
	}
	return *(v.displayName)
}

func (v *Group) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.externalID == nil {
		return ""
	}
	return *(v.externalID)
}

func (v *Group) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *Group) Members() []*User {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.members
}

func (v *Group) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta
}

func (v *Group) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas
}

func (v *Group) MarshalJSON() ([]byte, error) {
	type pair struct {
		Key   string
		Value interface{}
	}
	var pairs []pair
	if v.displayName != nil {
		pairs = append(pairs, pair{Key: "displayName", Value: *(v.displayName)})
	}
	if v.externalID != nil {
		pairs = append(pairs, pair{Key: "externalId", Value: *(v.externalID)})
	}
	if v.id != nil {
		pairs = append(pairs, pair{Key: "id", Value: *(v.id)})
	}
	if v.members != nil {
		pairs = append(pairs, pair{Key: "members", Value: v.members})
	}
	if v.meta != nil {
		pairs = append(pairs, pair{Key: "meta", Value: v.meta})
	}
	if v.schemas != nil {
		pairs = append(pairs, pair{Key: "schemas", Value: v.schemas})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})

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

func (v *Group) Get(name string, options ...GetOption) (interface{}, bool) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	var ext string
	for _, option := range options {
		switch option.Ident() {
		case identExtension{}:
			ext = option.Value().(string)
		}
	}
	switch name {
	case "displayName":
		if v.displayName == nil {
			return nil, false
		}
		return *(v.displayName), true
	case "externalId":
		if v.externalID == nil {
			return nil, false
		}
		return *(v.externalID), true
	case "id":
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case "members":
		if v.members == nil {
			return nil, false
		}
		return v.members, true
	case "meta":
		if v.meta == nil {
			return nil, false
		}
		return v.meta, true
	case "schemas":
		if v.schemas == nil {
			return nil, false
		}
		return v.schemas, true
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

func (v *Group) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case "displayName":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "displayName", but got %T`, value)
		}
		v.displayName = &tmp
		return nil
	case "externalId":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "externalId", but got %T`, value)
		}
		v.externalID = &tmp
		return nil
	case "id":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case "members":
		var tmp []*User
		tmp, ok := value.([]*User)
		if !ok {
			return fmt.Errorf(`expected []*User for field "members", but got %T`, value)
		}
		v.members = tmp
		return nil
	case "meta":
		var tmp *Meta
		tmp, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected *Meta for field "meta", but got %T`, value)
		}
		v.meta = tmp
		return nil
	case "schemas":
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "schemas", but got %T`, value)
		}
		v.schemas = tmp
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

func (v *Group) UnmarshalJSON(data []byte) error {
	v.displayName = nil
	v.externalID = nil
	v.id = nil
	v.members = nil
	v.meta = nil
	v.schemas = nil
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
			case "displayName":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "displayName": %w`, err)
				}
				v.displayName = &x
			case "externalId":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "externalId": %w`, err)
				}
				v.externalID = &x
			case "id":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case "members":
				var x []*User
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "members": %w`, err)
				}
				v.members = x
			case "meta":
				var x *Meta
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "meta": %w`, err)
				}
				v.meta = x
			case "schemas":
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = x
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

type GroupBuilder struct {
	mu        sync.Mutex
	err       error
	validator GroupValidator
	object    *Group
}

func (b *Builder) Group() *GroupBuilder {
	return &GroupBuilder{}
}

func (b *GroupBuilder) DisplayName(v string) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Group{}
	}
	if err := b.object.Set("displayName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) ExternalID(v string) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Group{}
	}
	if err := b.object.Set("externalId", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) ID(v string) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Group{}
	}
	if err := b.object.Set("id", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Members(v ...*User) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Group{}
	}
	if err := b.object.Set("members", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Meta(v *Meta) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Group{}
	}
	if err := b.object.Set("meta", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Schemas(v ...string) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Group{}
	}
	if err := b.object.Set("schemas", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Extension(uri string, value interface{}) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Group{}
	}
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Validator(v GroupValidator) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *GroupBuilder) Build() (*Group, error) {
	object := b.object
	validator := b.validator
	b.object = nil
	b.validator = nil
	if object == nil {
		return nil, fmt.Errorf("resource.GroupBuilder: object was not initialized")
	}
	if err := b.err; err != nil {
		return nil, err
	}
	if validator == nil {
		validator = DefaultGroupValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *GroupBuilder) MustBuild() *Group {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
