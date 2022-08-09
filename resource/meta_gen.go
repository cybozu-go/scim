package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"time"
)

const (
	MetaCreatedKey      = "created"
	MetaLastModifiedKey = "lastModified"
	MetaLocationKey     = "location"
	MetaResourceTypeKey = "resourceType"
	MetaVersionKey      = "version"
)

// Meta represents the `meta` field included in SCIM responses.
// See https://datatracker.ietf.org/doc/html/rfc7643#section-3.1 for details
type Meta struct {
	created       *time.Time
	lastModified  *time.Time
	location      *string
	resourceType  *string
	version       *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type MetaValidator interface {
	Validate(*Meta) error
}

type MetaValidateFunc func(v *Meta) error

func (f MetaValidateFunc) Validate(v *Meta) error {
	return f(v)
}

var DefaultMetaValidator MetaValidator = MetaValidateFunc(func(v *Meta) error {
	return nil
})

func (v *Meta) HasCreated() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.created != nil
}

func (v *Meta) Created() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.created == nil {
		return time.Time{}
	}
	return *(v.created)
}

func (v *Meta) HasLastModified() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.lastModified != nil
}

func (v *Meta) LastModified() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.lastModified == nil {
		return time.Time{}
	}
	return *(v.lastModified)
}

func (v *Meta) HasLocation() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.location != nil
}

func (v *Meta) Location() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.location == nil {
		return ""
	}
	return *(v.location)
}

func (v *Meta) HasResourceType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.resourceType != nil
}

func (v *Meta) ResourceType() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.resourceType == nil {
		return ""
	}
	return *(v.resourceType)
}

func (v *Meta) HasVersion() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.version != nil
}

func (v *Meta) Version() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.version == nil {
		return ""
	}
	return *(v.version)
}

func (v *Meta) makePairs() []pair {
	pairs := make([]pair, 0, 5)
	if v.created != nil {
		pairs = append(pairs, pair{Key: "created", Value: *(v.created)})
	}
	if v.lastModified != nil {
		pairs = append(pairs, pair{Key: "lastModified", Value: *(v.lastModified)})
	}
	if v.location != nil {
		pairs = append(pairs, pair{Key: "location", Value: *(v.location)})
	}
	if v.resourceType != nil {
		pairs = append(pairs, pair{Key: "resourceType", Value: *(v.resourceType)})
	}
	if v.version != nil {
		pairs = append(pairs, pair{Key: "version", Value: *(v.version)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *Meta) MarshalJSON() ([]byte, error) {
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

func (v *Meta) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case MetaCreatedKey:
		if v.created == nil {
			return nil, false
		}
		return *(v.created), true
	case MetaLastModifiedKey:
		if v.lastModified == nil {
			return nil, false
		}
		return *(v.lastModified), true
	case MetaLocationKey:
		if v.location == nil {
			return nil, false
		}
		return *(v.location), true
	case MetaResourceTypeKey:
		if v.resourceType == nil {
			return nil, false
		}
		return *(v.resourceType), true
	case MetaVersionKey:
		if v.version == nil {
			return nil, false
		}
		return *(v.version), true
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

func (v *Meta) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case MetaCreatedKey:
		var tmp time.Time
		tmp, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf(`expected time.Time for field "created", but got %T`, value)
		}
		v.created = &tmp
		return nil
	case MetaLastModifiedKey:
		var tmp time.Time
		tmp, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf(`expected time.Time for field "lastModified", but got %T`, value)
		}
		v.lastModified = &tmp
		return nil
	case MetaLocationKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "location", but got %T`, value)
		}
		v.location = &tmp
		return nil
	case MetaResourceTypeKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "resourceType", but got %T`, value)
		}
		v.resourceType = &tmp
		return nil
	case MetaVersionKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "version", but got %T`, value)
		}
		v.version = &tmp
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

func (v *Meta) Clone() *Meta {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &Meta{
		created:      v.created,
		lastModified: v.lastModified,
		location:     v.location,
		resourceType: v.resourceType,
		version:      v.version,
	}
}

func (v *Meta) UnmarshalJSON(data []byte) error {
	v.created = nil
	v.lastModified = nil
	v.location = nil
	v.resourceType = nil
	v.version = nil
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
			case MetaCreatedKey:
				var x time.Time
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "created": %w`, err)
				}
				v.created = &x
			case MetaLastModifiedKey:
				var x time.Time
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "lastModified": %w`, err)
				}
				v.lastModified = &x
			case MetaLocationKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "location": %w`, err)
				}
				v.location = &x
			case MetaResourceTypeKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "resourceType": %w`, err)
				}
				v.resourceType = &x
			case MetaVersionKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "version": %w`, err)
				}
				v.version = &x
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

func (v *Meta) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// MetaBuilder creates a Meta resource
type MetaBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator MetaValidator
	object    *Meta
}

func (b *Builder) Meta() *MetaBuilder {
	return NewMetaBuilder()
}

func NewMetaBuilder() *MetaBuilder {
	var b MetaBuilder
	b.init()
	return &b
}

func (b *MetaBuilder) From(in *Meta) *MetaBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *MetaBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &Meta{}
}

func (b *MetaBuilder) Created(v time.Time) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("created", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) LastModified(v time.Time) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("lastModified", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) Location(v string) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("location", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) ResourceType(v string) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("resourceType", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) Version(v string) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("version", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) Validator(v MetaValidator) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *MetaBuilder) Build() (*Meta, error) {
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
		return nil, fmt.Errorf("resource.MetaBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultMetaValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *MetaBuilder) MustBuild() *Meta {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
