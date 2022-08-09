package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	PatchOperationExternalIDKey = "externalId"
	PatchOperationIDKey         = "id"
	PatchOperationMetaKey       = "meta"
	PatchOperationOpKey         = "op"
	PatchOperationPathKey       = "path"
	PatchOperationValueKey      = "value"
)

type PatchOperation struct {
	externalID    *string
	id            *string
	meta          *Meta
	op            *PatchOperationType
	path          *string
	value         *json.RawMessage
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type PatchOperationValidator interface {
	Validate(*PatchOperation) error
}

type PatchOperationValidateFunc func(v *PatchOperation) error

func (f PatchOperationValidateFunc) Validate(v *PatchOperation) error {
	return f(v)
}

var DefaultPatchOperationValidator PatchOperationValidator = PatchOperationValidateFunc(func(v *PatchOperation) error {
	return nil
})

func (v *PatchOperation) HasExternalID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.externalID != nil
}

func (v *PatchOperation) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.externalID == nil {
		return ""
	}
	return *(v.externalID)
}

func (v *PatchOperation) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *PatchOperation) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *PatchOperation) HasMeta() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta != nil
}

func (v *PatchOperation) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta
}

func (v *PatchOperation) HasOp() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.op != nil
}

func (v *PatchOperation) Op() PatchOperationType {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.op == nil {
		return PatchInvalid
	}
	return *(v.op)
}

func (v *PatchOperation) HasPath() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.path != nil
}

func (v *PatchOperation) Path() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.path == nil {
		return ""
	}
	return *(v.path)
}

func (v *PatchOperation) HasValue() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.value != nil
}

func (v *PatchOperation) Value() json.RawMessage {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.value == nil {
		return nil
	}
	return *(v.value)
}

func (v *PatchOperation) makePairs() []pair {
	pairs := make([]pair, 0, 6)
	if v.externalID != nil {
		pairs = append(pairs, pair{Key: "externalId", Value: *(v.externalID)})
	}
	if v.id != nil {
		pairs = append(pairs, pair{Key: "id", Value: *(v.id)})
	}
	if v.meta != nil {
		pairs = append(pairs, pair{Key: "meta", Value: v.meta})
	}
	if v.op != nil {
		pairs = append(pairs, pair{Key: "op", Value: *(v.op)})
	}
	if v.path != nil {
		pairs = append(pairs, pair{Key: "path", Value: *(v.path)})
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

func (v *PatchOperation) MarshalJSON() ([]byte, error) {
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

func (v *PatchOperation) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case PatchOperationExternalIDKey:
		if v.externalID == nil {
			return nil, false
		}
		return *(v.externalID), true
	case PatchOperationIDKey:
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case PatchOperationMetaKey:
		if v.meta == nil {
			return nil, false
		}
		return v.meta, true
	case PatchOperationOpKey:
		if v.op == nil {
			return nil, false
		}
		return *(v.op), true
	case PatchOperationPathKey:
		if v.path == nil {
			return nil, false
		}
		return *(v.path), true
	case PatchOperationValueKey:
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

func (v *PatchOperation) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case PatchOperationExternalIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "externalId", but got %T`, value)
		}
		v.externalID = &tmp
		return nil
	case PatchOperationIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case PatchOperationMetaKey:
		var tmp *Meta
		tmp, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected *Meta for field "meta", but got %T`, value)
		}
		v.meta = tmp
		return nil
	case PatchOperationOpKey:
		var tmp PatchOperationType
		tmp, ok := value.(PatchOperationType)
		if !ok {
			return fmt.Errorf(`expected PatchOperationType for field "op", but got %T`, value)
		}
		v.op = &tmp
		return nil
	case PatchOperationPathKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "path", but got %T`, value)
		}
		v.path = &tmp
		return nil
	case PatchOperationValueKey:
		var tmp json.RawMessage
		tmp, ok := value.(json.RawMessage)
		if !ok {
			return fmt.Errorf(`expected json.RawMessage for field "value", but got %T`, value)
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

func (v *PatchOperation) Clone() *PatchOperation {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &PatchOperation{
		externalID: v.externalID,
		id:         v.id,
		meta:       v.meta,
		op:         v.op,
		path:       v.path,
		value:      v.value,
	}
}

func (v *PatchOperation) UnmarshalJSON(data []byte) error {
	v.externalID = nil
	v.id = nil
	v.meta = nil
	v.op = nil
	v.path = nil
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
			case PatchOperationExternalIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "externalId": %w`, err)
				}
				v.externalID = &x
			case PatchOperationIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case PatchOperationMetaKey:
				var x *Meta
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "meta": %w`, err)
				}
				v.meta = x
			case PatchOperationOpKey:
				var x PatchOperationType
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "op": %w`, err)
				}
				v.op = &x
			case PatchOperationPathKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "path": %w`, err)
				}
				v.path = &x
			case PatchOperationValueKey:
				var x json.RawMessage
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

func (v *PatchOperation) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// PatchOperationBuilder creates a PatchOperation resource
type PatchOperationBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator PatchOperationValidator
	object    *PatchOperation
}

func (b *Builder) PatchOperation() *PatchOperationBuilder {
	return NewPatchOperationBuilder()
}

func NewPatchOperationBuilder() *PatchOperationBuilder {
	var b PatchOperationBuilder
	b.init()
	return &b
}

func (b *PatchOperationBuilder) From(in *PatchOperation) *PatchOperationBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *PatchOperationBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &PatchOperation{}
}

func (b *PatchOperationBuilder) ExternalID(v string) *PatchOperationBuilder {
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

func (b *PatchOperationBuilder) ID(v string) *PatchOperationBuilder {
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

func (b *PatchOperationBuilder) Meta(v *Meta) *PatchOperationBuilder {
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

func (b *PatchOperationBuilder) Op(v PatchOperationType) *PatchOperationBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("op", v); err != nil {
		b.err = err
	}
	return b
}

func (b *PatchOperationBuilder) Path(v string) *PatchOperationBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("path", v); err != nil {
		b.err = err
	}
	return b
}

func (b *PatchOperationBuilder) Validator(v PatchOperationValidator) *PatchOperationBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *PatchOperationBuilder) Build() (*PatchOperation, error) {
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
		return nil, fmt.Errorf("resource.PatchOperationBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultPatchOperationValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *PatchOperationBuilder) MustBuild() *PatchOperation {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
