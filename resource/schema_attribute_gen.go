package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	schemaAttributeCanonicalValuesJSONKey = "canonicalValues"
	schemaAttributeCaseExactJSONKey       = "caseExact"
	schemaAttributeDescriptionJSONKey     = "description"
	schemaAttributeFallbackTypeJSONKey    = "fallbackType"
	schemaAttributeMultiValuedJSONKey     = "multiValued"
	schemaAttributeMutabilityJSONKey      = "mutability"
	schemaAttributeNameJSONKey            = "name"
	schemaAttributeReferenceTypesJSONKey  = "referenceTypes"
	schemaAttributeRequiredJSONKey        = "required"
	schemaAttributeReturnedJSONKey        = "returned"
	schemaAttributeSubAttributesJSONKey   = "subAttributes"
	schemaAttributeTypeJSONKey            = "type"
	schemaAttributeUniquenessJSONKey      = "uniqueness"
)

type SchemaAttribute struct {
	canonicalValues []interface{}
	caseExact       *bool
	description     *string
	fallbackType    *SchemaAttribute
	multiValued     *bool
	mutability      *Mutability
	name            *string
	referenceTypes  []string
	required        *bool
	returned        *Returned
	subAttributes   []*SchemaAttribute
	typ             *string
	uniqueness      *Uniqueness
	privateParams   map[string]interface{}
	mu              sync.RWMutex
}

type SchemaAttributeValidator interface {
	Validate(*SchemaAttribute) error
}

type SchemaAttributeValidateFunc func(v *SchemaAttribute) error

func (f SchemaAttributeValidateFunc) Validate(v *SchemaAttribute) error {
	return f(v)
}

var DefaultSchemaAttributeValidator SchemaAttributeValidator = SchemaAttributeValidateFunc(func(v *SchemaAttribute) error {
	return nil
})

func (v *SchemaAttribute) CanonicalValues() []interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.canonicalValues
}

func (v *SchemaAttribute) CaseExact() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.caseExact == nil {
		return false
	}
	return *(v.caseExact)
}

func (v *SchemaAttribute) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.description == nil {
		return ""
	}
	return *(v.description)
}

func (v *SchemaAttribute) FallbackType() *SchemaAttribute {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.fallbackType
}

func (v *SchemaAttribute) MultiValued() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.multiValued == nil {
		return false
	}
	return *(v.multiValued)
}

func (v *SchemaAttribute) Mutability() Mutability {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.mutability == nil {
		return MutReadOnly
	}
	return *(v.mutability)
}

func (v *SchemaAttribute) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.name == nil {
		return ""
	}
	return *(v.name)
}

func (v *SchemaAttribute) ReferenceTypes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.referenceTypes
}

func (v *SchemaAttribute) Required() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.required == nil {
		return false
	}
	return *(v.required)
}

func (v *SchemaAttribute) Returned() Returned {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.returned == nil {
		return ReturnedAlways
	}
	return *(v.returned)
}

func (v *SchemaAttribute) SubAttributes() []*SchemaAttribute {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.subAttributes
}

func (v *SchemaAttribute) Type() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.typ == nil {
		return ""
	}
	return *(v.typ)
}

func (v *SchemaAttribute) Uniqueness() Uniqueness {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.uniqueness == nil {
		return UniqNone
	}
	return *(v.uniqueness)
}

