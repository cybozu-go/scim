// Generated by "sketch" utility. DO NOT EDIT
package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

func init() {
	Register("SchemaAttribute", "", SchemaAttribute{})
}

type SchemaAttribute struct {
	mu              sync.RWMutex
	canonicalValues *string
	caseExact       *bool
	description     *string
	multiValued     *bool
	mutability      *Mutability
	name            *string
	referenceTypes  []string
	required        *bool
	returned        *Returned
	subAttributes   []*SchemaAttribute
	typ             *DataType
	uniqueness      *Uniqueness
	extra           map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	SchemaAttributeCanonicalValuesKey = "canonicalValues"
	SchemaAttributeCaseExactKey       = "caseExact"
	SchemaAttributeDescriptionKey     = "description"
	SchemaAttributeMultiValuedKey     = "multiValued"
	SchemaAttributeMutabilityKey      = "mutability"
	SchemaAttributeNameKey            = "name"
	SchemaAttributeReferenceTypesKey  = "referenceTypes"
	SchemaAttributeRequiredKey        = "required"
	SchemaAttributeReturnedKey        = "returned"
	SchemaAttributeSubAttributesKey   = "subAttributes"
	SchemaAttributeTypeKey            = "type"
	SchemaAttributeUniquenessKey      = "uniqueness"
)

