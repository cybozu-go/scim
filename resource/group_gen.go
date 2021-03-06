package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	GroupDisplayNameKey = "displayName"
	GroupExternalIDKey  = "externalId"
	GroupIDKey          = "id"
	GroupMembersKey     = "members"
	GroupMetaKey        = "meta"
	GroupSchemasKey     = "schemas"
)

const GroupSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:Group"

func init() {
	RegisterExtension(GroupSchemaURI, Group{})
}

type Group struct {
	displayName   *string
	externalID    *string
	id            *string
	members       []*GroupMember
	meta          *Meta
	schemas       schemas
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

var DefaultGroupValidator GroupValidator = GroupValidateFunc(func(v *Group) error {
	return nil
})

func (v *Group) HasDisplayName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.displayName != nil
}

func (v *Group) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.displayName == nil {
		return ""
	}
	return *(v.displayName)
}

func (v *Group) HasExternalID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.externalID != nil
}

func (v *Group) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.externalID == nil {
		return ""
	}
	return *(v.externalID)
}

func (v *Group) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *Group) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *Group) HasMembers() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.members != nil
}

func (v *Group) Members() []*GroupMember {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.members
}

func (v *Group) HasMeta() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta != nil
}

func (v *Group) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta
}

func (v *Group) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *Group) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *Group) makePairs() []pair {
	pairs := make([]pair, 0, 6)
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
	return pairs
}

func (v *Group) MarshalJSON() ([]byte, error) {
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

func (v *Group) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case GroupDisplayNameKey:
		if v.displayName == nil {
			return nil, false
		}
		return *(v.displayName), true
	case GroupExternalIDKey:
		if v.externalID == nil {
			return nil, false
		}
		return *(v.externalID), true
	case GroupIDKey:
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case GroupMembersKey:
		if v.members == nil {
			return nil, false
		}
		return v.members, true
	case GroupMetaKey:
		if v.meta == nil {
			return nil, false
		}
		return v.meta, true
	case GroupSchemasKey:
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
	case GroupDisplayNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "displayName", but got %T`, value)
		}
		v.displayName = &tmp
		return nil
	case GroupExternalIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "externalId", but got %T`, value)
		}
		v.externalID = &tmp
		return nil
	case GroupIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case GroupMembersKey:
		var tmp []*GroupMember
		tmp, ok := value.([]*GroupMember)
		if !ok {
			return fmt.Errorf(`expected []*GroupMember for field "members", but got %T`, value)
		}
		v.members = tmp
		return nil
	case GroupMetaKey:
		var tmp *Meta
		tmp, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected *Meta for field "meta", but got %T`, value)
		}
		v.meta = tmp
		return nil
	case GroupSchemasKey:
		var tmp schemas
		tmp, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected schemas for field "schemas", but got %T`, value)
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

func (v *Group) Clone() *Group {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &Group{
		displayName: v.displayName,
		externalID:  v.externalID,
		id:          v.id,
		members:     v.members,
		meta:        v.meta,
		schemas:     v.schemas,
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
			case GroupDisplayNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "displayName": %w`, err)
				}
				v.displayName = &x
			case GroupExternalIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "externalId": %w`, err)
				}
				v.externalID = &x
			case GroupIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case GroupMembersKey:
				var x []*GroupMember
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "members": %w`, err)
				}
				v.members = x
			case GroupMetaKey:
				var x *Meta
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "meta": %w`, err)
				}
				v.meta = x
			case GroupSchemasKey:
				var x schemas
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

func (v *Group) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type GroupBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator GroupValidator
	object    *Group
}

func (b *Builder) Group() *GroupBuilder {
	return NewGroupBuilder()
}

func NewGroupBuilder() *GroupBuilder {
	var b GroupBuilder
	b.init()
	return &b
}

func (b *GroupBuilder) From(in *Group) *GroupBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *GroupBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &Group{}

	b.object.schemas = make(schemas)
	b.object.schemas.Add(GroupSchemaURI)
}

func (b *GroupBuilder) DisplayName(v string) *GroupBuilder {
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

func (b *GroupBuilder) ExternalID(v string) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("externalId", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) ID(v string) *GroupBuilder {
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

func (b *GroupBuilder) Members(v ...*GroupMember) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("members", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Meta(v *Meta) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("meta", v); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Schemas(v ...string) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	for _, schema := range v {
		b.object.schemas.Add(schema)
	}
	return b
}

func (b *GroupBuilder) Extension(uri string, value interface{}) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (b *GroupBuilder) Validator(v GroupValidator) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *GroupBuilder) Build() (*Group, error) {
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
		return nil, fmt.Errorf("resource.GroupBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultGroupValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
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