func (v *SchemaAttribute) makePairs() []pair {
	pairs := make([]pair, 0, 13)
	if v.canonicalValues != nil {
		pairs = append(pairs, pair{Key: "canonicalValues", Value: v.canonicalValues})
	}
	if v.caseExact != nil {
		pairs = append(pairs, pair{Key: "caseExact", Value: *(v.caseExact)})
	}
	if v.description != nil {
		pairs = append(pairs, pair{Key: "description", Value: *(v.description)})
	}
	if v.fallbackType != nil {
		pairs = append(pairs, pair{Key: "fallbackType", Value: v.fallbackType})
	}
	if v.multiValued != nil {
		pairs = append(pairs, pair{Key: "multiValued", Value: *(v.multiValued)})
	}
	if v.mutability != nil {
		pairs = append(pairs, pair{Key: "mutability", Value: *(v.mutability)})
	}
	if v.name != nil {
		pairs = append(pairs, pair{Key: "name", Value: *(v.name)})
	}
	if v.referenceTypes != nil {
		pairs = append(pairs, pair{Key: "referenceTypes", Value: v.referenceTypes})
	}
	if v.required != nil {
		pairs = append(pairs, pair{Key: "required", Value: *(v.required)})
	}
	if v.returned != nil {
		pairs = append(pairs, pair{Key: "returned", Value: *(v.returned)})
	}
	if v.subAttributes != nil {
		pairs = append(pairs, pair{Key: "subAttributes", Value: v.subAttributes})
	}
	if v.typ != nil {
		pairs = append(pairs, pair{Key: "type", Value: *(v.typ)})
	}
	if v.uniqueness != nil {
		pairs = append(pairs, pair{Key: "uniqueness", Value: *(v.uniqueness)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *SchemaAttribute) MarshalJSON() ([]byte, error) {
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

func (v *SchemaAttribute) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case schemaAttributeCanonicalValuesJSONKey:
		if v.canonicalValues == nil {
			return nil, false
		}
		return v.canonicalValues, true
	case schemaAttributeCaseExactJSONKey:
		if v.caseExact == nil {
			return nil, false
		}
		return *(v.caseExact), true
	case schemaAttributeDescriptionJSONKey:
		if v.description == nil {
			return nil, false
		}
		return *(v.description), true
	case schemaAttributeFallbackTypeJSONKey:
		if v.fallbackType == nil {
			return nil, false
		}
		return v.fallbackType, true
	case schemaAttributeMultiValuedJSONKey:
		if v.multiValued == nil {
			return nil, false
		}
		return *(v.multiValued), true
	case schemaAttributeMutabilityJSONKey:
		if v.mutability == nil {
			return nil, false
		}
		return *(v.mutability), true
	case schemaAttributeNameJSONKey:
		if v.name == nil {
			return nil, false
		}
		return *(v.name), true
	case schemaAttributeReferenceTypesJSONKey:
		if v.referenceTypes == nil {
			return nil, false
		}
		return v.referenceTypes, true
	case schemaAttributeRequiredJSONKey:
		if v.required == nil {
			return nil, false
		}
		return *(v.required), true
	case schemaAttributeReturnedJSONKey:
		if v.returned == nil {
			return nil, false
		}
		return *(v.returned), true
	case schemaAttributeSubAttributesJSONKey:
		if v.subAttributes == nil {
			return nil, false
		}
		return v.subAttributes, true
	case schemaAttributeTypeJSONKey:
		if v.typ == nil {
			return nil, false
		}
		return *(v.typ), true
	case schemaAttributeUniquenessJSONKey:
		if v.uniqueness == nil {
			return nil, false
		}
		return *(v.uniqueness), true
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

func (v *SchemaAttribute) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case schemaAttributeCanonicalValuesJSONKey:
		var tmp []interface{}
		tmp, ok := value.([]interface{})
		if !ok {
			return fmt.Errorf(`expected []interface{} for field "canonicalValues", but got %T`, value)
		}
		v.canonicalValues = tmp
		return nil
	case schemaAttributeCaseExactJSONKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "caseExact", but got %T`, value)
		}
		v.caseExact = &tmp
		return nil
	case schemaAttributeDescriptionJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "description", but got %T`, value)
		}
		v.description = &tmp
		return nil
	case schemaAttributeFallbackTypeJSONKey:
		var tmp *SchemaAttribute
		tmp, ok := value.(*SchemaAttribute)
		if !ok {
			return fmt.Errorf(`expected *SchemaAttribute for field "fallbackType", but got %T`, value)
		}
		v.fallbackType = tmp
		return nil
	case schemaAttributeMultiValuedJSONKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "multiValued", but got %T`, value)
		}
		v.multiValued = &tmp
		return nil
	case schemaAttributeMutabilityJSONKey:
		var tmp Mutability
		tmp, ok := value.(Mutability)
		if !ok {
			return fmt.Errorf(`expected Mutability for field "mutability", but got %T`, value)
		}
		v.mutability = &tmp
		return nil
	case schemaAttributeNameJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "name", but got %T`, value)
		}
		v.name = &tmp
		return nil
	case schemaAttributeReferenceTypesJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "referenceTypes", but got %T`, value)
		}
		v.referenceTypes = tmp
		return nil
	case schemaAttributeRequiredJSONKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "required", but got %T`, value)
		}
		v.required = &tmp
		return nil
	case schemaAttributeReturnedJSONKey:
		var tmp Returned
		tmp, ok := value.(Returned)
		if !ok {
			return fmt.Errorf(`expected Returned for field "returned", but got %T`, value)
		}
		v.returned = &tmp
		return nil
	case schemaAttributeSubAttributesJSONKey:
		var tmp []*SchemaAttribute
		tmp, ok := value.([]*SchemaAttribute)
		if !ok {
			return fmt.Errorf(`expected []*SchemaAttribute for field "subAttributes", but got %T`, value)
		}
		v.subAttributes = tmp
		return nil
	case schemaAttributeTypeJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "type", but got %T`, value)
		}
		v.typ = &tmp
		return nil
	case schemaAttributeUniquenessJSONKey:
		var tmp Uniqueness
		tmp, ok := value.(Uniqueness)
		if !ok {
			return fmt.Errorf(`expected Uniqueness for field "uniqueness", but got %T`, value)
		}
		v.uniqueness = &tmp
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

func (v *SchemaAttribute) UnmarshalJSON(data []byte) error {
	v.canonicalValues = nil
	v.caseExact = nil
	v.description = nil
	v.fallbackType = nil
	v.multiValued = nil
	v.mutability = nil
	v.name = nil
	v.referenceTypes = nil
	v.required = nil
	v.returned = nil
	v.subAttributes = nil
	v.typ = nil
	v.uniqueness = nil
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
			case schemaAttributeCanonicalValuesJSONKey:
				var x []interface{}
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "canonicalValues": %w`, err)
				}
				v.canonicalValues = x
			case schemaAttributeCaseExactJSONKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "caseExact": %w`, err)
				}
				v.caseExact = &x
			case schemaAttributeDescriptionJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "description": %w`, err)
				}
				v.description = &x
			case schemaAttributeFallbackTypeJSONKey:
				var x *SchemaAttribute
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "fallbackType": %w`, err)
				}
				v.fallbackType = x
			case schemaAttributeMultiValuedJSONKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "multiValued": %w`, err)
				}
				v.multiValued = &x
			case schemaAttributeMutabilityJSONKey:
				var x Mutability
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "mutability": %w`, err)
				}
				v.mutability = &x
			case schemaAttributeNameJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "name": %w`, err)
				}
				v.name = &x
			case schemaAttributeReferenceTypesJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "referenceTypes": %w`, err)
				}
				v.referenceTypes = x
			case schemaAttributeRequiredJSONKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "required": %w`, err)
				}
				v.required = &x
			case schemaAttributeReturnedJSONKey:
				var x Returned
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "returned": %w`, err)
				}
				v.returned = &x
			case schemaAttributeSubAttributesJSONKey:
				var x []*SchemaAttribute
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "subAttributes": %w`, err)
				}
				v.subAttributes = x
			case schemaAttributeTypeJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "type": %w`, err)
				}
				v.typ = &x
			case schemaAttributeUniquenessJSONKey:
				var x Uniqueness
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "uniqueness": %w`, err)
				}
				v.uniqueness = &x
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