// Get retrieves the value associated with a key
func (v *SchemaAttribute) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case SchemaAttributeCanonicalValuesKey:
		if val := v.canonicalValues; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeCaseExactKey:
		if val := v.caseExact; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeDescriptionKey:
		if val := v.description; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeMultiValuedKey:
		if val := v.multiValued; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeMutabilityKey:
		if val := v.mutability; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeNameKey:
		if val := v.name; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeReferenceTypesKey:
		if val := v.referenceTypes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case SchemaAttributeRequiredKey:
		if val := v.required; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeReturnedKey:
		if val := v.returned; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeSubAttributesKey:
		if val := v.subAttributes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case SchemaAttributeTypeKey:
		if val := v.typ; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SchemaAttributeUniquenessKey:
		if val := v.uniqueness; val != nil {
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
func (v *SchemaAttribute) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case SchemaAttributeCanonicalValuesKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field canonicalValues, got %T`, value)
		}
		v.canonicalValues = &converted
	case SchemaAttributeCaseExactKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field caseExact, got %T`, value)
		}
		v.caseExact = &converted
	case SchemaAttributeDescriptionKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field description, got %T`, value)
		}
		v.description = &converted
	case SchemaAttributeMultiValuedKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field multiValued, got %T`, value)
		}
		v.multiValued = &converted
	case SchemaAttributeMutabilityKey:
		converted, ok := value.(Mutability)
		if !ok {
			return fmt.Errorf(`expected value of type Mutability for field mutability, got %T`, value)
		}
		v.mutability = &converted
	case SchemaAttributeNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field name, got %T`, value)
		}
		v.name = &converted
	case SchemaAttributeReferenceTypesKey:
		converted, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected value of type []string for field referenceTypes, got %T`, value)
		}
		v.referenceTypes = converted
	case SchemaAttributeRequiredKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field required, got %T`, value)
		}
		v.required = &converted
	case SchemaAttributeReturnedKey:
		converted, ok := value.(Returned)
		if !ok {
			return fmt.Errorf(`expected value of type Returned for field returned, got %T`, value)
		}
		v.returned = &converted
	case SchemaAttributeSubAttributesKey:
		converted, ok := value.([]*SchemaAttribute)
		if !ok {
			return fmt.Errorf(`expected value of type []*SchemaAttribute for field subAttributes, got %T`, value)
		}
		v.subAttributes = converted
	case SchemaAttributeTypeKey:
		converted, ok := value.(DataType)
		if !ok {
			return fmt.Errorf(`expected value of type DataType for field type, got %T`, value)
		}
		v.typ = &converted
	case SchemaAttributeUniquenessKey:
		converted, ok := value.(Uniqueness)
		if !ok {
			return fmt.Errorf(`expected value of type Uniqueness for field uniqueness, got %T`, value)
		}
		v.uniqueness = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *SchemaAttribute) HasCanonicalValues() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.canonicalValues != nil
}

func (v *SchemaAttribute) HasCaseExact() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.caseExact != nil
}

func (v *SchemaAttribute) HasDescription() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.description != nil
}

func (v *SchemaAttribute) HasMultiValued() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.multiValued != nil
}

func (v *SchemaAttribute) HasMutability() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.mutability != nil
}

func (v *SchemaAttribute) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

func (v *SchemaAttribute) HasReferenceTypes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.referenceTypes != nil
}

func (v *SchemaAttribute) HasRequired() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.required != nil
}

func (v *SchemaAttribute) HasReturned() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.returned != nil
}

func (v *SchemaAttribute) HasSubAttributes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.subAttributes != nil
}

func (v *SchemaAttribute) HasType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.typ != nil
}

func (v *SchemaAttribute) HasUniqueness() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.uniqueness != nil
}

func (v *SchemaAttribute) CanonicalValues() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.canonicalValues; val != nil {
		return *val
	}
	return ""
}

func (v *SchemaAttribute) CaseExact() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.caseExact; val != nil {
		return *val
	}
	return false
}

func (v *SchemaAttribute) Description() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.description; val != nil {
		return *val
	}
	return ""
}

func (v *SchemaAttribute) MultiValued() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.multiValued; val != nil {
		return *val
	}
	return false
}

func (v *SchemaAttribute) Mutability() Mutability {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.mutability; val != nil {
		return *val
	}
	return MutReadOnly
}

func (v *SchemaAttribute) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.name; val != nil {
		return *val
	}
	return ""
}

func (v *SchemaAttribute) ReferenceTypes() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.referenceTypes; val != nil {
		return val
	}
	return []string(nil)
}

func (v *SchemaAttribute) Required() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.required; val != nil {
		return *val
	}
	return false
}

func (v *SchemaAttribute) Returned() Returned {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.returned; val != nil {
		return *val
	}
	return ReturnedNever
}

func (v *SchemaAttribute) SubAttributes() []*SchemaAttribute {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.subAttributes; val != nil {
		return val
	}
	return nil
}

func (v *SchemaAttribute) Type() DataType {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.typ; val != nil {
		return *val
	}
	return InvalidDataType
}

func (v *SchemaAttribute) Uniqueness() Uniqueness {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.uniqueness; val != nil {
		return *val
	}
	return UniqNone
}

// Remove removes the value associated with a key
func (v *SchemaAttribute) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case SchemaAttributeCanonicalValuesKey:
		v.canonicalValues = nil
	case SchemaAttributeCaseExactKey:
		v.caseExact = nil
	case SchemaAttributeDescriptionKey:
		v.description = nil
	case SchemaAttributeMultiValuedKey:
		v.multiValued = nil
	case SchemaAttributeMutabilityKey:
		v.mutability = nil
	case SchemaAttributeNameKey:
		v.name = nil
	case SchemaAttributeReferenceTypesKey:
		v.referenceTypes = nil
	case SchemaAttributeRequiredKey:
		v.required = nil
	case SchemaAttributeReturnedKey:
		v.returned = nil
	case SchemaAttributeSubAttributesKey:
		v.subAttributes = nil
	case SchemaAttributeTypeKey:
		v.typ = nil
	case SchemaAttributeUniquenessKey:
		v.uniqueness = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *SchemaAttribute) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 12)
	if val := v.canonicalValues; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeCanonicalValuesKey, Value: *val})
	}
	if val := v.caseExact; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeCaseExactKey, Value: *val})
	}
	if val := v.description; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeDescriptionKey, Value: *val})
	}
	if val := v.multiValued; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeMultiValuedKey, Value: *val})
	}
	if val := v.mutability; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeMutabilityKey, Value: *val})
	}
	if val := v.name; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeNameKey, Value: *val})
	}
	if val := v.referenceTypes; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeReferenceTypesKey, Value: val})
	}
	if val := v.required; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeRequiredKey, Value: *val})
	}
	if val := v.returned; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeReturnedKey, Value: *val})
	}
	if val := v.subAttributes; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeSubAttributesKey, Value: val})
	}
	if val := v.typ; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeTypeKey, Value: *val})
	}
	if val := v.uniqueness; val != nil {
		pairs = append(pairs, &fieldPair{Name: SchemaAttributeUniquenessKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes SchemaAttribute into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *SchemaAttribute) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into SchemaAttribute.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *SchemaAttribute) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.canonicalValues = nil
	v.caseExact = nil
	v.description = nil
	v.multiValued = nil
	v.mutability = nil
	v.name = nil
	v.referenceTypes = nil
	v.required = nil
	v.returned = nil
	v.subAttributes = nil
	v.typ = nil
	v.uniqueness = nil

	dec := json.NewDecoder(bytes.NewReader(data))
	var extra map[string]interface{}

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
			case SchemaAttributeCanonicalValuesKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeCanonicalValuesKey, err)
				}
				v.canonicalValues = &val
			case SchemaAttributeCaseExactKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeCaseExactKey, err)
				}
				v.caseExact = &val
			case SchemaAttributeDescriptionKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeDescriptionKey, err)
				}
				v.description = &val
			case SchemaAttributeMultiValuedKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeMultiValuedKey, err)
				}
				v.multiValued = &val
			case SchemaAttributeMutabilityKey:
				var val Mutability
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeMutabilityKey, err)
				}
				v.mutability = &val
			case SchemaAttributeNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeNameKey, err)
				}
				v.name = &val
			case SchemaAttributeReferenceTypesKey:
				var val []string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeReferenceTypesKey, err)
				}
				v.referenceTypes = val
			case SchemaAttributeRequiredKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeRequiredKey, err)
				}
				v.required = &val
			case SchemaAttributeReturnedKey:
				var val Returned
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeReturnedKey, err)
				}
				v.returned = &val
			case SchemaAttributeSubAttributesKey:
				var val []*SchemaAttribute
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeSubAttributesKey, err)
				}
				v.subAttributes = val
			case SchemaAttributeTypeKey:
				var val DataType
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeTypeKey, err)
				}
				v.typ = &val
			case SchemaAttributeUniquenessKey:
				var val Uniqueness
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SchemaAttributeUniquenessKey, err)
				}
				v.uniqueness = &val
			default:
				var val interface{}
				if err := extraFieldsDecoder(tok, dec, &val); err != nil {
					return err
				}
				if extra == nil {
					extra = make(map[string]interface{})
				}
				extra[tok] = val
			}
		}
	}

	if extra != nil {
		v.extra = extra
	}
	return nil
}

type SchemaAttributeBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *SchemaAttribute
}

// NewSchemaAttributeBuilder creates a new SchemaAttributeBuilder instance.
// SchemaAttributeBuilder is safe to be used uninitialized as well.
func NewSchemaAttributeBuilder() *SchemaAttributeBuilder {
	return &SchemaAttributeBuilder{}
}

func (b *SchemaAttributeBuilder) initialize() {
	b.err = nil
	b.object = &SchemaAttribute{}
}
func (b *SchemaAttributeBuilder) CanonicalValues(in string) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeCanonicalValuesKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) CaseExact(in bool) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeCaseExactKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) Description(in string) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeDescriptionKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) MultiValued(in bool) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeMultiValuedKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) Mutability(in Mutability) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeMutabilityKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) Name(in string) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeNameKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) ReferenceTypes(in ...string) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeReferenceTypesKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) Required(in bool) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeRequiredKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) Returned(in Returned) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeReturnedKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) SubAttributes(in ...*SchemaAttribute) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeSubAttributesKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) Type(in DataType) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeTypeKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *SchemaAttributeBuilder) Uniqueness(in Uniqueness) *SchemaAttributeBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(SchemaAttributeUniquenessKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *SchemaAttributeBuilder) Build() (*SchemaAttribute, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err := b.err; err != nil {
		return nil, err
	}
	if b.object.multiValued == nil {
		return nil, fmt.Errorf("required field 'MultiValued' not initialized")
	}
	if b.object.typ == nil {
		return nil, fmt.Errorf("required field 'Type' not initialized")
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *SchemaAttributeBuilder) MustBuild() *SchemaAttribute {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *SchemaAttributeBuilder) From(in *SchemaAttribute) *SchemaAttributeBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *SchemaAttribute) Clone() *SchemaAttribute {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &SchemaAttribute{
		canonicalValues: v.canonicalValues,
		caseExact:       v.caseExact,
		description:     v.description,
		multiValued:     v.multiValued,
		mutability:      v.mutability,
		name:            v.name,
		referenceTypes:  v.referenceTypes,
		required:        v.required,
		returned:        v.returned,
		subAttributes:   v.subAttributes,
		typ:             v.typ,
		uniqueness:      v.uniqueness,
	}
}

func (v *SchemaAttribute) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *SchemaAttribute) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) SchemaAttribute() *SchemaAttributeBuilder {
	return &SchemaAttributeBuilder{}
}
