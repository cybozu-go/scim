package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"time"
)

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

var DefaultMetaValidator MetaValidator

func (v *Meta) Created() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.created == nil {
		return time.Time{}
	}
	return *(v.created)
}

func (v *Meta) LastModified() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.lastModified == nil {
		return time.Time{}
	}
	return *(v.lastModified)
}

func (v *Meta) Location() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.location == nil {
		return ""
	}
	return *(v.location)
}

func (v *Meta) ResourceType() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.resourceType == nil {
		return ""
	}
	return *(v.resourceType)
}

func (v *Meta) Version() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.version == nil {
		return ""
	}
	return *(v.version)
}

func (v *Meta) MarshalJSON() ([]byte, error) {
	type pair struct {
		Key   string
		Value interface{}
	}
	var pairs []pair
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
	for _, option := range options {
		switch option.Ident() {
		case identExtension{}:
			ext = option.Value().(string)
		}
	}
	switch name {
	case "created":
		if v.created == nil {
			return nil, false
		}
		return *(v.created), true
	case "lastModified":
		if v.lastModified == nil {
			return nil, false
		}
		return *(v.lastModified), true
	case "location":
		if v.location == nil {
			return nil, false
		}
		return *(v.location), true
	case "resourceType":
		if v.resourceType == nil {
			return nil, false
		}
		return *(v.resourceType), true
	case "version":
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
	case "created":
		var tmp time.Time
		tmp, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf(`expected time.Time for field "created", but got %T`, value)
		}
		v.created = &tmp
		return nil
	case "lastModified":
		var tmp time.Time
		tmp, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf(`expected time.Time for field "lastModified", but got %T`, value)
		}
		v.lastModified = &tmp
		return nil
	case "location":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "location", but got %T`, value)
		}
		v.location = &tmp
		return nil
	case "resourceType":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "resourceType", but got %T`, value)
		}
		v.resourceType = &tmp
		return nil
	case "version":
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
			case "created":
				var x time.Time
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "created": %w`, err)
				}
				v.created = &x
			case "lastModified":
				var x time.Time
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "lastModified": %w`, err)
				}
				v.lastModified = &x
			case "location":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "location": %w`, err)
				}
				v.location = &x
			case "resourceType":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "resourceType": %w`, err)
				}
				v.resourceType = &x
			case "version":
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

type MetaBuilder struct {
	mu        sync.Mutex
	err       error
	validator MetaValidator
	object    *Meta
}

func (b *Builder) Meta() *MetaBuilder {
	return &MetaBuilder{}
}

func (b *MetaBuilder) Created(v time.Time) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Meta{}
	}
	if err := b.object.Set("created", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) LastModified(v time.Time) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Meta{}
	}
	if err := b.object.Set("lastModified", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) Location(v string) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Meta{}
	}
	if err := b.object.Set("location", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) ResourceType(v string) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Meta{}
	}
	if err := b.object.Set("resourceType", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) Version(v string) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Meta{}
	}
	if err := b.object.Set("version", v); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) Extension(uri string, value interface{}) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &Meta{}
	}
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (b *MetaBuilder) Validator(v MetaValidator) *MetaBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *MetaBuilder) Build() (*Meta, error) {
	object := b.object
	validator := b.validator
	b.object = nil
	b.validator = nil
	if object == nil {
		return nil, fmt.Errorf("resource.MetaBuilder: object was not initialized")
	}
	if err := b.err; err != nil {
		return nil, err
	}
	if validator == nil {
		validator = DefaultMetaValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
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
