package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const EnterpriseUserSchemaURI = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"

func init() {
	RegisterExtension(EnterpriseUserSchemaURI, EnterpriseUser{})
}

type EnterpriseUser struct {
	mu             sync.RWMutex
	costCenter     *string
	department     *string
	division       *string
	employeeNumber *string
	manager        *EnterpriseManager
	organization   *string
	schemas        *schemas
	extra          map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	EnterpriseUserCostCenterKey     = "costCenter"
	EnterpriseUserDepartmentKey     = "department"
	EnterpriseUserDivisionKey       = "division"
	EnterpriseUserEmployeeNumberKey = "employeeNumber"
	EnterpriseUserManagerKey        = "manager"
	EnterpriseUserOrganizationKey   = "organization"
	EnterpriseUserSchemasKey        = "schemas"
)

// Get retrieves the value associated with a key
func (v *EnterpriseUser) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case EnterpriseUserCostCenterKey:
		if val := v.costCenter; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EnterpriseUserDepartmentKey:
		if val := v.department; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EnterpriseUserDivisionKey:
		if val := v.division; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EnterpriseUserEmployeeNumberKey:
		if val := v.employeeNumber; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EnterpriseUserManagerKey:
		if val := v.manager; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case EnterpriseUserOrganizationKey:
		if val := v.organization; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EnterpriseUserSchemasKey:
		if val := v.schemas; val != nil {
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
func (v *EnterpriseUser) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case EnterpriseUserCostCenterKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field costCenter, got %T`, value)
		}
		v.costCenter = &converted
	case EnterpriseUserDepartmentKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field department, got %T`, value)
		}
		v.department = &converted
	case EnterpriseUserDivisionKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field division, got %T`, value)
		}
		v.division = &converted
	case EnterpriseUserEmployeeNumberKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field employeeNumber, got %T`, value)
		}
		v.employeeNumber = &converted
	case EnterpriseUserManagerKey:
		converted, ok := value.(*EnterpriseManager)
		if !ok {
			return fmt.Errorf(`expected value of type *EnterpriseManager for field manager, got %T`, value)
		}
		v.manager = converted
	case EnterpriseUserOrganizationKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field organization, got %T`, value)
		}
		v.organization = &converted
	case EnterpriseUserSchemasKey:
		converted, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected value of type schemas for field schemas, got %T`, value)
		}
		v.schemas = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *EnterpriseUser) HasCostCenter() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.costCenter != nil
}

func (v *EnterpriseUser) HasDepartment() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.department != nil
}

func (v *EnterpriseUser) HasDivision() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.division != nil
}

func (v *EnterpriseUser) HasEmployeeNumber() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.employeeNumber != nil
}

func (v *EnterpriseUser) HasManager() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.manager != nil
}

func (v *EnterpriseUser) HasOrganization() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.organization != nil
}

func (v *EnterpriseUser) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *EnterpriseUser) CostCenter() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.costCenter; val != nil {
		return *val
	}
	return ""
}

func (v *EnterpriseUser) Department() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.department; val != nil {
		return *val
	}
	return ""
}

func (v *EnterpriseUser) Division() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.division; val != nil {
		return *val
	}
	return ""
}

func (v *EnterpriseUser) EmployeeNumber() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.employeeNumber; val != nil {
		return *val
	}
	return ""
}

func (v *EnterpriseUser) Manager() *EnterpriseManager {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.manager; val != nil {
		return val
	}
	return nil
}

func (v *EnterpriseUser) Organization() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.organization; val != nil {
		return *val
	}
	return ""
}

func (v *EnterpriseUser) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

