package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
)

const ServiceProviderConfigSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"

func init() {
	RegisterExtension(ServiceProviderConfigSchemaURI, ServiceProviderConfig{})
}

type ServiceProviderConfig struct {
	mu                    sync.RWMutex
	authenticationSchemes []*AuthenticationScheme
	bulkSupport           *BulkSupport
	changePassword        *GenericSupport
	documentationUri      *string
	etag                  *GenericSupport
	filter                *FilterSupport
	patch                 *GenericSupport
	schemas               *schemas
	sort                  *GenericSupport
	extra                 map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	ServiceProviderConfigAuthenticationSchemesKey = "authenticationSchemes"
	ServiceProviderConfigBulkSupportKey           = "bulkSupport"
	ServiceProviderConfigChangePasswordKey        = "changePassword"
	ServiceProviderConfigDocumentationURIKey      = "documentationUri"
	ServiceProviderConfigETagKey                  = "etag"
	ServiceProviderConfigFilterKey                = "filter"
	ServiceProviderConfigPatchKey                 = "patch"
	ServiceProviderConfigSchemasKey               = "schemas"
	ServiceProviderConfigSortKey                  = "sort"
)

// Get retrieves the value associated with a key
func (v *ServiceProviderConfig) Get(key string, dst interface{}) error {
	switch key {
	case ServiceProviderConfigAuthenticationSchemesKey:
		if val := v.authenticationSchemes; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ServiceProviderConfigBulkSupportKey:
		if val := v.bulkSupport; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ServiceProviderConfigChangePasswordKey:
		if val := v.changePassword; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ServiceProviderConfigDocumentationURIKey:
		if val := v.documentationUri; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ServiceProviderConfigETagKey:
		if val := v.etag; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ServiceProviderConfigFilterKey:
		if val := v.filter; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ServiceProviderConfigPatchKey:
		if val := v.patch; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case ServiceProviderConfigSchemasKey:
		if val := v.schemas; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ServiceProviderConfigSortKey:
		if val := v.sort; val != nil {
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
func (v *ServiceProviderConfig) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case ServiceProviderConfigAuthenticationSchemesKey:
		converted, ok := value.([]*AuthenticationScheme)
		if !ok {
			return fmt.Errorf(`expected value of type []*AuthenticationScheme for field authenticationSchemes, got %T`, value)
		}
		v.authenticationSchemes = converted
	case ServiceProviderConfigBulkSupportKey:
		converted, ok := value.(*BulkSupport)
		if !ok {
			return fmt.Errorf(`expected value of type *BulkSupport for field bulkSupport, got %T`, value)
		}
		v.bulkSupport = converted
	case ServiceProviderConfigChangePasswordKey:
		converted, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected value of type *GenericSupport for field changePassword, got %T`, value)
		}
		v.changePassword = converted
	case ServiceProviderConfigDocumentationURIKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field documentationUri, got %T`, value)
		}
		v.documentationUri = &converted
	case ServiceProviderConfigETagKey:
		converted, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected value of type *GenericSupport for field etag, got %T`, value)
		}
		v.etag = converted
	case ServiceProviderConfigFilterKey:
		converted, ok := value.(*FilterSupport)
		if !ok {
			return fmt.Errorf(`expected value of type *FilterSupport for field filter, got %T`, value)
		}
		v.filter = converted
	case ServiceProviderConfigPatchKey:
		converted, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected value of type *GenericSupport for field patch, got %T`, value)
		}
		v.patch = converted
	case ServiceProviderConfigSchemasKey:
		converted, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected value of type schemas for field schemas, got %T`, value)
		}
		v.schemas = &converted
	case ServiceProviderConfigSortKey:
		converted, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected value of type *GenericSupport for field sort, got %T`, value)
		}
		v.sort = converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *ServiceProviderConfig) HasAuthenticationSchemes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.authenticationSchemes != nil
}

func (v *ServiceProviderConfig) HasBulkSupport() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.bulkSupport != nil
}

func (v *ServiceProviderConfig) HasChangePassword() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.changePassword != nil
}

func (v *ServiceProviderConfig) HasDocumentationURI() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.documentationUri != nil
}

func (v *ServiceProviderConfig) HasETag() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.etag != nil
}

func (v *ServiceProviderConfig) HasFilter() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.filter != nil
}

func (v *ServiceProviderConfig) HasPatch() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.patch != nil
}

func (v *ServiceProviderConfig) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas != nil
}

func (v *ServiceProviderConfig) HasSort() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sort != nil
}

func (v *ServiceProviderConfig) AuthenticationSchemes() []*AuthenticationScheme {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.authenticationSchemes; val != nil {
		return val
	}
	return nil
}

func (v *ServiceProviderConfig) BulkSupport() *BulkSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.bulkSupport; val != nil {
		return val
	}
	return nil
}

func (v *ServiceProviderConfig) ChangePassword() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.changePassword; val != nil {
		return val
	}
	return nil
}

func (v *ServiceProviderConfig) DocumentationURI() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.documentationUri; val != nil {
		return *val
	}
	return ""
}

func (v *ServiceProviderConfig) ETag() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.etag; val != nil {
		return val
	}
	return nil
}

func (v *ServiceProviderConfig) Filter() *FilterSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.filter; val != nil {
		return val
	}
	return nil
}

func (v *ServiceProviderConfig) Patch() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.patch; val != nil {
		return val
	}
	return nil
}

