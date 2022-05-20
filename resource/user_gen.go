package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	userActiveJSONKey            = "active"
	userAddressesJSONKey         = "addresses"
	userDisplayNameJSONKey       = "displayName"
	userEmailsJSONKey            = "emails"
	userEntitlementsJSONKey      = "entitlements"
	userExternalIDJSONKey        = "externalId"
	userGroupsJSONKey            = "groups"
	userIDJSONKey                = "id"
	userIMSJSONKey               = "ims"
	userMetaJSONKey              = "meta"
	userNameJSONKey              = "name"
	userNickNameJSONKey          = "nickName"
	userPasswordJSONKey          = "password"
	userPhoneNumbersJSONKey      = "phoneNumbers"
	userPreferredLanguageJSONKey = "preferredLanguage"
	userProfileURLJSONKey        = "profileUrl"
	userRolesJSONKey             = "roles"
	userSchemasJSONKey           = "schemas"
	userTimezoneJSONKey          = "timezone"
	userTitleJSONKey             = "title"
	userUserNameJSONKey          = "userName"
	userUserTypeJSONKey          = "userType"
	userX509CertificatesJSONKey  = "x509Certificates"
)

const UserSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:User"

func init() {
	RegisterExtension(UserSchemaURI, User{})
}

type User struct {
	active            *bool
	addresses         []string
	displayName       *string
	emails            []*Email
	entitlements      []string
	externalID        *string
	groups            []string
	id                *string
	ims               []string
	meta              *Meta
	name              *Names
	nickName          *string
	password          *string
	phoneNumbers      []string
	preferredLanguage *string
	profileURL        *string
	roles             []string
	schemas           []string
	timezone          *string
	title             *string
	userName          *string
	userType          *string
	x509Certificates  []string
	privateParams     map[string]interface{}
	mu                sync.RWMutex
}

type UserValidator interface {
	Validate(*User) error
}

type UserValidateFunc func(v *User) error

func (f UserValidateFunc) Validate(v *User) error {
	return f(v)
}

var DefaultUserValidator UserValidator

func (v *User) Active() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.active == nil {
		return false
	}
	return *(v.active)
}

func (v *User) Addresses() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.addresses
}

func (v *User) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.displayName == nil {
		return ""
	}
	return *(v.displayName)
}

func (v *User) Emails() []*Email {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.emails
}

func (v *User) Entitlements() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.entitlements
}

func (v *User) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.externalID == nil {
		return ""
	}
	return *(v.externalID)
}

func (v *User) Groups() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.groups
}

func (v *User) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *User) IMS() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ims
}

func (v *User) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta
}

func (v *User) Name() *Names {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name
}

func (v *User) NickName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.nickName == nil {
		return ""
	}
	return *(v.nickName)
}

func (v *User) Password() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.password == nil {
		return ""
	}
	return *(v.password)
}

func (v *User) PhoneNumbers() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.phoneNumbers
}

func (v *User) PreferredLanguage() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.preferredLanguage == nil {
		return ""
	}
	return *(v.preferredLanguage)
}

func (v *User) ProfileURL() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.profileURL == nil {
		return ""
	}
	return *(v.profileURL)
}

func (v *User) Roles() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.roles
}

func (v *User) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas
}

func (v *User) Timezone() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.timezone == nil {
		return ""
	}
	return *(v.timezone)
}

func (v *User) Title() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.title == nil {
		return ""
	}
	return *(v.title)
}

func (v *User) UserName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.userName == nil {
		return ""
	}
	return *(v.userName)
}

func (v *User) UserType() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.userType == nil {
		return ""
	}
	return *(v.userType)
}

func (v *User) X509Certificates() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509Certificates
}

