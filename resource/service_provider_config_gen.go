package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	ServiceProviderConfigAuthenticationSchemesKey = "authenticationSchemes"
	ServiceProviderConfigBulkKey                  = "bulk"
	ServiceProviderConfigChangePasswordKey        = "changePassword"
	ServiceProviderConfigDocumentationURIKey      = "documentationUri"
	ServiceProviderConfigEtagKey                  = "etag"
	ServiceProviderConfigFilterKey                = "filter"
	ServiceProviderConfigPatchKey                 = "patch"
	ServiceProviderConfigSchemasKey               = "schemas"
	ServiceProviderConfigSortKey                  = "sort"
)

const ServiceProviderConfigSchemaURI = "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"

func init() {
	RegisterExtension(ServiceProviderConfigSchemaURI, ServiceProviderConfig{})
}

type ServiceProviderConfig struct {
	authenticationSchemes []AuthenticationScheme
	bulk                  *BulkSupport
	changePassword        *GenericSupport
	documentationURI      *string
	etag                  *GenericSupport
	filter                *FilterSupport
	patch                 *GenericSupport
	schemas               schemas
	sort                  *GenericSupport
	privateParams         map[string]interface{}
	mu                    sync.RWMutex
}

type ServiceProviderConfigValidator interface {
	Validate(*ServiceProviderConfig) error
}

type ServiceProviderConfigValidateFunc func(v *ServiceProviderConfig) error

func (f ServiceProviderConfigValidateFunc) Validate(v *ServiceProviderConfig) error {
	return f(v)
}

var DefaultServiceProviderConfigValidator ServiceProviderConfigValidator = ServiceProviderConfigValidateFunc(func(v *ServiceProviderConfig) error {
	if v.authenticationSchemes == nil {
		return fmt.Errorf(`required field "authenticationSchemes" is missing in "ServiceProviderConfig"`)
	}
	if v.bulk == nil {
		return fmt.Errorf(`required field "bulk" is missing in "ServiceProviderConfig"`)
	}
	if v.changePassword == nil {
		return fmt.Errorf(`required field "changePassword" is missing in "ServiceProviderConfig"`)
	}
	if v.etag == nil {
		return fmt.Errorf(`required field "etag" is missing in "ServiceProviderConfig"`)
	}
	if v.filter == nil {
		return fmt.Errorf(`required field "filter" is missing in "ServiceProviderConfig"`)
	}
	if v.patch == nil {
		return fmt.Errorf(`required field "patch" is missing in "ServiceProviderConfig"`)
	}
	if v.sort == nil {
		return fmt.Errorf(`required field "sort" is missing in "ServiceProviderConfig"`)
	}
	return nil
})

func (v *ServiceProviderConfig) HasAuthenticationSchemes() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.authenticationSchemes != nil
}

func (v *ServiceProviderConfig) AuthenticationSchemes() []AuthenticationScheme {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.authenticationSchemes
}

func (v *ServiceProviderConfig) HasBulk() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.bulk != nil
}

func (v *ServiceProviderConfig) Bulk() *BulkSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.bulk
}

func (v *ServiceProviderConfig) HasChangePassword() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.changePassword != nil
}

func (v *ServiceProviderConfig) ChangePassword() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.changePassword
}

func (v *ServiceProviderConfig) HasDocumentationURI() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.documentationURI != nil
}

func (v *ServiceProviderConfig) DocumentationURI() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.documentationURI == nil {
		return ""
	}
	return *(v.documentationURI)
}

func (v *ServiceProviderConfig) HasEtag() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.etag != nil
}

func (v *ServiceProviderConfig) Etag() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.etag
}

func (v *ServiceProviderConfig) HasFilter() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.filter != nil
}

func (v *ServiceProviderConfig) Filter() *FilterSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.filter
}

func (v *ServiceProviderConfig) HasPatch() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.patch != nil
}

func (v *ServiceProviderConfig) Patch() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.patch
}

func (v *ServiceProviderConfig) HasSchemas() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return true
}

func (v *ServiceProviderConfig) Schemas() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.schemas.List()
}

func (v *ServiceProviderConfig) HasSort() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sort != nil
}

func (v *ServiceProviderConfig) Sort() *GenericSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.sort
}