func (v *ServiceProviderConfig) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.schemas; val != nil {
		return val.Get()
	}
	return nil
}

func (v *ServiceProviderConfig) Sort() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.sort; val != nil {
		return val
	}
	return nil
}

// Remove removes the value associated with a key
func (v *ServiceProviderConfig) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case ServiceProviderConfigAuthenticationSchemesKey:
		v.authenticationSchemes = nil
	case ServiceProviderConfigBulkSupportKey:
		v.bulkSupport = nil
	case ServiceProviderConfigChangePasswordKey:
		v.changePassword = nil
	case ServiceProviderConfigDocumentationURIKey:
		v.documentationUri = nil
	case ServiceProviderConfigETagKey:
		v.etag = nil
	case ServiceProviderConfigFilterKey:
		v.filter = nil
	case ServiceProviderConfigPatchKey:
		v.patch = nil
	case ServiceProviderConfigSchemasKey:
		v.schemas = nil
	case ServiceProviderConfigSortKey:
		v.sort = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *ServiceProviderConfig) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 9)
	if val := v.authenticationSchemes; len(val) > 0 {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigAuthenticationSchemesKey, Value: val})
	}
	if val := v.bulkSupport; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigBulkSupportKey, Value: val})
	}
	if val := v.changePassword; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigChangePasswordKey, Value: val})
	}
	if val := v.documentationUri; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigDocumentationURIKey, Value: *val})
	}
	if val := v.etag; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigETagKey, Value: val})
	}
	if val := v.filter; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigFilterKey, Value: val})
	}
	if val := v.patch; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigPatchKey, Value: val})
	}
	if val := v.schemas; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigSchemasKey, Value: *val})
	}
	if val := v.sort; val != nil {
		pairs = append(pairs, &fieldPair{Name: ServiceProviderConfigSortKey, Value: val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

// MarshalJSON serializes ServiceProviderConfig into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *ServiceProviderConfig) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into ServiceProviderConfig.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *ServiceProviderConfig) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.authenticationSchemes = nil
	v.bulkSupport = nil
	v.changePassword = nil
	v.documentationUri = nil
	v.etag = nil
	v.filter = nil
	v.patch = nil
	v.schemas = nil
	v.sort = nil

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
			case ServiceProviderConfigAuthenticationSchemesKey:
				var val []*AuthenticationScheme
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigAuthenticationSchemesKey, err)
				}
				v.authenticationSchemes = val
			case ServiceProviderConfigBulkSupportKey:
				var val *BulkSupport
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigBulkSupportKey, err)
				}
				v.bulkSupport = val
			case ServiceProviderConfigChangePasswordKey:
				var val *GenericSupport
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigChangePasswordKey, err)
				}
				v.changePassword = val
			case ServiceProviderConfigDocumentationURIKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigDocumentationURIKey, err)
				}
				v.documentationUri = &val
			case ServiceProviderConfigETagKey:
				var val *GenericSupport
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigETagKey, err)
				}
				v.etag = val
			case ServiceProviderConfigFilterKey:
				var val *FilterSupport
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigFilterKey, err)
				}
				v.filter = val
			case ServiceProviderConfigPatchKey:
				var val *GenericSupport
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigPatchKey, err)
				}
				v.patch = val
			case ServiceProviderConfigSchemasKey:
				var val schemas
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigSchemasKey, err)
				}
				v.schemas = &val
			case ServiceProviderConfigSortKey:
				var val *GenericSupport
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ServiceProviderConfigSortKey, err)
				}
				v.sort = val
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

type ServiceProviderConfigBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *ServiceProviderConfig
}

// NewServiceProviderConfigBuilder creates a new ServiceProviderConfigBuilder instance.
// ServiceProviderConfigBuilder is safe to be used uninitialized as well.
func NewServiceProviderConfigBuilder() *ServiceProviderConfigBuilder {
	return &ServiceProviderConfigBuilder{}
}

func (b *ServiceProviderConfigBuilder) initialize() {
	b.err = nil
	b.object = &ServiceProviderConfig{}
}
func (b *ServiceProviderConfigBuilder) AuthenticationSchemes(in ...*AuthenticationScheme) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigAuthenticationSchemesKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) BulkSupport(in *BulkSupport) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigBulkSupportKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) ChangePassword(in *GenericSupport) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigChangePasswordKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) DocumentationURI(in string) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigDocumentationURIKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) ETag(in *GenericSupport) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigETagKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) Filter(in *FilterSupport) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigFilterKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) Patch(in *GenericSupport) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigPatchKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) Schemas(in ...string) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigSchemasKey, in)
	return b
}
func (b *ServiceProviderConfigBuilder) Sort(in *GenericSupport) *ServiceProviderConfigBuilder {
	b.once.Do(b.initialize)
	_ = b.object.Set(ServiceProviderConfigSortKey, in)
	return b
}

func (b *ServiceProviderConfigBuilder) Build() (*ServiceProviderConfig, error) {
	err := b.err
	if err != nil {
		return nil, err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *ServiceProviderConfigBuilder) MustBuild() *ServiceProviderConfig {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (v *ServiceProviderConfig) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *ServiceProviderConfig) GetExtension(name, uri string, dst interface{}) error {
	if uri == "" {
		return v.Get(name, dst)
	}
	var ext interface{}
	if err := v.Get(uri, ext); err != nil {
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

func (b *Builder) ServiceProviderConfig() *ServiceProviderConfigBuilder {
	return &ServiceProviderConfigBuilder{}
}
