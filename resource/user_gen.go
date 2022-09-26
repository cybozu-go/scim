package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const UserSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:User"

func init() {
	RegisterExtension(UserSchemaURI, User{})
}

// User represents a User resource as defined in the SCIM RFC
type User struct {
	mu                sync.RWMutex
	active            *bool
	addresses         []*Address
	displayName       *string
	emails            []*Email
	entitlements      []*Entitlement
	externalId        *string
	groups            []*GroupMember
	id                *string
	ims               []*IMS
	locale            *string
	meta              *Meta
	name              *Names
	nickName          *string
	password          *string
	phoneNumbers      []*PhoneNumber
	photos            []*Photo
	preferredLanguage *string
	profileUrl        *string
	roles             []*Role
	schemas           *schemas
	timezone          *string
	title             *string
	userName          *string
	userType          *string
	x509Certificates  []*X509Certificate
	extra             map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	UserActiveKey            = "active"
	UserAddressesKey         = "addresses"
	UserDisplayNameKey       = "displayName"
	UserEmailsKey            = "emails"
	UserEntitlementsKey      = "entitlements"
	UserExternalIDKey        = "externalId"
	UserGroupsKey            = "groups"
	UserIDKey                = "id"
	UserIMSKey               = "ims"
	UserLocaleKey            = "locale"
	UserMetaKey              = "meta"
	UserNameKey              = "name"
	UserNickNameKey          = "nickName"
	UserPasswordKey          = "password"
	UserPhoneNumbersKey      = "phoneNumbers"
	UserPhotosKey            = "photos"
	UserPreferredLanguageKey = "preferredLanguage"
	UserProfileURLKey        = "profileUrl"
	UserRolesKey             = "roles"
	UserSchemasKey           = "schemas"
	UserTimezoneKey          = "timezone"
	UserTitleKey             = "title"
	UserUserNameKey          = "userName"
	UserUserTypeKey          = "userType"
	UserX509CertificatesKey  = "x509Certificates"
)

