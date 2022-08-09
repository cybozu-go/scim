package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	EnterpriseUserCostCenterKey     = "costCenter"
	EnterpriseUserDepartmentKey     = "department"
	EnterpriseUserDivisionKey       = "division"
	EnterpriseUserEmployeeNumberKey = "employeeNumber"
	EnterpriseUserManagerKey        = "manager"
	EnterpriseUserOrganizationKey   = "organization"
	EnterpriseUserSchemasKey        = "schemas"
)

const EnterpriseUserSchemaURI = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"

func init() {
	RegisterExtension(EnterpriseUserSchemaURI, EnterpriseUser{})
}

type EnterpriseUser struct {
	costCenter     *string
	department     *string
	division       *string
	employeeNumber *string
	manager        *EnterpriseManager
	organization   *string
	schemas        schemas
	privateParams  map[string]interface{}
	mu             sync.RWMutex
}

type EnterpriseUserValidator interface {
	Validate(*EnterpriseUser) error
}

type EnterpriseUserValidateFunc func(v *EnterpriseUser) error

func (f EnterpriseUserValidateFunc) Validate(v *EnterpriseUser) error {
	return f(v)
}

var DefaultEnterpriseUserValidator EnterpriseUserValidator = EnterpriseUserValidateFunc(func(v *EnterpriseUser) error {
	return nil
})

func (v *EnterpriseUser) HasCostCenter() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.costCenter != nil
}

func (v *EnterpriseUser) CostCenter() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.costCenter == nil {
		return ""
	}
	return *(v.costCenter)
}

func (v *EnterpriseUser) HasDepartment() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.department != nil
}

func (v *EnterpriseUser) Department() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.department == nil {
		return ""
	}
	return *(v.department)
}

func (v *EnterpriseUser) HasDivision() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.division != nil
}

func (v *EnterpriseUser) Division() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.division == nil {
		return ""
	}
	return *(v.division)
}

func (v *EnterpriseUser) HasEmployeeNumber() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.employeeNumber != nil
}

func (v *EnterpriseUser) EmployeeNumber() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.employeeNumber == nil {
		return ""
	}
	return *(v.employeeNumber)
}

func (v *EnterpriseUser) HasManager() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.manager != nil
}

func (v *EnterpriseUser) Manager() *EnterpriseManager {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.manager
}

func (v *EnterpriseUser) HasOrganization() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.organization != nil
}

func (v *EnterpriseUser) Organization() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.organization == nil {
		return ""
	}
	return *(v.organization)
}