// Remove removes the value associated with a key
func (v *EnterpriseUser) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case EnterpriseUserCostCenterKey:
		v.costCenter = nil
	case EnterpriseUserDepartmentKey:
		v.department = nil
	case EnterpriseUserDivisionKey:
		v.division = nil
	case EnterpriseUserEmployeeNumberKey:
		v.employeeNumber = nil
	case EnterpriseUserManagerKey:
		v.manager = nil
	case EnterpriseUserOrganizationKey:
		v.organization = nil
	case EnterpriseUserSchemasKey:
		v.schemas = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *EnterpriseUser) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 7)
	if val := v.costCenter; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseUserCostCenterKey, Value: *val})
	}
	if val := v.department; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseUserDepartmentKey, Value: *val})
	}
	if val := v.division; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseUserDivisionKey, Value: *val})
	}
	if val := v.employeeNumber; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseUserEmployeeNumberKey, Value: *val})
	}
	if val := v.manager; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseUserManagerKey, Value: val})
	}
	if val := v.organization; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseUserOrganizationKey, Value: *val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: EnterpriseUserSchemasKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes EnterpriseUser into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *EnterpriseUser) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into EnterpriseUser.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *EnterpriseUser) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.costCenter = nil
	v.department = nil
	v.division = nil
	v.employeeNumber = nil
	v.manager = nil
	v.organization = nil
	v.schemas = nil

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
			case EnterpriseUserCostCenterKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseUserCostCenterKey, err)
				}
				v.costCenter = &val
			case EnterpriseUserDepartmentKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseUserDepartmentKey, err)
				}
				v.department = &val
			case EnterpriseUserDivisionKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseUserDivisionKey, err)
				}
				v.division = &val
			case EnterpriseUserEmployeeNumberKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseUserEmployeeNumberKey, err)
				}
				v.employeeNumber = &val
			case EnterpriseUserManagerKey:
				var val *EnterpriseManager
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseUserManagerKey, err)
				}
				v.manager = val
			case EnterpriseUserOrganizationKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseUserOrganizationKey, err)
				}
				v.organization = &val
			case EnterpriseUserSchemasKey:
				var val schemas
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EnterpriseUserSchemasKey, err)
				}
				v.schemas = &val
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

type EnterpriseUserBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *EnterpriseUser
}

// NewEnterpriseUserBuilder creates a new EnterpriseUserBuilder instance.
// EnterpriseUserBuilder is safe to be used uninitialized as well.
func NewEnterpriseUserBuilder() *EnterpriseUserBuilder {
	return &EnterpriseUserBuilder{}
}

func (b *EnterpriseUserBuilder) initialize() {
	b.err = nil
	b.object = &EnterpriseUser{}
}
func (b *EnterpriseUserBuilder) CostCenter(in string) *EnterpriseUserBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(EnterpriseUserCostCenterKey, in)
	return b
}
func (b *EnterpriseUserBuilder) Department(in string) *EnterpriseUserBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(EnterpriseUserDepartmentKey, in)
	return b
}
func (b *EnterpriseUserBuilder) Division(in string) *EnterpriseUserBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(EnterpriseUserDivisionKey, in)
	return b
}
func (b *EnterpriseUserBuilder) EmployeeNumber(in string) *EnterpriseUserBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(EnterpriseUserEmployeeNumberKey, in)
	return b
}
func (b *EnterpriseUserBuilder) Manager(in *EnterpriseManager) *EnterpriseUserBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(EnterpriseUserManagerKey, in)
	return b
}
func (b *EnterpriseUserBuilder) Organization(in string) *EnterpriseUserBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(EnterpriseUserOrganizationKey, in)
	return b
}
func (b *EnterpriseUserBuilder) Schemas(in ...string) *EnterpriseUserBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = b.object.Set(EnterpriseUserSchemasKey, in)
	return b
}

func (b *EnterpriseUserBuilder) Build() (*EnterpriseUser, error) {
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

func (b *EnterpriseUserBuilder) MustBuild() *EnterpriseUser {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *EnterpriseUserBuilder) From(in *EnterpriseUser) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.object = in.Clone()
	return b
}

func (b *EnterpriseUserBuilder) Extension(uri string, value interface{}) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}
	if b.object.schemas == nil {
		b.object.schemas = &schemas{}
		b.object.schemas.Add(EnterpriseUserSchemaURI)
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (v *EnterpriseUser) Clone() *EnterpriseUser {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &EnterpriseUser{
		costCenter:     v.costCenter,
		department:     v.department,
		division:       v.division,
		employeeNumber: v.employeeNumber,
		manager:        v.manager,
		organization:   v.organization,
		schemas:        v.schemas,
	}
}

func (v *EnterpriseUser) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *EnterpriseUser) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) EnterpriseUser() *EnterpriseUserBuilder {
	return &EnterpriseUserBuilder{}
}