func (v *User) makePairs() []pair {
	pairs := make([]pair, 0, 23)
	if v.active != nil {
		pairs = append(pairs, pair{Key: "active", Value: *(v.active)})
	}
	if v.addresses != nil {
		pairs = append(pairs, pair{Key: "addresses", Value: v.addresses})
	}
	if v.displayName != nil {
		pairs = append(pairs, pair{Key: "displayName", Value: *(v.displayName)})
	}
	if v.emails != nil {
		pairs = append(pairs, pair{Key: "emails", Value: v.emails})
	}
	if v.entitlements != nil {
		pairs = append(pairs, pair{Key: "entitlements", Value: v.entitlements})
	}
	if v.externalID != nil {
		pairs = append(pairs, pair{Key: "externalId", Value: *(v.externalID)})
	}
	if v.groups != nil {
		pairs = append(pairs, pair{Key: "groups", Value: v.groups})
	}
	if v.id != nil {
		pairs = append(pairs, pair{Key: "id", Value: *(v.id)})
	}
	if v.ims != nil {
		pairs = append(pairs, pair{Key: "ims", Value: v.ims})
	}
	if v.meta != nil {
		pairs = append(pairs, pair{Key: "meta", Value: v.meta})
	}
	if v.name != nil {
		pairs = append(pairs, pair{Key: "name", Value: v.name})
	}
	if v.nickName != nil {
		pairs = append(pairs, pair{Key: "nickName", Value: *(v.nickName)})
	}
	if v.password != nil {
		pairs = append(pairs, pair{Key: "password", Value: *(v.password)})
	}
	if v.phoneNumbers != nil {
		pairs = append(pairs, pair{Key: "phoneNumbers", Value: v.phoneNumbers})
	}
	if v.preferredLanguage != nil {
		pairs = append(pairs, pair{Key: "preferredLanguage", Value: *(v.preferredLanguage)})
	}
	if v.profileURL != nil {
		pairs = append(pairs, pair{Key: "profileUrl", Value: *(v.profileURL)})
	}
	if v.roles != nil {
		pairs = append(pairs, pair{Key: "roles", Value: v.roles})
	}
	if v.schemas != nil {
		pairs = append(pairs, pair{Key: "schemas", Value: v.schemas})
	}
	if v.timezone != nil {
		pairs = append(pairs, pair{Key: "timezone", Value: *(v.timezone)})
	}
	if v.title != nil {
		pairs = append(pairs, pair{Key: "title", Value: *(v.title)})
	}
	if v.userName != nil {
		pairs = append(pairs, pair{Key: "userName", Value: *(v.userName)})
	}
	if v.userType != nil {
		pairs = append(pairs, pair{Key: "userType", Value: *(v.userType)})
	}
	if v.x509Certificates != nil {
		pairs = append(pairs, pair{Key: "x509Certificates", Value: v.x509Certificates})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *User) MarshalJSON() ([]byte, error) {
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

func (v *User) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case userActiveJSONKey:
		if v.active == nil {
			return nil, false
		}
		return *(v.active), true
	case userAddressesJSONKey:
		if v.addresses == nil {
			return nil, false
		}
		return v.addresses, true
	case userDisplayNameJSONKey:
		if v.displayName == nil {
			return nil, false
		}
		return *(v.displayName), true
	case userEmailsJSONKey:
		if v.emails == nil {
			return nil, false
		}
		return v.emails, true
	case userEntitlementsJSONKey:
		if v.entitlements == nil {
			return nil, false
		}
		return v.entitlements, true
	case userExternalIDJSONKey:
		if v.externalID == nil {
			return nil, false
		}
		return *(v.externalID), true
	case userGroupsJSONKey:
		if v.groups == nil {
			return nil, false
		}
		return v.groups, true
	case userIDJSONKey:
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case userIMSJSONKey:
		if v.ims == nil {
			return nil, false
		}
		return v.ims, true
	case userMetaJSONKey:
		if v.meta == nil {
			return nil, false
		}
		return v.meta, true
	case userNameJSONKey:
		if v.name == nil {
			return nil, false
		}
		return v.name, true
	case userNickNameJSONKey:
		if v.nickName == nil {
			return nil, false
		}
		return *(v.nickName), true
	case userPasswordJSONKey:
		if v.password == nil {
			return nil, false
		}
		return *(v.password), true
	case userPhoneNumbersJSONKey:
		if v.phoneNumbers == nil {
			return nil, false
		}
		return v.phoneNumbers, true
	case userPreferredLanguageJSONKey:
		if v.preferredLanguage == nil {
			return nil, false
		}
		return *(v.preferredLanguage), true
	case userProfileURLJSONKey:
		if v.profileURL == nil {
			return nil, false
		}
		return *(v.profileURL), true
	case userRolesJSONKey:
		if v.roles == nil {
			return nil, false
		}
		return v.roles, true
	case userSchemasJSONKey:
		if v.schemas == nil {
			return nil, false
		}
		return v.schemas, true
	case userTimezoneJSONKey:
		if v.timezone == nil {
			return nil, false
		}
		return *(v.timezone), true
	case userTitleJSONKey:
		if v.title == nil {
			return nil, false
		}
		return *(v.title), true
	case userUserNameJSONKey:
		if v.userName == nil {
			return nil, false
		}
		return *(v.userName), true
	case userUserTypeJSONKey:
		if v.userType == nil {
			return nil, false
		}
		return *(v.userType), true
	case userX509CertificatesJSONKey:
		if v.x509Certificates == nil {
			return nil, false
		}
		return v.x509Certificates, true
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

func (v *User) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case userActiveJSONKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "active", but got %T`, value)
		}
		v.active = &tmp
		return nil
	case userAddressesJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "addresses", but got %T`, value)
		}
		v.addresses = tmp
		return nil
	case userDisplayNameJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "displayName", but got %T`, value)
		}
		v.displayName = &tmp
		return nil
	case userEmailsJSONKey:
		var tmp []*Email
		tmp, ok := value.([]*Email)
		if !ok {
			return fmt.Errorf(`expected []*Email for field "emails", but got %T`, value)
		}
		v.emails = tmp
		return nil
	case userEntitlementsJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "entitlements", but got %T`, value)
		}
		v.entitlements = tmp
		return nil
	case userExternalIDJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "externalId", but got %T`, value)
		}
		v.externalID = &tmp
		return nil
	case userGroupsJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "groups", but got %T`, value)
		}
		v.groups = tmp
		return nil
	case userIDJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case userIMSJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "ims", but got %T`, value)
		}
		v.ims = tmp
		return nil
	case userMetaJSONKey:
		var tmp *Meta
		tmp, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected *Meta for field "meta", but got %T`, value)
		}
		v.meta = tmp
		return nil
	case userNameJSONKey:
		var tmp *Names
		tmp, ok := value.(*Names)
		if !ok {
			return fmt.Errorf(`expected *Names for field "name", but got %T`, value)
		}
		v.name = tmp
		return nil
	case userNickNameJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "nickName", but got %T`, value)
		}
		v.nickName = &tmp
		return nil
	case userPasswordJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "password", but got %T`, value)
		}
		v.password = &tmp
		return nil
	case userPhoneNumbersJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "phoneNumbers", but got %T`, value)
		}
		v.phoneNumbers = tmp
		return nil
	case userPreferredLanguageJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "preferredLanguage", but got %T`, value)
		}
		v.preferredLanguage = &tmp
		return nil
	case userProfileURLJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "profileUrl", but got %T`, value)
		}
		v.profileURL = &tmp
		return nil
	case userRolesJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "roles", but got %T`, value)
		}
		v.roles = tmp
		return nil
	case userSchemasJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "schemas", but got %T`, value)
		}
		v.schemas = tmp
		return nil
	case userTimezoneJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "timezone", but got %T`, value)
		}
		v.timezone = &tmp
		return nil
	case userTitleJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "title", but got %T`, value)
		}
		v.title = &tmp
		return nil
	case userUserNameJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "userName", but got %T`, value)
		}
		v.userName = &tmp
		return nil
	case userUserTypeJSONKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "userType", but got %T`, value)
		}
		v.userType = &tmp
		return nil
	case userX509CertificatesJSONKey:
		var tmp []string
		tmp, ok := value.([]string)
		if !ok {
			return fmt.Errorf(`expected []string for field "x509Certificates", but got %T`, value)
		}
		v.x509Certificates = tmp
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

func (v *User) UnmarshalJSON(data []byte) error {
	v.active = nil
	v.addresses = nil
	v.displayName = nil
	v.emails = nil
	v.entitlements = nil
	v.externalID = nil
	v.groups = nil
	v.id = nil
	v.ims = nil
	v.meta = nil
	v.name = nil
	v.nickName = nil
	v.password = nil
	v.phoneNumbers = nil
	v.preferredLanguage = nil
	v.profileURL = nil
	v.roles = nil
	v.schemas = nil
	v.timezone = nil
	v.title = nil
	v.userName = nil
	v.userType = nil
	v.x509Certificates = nil
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
			case userActiveJSONKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "active": %w`, err)
				}
				v.active = &x
			case userAddressesJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "addresses": %w`, err)
				}
				v.addresses = x
			case userDisplayNameJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "displayName": %w`, err)
				}
				v.displayName = &x
			case userEmailsJSONKey:
				var x []*Email
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "emails": %w`, err)
				}
				v.emails = x
			case userEntitlementsJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "entitlements": %w`, err)
				}
				v.entitlements = x
			case userExternalIDJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "externalId": %w`, err)
				}
				v.externalID = &x
			case userGroupsJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "groups": %w`, err)
				}
				v.groups = x
			case userIDJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case userIMSJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "ims": %w`, err)
				}
				v.ims = x
			case userMetaJSONKey:
				var x *Meta
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "meta": %w`, err)
				}
				v.meta = x
			case userNameJSONKey:
				var x *Names
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "name": %w`, err)
				}
				v.name = x
			case userNickNameJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "nickName": %w`, err)
				}
				v.nickName = &x
			case userPasswordJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "password": %w`, err)
				}
				v.password = &x
			case userPhoneNumbersJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "phoneNumbers": %w`, err)
				}
				v.phoneNumbers = x
			case userPreferredLanguageJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "preferredLanguage": %w`, err)
				}
				v.preferredLanguage = &x
			case userProfileURLJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "profileUrl": %w`, err)
				}
				v.profileURL = &x
			case userRolesJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "roles": %w`, err)
				}
				v.roles = x
			case userSchemasJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = x
			case userTimezoneJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "timezone": %w`, err)
				}
				v.timezone = &x
			case userTitleJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "title": %w`, err)
				}
				v.title = &x
			case userUserNameJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "userName": %w`, err)
				}
				v.userName = &x
			case userUserTypeJSONKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "userType": %w`, err)
				}
				v.userType = &x
			case userX509CertificatesJSONKey:
				var x []string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "x509Certificates": %w`, err)
				}
				v.x509Certificates = x
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