func (v *ServiceProviderConfig) makePairs() []pair {
	pairs := make([]pair, 0, 9)
	if v.authenticationSchemes != nil {
		pairs = append(pairs, pair{Key: "authenticationSchemes", Value: v.authenticationSchemes})
	}
	if v.bulk != nil {
		pairs = append(pairs, pair{Key: "bulk", Value: v.bulk})
	}
	if v.changePassword != nil {
		pairs = append(pairs, pair{Key: "changePassword", Value: v.changePassword})
	}
	if v.documentationURI != nil {
		pairs = append(pairs, pair{Key: "documentationUri", Value: *(v.documentationURI)})
	}
	if v.etag != nil {
		pairs = append(pairs, pair{Key: "etag", Value: v.etag})
	}
	if v.filter != nil {
		pairs = append(pairs, pair{Key: "filter", Value: v.filter})
	}
	if v.patch != nil {
		pairs = append(pairs, pair{Key: "patch", Value: v.patch})
	}
	if v.schemas != nil {
		pairs = append(pairs, pair{Key: "schemas", Value: v.schemas})
	}
	if v.sort != nil {
		pairs = append(pairs, pair{Key: "sort", Value: v.sort})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *ServiceProviderConfig) MarshalJSON() ([]byte, error) {
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

func (v *ServiceProviderConfig) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case ServiceProviderConfigAuthenticationSchemesKey:
		if v.authenticationSchemes == nil {
			return nil, false
		}
		return v.authenticationSchemes, true
	case ServiceProviderConfigBulkKey:
		if v.bulk == nil {
			return nil, false
		}
		return v.bulk, true
	case ServiceProviderConfigChangePasswordKey:
		if v.changePassword == nil {
			return nil, false
		}
		return v.changePassword, true
	case ServiceProviderConfigDocumentationURIKey:
		if v.documentationURI == nil {
			return nil, false
		}
		return *(v.documentationURI), true
	case ServiceProviderConfigEtagKey:
		if v.etag == nil {
			return nil, false
		}
		return v.etag, true
	case ServiceProviderConfigFilterKey:
		if v.filter == nil {
			return nil, false
		}
		return v.filter, true
	case ServiceProviderConfigPatchKey:
		if v.patch == nil {
			return nil, false
		}
		return v.patch, true
	case ServiceProviderConfigSchemasKey:
		if v.schemas == nil {
			return nil, false
		}
		return v.schemas, true
	case ServiceProviderConfigSortKey:
		if v.sort == nil {
			return nil, false
		}
		return v.sort, true
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

func (v *ServiceProviderConfig) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case ServiceProviderConfigAuthenticationSchemesKey:
		var tmp []AuthenticationScheme
		tmp, ok := value.([]AuthenticationScheme)
		if !ok {
			return fmt.Errorf(`expected []AuthenticationScheme for field "authenticationSchemes", but got %T`, value)
		}
		v.authenticationSchemes = tmp
		return nil
	case ServiceProviderConfigBulkKey:
		var tmp *BulkSupport
		tmp, ok := value.(*BulkSupport)
		if !ok {
			return fmt.Errorf(`expected *BulkSupport for field "bulk", but got %T`, value)
		}
		v.bulk = tmp
		return nil
	case ServiceProviderConfigChangePasswordKey:
		var tmp *GenericSupport
		tmp, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected *GenericSupport for field "changePassword", but got %T`, value)
		}
		v.changePassword = tmp
		return nil
	case ServiceProviderConfigDocumentationURIKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "documentationUri", but got %T`, value)
		}
		v.documentationURI = &tmp
		return nil
	case ServiceProviderConfigEtagKey:
		var tmp *GenericSupport
		tmp, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected *GenericSupport for field "etag", but got %T`, value)
		}
		v.etag = tmp
		return nil
	case ServiceProviderConfigFilterKey:
		var tmp *FilterSupport
		tmp, ok := value.(*FilterSupport)
		if !ok {
			return fmt.Errorf(`expected *FilterSupport for field "filter", but got %T`, value)
		}
		v.filter = tmp
		return nil
	case ServiceProviderConfigPatchKey:
		var tmp *GenericSupport
		tmp, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected *GenericSupport for field "patch", but got %T`, value)
		}
		v.patch = tmp
		return nil
	case ServiceProviderConfigSchemasKey:
		var tmp schemas
		tmp, ok := value.(schemas)
		if !ok {
			return fmt.Errorf(`expected schemas for field "schemas", but got %T`, value)
		}
		v.schemas = tmp
		return nil
	case ServiceProviderConfigSortKey:
		var tmp *GenericSupport
		tmp, ok := value.(*GenericSupport)
		if !ok {
			return fmt.Errorf(`expected *GenericSupport for field "sort", but got %T`, value)
		}
		v.sort = tmp
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

