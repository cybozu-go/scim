package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/lestrrat-go/blackmagic"
)

// Meta represents the `meta` field included in SCIM responses. See https://datatracker.ietf.org/doc/html/rfc7643#section-3.1 for details
type Meta struct {
	mu           sync.RWMutex
	resourceType *string
	location     *string
	version      *string
	created      *time.Time
	lastModified *time.Time
	extra        map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	MetaResourceTypeKey = "resourceType"
	MetaLocationKey     = "location"
	MetaVersionKey      = "version"
	MetaCreatedKey      = "created"
	MetaLastModifiedKey = "lastModified"
)

// Get retrieves the value associated with a key
func (v *Meta) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case MetaResourceTypeKey:
		if val := v.resourceType; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case MetaLocationKey:
		if val := v.location; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case MetaVersionKey:
		if val := v.version; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case MetaCreatedKey:
		if val := v.created; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case MetaLastModifiedKey:
		if val := v.lastModified; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
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
func (v *Meta) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case MetaResourceTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field resourceType, got %T`, value)
		}
		v.resourceType = &converted
	case MetaLocationKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field location, got %T`, value)
		}
		v.location = &converted
	case MetaVersionKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field version, got %T`, value)
		}
		v.version = &converted
	case MetaCreatedKey:
		converted, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf(`expected value of type time.Time for field created, got %T`, value)
		}
		v.created = &converted
	case MetaLastModifiedKey:
		converted, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf(`expected value of type time.Time for field lastModified, got %T`, value)
		}
		v.lastModified = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *Meta) HasResourceType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.resourceType != nil
}

func (v *Meta) HasLocation() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.location != nil
}

func (v *Meta) HasVersion() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.version != nil
}

func (v *Meta) HasCreated() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.created != nil
}

func (v *Meta) HasLastModified() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.lastModified != nil
}

func (v *Meta) ResourceType() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.resourceType; val != nil {
		return *val
	}
	return ""
}

func (v *Meta) Location() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.location; val != nil {
		return *val
	}
	return ""
}

func (v *Meta) Version() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.version; val != nil {
		return *val
	}
	return ""
}

func (v *Meta) Created() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.created; val != nil {
		return *val
	}
	return time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func (v *Meta) LastModified() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.lastModified; val != nil {
		return *val
	}
	return time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// Remove removes the value associated with a key
func (v *Meta) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case MetaResourceTypeKey:
		v.resourceType = nil
	case MetaLocationKey:
		v.location = nil
	case MetaVersionKey:
		v.version = nil
	case MetaCreatedKey:
		v.created = nil
	case MetaLastModifiedKey:
		v.lastModified = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *Meta) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 5)
	if val := v.resourceType; val != nil {
		pairs = append(pairs, &fieldPair{Name: MetaResourceTypeKey, Value: *val})
	}
	if val := v.location; val != nil {
		pairs = append(pairs, &fieldPair{Name: MetaLocationKey, Value: *val})
	}
	if val := v.version; val != nil {
		pairs = append(pairs, &fieldPair{Name: MetaVersionKey, Value: *val})
	}
	if val := v.created; val != nil {
		pairs = append(pairs, &fieldPair{Name: MetaCreatedKey, Value: *val})
	}
	if val := v.lastModified; val != nil {
		pairs = append(pairs, &fieldPair{Name: MetaLastModifiedKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes Meta into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *Meta) MarshalJSON() ([]byte, error) {
	pairs := v.makePairs()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, pair := range pairs {
		if i > 0 {
			buf.WriteByte(',')
		}
		if err := enc.Encode(pair.Name); err != nil {
			return nil, fmt.Errorf(`failed to encode map key name: %w`, err)
		}
		buf.WriteByte(':')
		if err := enc.Encode(pair.Value); err != nil {
			return nil, fmt.Errorf(`failed to encode map value for %q: %w`, pair.Name, err)
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON deserializes a piece of JSON data into Meta.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *Meta) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.resourceType = nil
	v.location = nil
	v.version = nil
	v.created = nil
	v.lastModified = nil

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
			case MetaResourceTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, MetaResourceTypeKey, err)
				}
				v.resourceType = &val
			case MetaLocationKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, MetaLocationKey, err)
				}
				v.location = &val
			case MetaVersionKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, MetaVersionKey, err)
				}
				v.version = &val
			case MetaCreatedKey:
				var val time.Time
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, MetaCreatedKey, err)
				}
				v.created = &val
			case MetaLastModifiedKey:
				var val time.Time
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, MetaLastModifiedKey, err)
				}
				v.lastModified = &val
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

type MetaBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *Meta
}

// NewMetaBuilder creates a new MetaBuilder instance.
// MetaBuilder is safe to be used uninitialized as well.
func NewMetaBuilder() *MetaBuilder {
	return &MetaBuilder{}
}

func (b *MetaBuilder) initialize() {
	b.err = nil
	b.object = &Meta{}
}
func (b *MetaBuilder) ResourceType(in string) *MetaBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(MetaResourceTypeKey, in)
	return b
}
func (b *MetaBuilder) Location(in string) *MetaBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(MetaLocationKey, in)
	return b
}
func (b *MetaBuilder) Version(in string) *MetaBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(MetaVersionKey, in)
	return b
}
func (b *MetaBuilder) Created(in time.Time) *MetaBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(MetaCreatedKey, in)
	return b
}
func (b *MetaBuilder) LastModified(in time.Time) *MetaBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(MetaLastModifiedKey, in)
	return b
}

func (b *MetaBuilder) Build() (*Meta, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	err := b.err
	if err != nil {
		return nil, err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *MetaBuilder) MustBuild() *Meta {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (v *Meta) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *Meta) GetExtension(name, uri string, dst interface{}) error {
	if uri == "" {
		return v.Get(name, dst)
	}
	var ext interface{}
	if err := v.Get(uri, &ext); err != nil {
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

func (b *Builder) Meta() *MetaBuilder {
	return &MetaBuilder{}
}
