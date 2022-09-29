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
	Register("FilterSupport", "", FilterSupport{})
	RegisterBuilder("FilterSupport", "", FilterSupportBuilder{})
}

type FilterSupport struct {
	mu         sync.RWMutex
	maxResults *int
	supported  *bool
	extra      map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	FilterSupportMaxResultsKey = "maxResults"
	FilterSupportSupportedKey  = "supported"
)

// Get retrieves the value associated with a key
func (v *FilterSupport) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	switch key {
	case FilterSupportMaxResultsKey:
		if val := v.maxResults; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case FilterSupportSupportedKey:
		if val := v.supported; val != nil {
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
func (v *FilterSupport) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case FilterSupportMaxResultsKey:
		converted, ok := value.(int)
		if !ok {
			return fmt.Errorf(`expected value of type int for field maxResults, got %T`, value)
		}
		v.maxResults = &converted
	case FilterSupportSupportedKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field supported, got %T`, value)
		}
		v.supported = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}
func (v *FilterSupport) HasMaxResults() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.maxResults != nil
}

func (v *FilterSupport) HasSupported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.supported != nil
}

func (v *FilterSupport) MaxResults() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.maxResults; val != nil {
		return *val
	}
	return 0
}

func (v *FilterSupport) Supported() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.supported; val != nil {
		return *val
	}
	return false
}

// Remove removes the value associated with a key
func (v *FilterSupport) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case FilterSupportMaxResultsKey:
		v.maxResults = nil
	case FilterSupportSupportedKey:
		v.supported = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *FilterSupport) makePairs() []*fieldPair {
	pairs := make([]*fieldPair, 0, 2)
	if val := v.maxResults; val != nil {
		pairs = append(pairs, &fieldPair{Name: FilterSupportMaxResultsKey, Value: *val})
	}
	if val := v.supported; val != nil {
		pairs = append(pairs, &fieldPair{Name: FilterSupportSupportedKey, Value: *val})
	}

	for key, val := range v.extra {
		pairs = append(pairs, &fieldPair{Name: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})
	return pairs
}

func (v *FilterSupport) Clone() *FilterSupport {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return &FilterSupport{
		maxResults: v.maxResults,
		supported:  v.supported,
	}
}

// MarshalJSON serializes FilterSupport into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *FilterSupport) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into FilterSupport.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *FilterSupport) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.maxResults = nil
	v.supported = nil

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
			case FilterSupportMaxResultsKey:
				var val int
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, FilterSupportMaxResultsKey, err)
				}
				v.maxResults = &val
			case FilterSupportSupportedKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, FilterSupportSupportedKey, err)
				}
				v.supported = &val
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

type FilterSupportBuilder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *FilterSupport
}

// NewFilterSupportBuilder creates a new FilterSupportBuilder instance.
// FilterSupportBuilder is safe to be used uninitialized as well.
func NewFilterSupportBuilder() *FilterSupportBuilder {
	return &FilterSupportBuilder{}
}

func (b *FilterSupportBuilder) initialize() {
	b.err = nil
	b.object = &FilterSupport{}
}
func (b *FilterSupportBuilder) MaxResults(in int) *FilterSupportBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(FilterSupportMaxResultsKey, in); err != nil {
		b.err = err
	}
	return b
}
func (b *FilterSupportBuilder) Supported(in bool) *FilterSupportBuilder {
	b.once.Do(b.initialize)
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.err != nil {
		return b
	}

	if err := b.object.Set(FilterSupportSupportedKey, in); err != nil {
		b.err = err
	}
	return b
}

func (b *FilterSupportBuilder) Build() (*FilterSupport, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err := b.err; err != nil {
		return nil, err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}

func (b *FilterSupportBuilder) MustBuild() *FilterSupport {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

func (b *FilterSupportBuilder) From(in *FilterSupport) *FilterSupportBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.initialize)
	b.object = in.Clone()
	return b
}

func (v *FilterSupport) AsMap(dst map[string]interface{}) error {
	for _, pair := range v.makePairs() {
		dst[pair.Name] = pair.Value
	}
	return nil
}

// GetExtension takes into account extension uri, and fetches
// the specified attribute from the extension object
func (v *FilterSupport) GetExtension(name, uri string, dst interface{}) error {
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

func (b *Builder) FilterSupport() *FilterSupportBuilder {
	return &FilterSupportBuilder{}
}
