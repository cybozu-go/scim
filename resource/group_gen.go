package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const GroupSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:Group"

func init() {
	RegisterExtension(GroupSchemaURI, Group{})
}

type Group struct {
	mu          sync.RWMutex
	displayName *string
	externalId  *string
	id          *string
	members     []*GroupMember
	schemas     *schemas
	meta        *Meta
	extra       map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	GroupDisplayNameKey = "displayName"
	GroupExternalIDKey  = "externalId"
	GroupIDKey          = "id"
	GroupMembersKey     = "members"
	GroupSchemasKey     = "schemas"
	GroupMetaKey        = "meta"
)

// Get retrieves the value associated with a key
func (v *Group) Get(key string, dst interface{}) error {
	switch key {
	case GroupDisplayNameKey:
		if val := v.displayName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GroupExternalIDKey:
		if val := v.externalId; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GroupIDKey:
		if val := v.id; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GroupMembersKey:
		if val := v.members; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case GroupSchemasKey:
		if val := v.schemas; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GroupMetaKey:
		if val := v.meta; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	default:
		if v.extra != nil {
			val, ok := v.extra[key]
			if ok {
				return blackmagic.AssignIfCompatible(dst, val)
			}
		}
	}
	return fmt.Errorf(`no such key %q`, key)
}

// Set sets the value of the specified field. The name must be a JSON
// field name, not the Go name
func (v *Group) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case GroupDisplayNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field displayName, got %T`, value)
		}
		v.displayName = &converted
	case GroupExternalIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field externalId, got %T`, value)
		}
		v.externalId = &converted
	case GroupIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field id, got %T`, value)
		}
		v.id = &converted
	case GroupMembersKey:
		converted, ok := value.([]*GroupMember)
		if !ok {
			return fmt.Errorf(`expected value of type []*GroupMember for field members, got %T`, value)
		}
		v.members = converted
	case GroupSchemasKey:
		converted, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected value of type schemas for field schemas, got %T`, value)
		}
		v.schemas = &converted
	case GroupMetaKey:
		converted, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected value of type *Meta for field meta, got %T`, value)
		}
		v.meta = converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *Group) HasDisplayName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.displayName != nil
}

func (v *Group) HasExternalID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.externalId != nil
}

func (v *Group) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *Group) HasMembers() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.members != nil
}

func (v *Group) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *Group) HasMeta() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta != nil
}

func (v *Group) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.displayName; val != nil {
		return *val
	}
	return ""
}

func (v *Group) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.externalId; val != nil {
		return *val
	}
	return ""
}

func (v *Group) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.id; val != nil {
		return *val
	}
	return ""
}

func (v *Group) Members() []*GroupMember {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.members; val != nil {
		return val
	}
	return nil
}

func (v *Group) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

func (v *Group) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.meta; val != nil {
		return val
	}
	return nil
}

// Remove removes the value associated with a key
func (v *Group) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case GroupDisplayNameKey:
		v.displayName = nil
	case GroupExternalIDKey:
		v.externalId = nil
	case GroupIDKey:
		v.id = nil
	case GroupMembersKey:
		v.members = nil
	case GroupSchemasKey:
		v.schemas = nil
	case GroupMetaKey:
		v.meta = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *Group) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 6)
	if val := v.displayName; val != nil {
		pairs = append(pairs, &fieldPair{Name: GroupDisplayNameKey, Value: *val})
	}
	if val := v.externalId; val != nil {
		pairs = append(pairs, &fieldPair{Name: GroupExternalIDKey, Value: *val})
	}
	if val := v.id; val != nil {
		pairs = append(pairs, &fieldPair{Name: GroupIDKey, Value: *val})
	}
	if val := v.members; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: GroupMembersKey, Value: val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: GroupSchemasKey, Value: *val})
	}
	if val := v.meta; val != nil {
		pairs = append(pairs, &fieldPair{Name: GroupMetaKey, Value: val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes Group into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *Group) MarshalJSON() ([]byte, error) {
	pairs := v.makePairs()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, pair := range pairs {
		if i > 0 {
			buf.WriteByte(',')
		}
		enc.Encode(pair.Name)
		buf.WriteByte(':')
		enc.Encode(pair.Value)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON deserializes a piece of JSON data into Group.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *Group) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.displayName = nil
	v.externalId = nil
	v.id = nil
	v.members = nil
	v.schemas = nil
	v.meta = nil

	dec := json.NewDecoder(bytes.NewReader(data))

LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading JSON token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			if tok == '}' { // end of object
				break LOOP
			}
			// we should only get into this clause at the very beginning, and just once
			if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string:
			switch tok {
			case GroupDisplayNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupDisplayNameKey, err)
				}
				v.displayName = &val
			case GroupExternalIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupExternalIDKey, err)
				}
				v.externalId = &val
			case GroupIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupIDKey, err)
				}
				v.id = &val
			case GroupMembersKey:
				var val []*GroupMember
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupMembersKey, err)
				}
				v.members = val
			case GroupSchemasKey:
				var val schemas
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupSchemasKey, err)
				}
				v.schemas = &val
			case GroupMetaKey:
				var val *Meta
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GroupMetaKey, err)
				}
				v.meta = val
			default:
				var val interface{}
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, tok, err)
				}
				if v.extra == nil {
					v.extra = make(map[string]interface{})
				}
				v.extra[tok] = val
			}
		}
	}
	return nil
}

type GroupBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *Group
}

// NewGroupBuilder creates a new GroupBuilder instance.
// GroupBuilder is safe to be used uninitialized as well.
func NewGroupBuilder() *GroupBuilder {
	return &GroupBuilder{}
}

func (b *GroupBuilder) initialize() {
	b.err = nil
	b.object = &Group{}
}
func (b *GroupBuilder) DisplayName(in string) *GroupBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(GroupDisplayNameKey, in)
	return b
}
func (b *GroupBuilder) ExternalID(in string) *GroupBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(GroupExternalIDKey, in)
	return b
}
func (b *GroupBuilder) ID(in string) *GroupBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(GroupIDKey, in)
	return b
}
func (b *GroupBuilder) Members(in ...*GroupMember) *GroupBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(GroupMembersKey, in)
	return b
}
func (b *GroupBuilder) Schemas(in ...string) *GroupBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(GroupSchemasKey, in)
	return b
}
func (b *GroupBuilder) Meta(in *Meta) *GroupBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(GroupMetaKey, in)
	return b
}

func (b *GroupBuilder) Build() (*Group, error) {
	err := b.err
	if err != nil {
		return nil, err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *GroupBuilder) MustBuild() *Group {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *GroupBuilder) Extension(uri string, value interface{}) *GroupBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}
	if b.object.schemas == nil {
		b.object.schemas = &schemas{}
		b.object.schemas.Add(GroupSchemaURI)
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (v *Group) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *Group) GetExtension(name, uri string, dst interface{}) error {
	if uri == "" {
		return v.Get(name, dst)
	}
	var ext interface{}
	if err := v.Get(uri, ext); err != nil {
		return fmt.Errorf(`failed to fetch extension %q: %w`, uri, err)
	}

	getter, ok := ext.(interface {
		Get(string, interface{}) error
	})
	if !ok {
		return fmt.Errorf(`extension does not implement Get(string, interface{}) error`)
	}
	return getter.Get(name, dst)
}

func (b *Builder) Group() *GroupBuilder {
	return &GroupBuilder{}
}