func (v *EnterpriseUser) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *EnterpriseUser) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *EnterpriseUser) makePairs() []pair {
	pairs := make([]pair, 0, 7)
	if v.costCenter != nil {
		pairs = append(pairs, pair{Key: "costCenter", Value: *(v.costCenter)})
	}
	if v.department != nil {
		pairs = append(pairs, pair{Key: "department", Value: *(v.department)})
	}
	if v.division != nil {
		pairs = append(pairs, pair{Key: "division", Value: *(v.division)})
	}
	if v.employeeNumber != nil {
		pairs = append(pairs, pair{Key: "employeeNumber", Value: *(v.employeeNumber)})
	}
	if v.manager != nil {
		pairs = append(pairs, pair{Key: "manager", Value: v.manager})
	}
	if v.organization != nil {
		pairs = append(pairs, pair{Key: "organization", Value: *(v.organization)})
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

func (v *EnterpriseUser) MarshalJSON() ([]byte, error) {
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

func (v *EnterpriseUser) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case EnterpriseUserCostCenterKey:
		if v.costCenter == nil {
			return nil, false
		}
		return *(v.costCenter), true
	case EnterpriseUserDepartmentKey:
		if v.department == nil {
			return nil, false
		}
		return *(v.department), true
	case EnterpriseUserDivisionKey:
		if v.division == nil {
			return nil, false
		}
		return *(v.division), true
	case EnterpriseUserEmployeeNumberKey:
		if v.employeeNumber == nil {
			return nil, false
		}
		return *(v.employeeNumber), true
	case EnterpriseUserManagerKey:
		if v.manager == nil {
			return nil, false
		}
		return v.manager, true
	case EnterpriseUserOrganizationKey:
		if v.organization == nil {
			return nil, false
		}
		return *(v.organization), true
	case EnterpriseUserSchemasKey:
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

func (v *EnterpriseUser) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case EnterpriseUserCostCenterKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "costCenter", but got %T`, value)
		}
		v.costCenter = &tmp
		return nil
	case EnterpriseUserDepartmentKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "department", but got %T`, value)
		}
		v.department = &tmp
		return nil
	case EnterpriseUserDivisionKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "division", but got %T`, value)
		}
		v.division = &tmp
		return nil
	case EnterpriseUserEmployeeNumberKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "employeeNumber", but got %T`, value)
		}
		v.employeeNumber = &tmp
		return nil
	case EnterpriseUserManagerKey:
		var tmp *EnterpriseManager
		tmp, ok := value.(*EnterpriseManager)
		if !ok {
			return fmt.Errorf(`expected *EnterpriseManager for field "manager", but got %T`, value)
		}
		v.manager = tmp
		return nil
	case EnterpriseUserOrganizationKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "organization", but got %T`, value)
		}
		v.organization = &tmp
		return nil
	case EnterpriseUserSchemasKey:
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

func (v *EnterpriseUser) Clone() *EnterpriseUser {
	v.mu.Lock()
	defer v.mu.Unlock()
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

func (v *EnterpriseUser) UnmarshalJSON(data []byte) error {
	v.costCenter = nil
	v.department = nil
	v.division = nil
	v.employeeNumber = nil
	v.manager = nil
	v.organization = nil
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
			case EnterpriseUserCostCenterKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "costCenter": %w`, err)
				}
				v.costCenter = &x
			case EnterpriseUserDepartmentKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "department": %w`, err)
				}
				v.department = &x
			case EnterpriseUserDivisionKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "division": %w`, err)
				}
				v.division = &x
			case EnterpriseUserEmployeeNumberKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "employeeNumber": %w`, err)
				}
				v.employeeNumber = &x
			case EnterpriseUserManagerKey:
				var x *EnterpriseManager
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "manager": %w`, err)
				}
				v.manager = x
			case EnterpriseUserOrganizationKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "organization": %w`, err)
				}
				v.organization = &x
			case EnterpriseUserSchemasKey:
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

func (v *EnterpriseUser) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// EnterpriseUserBuilder creates a EnterpriseUser resource
type EnterpriseUserBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator EnterpriseUserValidator
	object    *EnterpriseUser
}

func (b *Builder) EnterpriseUser() *EnterpriseUserBuilder {
	return NewEnterpriseUserBuilder()
}

func NewEnterpriseUserBuilder() *EnterpriseUserBuilder {
	var b EnterpriseUserBuilder
	b.init()
	return &b
}

func (b *EnterpriseUserBuilder) From(in *EnterpriseUser) *EnterpriseUserBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *EnterpriseUserBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &EnterpriseUser{}
}

func (b *EnterpriseUserBuilder) CostCenter(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("costCenter", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Department(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("department", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Division(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("division", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) EmployeeNumber(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("employeeNumber", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Manager(v *EnterpriseManager) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("manager", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Organization(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("organization", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Schemas(v ...string) *EnterpriseUserBuilder {
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

func (b *EnterpriseUserBuilder) Extension(uri string, value interface{}) *EnterpriseUserBuilder {
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

func (b *EnterpriseUserBuilder) Validator(v EnterpriseUserValidator) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *EnterpriseUserBuilder) Build() (*EnterpriseUser, error) {
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
		return nil, fmt.Errorf("resource.EnterpriseUserBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultEnterpriseUserValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *EnterpriseUserBuilder) MustBuild() *EnterpriseUser {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