// Get retrieves the value associated with a key
func (v *User) Get(key string, dst interface{}) error {
	switch key {
	case UserActiveKey:
		if val := v.active; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserAddressesKey:
		if val := v.addresses; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserDisplayNameKey:
		if val := v.displayName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserEmailsKey:
		if val := v.emails; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserEntitlementsKey:
		if val := v.entitlements; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserExternalIDKey:
		if val := v.externalId; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserGroupsKey:
		if val := v.groups; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserIDKey:
		if val := v.id; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserIMSKey:
		if val := v.ims; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserLocaleKey:
		if val := v.locale; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserMetaKey:
		if val := v.meta; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserNameKey:
		if val := v.name; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserNickNameKey:
		if val := v.nickName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserPasswordKey:
		if val := v.password; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserPhoneNumbersKey:
		if val := v.phoneNumbers; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserPhotosKey:
		if val := v.photos; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserPreferredLanguageKey:
		if val := v.preferredLanguage; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserProfileURLKey:
		if val := v.profileUrl; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserRolesKey:
		if val := v.roles; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case UserSchemasKey:
		if val := v.schemas; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserTimezoneKey:
		if val := v.timezone; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserTitleKey:
		if val := v.title; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserUserNameKey:
		if val := v.userName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserUserTypeKey:
		if val := v.userType; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UserX509CertificatesKey:
		if val := v.x509Certificates; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
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
func (v *User) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case UserActiveKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field active, got %T`, value)
		}
		v.active = &converted
	case UserAddressesKey:
		converted, ok := value.([]*Address)
		if !ok {
			return fmt.Errorf(`expected value of type []*Address for field addresses, got %T`, value)
		}
		v.addresses = converted
	case UserDisplayNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field displayName, got %T`, value)
		}
		v.displayName = &converted
	case UserEmailsKey:
		converted, ok := value.([]*Email)
		if !ok {
			return fmt.Errorf(`expected value of type []*Email for field emails, got %T`, value)
		}
		v.emails = converted
	case UserEntitlementsKey:
		converted, ok := value.([]*Entitlement)
		if !ok {
			return fmt.Errorf(`expected value of type []*Entitlement for field entitlements, got %T`, value)
		}
		v.entitlements = converted
	case UserExternalIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field externalId, got %T`, value)
		}
		v.externalId = &converted
	case UserGroupsKey:
		converted, ok := value.([]*GroupMember)
		if !ok {
			return fmt.Errorf(`expected value of type []*GroupMember for field groups, got %T`, value)
		}
		v.groups = converted
	case UserIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field id, got %T`, value)
		}
		v.id = &converted
	case UserIMSKey:
		converted, ok := value.([]*IMS)
		if !ok {
			return fmt.Errorf(`expected value of type []*IMS for field ims, got %T`, value)
		}
		v.ims = converted
	case UserLocaleKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field locale, got %T`, value)
		}
		v.locale = &converted
	case UserMetaKey:
		converted, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected value of type *Meta for field meta, got %T`, value)
		}
		v.meta = converted
	case UserNameKey:
		converted, ok := value.(*Names)
		if !ok {
			return fmt.Errorf(`expected value of type *Names for field name, got %T`, value)
		}
		v.name = converted
	case UserNickNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field nickName, got %T`, value)
		}
		v.nickName = &converted
	case UserPasswordKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field password, got %T`, value)
		}
		v.password = &converted
	case UserPhoneNumbersKey:
		converted, ok := value.([]*PhoneNumber)
		if !ok {
			return fmt.Errorf(`expected value of type []*PhoneNumber for field phoneNumbers, got %T`, value)
		}
		v.phoneNumbers = converted
	case UserPhotosKey:
		converted, ok := value.([]*Photo)
		if !ok {
			return fmt.Errorf(`expected value of type []*Photo for field photos, got %T`, value)
		}
		v.photos = converted
	case UserPreferredLanguageKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field preferredLanguage, got %T`, value)
		}
		v.preferredLanguage = &converted
	case UserProfileURLKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field profileUrl, got %T`, value)
		}
		v.profileUrl = &converted
	case UserRolesKey:
		converted, ok := value.([]*Role)
		if !ok {
			return fmt.Errorf(`expected value of type []*Role for field roles, got %T`, value)
		}
		v.roles = converted
	case UserSchemasKey:
		converted, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected value of type schemas for field schemas, got %T`, value)
		}
		v.schemas = &converted
	case UserTimezoneKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field timezone, got %T`, value)
		}
		v.timezone = &converted
	case UserTitleKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field title, got %T`, value)
		}
		v.title = &converted
	case UserUserNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field userName, got %T`, value)
		}
		v.userName = &converted
	case UserUserTypeKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field userType, got %T`, value)
		}
		v.userType = &converted
	case UserX509CertificatesKey:
		converted, ok := value.([]*X509Certificate)
		if !ok {
			return fmt.Errorf(`expected value of type []*X509Certificate for field x509Certificates, got %T`, value)
		}
		v.x509Certificates = converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *User) HasActive() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.active != nil
}

func (v *User) HasAddresses() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.addresses != nil
}

func (v *User) HasDisplayName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.displayName != nil
}

func (v *User) HasEmails() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.emails != nil
}

func (v *User) HasEntitlements() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.entitlements != nil
}

func (v *User) HasExternalID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.externalId != nil
}

func (v *User) HasGroups() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.groups != nil
}

func (v *User) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *User) HasIMS() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ims != nil
}

func (v *User) HasLocale() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.locale != nil
}

func (v *User) HasMeta() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta != nil
}

func (v *User) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

func (v *User) HasNickName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.nickName != nil
}

func (v *User) HasPassword() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.password != nil
}

func (v *User) HasPhoneNumbers() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.phoneNumbers != nil
}

func (v *User) HasPhotos() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.photos != nil
}

func (v *User) HasPreferredLanguage() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.preferredLanguage != nil
}

func (v *User) HasProfileURL() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.profileUrl != nil
}

func (v *User) HasRoles() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.roles != nil
}

func (v *User) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *User) HasTimezone() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.timezone != nil
}

func (v *User) HasTitle() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.title != nil
}

func (v *User) HasUserName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.userName != nil
}

func (v *User) HasUserType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.userType != nil
}

func (v *User) HasX509Certificates() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509Certificates != nil
}

func (v *User) Active() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.active; val != nil {
		return *val
	}
	return false
}

func (v *User) Addresses() []*Address {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.addresses; val != nil {
		return val
	}
	return nil
}

func (v *User) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.displayName; val != nil {
		return *val
	}
	return ""
}

func (v *User) Emails() []*Email {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.emails; val != nil {
		return val
	}
	return nil
}

func (v *User) Entitlements() []*Entitlement {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.entitlements; val != nil {
		return val
	}
	return nil
}

func (v *User) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.externalId; val != nil {
		return *val
	}
	return ""
}

func (v *User) Groups() []*GroupMember {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.groups; val != nil {
		return val
	}
	return nil
}

func (v *User) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.id; val != nil {
		return *val
	}
	return ""
}

func (v *User) IMS() []*IMS {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.ims; val != nil {
		return val
	}
	return nil
}

func (v *User) Locale() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.locale; val != nil {
		return *val
	}
	return ""
}

func (v *User) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.meta; val != nil {
		return val
	}
	return nil
}

func (v *User) Name() *Names {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.name; val != nil {
		return val
	}
	return nil
}

func (v *User) NickName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.nickName; val != nil {
		return *val
	}
	return ""
}

func (v *User) Password() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.password; val != nil {
		return *val
	}
	return ""
}

func (v *User) PhoneNumbers() []*PhoneNumber {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.phoneNumbers; val != nil {
		return val
	}
	return nil
}

func (v *User) Photos() []*Photo {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.photos; val != nil {
		return val
	}
	return nil
}

func (v *User) PreferredLanguage() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.preferredLanguage; val != nil {
		return *val
	}
	return ""
}

func (v *User) ProfileURL() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.profileUrl; val != nil {
		return *val
	}
	return ""
}

func (v *User) Roles() []*Role {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.roles; val != nil {
		return val
	}
	return nil
}

func (v *User) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

func (v *User) Timezone() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.timezone; val != nil {
		return *val
	}
	return ""
}

func (v *User) Title() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.title; val != nil {
		return *val
	}
	return ""
}

func (v *User) UserName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.userName; val != nil {
		return *val
	}
	return ""
}

func (v *User) UserType() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.userType; val != nil {
		return *val
	}
	return ""
}

func (v *User) X509Certificates() []*X509Certificate {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.x509Certificates; val != nil {
		return val
	}
	return nil
}

// Remove removes the value associated with a key
func (v *User) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case UserActiveKey:
		v.active = nil
	case UserAddressesKey:
		v.addresses = nil
	case UserDisplayNameKey:
		v.displayName = nil
	case UserEmailsKey:
		v.emails = nil
	case UserEntitlementsKey:
		v.entitlements = nil
	case UserExternalIDKey:
		v.externalId = nil
	case UserGroupsKey:
		v.groups = nil
	case UserIDKey:
		v.id = nil
	case UserIMSKey:
		v.ims = nil
	case UserLocaleKey:
		v.locale = nil
	case UserMetaKey:
		v.meta = nil
	case UserNameKey:
		v.name = nil
	case UserNickNameKey:
		v.nickName = nil
	case UserPasswordKey:
		v.password = nil
	case UserPhoneNumbersKey:
		v.phoneNumbers = nil
	case UserPhotosKey:
		v.photos = nil
	case UserPreferredLanguageKey:
		v.preferredLanguage = nil
	case UserProfileURLKey:
		v.profileUrl = nil
	case UserRolesKey:
		v.roles = nil
	case UserSchemasKey:
		v.schemas = nil
	case UserTimezoneKey:
		v.timezone = nil
	case UserTitleKey:
		v.title = nil
	case UserUserNameKey:
		v.userName = nil
	case UserUserTypeKey:
		v.userType = nil
	case UserX509CertificatesKey:
		v.x509Certificates = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *User) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 25)
	if val := v.active; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserActiveKey, Value: *val})
	}
	if val := v.addresses; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserAddressesKey, Value: val})
	}
	if val := v.displayName; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserDisplayNameKey, Value: *val})
	}
	if val := v.emails; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserEmailsKey, Value: val})
	}
	if val := v.entitlements; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserEntitlementsKey, Value: val})
	}
	if val := v.externalId; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserExternalIDKey, Value: *val})
	}
	if val := v.groups; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserGroupsKey, Value: val})
	}
	if val := v.id; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserIDKey, Value: *val})
	}
	if val := v.ims; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserIMSKey, Value: val})
	}
	if val := v.locale; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserLocaleKey, Value: *val})
	}
	if val := v.meta; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserMetaKey, Value: val})
	}
	if val := v.name; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserNameKey, Value: val})
	}
	if val := v.nickName; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserNickNameKey, Value: *val})
	}
	if val := v.password; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserPasswordKey, Value: *val})
	}
	if val := v.phoneNumbers; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserPhoneNumbersKey, Value: val})
	}
	if val := v.photos; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserPhotosKey, Value: val})
	}
	if val := v.preferredLanguage; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserPreferredLanguageKey, Value: *val})
	}
	if val := v.profileUrl; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserProfileURLKey, Value: *val})
	}
	if val := v.roles; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserRolesKey, Value: val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserSchemasKey, Value: *val})
	}
	if val := v.timezone; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserTimezoneKey, Value: *val})
	}
	if val := v.title; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserTitleKey, Value: *val})
	}
	if val := v.userName; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserUserNameKey, Value: *val})
	}
	if val := v.userType; val != nil {
		pairs = append(pairs, &fieldPair{Name: UserUserTypeKey, Value: *val})
	}
	if val := v.x509Certificates; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: UserX509CertificatesKey, Value: val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes User into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *User) MarshalJSON() ([]byte, error) {
	pairs := v.makePairs()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, pair := range pairs {
		if i > 0 {
			buf.WriteByte(',')
		}
		enc.Encode(pair.Name)
		buf.WriteByte(':')
		enc.Encode(pair.Value)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON deserializes a piece of JSON data into User.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *User) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.active = nil
	v.addresses = nil
	v.displayName = nil
	v.emails = nil
	v.entitlements = nil
	v.externalId = nil
	v.groups = nil
	v.id = nil
	v.ims = nil
	v.locale = nil
	v.meta = nil
	v.name = nil
	v.nickName = nil
	v.password = nil
	v.phoneNumbers = nil
	v.photos = nil
	v.preferredLanguage = nil
	v.profileUrl = nil
	v.roles = nil
	v.schemas = nil
	v.timezone = nil
	v.title = nil
	v.userName = nil
	v.userType = nil
	v.x509Certificates = nil

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
			case UserActiveKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserActiveKey, err)
				}
				v.active = &val
			case UserAddressesKey:
				var val []*Address
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserAddressesKey, err)
				}
				v.addresses = val
			case UserDisplayNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserDisplayNameKey, err)
				}
				v.displayName = &val
			case UserEmailsKey:
				var val []*Email
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserEmailsKey, err)
				}
				v.emails = val
			case UserEntitlementsKey:
				var val []*Entitlement
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserEntitlementsKey, err)
				}
				v.entitlements = val
			case UserExternalIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserExternalIDKey, err)
				}
				v.externalId = &val
			case UserGroupsKey:
				var val []*GroupMember
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserGroupsKey, err)
				}
				v.groups = val
			case UserIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserIDKey, err)
				}
				v.id = &val
			case UserIMSKey:
				var val []*IMS
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserIMSKey, err)
				}
				v.ims = val
			case UserLocaleKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserLocaleKey, err)
				}
				v.locale = &val
			case UserMetaKey:
				var val *Meta
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserMetaKey, err)
				}
				v.meta = val
			case UserNameKey:
				var val *Names
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserNameKey, err)
				}
				v.name = val
			case UserNickNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserNickNameKey, err)
				}
				v.nickName = &val
			case UserPasswordKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserPasswordKey, err)
				}
				v.password = &val
			case UserPhoneNumbersKey:
				var val []*PhoneNumber
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserPhoneNumbersKey, err)
				}
				v.phoneNumbers = val
			case UserPhotosKey:
				var val []*Photo
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserPhotosKey, err)
				}
				v.photos = val
			case UserPreferredLanguageKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserPreferredLanguageKey, err)
				}
				v.preferredLanguage = &val
			case UserProfileURLKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserProfileURLKey, err)
				}
				v.profileUrl = &val
			case UserRolesKey:
				var val []*Role
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserRolesKey, err)
				}
				v.roles = val
			case UserSchemasKey:
				var val schemas
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserSchemasKey, err)
				}
				v.schemas = &val
			case UserTimezoneKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserTimezoneKey, err)
				}
				v.timezone = &val
			case UserTitleKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserTitleKey, err)
				}
				v.title = &val
			case UserUserNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserUserNameKey, err)
				}
				v.userName = &val
			case UserUserTypeKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserUserTypeKey, err)
				}
				v.userType = &val
			case UserX509CertificatesKey:
				var val []*X509Certificate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UserX509CertificatesKey, err)
				}
				v.x509Certificates = val
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

type UserBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *User
}

// NewUserBuilder creates a new UserBuilder instance.
// UserBuilder is safe to be used uninitialized as well.
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) initialize() {
	b.err = nil
	b.object = &User{}
}
func (b *UserBuilder) Active(in bool) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserActiveKey, in)
	return b
}
func (b *UserBuilder) Addresses(in ...*Address) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserAddressesKey, in)
	return b
}
func (b *UserBuilder) DisplayName(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserDisplayNameKey, in)
	return b
}
func (b *UserBuilder) Emails(in ...*Email) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserEmailsKey, in)
	return b
}
func (b *UserBuilder) Entitlements(in ...*Entitlement) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserEntitlementsKey, in)
	return b
}
func (b *UserBuilder) ExternalID(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserExternalIDKey, in)
	return b
}
func (b *UserBuilder) Groups(in ...*GroupMember) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserGroupsKey, in)
	return b
}
func (b *UserBuilder) ID(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserIDKey, in)
	return b
}
func (b *UserBuilder) IMS(in ...*IMS) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserIMSKey, in)
	return b
}
func (b *UserBuilder) Locale(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserLocaleKey, in)
	return b
}
func (b *UserBuilder) Meta(in *Meta) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserMetaKey, in)
	return b
}
func (b *UserBuilder) Name(in *Names) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserNameKey, in)
	return b
}
func (b *UserBuilder) NickName(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserNickNameKey, in)
	return b
}
func (b *UserBuilder) Password(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserPasswordKey, in)
	return b
}
func (b *UserBuilder) PhoneNumbers(in ...*PhoneNumber) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserPhoneNumbersKey, in)
	return b
}
func (b *UserBuilder) Photos(in ...*Photo) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserPhotosKey, in)
	return b
}
func (b *UserBuilder) PreferredLanguage(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserPreferredLanguageKey, in)
	return b
}
func (b *UserBuilder) ProfileURL(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserProfileURLKey, in)
	return b
}
func (b *UserBuilder) Roles(in ...*Role) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserRolesKey, in)
	return b
}
func (b *UserBuilder) Schemas(in ...string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserSchemasKey, in)
	return b
}
func (b *UserBuilder) Timezone(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserTimezoneKey, in)
	return b
}
func (b *UserBuilder) Title(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserTitleKey, in)
	return b
}
func (b *UserBuilder) UserName(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserUserNameKey, in)
	return b
}
func (b *UserBuilder) UserType(in string) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserUserTypeKey, in)
	return b
}
func (b *UserBuilder) X509Certificates(in ...*X509Certificate) *UserBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(UserX509CertificatesKey, in)
	return b
}

func (b *UserBuilder) Build() (*User, error) {
	err := b.err
	if err != nil {
		return nil, err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *UserBuilder) MustBuild() *User {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *UserBuilder) Extension(uri string, value interface{}) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}
	if b.object.schemas == nil {
		b.object.schemas = &schemas{}
		b.object.schemas.Add(UserSchemaURI)
	}
	b.object.schemas.Add(uri)
	if err := b.object.Set(uri, value); err != nil {
		b.err = err
	}
	return b
}

func (v *User) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *User) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) User() *UserBuilder {
	return &UserBuilder{}
}