func (v *User) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type UserBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator UserValidator
	object    *User
}

func (b *Builder) User() *UserBuilder {
	return NewUserBuilder()
}

func NewUserBuilder() *UserBuilder {
	var b UserBuilder
	b.init()
	return &b
}

func (b *UserBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &User{}

	b.object.schemas = []string{UserSchemaURI}
}

func (b *UserBuilder) Active(v bool) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("active", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Addresses(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("addresses", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) DisplayName(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("displayName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Emails(v ...*Email) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("emails", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Entitlements(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("entitlements", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) ExternalID(v string) *UserBuilder {
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

func (b *UserBuilder) Groups(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("groups", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) ID(v string) *UserBuilder {
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

func (b *UserBuilder) IMS(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("ims", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Meta(v *Meta) *UserBuilder {
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

func (b *UserBuilder) Name(v *Names) *UserBuilder {
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

func (b *UserBuilder) NickName(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("nickName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Password(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("password", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) PhoneNumbers(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("phoneNumbers", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) PreferredLanguage(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("preferredLanguage", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) ProfileURL(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("profileUrl", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Roles(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("roles", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Schemas(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("schemas", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Timezone(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("timezone", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Title(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("title", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) UserName(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("userName", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) UserType(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("userType", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) X509Certificates(v ...string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("x509Certificates", v); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Extension(uri string, value interface{}) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (b *UserBuilder) Validator(v UserValidator) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *UserBuilder) Build() (*User, error) {
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
		return nil, fmt.Errorf("resource.UserBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultUserValidator
	}
	if validator != nil {
		if err := validator.Validate(object); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func (b *UserBuilder) MustBuild() *User {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
