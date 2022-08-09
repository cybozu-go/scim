package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

// JSON key names for User resource
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

const UserSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:User"

func init() {
	RegisterExtension(UserSchemaURI, User{})
}

// User represents a user SCIM resource.
type User struct {
	active            *bool
	addresses         []*Address
	displayName       *string
	emails            []*Email
	entitlements      []*Entitlement
	externalID        *string
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
	profileURL        *string
	roles             []*Role
	schemas           schemas
	timezone          *string
	title             *string
	userName          *string
	userType          *string
	x509Certificates  []*X509Certificate
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

var DefaultUserValidator UserValidator = UserValidateFunc(func(v *User) error {
	if v.userName == nil {
		return fmt.Errorf(`required field "userName" is missing in "User"`)
	}
	return nil
})

func (v *User) HasActive() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.active != nil
}

func (v *User) Active() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.active == nil {
		return false
	}
	return *(v.active)
}

func (v *User) HasAddresses() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.addresses != nil
}

func (v *User) Addresses() []*Address {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.addresses
}

func (v *User) HasDisplayName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.displayName != nil
}

func (v *User) DisplayName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.displayName == nil {
		return ""
	}
	return *(v.displayName)
}

func (v *User) HasEmails() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.emails != nil
}

func (v *User) Emails() []*Email {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.emails
}

func (v *User) HasEntitlements() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.entitlements != nil
}

func (v *User) Entitlements() []*Entitlement {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.entitlements
}

func (v *User) HasExternalID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.externalID != nil
}

func (v *User) ExternalID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.externalID == nil {
		return ""
	}
	return *(v.externalID)
}

func (v *User) HasGroups() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.groups != nil
}

func (v *User) Groups() []*GroupMember {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.groups
}

func (v *User) HasID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.id != nil
}

func (v *User) ID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.id == nil {
		return ""
	}
	return *(v.id)
}

func (v *User) HasIMS() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ims != nil
}

func (v *User) IMS() []*IMS {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.ims
}

func (v *User) HasLocale() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.locale != nil
}

func (v *User) Locale() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.locale == nil {
		return ""
	}
	return *(v.locale)
}

func (v *User) HasMeta() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta != nil
}

func (v *User) Meta() *Meta {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.meta
}

func (v *User) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

func (v *User) Name() *Names {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name
}

func (v *User) HasNickName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.nickName != nil
}

func (v *User) NickName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.nickName == nil {
		return ""
	}
	return *(v.nickName)
}

func (v *User) HasPassword() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.password != nil
}

func (v *User) Password() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.password == nil {
		return ""
	}
	return *(v.password)
}

func (v *User) HasPhoneNumbers() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.phoneNumbers != nil
}

func (v *User) PhoneNumbers() []*PhoneNumber {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.phoneNumbers
}

func (v *User) HasPhotos() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.photos != nil
}

func (v *User) Photos() []*Photo {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.photos
}

func (v *User) HasPreferredLanguage() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.preferredLanguage != nil
}

func (v *User) PreferredLanguage() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.preferredLanguage == nil {
		return ""
	}
	return *(v.preferredLanguage)
}

func (v *User) HasProfileURL() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.profileURL != nil
}

func (v *User) ProfileURL() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.profileURL == nil {
		return ""
	}
	return *(v.profileURL)
}

func (v *User) HasRoles() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.roles != nil
}

func (v *User) Roles() []*Role {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.roles
}

func (v *User) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *User) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *User) HasTimezone() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.timezone != nil
}

func (v *User) Timezone() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.timezone == nil {
		return ""
	}
	return *(v.timezone)
}

func (v *User) HasTitle() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.title != nil
}

func (v *User) Title() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.title == nil {
		return ""
	}
	return *(v.title)
}

func (v *User) HasUserName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.userName != nil
}

func (v *User) UserName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.userName == nil {
		return ""
	}
	return *(v.userName)
}

func (v *User) HasUserType() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.userType != nil
}

func (v *User) UserType() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.userType == nil {
		return ""
	}
	return *(v.userType)
}

func (v *User) HasX509Certificates() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509Certificates != nil
}