func (v *ServiceProviderConfig) Clone() *ServiceProviderConfig {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &ServiceProviderConfig{
		authenticationSchemes: v.authenticationSchemes,
		bulk:                  v.bulk,
		changePassword:        v.changePassword,
		documentationURI:      v.documentationURI,
		etag:                  v.etag,
		filter:                v.filter,
		patch:                 v.patch,
		schemas:               v.schemas,
		sort:                  v.sort,
	}
}

func (v *ServiceProviderConfig) UnmarshalJSON(data []byte) error {
	v.authenticationSchemes = nil
	v.bulk = nil
	v.changePassword = nil
	v.documentationURI = nil
	v.etag = nil
	v.filter = nil
	v.patch = nil
	v.schemas = nil
	v.sort = nil
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
			case ServiceProviderConfigAuthenticationSchemesKey:
				var x []AuthenticationScheme
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "authenticationSchemes": %w`, err)
				}
				v.authenticationSchemes = x
			case ServiceProviderConfigBulkKey:
				var x *BulkSupport
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "bulk": %w`, err)
				}
				v.bulk = x
			case ServiceProviderConfigChangePasswordKey:
				var x *GenericSupport
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "changePassword": %w`, err)
				}
				v.changePassword = x
			case ServiceProviderConfigDocumentationURIKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "documentationUri": %w`, err)
				}
				v.documentationURI = &x
			case ServiceProviderConfigEtagKey:
				var x *GenericSupport
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "etag": %w`, err)
				}
				v.etag = x
			case ServiceProviderConfigFilterKey:
				var x *FilterSupport
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "filter": %w`, err)
				}
				v.filter = x
			case ServiceProviderConfigPatchKey:
				var x *GenericSupport
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "patch": %w`, err)
				}
				v.patch = x
			case ServiceProviderConfigSchemasKey:
				var x schemas
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = x
			case ServiceProviderConfigSortKey:
				var x *GenericSupport
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "sort": %w`, err)
				}
				v.sort = x
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

func (v *ServiceProviderConfig) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

type ServiceProviderConfigBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator ServiceProviderConfigValidator
	object    *ServiceProviderConfig
}

func (b *Builder) ServiceProviderConfig() *ServiceProviderConfigBuilder {
	return NewServiceProviderConfigBuilder()
}

func NewServiceProviderConfigBuilder() *ServiceProviderConfigBuilder {
	var b ServiceProviderConfigBuilder
	b.init()
	return &b
}

func (b *ServiceProviderConfigBuilder) From(in *ServiceProviderConfig) *ServiceProviderConfigBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *ServiceProviderConfigBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &ServiceProviderConfig{}
}

func (b *ServiceProviderConfigBuilder) AuthenticationSchemes(v ...AuthenticationScheme) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("authenticationSchemes", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) Bulk(v *BulkSupport) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("bulk", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) ChangePassword(v *GenericSupport) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("changePassword", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) DocumentationURI(v string) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("documentationUri", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) Etag(v *GenericSupport) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("etag", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) Filter(v *FilterSupport) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("filter", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) Patch(v *GenericSupport) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("patch", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) Schemas(v ...string) *ServiceProviderConfigBuilder {
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

func (b *ServiceProviderConfigBuilder) Sort(v *GenericSupport) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("sort", v); err != nil {
		b.err = err
	}
	return b
}

func (b *ServiceProviderConfigBuilder) Extension(uri string, value interface{}) *ServiceProviderConfigBuilder {
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

func (b *ServiceProviderConfigBuilder) Validator(v ServiceProviderConfigValidator) *ServiceProviderConfigBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *ServiceProviderConfigBuilder) Build() (*ServiceProviderConfig, error) {
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
		return nil, fmt.Errorf("resource.ServiceProviderConfigBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultServiceProviderConfigValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *ServiceProviderConfigBuilder) MustBuild() *ServiceProviderConfig {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
