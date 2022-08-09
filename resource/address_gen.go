package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

const (
	AddressCountryKey       = "country"
	AddressFormattedKey     = "formatted"
	AddressLocalityKey      = "locality"
	AddressPostalCodeKey    = "postalCode"
	AddressRegionKey        = "region"
	AddressStreetAddressKey = "streetAddress"
)

type Address struct {
	country       *string
	formatted     *string
	locality      *string
	postalCode    *string
	region        *string
	streetAddress *string
	privateParams map[string]interface{}
	mu            sync.RWMutex
}

type AddressValidator interface {
	Validate(*Address) error
}

type AddressValidateFunc func(v *Address) error

func (f AddressValidateFunc) Validate(v *Address) error {
	return f(v)
}

var DefaultAddressValidator AddressValidator = AddressValidateFunc(func(v *Address) error {
	return nil
})

func (v *Address) HasCountry() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.country != nil
}

func (v *Address) Country() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.country == nil {
		return ""
	}
	return *(v.country)
}

func (v *Address) HasFormatted() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.formatted != nil
}

func (v *Address) Formatted() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.formatted == nil {
		return ""
	}
	return *(v.formatted)
}

func (v *Address) HasLocality() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.locality != nil
}

func (v *Address) Locality() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.locality == nil {
		return ""
	}
	return *(v.locality)
}

func (v *Address) HasPostalCode() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.postalCode != nil
}

func (v *Address) PostalCode() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.postalCode == nil {
		return ""
	}
	return *(v.postalCode)
}

func (v *Address) HasRegion() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.region != nil
}

func (v *Address) Region() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.region == nil {
		return ""
	}
	return *(v.region)
}

func (v *Address) HasStreetAddress() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.streetAddress != nil
}

func (v *Address) StreetAddress() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.streetAddress == nil {
		return ""
	}
	return *(v.streetAddress)
}

func (v *Address) makePairs() []pair {
	pairs := make([]pair, 0, 6)
	if v.country != nil {
		pairs = append(pairs, pair{Key: "country", Value: *(v.country)})
	}
	if v.formatted != nil {
		pairs = append(pairs, pair{Key: "formatted", Value: *(v.formatted)})
	}
	if v.locality != nil {
		pairs = append(pairs, pair{Key: "locality", Value: *(v.locality)})
	}
	if v.postalCode != nil {
		pairs = append(pairs, pair{Key: "postalCode", Value: *(v.postalCode)})
	}
	if v.region != nil {
		pairs = append(pairs, pair{Key: "region", Value: *(v.region)})
	}
	if v.streetAddress != nil {
		pairs = append(pairs, pair{Key: "streetAddress", Value: *(v.streetAddress)})
	}
	for k, v := range v.privateParams {
		pairs = append(pairs, pair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return pairs
}

func (v *Address) MarshalJSON() ([]byte, error) {
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

func (v *Address) Get(name string, options ...GetOption) (interface{}, bool) {
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
	case AddressCountryKey:
		if v.country == nil {
			return nil, false
		}
		return *(v.country), true
	case AddressFormattedKey:
		if v.formatted == nil {
			return nil, false
		}
		return *(v.formatted), true
	case AddressLocalityKey:
		if v.locality == nil {
			return nil, false
		}
		return *(v.locality), true
	case AddressPostalCodeKey:
		if v.postalCode == nil {
			return nil, false
		}
		return *(v.postalCode), true
	case AddressRegionKey:
		if v.region == nil {
			return nil, false
		}
		return *(v.region), true
	case AddressStreetAddressKey:
		if v.streetAddress == nil {
			return nil, false
		}
		return *(v.streetAddress), true
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

func (v *Address) Set(name string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch name {
	case AddressCountryKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "country", but got %T`, value)
		}
		v.country = &tmp
		return nil
	case AddressFormattedKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "formatted", but got %T`, value)
		}
		v.formatted = &tmp
		return nil
	case AddressLocalityKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "locality", but got %T`, value)
		}
		v.locality = &tmp
		return nil
	case AddressPostalCodeKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "postalCode", but got %T`, value)
		}
		v.postalCode = &tmp
		return nil
	case AddressRegionKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "region", but got %T`, value)
		}
		v.region = &tmp
		return nil
	case AddressStreetAddressKey:
		var tmp string
		tmp, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected string for field "streetAddress", but got %T`, value)
		}
		v.streetAddress = &tmp
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

func (v *Address) Clone() *Address {
	v.mu.Lock()
	defer v.mu.Unlock()
	return &Address{
		country:       v.country,
		formatted:     v.formatted,
		locality:      v.locality,
		postalCode:    v.postalCode,
		region:        v.region,
		streetAddress: v.streetAddress,
	}
}

func (v *Address) UnmarshalJSON(data []byte) error {
	v.country = nil
	v.formatted = nil
	v.locality = nil
	v.postalCode = nil
	v.region = nil
	v.streetAddress = nil
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
			case AddressCountryKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "country": %w`, err)
				}
				v.country = &x
			case AddressFormattedKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "formatted": %w`, err)
				}
				v.formatted = &x
			case AddressLocalityKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "locality": %w`, err)
				}
				v.locality = &x
			case AddressPostalCodeKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "postalCode": %w`, err)
				}
				v.postalCode = &x
			case AddressRegionKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "region": %w`, err)
				}
				v.region = &x
			case AddressStreetAddressKey:
				var x string
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "streetAddress": %w`, err)
				}
				v.streetAddress = &x
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

func (v *Address) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Key] = pair.Value
	}
	return nil
}

// AddressBuilder creates a Address resource
type AddressBuilder struct {
	once      sync.Once
	mu        sync.Mutex
	err       error
	validator AddressValidator
	object    *Address
}

func (b *Builder) Address() *AddressBuilder {
	return NewAddressBuilder()
}

func NewAddressBuilder() *AddressBuilder {
	var b AddressBuilder
	b.init()
	return &b
}

func (b *AddressBuilder) From(in *Address) *AddressBuilder {
	b.once.Do(b.init)
	b.object = in.Clone()
	return b
}

func (b *AddressBuilder) init() {
	b.err = nil
	b.validator = nil
	b.object = &Address{}
}

func (b *AddressBuilder) Country(v string) *AddressBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("country", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AddressBuilder) Formatted(v string) *AddressBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("formatted", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AddressBuilder) Locality(v string) *AddressBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("locality", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AddressBuilder) PostalCode(v string) *AddressBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("postalCode", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AddressBuilder) Region(v string) *AddressBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("region", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AddressBuilder) StreetAddress(v string) *AddressBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	if err := b.object.Set("streetAddress", v); err != nil {
		b.err = err
	}
	return b
}

func (b *AddressBuilder) Validator(v AddressValidator) *AddressBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}
	b.validator = v
	return b
}

func (b *AddressBuilder) Build() (*Address, error) {
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
		return nil, fmt.Errorf("resource.AddressBuilder: object was not initialized")
	}
	if validator == nil {
		validator = DefaultAddressValidator
	}
	if err := validator.Validate(object); err != nil {
		return nil, err
	}
	return object, nil
}

func (b *AddressBuilder) MustBuild() *Address {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}