func (v *User) X509Certificates() []*X509Certificate {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509Certificates
}

func (v *User) makePairs() []pair {
	pairs := make([]pair, 0, 25)
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
	if v.locale != nil {
		pairs = append(pairs, pair{Key: "locale", Value: *(v.locale)})
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
	if v.photos != nil {
		pairs = append(pairs, pair{Key: "photos", Value: v.photos})
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
	case UserActiveKey:
		if v.active == nil {
			return nil, false
		}
		return *(v.active), true
	case UserAddressesKey:
		if v.addresses == nil {
			return nil, false
		}
		return v.addresses, true
	case UserDisplayNameKey:
		if v.displayName == nil {
			return nil, false
		}
		return *(v.displayName), true
	case UserEmailsKey:
		if v.emails == nil {
			return nil, false
		}
		return v.emails, true
	case UserEntitlementsKey:
		if v.entitlements == nil {
			return nil, false
		}
		return v.entitlements, true
	case UserExternalIDKey:
		if v.externalID == nil {
			return nil, false
		}
		return *(v.externalID), true
	case UserGroupsKey:
		if v.groups == nil {
			return nil, false
		}
		return v.groups, true
	case UserIDKey:
		if v.id == nil {
			return nil, false
		}
		return *(v.id), true
	case UserIMSKey:
		if v.ims == nil {
			return nil, false
		}
		return v.ims, true
	case UserLocaleKey:
		if v.locale == nil {
			return nil, false
		}
		return *(v.locale), true
	case UserMetaKey:
		if v.meta == nil {
			return nil, false
		}
		return v.meta, true
	case UserNameKey:
		if v.name == nil {
			return nil, false
		}
		return v.name, true
	case UserNickNameKey:
		if v.nickName == nil {
			return nil, false
		}
		return *(v.nickName), true
	case UserPasswordKey:
		if v.password == nil {
			return nil, false
		}
		return *(v.password), true
	case UserPhoneNumbersKey:
		if v.phoneNumbers == nil {
			return nil, false
		}
		return v.phoneNumbers, true
	case UserPhotosKey:
		if v.photos == nil {
			return nil, false
		}
		return v.photos, true
	case UserPreferredLanguageKey:
		if v.preferredLanguage == nil {
			return nil, false
		}
		return *(v.preferredLanguage), true
	case UserProfileURLKey:
		if v.profileURL == nil {
			return nil, false
		}
		return *(v.profileURL), true
	case UserRolesKey:
		if v.roles == nil {
			return nil, false
		}
		return v.roles, true
	case UserSchemasKey:
		if v.schemas == nil {
			return nil, false
		}
		return v.schemas, true
	case UserTimezoneKey:
		if v.timezone == nil {
			return nil, false
		}
		return *(v.timezone), true
	case UserTitleKey:
		if v.title == nil {
			return nil, false
		}
		return *(v.title), true
	case UserUserNameKey:
		if v.userName == nil {
			return nil, false
		}
		return *(v.userName), true
	case UserUserTypeKey:
		if v.userType == nil {
			return nil, false
		}
		return *(v.userType), true
	case UserX509CertificatesKey:
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
	case UserActiveKey:
		var tmp bool
		tmp, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected bool for field "active", but got %T`, value)
		}
		v.active = &tmp
		return nil
	case UserAddressesKey:
		var tmp []*Address
		tmp, ok := value.([]*Address)
		if !ok {
			return fmt.Errorf(`expected []*Address for field "addresses", but got %T`, value)
		}
		v.addresses = tmp
		return nil
	case UserDisplayNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "displayName", but got %T`, value)
		}
		v.displayName = &tmp
		return nil
	case UserEmailsKey:
		var tmp []*Email
		tmp, ok := value.([]*Email)
		if !ok {
			return fmt.Errorf(`expected []*Email for field "emails", but got %T`, value)
		}
		v.emails = tmp
		return nil
	case UserEntitlementsKey:
		var tmp []*Entitlement
		tmp, ok := value.([]*Entitlement)
		if !ok {
			return fmt.Errorf(`expected []*Entitlement for field "entitlements", but got %T`, value)
		}
		v.entitlements = tmp
		return nil
	case UserExternalIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "externalId", but got %T`, value)
		}
		v.externalID = &tmp
		return nil
	case UserGroupsKey:
		var tmp []*GroupMember
		tmp, ok := value.([]*GroupMember)
		if !ok {
			return fmt.Errorf(`expected []*GroupMember for field "groups", but got %T`, value)
		}
		v.groups = tmp
		return nil
	case UserIDKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "id", but got %T`, value)
		}
		v.id = &tmp
		return nil
	case UserIMSKey:
		var tmp []*IMS
		tmp, ok := value.([]*IMS)
		if !ok {
			return fmt.Errorf(`expected []*IMS for field "ims", but got %T`, value)
		}
		v.ims = tmp
		return nil
	case UserLocaleKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "locale", but got %T`, value)
		}
		v.locale = &tmp
		return nil
	case UserMetaKey:
		var tmp *Meta
		tmp, ok := value.(*Meta)
		if !ok {
			return fmt.Errorf(`expected *Meta for field "meta", but got %T`, value)
		}
		v.meta = tmp
		return nil
	case UserNameKey:
		var tmp *Names
		tmp, ok := value.(*Names)
		if !ok {
			return fmt.Errorf(`expected *Names for field "name", but got %T`, value)
		}
		v.name = tmp
		return nil
	case UserNickNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "nickName", but got %T`, value)
		}
		v.nickName = &tmp
		return nil
	case UserPasswordKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "password", but got %T`, value)
		}
		v.password = &tmp
		return nil
	case UserPhoneNumbersKey:
		var tmp []*PhoneNumber
		tmp, ok := value.([]*PhoneNumber)
		if !ok {
			return fmt.Errorf(`expected []*PhoneNumber for field "phoneNumbers", but got %T`, value)
		}
		v.phoneNumbers = tmp
		return nil
	case UserPhotosKey:
		var tmp []*Photo
		tmp, ok := value.([]*Photo)
		if !ok {
			return fmt.Errorf(`expected []*Photo for field "photos", but got %T`, value)
		}
		v.photos = tmp
		return nil
	case UserPreferredLanguageKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "preferredLanguage", but got %T`, value)
		}
		v.preferredLanguage = &tmp
		return nil
	case UserProfileURLKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "profileUrl", but got %T`, value)
		}
		v.profileURL = &tmp
		return nil
	case UserRolesKey:
		var tmp []*Role
		tmp, ok := value.([]*Role)
		if !ok {
			return fmt.Errorf(`expected []*Role for field "roles", but got %T`, value)
		}
		v.roles = tmp
		return nil
	case UserSchemasKey:
		var tmp schemas
		tmp, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected schemas for field "schemas", but got %T`, value)
		}
		v.schemas = tmp
		return nil
	case UserTimezoneKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "timezone", but got %T`, value)
		}
		v.timezone = &tmp
		return nil
	case UserTitleKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "title", but got %T`, value)
		}
		v.title = &tmp
		return nil
	case UserUserNameKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "userName", but got %T`, value)
		}
		v.userName = &tmp
		return nil
	case UserUserTypeKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "userType", but got %T`, value)
		}
		v.userType = &tmp
		return nil
	case UserX509CertificatesKey:
		var tmp []*X509Certificate
		tmp, ok := value.([]*X509Certificate)
		if !ok {
			return fmt.Errorf(`expected []*X509Certificate for field "x509Certificates", but got %T`, value)
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

func (v *User) Clone() *User {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &User{
		active:            v.active,
		addresses:         v.addresses,
		displayName:       v.displayName,
		emails:            v.emails,
		entitlements:      v.entitlements,
		externalID:        v.externalID,
		groups:            v.groups,
		id:                v.id,
		ims:               v.ims,
		locale:            v.locale,
		meta:              v.meta,
		name:              v.name,
		nickName:          v.nickName,
		password:          v.password,
		phoneNumbers:      v.phoneNumbers,
		photos:            v.photos,
		preferredLanguage: v.preferredLanguage,
		profileURL:        v.profileURL,
		roles:             v.roles,
		schemas:           v.schemas,
		timezone:          v.timezone,
		title:             v.title,
		userName:          v.userName,
		userType:          v.userType,
		x509Certificates:  v.x509Certificates,
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
	v.locale = nil
	v.meta = nil
	v.name = nil
	v.nickName = nil
	v.password = nil
	v.phoneNumbers = nil
	v.photos = nil
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
			case UserActiveKey:
				var x bool
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "active": %w`, err)
				}
				v.active = &x
			case UserAddressesKey:
				var x []*Address
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "addresses": %w`, err)
				}
				v.addresses = x
			case UserDisplayNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "displayName": %w`, err)
				}
				v.displayName = &x
			case UserEmailsKey:
				var x []*Email
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "emails": %w`, err)
				}
				v.emails = x
			case UserEntitlementsKey:
				var x []*Entitlement
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "entitlements": %w`, err)
				}
				v.entitlements = x
			case UserExternalIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "externalId": %w`, err)
				}
				v.externalID = &x
			case UserGroupsKey:
				var x []*GroupMember
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "groups": %w`, err)
				}
				v.groups = x
			case UserIDKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "id": %w`, err)
				}
				v.id = &x
			case UserIMSKey:
				var x []*IMS
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "ims": %w`, err)
				}
				v.ims = x
			case UserLocaleKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "locale": %w`, err)
				}
				v.locale = &x
			case UserMetaKey:
				var x *Meta
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "meta": %w`, err)
				}
				v.meta = x
			case UserNameKey:
				var x *Names
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "name": %w`, err)
				}
				v.name = x
			case UserNickNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "nickName": %w`, err)
				}
				v.nickName = &x
			case UserPasswordKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "password": %w`, err)
				}
				v.password = &x
			case UserPhoneNumbersKey:
				var x []*PhoneNumber
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "phoneNumbers": %w`, err)
				}
				v.phoneNumbers = x
			case UserPhotosKey:
				var x []*Photo
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "photos": %w`, err)
				}
				v.photos = x
			case UserPreferredLanguageKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "preferredLanguage": %w`, err)
				}
				v.preferredLanguage = &x
			case UserProfileURLKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "profileUrl": %w`, err)
				}
				v.profileURL = &x
			case UserRolesKey:
				var x []*Role
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "roles": %w`, err)
				}
				v.roles = x
			case UserSchemasKey:
				var x schemas
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = x
			case UserTimezoneKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "timezone": %w`, err)
				}
				v.timezone = &x
			case UserTitleKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "title": %w`, err)
				}
				v.title = &x
			case UserUserNameKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "userName": %w`, err)
				}
				v.userName = &x
			case UserUserTypeKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "userType": %w`, err)
				}
				v.userType = &x
			case UserX509CertificatesKey:
				var x []*X509Certificate
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

// UserBuilder creates a User resource
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

func (b *UserBuilder) From(in *User) *UserBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *UserBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &User{}

	b.object.schemas = make(schemas)
	b.object.schemas.Add(UserSchemaURI)
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

func (b *UserBuilder) Addresses(v ...*Address) *UserBuilder {
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

func (b *UserBuilder) Entitlements(v ...*Entitlement) *UserBuilder {
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

func (b *UserBuilder) Groups(v ...*GroupMember) *UserBuilder {
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

func (b *UserBuilder) IMS(v ...*IMS) *UserBuilder {
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

func (b *UserBuilder) Locale(v string) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("locale", v); err != nil {
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

func (b *UserBuilder) PhoneNumbers(v ...*PhoneNumber) *UserBuilder {
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

func (b *UserBuilder) Photos(v ...*Photo) *UserBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("photos", v); err != nil {
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

func (b *UserBuilder) Roles(v ...*Role) *UserBuilder {
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
	for _, schema := range v {
		b.object.schemas.Add(schema)
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

func (b *UserBuilder) X509Certificates(v ...*X509Certificate) *UserBuilder {
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
	b.object.schemas.Add(uri)
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
	if err := validator.Validate(object); err != nil {
		return nil, err
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