func (v *SchemaAttribute) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type SchemaAttributeBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator SchemaAttributeValidator
	object    *SchemaAttribute
}

func (b *Builder) SchemaAttribute() *SchemaAttributeBuilder {
	return NewSchemaAttributeBuilder()
}

func NewSchemaAttributeBuilder() *SchemaAttributeBuilder {
	var b SchemaAttributeBuilder
	b.init()
	return &b
}

func (b *SchemaAttributeBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &SchemaAttribute{}
}

func (b *SchemaAttributeBuilder) CanonicalValues(v ...interface{}) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("canonicalValues", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) CaseExact(v bool) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("caseExact", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Description(v string) *SchemaAttributeBuilder {
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

func (b *SchemaAttributeBuilder) FallbackType(v *SchemaAttribute) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("fallbackType", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) MultiValued(v bool) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("multiValued", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Mutability(v Mutability) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("mutability", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Name(v string) *SchemaAttributeBuilder {
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

func (b *SchemaAttributeBuilder) ReferenceTypes(v ...string) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("referenceTypes", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Required(v bool) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("required", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Returned(v Returned) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("returned", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) SubAttributes(v ...*SchemaAttribute) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("subAttributes", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Type(v string) *SchemaAttributeBuilder {
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

func (b *SchemaAttributeBuilder) Uniqueness(v Uniqueness) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("uniqueness", v); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Validator(v SchemaAttributeValidator) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *SchemaAttributeBuilder) Build() (*SchemaAttribute, error) {
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
		return nil, fmt.Errorf("resource.SchemaAttributeBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultSchemaAttributeValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *SchemaAttributeBuilder) MustBuild() *SchemaAttribute {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
