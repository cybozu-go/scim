package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
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
	externalID     *string
	id             *string
	manager        *EnterpriseManager
	meta           *Meta
	organization   *string
	schemas        []string
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

var DefaultEnterpriseUserValidator EnterpriseUserValidator

func (v *EnterpriseUser) CostCenter() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.costCenter == nil {
		return ""
	}
	return *(v.costCenter)
}

func (v *EnterpriseUser) Department() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.department == nil {
		return ""
	}
	return *(v.department)
}

func (v *EnterpriseUser) Division() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.division == nil {
		return ""
	}
	return *(v.division)
}

func (v *EnterpriseUser) EmployeeNumber() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.employeeNumber == nil {
		return ""
	}
	return *(v.employeeNumber)
}

func (v *EnterpriseUser) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.externalID == nil {
		return ""
	}
	return *(v.externalID)
}

func (v *EnterpriseUser) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *EnterpriseUser) Manager() *EnterpriseManager {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.manager
}

func (v *EnterpriseUser) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta
}

func (v *EnterpriseUser) Organization() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.organization == nil {
		return ""
	}
	return *(v.organization)
}

func (v *EnterpriseUser) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas
}

func (v *EnterpriseUser) MarshalJSON() ([]byte, error) {
	type pair struct {
		Key   string
		Value interface{}
	}
	var pairs []pair
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
	if v.externalID != nil {
		pairs = append(pairs, pair{Key: "externalId", Value: *(v.externalID)})
	}
	if v.id != nil {
		pairs = append(pairs, pair{Key: "id", Value: *(v.id)})
	}
	if v.manager != nil {
		pairs = append(pairs, pair{Key: "manager", Value: v.manager})
	}
	if v.meta != nil {
		pairs = append(pairs, pair{Key: "meta", Value: v.meta})
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
	for _, option := range options {
		switch option.Ident() {
		case identExtension{}:
			ext = option.Value().(string)
		}
	}
	switch name {
	case "costCenter":
		if v.costCenter == nil {
			return nil, false
		}
		return *(v.costCenter), true
	case "department":
		if v.department == nil {
			return nil, false
		}
		return *(v.department), true
	case "division":
		if v.division == nil {
			return nil, false
		}
		return *(v.division), true
	case "employeeNumber":
		if v.employeeNumber == nil {
			return nil, false
		}
		return *(v.employeeNumber), true
	case "externalId":
		if v.externalID == nil {
			return nil, false
		}
		return *(v.externalID), true
	case "id":
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case "manager":
		if v.manager == nil {
			return nil, false
		}
		return v.manager, true
	case "meta":
		if v.meta == nil {
			return nil, false
		}
		return v.meta, true
	case "organization":
		if v.organization == nil {
			return nil, false
		}
		return *(v.organization), true
	case "schemas":
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
	case "costCenter":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "costCenter", but got %T`, value)
		}
		v.costCenter = &tmp
		return nil
	case "department":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "department", but got %T`, value)
		}
		v.department = &tmp
		return nil
	case "division":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "division", but got %T`, value)
		}
		v.division = &tmp
		return nil
	case "employeeNumber":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "employeeNumber", but got %T`, value)
		}
		v.employeeNumber = &tmp
		return nil
	case "externalId":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "externalId", but got %T`, value)
		}
		v.externalID = &tmp
		return nil
	case "id":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case "manager":
		var tmp *EnterpriseManager
		tmp, ok := value.(*EnterpriseManager)
		if !ok {
			return fmt.Errorf(`expected *EnterpriseManager for field "manager", but got %T`, value)
		}
		v.manager = tmp
		return nil
	case "meta":
		var tmp *Meta
		tmp, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected *Meta for field "meta", but got %T`, value)
		}
		v.meta = tmp
		return nil
	case "organization":
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "organization", but got %T`, value)
		}
		v.organization = &tmp
		return nil
	case "schemas":
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "schemas", but got %T`, value)
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

func (v *EnterpriseUser) UnmarshalJSON(data []byte) error {
	v.costCenter = nil
	v.department = nil
	v.division = nil
	v.employeeNumber = nil
	v.externalID = nil
	v.id = nil
	v.manager = nil
	v.meta = nil
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
			case "costCenter":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "costCenter": %w`, err)
				}
				v.costCenter = &x
			case "department":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "department": %w`, err)
				}
				v.department = &x
			case "division":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "division": %w`, err)
				}
				v.division = &x
			case "employeeNumber":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "employeeNumber": %w`, err)
				}
				v.employeeNumber = &x
			case "externalId":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "externalId": %w`, err)
				}
				v.externalID = &x
			case "id":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case "manager":
				var x *EnterpriseManager
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "manager": %w`, err)
				}
				v.manager = x
			case "meta":
				var x *Meta
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "meta": %w`, err)
				}
				v.meta = x
			case "organization":
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "organization": %w`, err)
				}
				v.organization = &x
			case "schemas":
				var x []string
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

type EnterpriseUserBuilder struct {
	mu        sync.Mutex
	err       error
	validator EnterpriseUserValidator
	object    *EnterpriseUser
}

func (b *Builder) EnterpriseUser() *EnterpriseUserBuilder {
	return &EnterpriseUserBuilder{}
}

func (b *EnterpriseUserBuilder) CostCenter(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("costCenter", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Department(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("department", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Division(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("division", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) EmployeeNumber(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("employeeNumber", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) ExternalID(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("externalId", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) ID(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("id", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Manager(v *EnterpriseManager) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("manager", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Meta(v *Meta) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("meta", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Organization(v string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("organization", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Schemas(v ...string) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set("schemas", v); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Extension(uri string, value interface{}) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	if b.object == nil {
		b.object = &EnterpriseUser{}
	}
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (b *EnterpriseUserBuilder) Validator(v EnterpriseUserValidator) *EnterpriseUserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *EnterpriseUserBuilder) Build() (*EnterpriseUser, error) {
	object := b.object
	validator := b.validator
	b.object = nil
	b.validator = nil
	if object == nil {
		return nil, fmt.Errorf("resource.EnterpriseUserBuilder: object was not initialized")
	}
	if err := b.err; err != nil {
		return nil, err
	}
	if validator == nil {
		validator = DefaultEnterpriseUserValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
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
