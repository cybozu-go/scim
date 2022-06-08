package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/lestrrat-go/codegen"
	"github.com/lestrrat-go/xstrings"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func yaml2json(fn string) ([]byte, error) {
	in, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf(`failed to open %q: %w`, fn, err)
	}
	defer in.Close()

	var v interface{}
	if err := yaml.NewDecoder(in).Decode(&v); err != nil {
		return nil, fmt.Errorf(`failed to decode %q: %w`, fn, err)
	}

	return json.Marshal(v)
}

func _main() error {
	codegen.RegisterZeroVal(`AuthenticationSchemeType`, `InvalidAuthenticationScheme`)
	codegen.RegisterZeroVal(`Mutability`, `MutReadOnly`)
	codegen.RegisterZeroVal(`Returned`, `ReturnedAlways`)
	codegen.RegisterZeroVal(`Uniqueness`, `UniqNone`)

	var objectsFile = flag.String("objects", "objects.yml", "")
	flag.Parse()
	jsonSrc, err := yaml2json(*objectsFile)
	if err != nil {
		return err
	}

	var def struct {
		Common  codegen.FieldList
		Objects []*codegen.Object `json:"objects"`
	}
	if err := json.NewDecoder(bytes.NewReader(jsonSrc)).Decode(&def); err != nil {
		return fmt.Errorf(`failed to decode %q: %w`, *objectsFile, err)
	}

	for _, object := range def.Objects {
		// Each object needs a common set of fields.
		if !object.Bool(`skipCommonFields`) {
			for _, commonField := range def.Common {
				object.AddField(commonField)
			}
		}
		if object.String(`schema`) != "" {
			// TODO: we needed codegen.FieldBulder
			var fl codegen.FieldList
			if err := json.Unmarshal([]byte(`[{"name":"schemas","type":"schemas"}]`), &fl); err != nil {
				return fmt.Errorf(`failed to unmarshal schemas field: %w`, err)
			}
			object.AddField(fl[0])
		}

		object.Organize()

		if err := generateObject(object); err != nil {
			return fmt.Errorf(`failed to generate object: %s`, err)
		}

		if err := generateEnt(object); err != nil {
			return fmt.Errorf(`faile dto generate ent adapter: %s`, err)
		}
	}

	return nil
}

func IsSlice(f codegen.Field) bool {
	return strings.HasPrefix(f.Type(), `[]`)
}

func IsPointer(f codegen.Field) bool {
	return strings.HasPrefix(f.Type(), `*`)
}

func IsIndirect(f codegen.Field) bool {
	s := f.Type()
	return s == `schemas` || strings.HasPrefix(s, `*`) || strings.HasPrefix(s, `[]`) || strings.HasSuffix(s, `List`)
}

func generateObject(object *codegen.Object) error {
	fmt.Printf("  ⌛ Generating %s...\n", object.Name(true))

	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)

	o.L(`package resource`)

	o.LL(`const (`)
	for _, f := range object.Fields() {
		o.L(`%s%sKey = %q`, object.Name(true), f.Name(true), f.JSON())
	}
	o.L(`)`)

	if schema := object.String(`schema`); schema != "" {
		o.LL(`const %sSchemaURI = %q`, object.Name(true), schema)
		o.L(`func init() {`)
		o.L(`RegisterExtension(%[1]sSchemaURI, %[1]s{})`, object.Name(true))
		o.L(`}`)
	}

	o.LL(`type %s struct {`, object.Name(true))
	for _, field := range object.Fields() {
		if IsIndirect(field) {
			o.L("%s %s", field.Name(false), field.Type())
		} else {
			o.L("%s *%s", field.Name(false), field.Type())
		}
	}
	o.L(`privateParams map[string]interface{}`)
	o.L(`mu sync.RWMutex`)
	o.L(`}`)

	o.LL(`type %sValidator interface {`, object.Name(true))
	o.L(`Validate(*%s) error`, object.Name(true))
	o.L(`}`)

	o.LL(`type %[1]sValidateFunc func(v *%[1]s) error`, object.Name(true))
	o.LL(`func (f %[1]sValidateFunc) Validate(v *%[1]s) error {`, object.Name(true))
	o.L(`return f(v)`)
	o.L(`}`)

	o.LL(`var Default%[1]sValidator %[1]sValidator = %[1]sValidateFunc(func(v *%[1]s) error {`, object.Name(true))
	for _, field := range object.Fields() {
		if field.IsRequired() {
			o.L(`if v.%s == nil {`, field.Name(false))
			o.L("return fmt.Errorf(`required field %q is missing in %q`)", field.JSON(), object.Name(true))
			o.L(`}`)
		}
	}
	o.L(`return nil`)
	o.L(`})`)

	for _, field := range object.Fields() {
		var rt string
		if field.Name(false) == `schemas` {
			rt = `[]string`
		} else {
			rt = field.Type()
		}

		o.LL(`func (v *%s) Has%s() bool {`, object.Name(true), field.Name(true))
		o.L(`v.mu.RLock()`)
		o.L(`defer v.mu.RUnlock()`)
		// schemas is a special case
		if field.Name(false) == `schemas` {
			o.L(`return true`)
		} else {
			o.L(`return v.%s != nil`, field.Name(false))
		}
		o.L(`}`)

		o.LL(`func (v *%s) %s() %s {`, object.Name(true), field.Name(true), rt)
		o.L(`v.mu.RLock()`)
		o.L(`defer v.mu.RUnlock()`)
		// schemas is a special case
		if field.Name(false) == `schemas` {
			o.L(`return v.%s.List()`, field.Name(false))
		} else if IsIndirect(field) {
			o.L(`return v.%s`, field.Name(false))
		} else {
			o.L(`if v.%s == nil {`, field.Name(false))
			o.L(`return %s`, codegen.ZeroVal(field.Type()))
			o.L(`}`)
			o.L(`return *(v.%s)`, field.Name(false))
		}
		o.L(`}`)
	}

	o.LL(`func (v *%s) makePairs() []pair {`, object.Name(true))
	o.L(`pairs := make([]pair, 0, %d)`, len(object.Fields()))
	for _, field := range object.Fields() {
		o.L(`if v.%s != nil {`, field.Name(false))
		if IsIndirect(field) {
			o.L(`pairs = append(pairs, pair{Key: %q, Value: v.%s})`, field.JSON(), field.Name(false))
		} else {
			o.L(`pairs = append(pairs, pair{Key: %q, Value: *(v.%s)})`, field.JSON(), field.Name(false))
		}
		o.L(`}`)
	}
	o.L(`for k, v := range v.privateParams {`)
	o.L(`pairs = append(pairs, pair{Key: k, Value: v})`)
	o.L(`}`)
	o.L(`sort.Slice(pairs, func(i, j int) bool {`)
	o.L(`return pairs[i].Key < pairs[j].Key`)
	o.L(`})`)
	o.L(`return pairs`)
	o.L(`}`)

	o.LL(`func (v *%s) MarshalJSON() ([]byte, error) {`, object.Name(true))
	o.L(`pairs := v.makePairs()`)
	o.LL(`var buf bytes.Buffer`)
	o.L(`enc := json.NewEncoder(&buf)`)
	o.L(`buf.WriteByte('{')`)
	o.L(`for i, pair := range pairs {`)
	o.L(`if i > 0 {`)
	o.L(`buf.WriteRune(',')`)
	o.L(`}`)
	o.L(`fmt.Fprintf(&buf, "%%q:", pair.Key)`)
	o.L(`if err := enc.Encode(pair.Value); err != nil {`)
	o.L(`return nil, fmt.Errorf("failed to encode value for key %%q: %%w", pair.Key, err)`)
	o.L(`}`)
	o.L(`}`)
	o.L(`buf.WriteByte('}')`)
	o.L(`return buf.Bytes(), nil`)
	o.L(`}`)

	o.LL(`func (v *%s) Get(name string, options ...GetOption) (interface{}, bool) {`, object.Name(true))
	o.L(`v.mu.RLock()`)
	o.L(`defer v.mu.RUnlock()`)
	o.LL(`var ext string`)
	o.L(`//nolint:forcetypeassert`)
	o.L(`for _, option := range options {`)
	o.L(`switch option.Ident() {`)
	o.L(`case identExtension{}:`)
	o.L(`ext = option.Value().(string)`)
	o.L(`}`)
	o.L(`}`)

	o.L(`switch name {`)
	for _, field := range object.Fields() {
		o.L(`case %s%sKey:`, object.Name(true), field.Name(true))
		o.L(`if v.%s == nil {`, field.Name(false))
		o.L(`return nil, false`)
		o.L(`}`)
		if IsIndirect(field) {
			o.L(`return v.%s, true`, field.Name(false))
		} else {
			o.L(`return *(v.%s), true`, field.Name(false))
		}
	}
	o.L(`default:`)
	o.L(`pp := v.privateParams`)
	o.L(`if pp == nil {`)
	o.L(`return nil, false`)
	o.L(`}`)
	o.L(`if ext == "" {`)
	o.L(`ret, ok := pp[name]`)
	o.L(`return ret, ok`)
	o.L(`}`)
	o.L(`obj, ok := pp[ext]`)
	o.L(`if !ok {`)
	o.L(`return nil, false`)
	o.L(`}`)
	o.L(`getter, ok := obj.(interface{ Get(string, ...GetOption) (interface{}, bool) })`)
	o.L(`if !ok {`)
	o.L(`return nil, false`)
	o.L(`}`)
	o.L(`return getter.Get(name)`)
	o.L(`}`)
	o.L(`}`)

	o.LL(`func (v *%s) Set(name string, value interface{}) error {`, object.Name(true))
	o.L(`v.mu.Lock()`)
	o.L(`defer v.mu.Unlock()`)
	o.L(`switch name {`)
	for _, field := range object.Fields() {
		o.L(`case %s%sKey:`, object.Name(true), field.Name(true))
		o.L(`var tmp %s`, field.Type())
		o.L(`tmp, ok := value.(%s)`, field.Type())
		o.L(`if !ok {`)
		o.L("return fmt.Errorf(`expected %s for field %q, but got %%T`, value)", field.Type(), field.JSON())
		o.L(`}`)
		if IsIndirect(field) {
			o.L(`v.%s = tmp`, field.Name(false))
		} else {
			o.L(`v.%s = &tmp`, field.Name(false))
		}
		o.L(`return nil`)
	}
	o.L(`default:`)
	o.L(`pp := v.privateParams`)
	o.L(`if pp == nil {`)
	o.L(`pp = make(map[string]interface{})`)
	o.L(`v.privateParams = pp`)
	o.L(`}`)
	o.L(`pp[name] = value`)
	o.L(`return nil`)
	o.L(`}`)
	o.L(`}`)

	o.LL(`func (v *%[1]s) Clone() *%[1]s {`, object.Name(true))
	o.L(`v.mu.Lock()`)
	o.L(`defer v.mu.Unlock()`)
	o.L(`return &%s{`, object.Name(true))
	for _, field := range object.Fields() {
		o.L(`%[1]s: v.%[1]s,`, field.Name(false))
	}
	o.L(`}`)
	o.L(`}`)

	o.LL(`func (v *%s) UnmarshalJSON(data []byte) error {`, object.Name(true))
	for _, field := range object.Fields() {
		if IsIndirect(field) {
			o.L(`v.%s = %s`, field.Name(false), codegen.ZeroVal(field.Type()))
		} else {
			o.L(`v.%s = nil`, field.Name(false))
		}
	}
	o.L(`v.privateParams = nil`)
	o.L(`dec := json.NewDecoder(bytes.NewReader(data))`)
	o.L(`{ // first token`)
	o.L(`tok, err := dec.Token()`)
	o.L(`if err != nil {`)
	o.L(`return fmt.Errorf("failed to read next token: %%s", err)`)
	o.L(`}`)
	o.L(`tok, ok := tok.(json.Delim)`)
	o.L(`if !ok {`)
	o.L(`return fmt.Errorf("expected first token to be '{', got %%c", tok)`)
	o.L(`}`)
	o.L(`}`)
	o.L(`var privateParams map[string]interface{}`)
	o.LL(`LOOP:`)
	o.L(`for {`)
	o.L(`tok, err := dec.Token()`)
	o.L(`if err != nil {`)
	o.L(`return fmt.Errorf("failed to read next token: %%s", err)`)
	o.L(`}`)
	o.L(`switch tok := tok.(type) {`)
	o.L(`case json.Delim:`)
	o.L(`if tok == '}' {`)
	o.L(`break LOOP`)
	o.L(`} else {`)
	o.L(`return fmt.Errorf("unexpected token %%c found", tok)`)
	o.L(`}`)
	o.L(`case string:`)
	o.L(`switch tok {`)
	for _, field := range object.Fields() {
		o.L(`case %s%sKey:`, object.Name(true), field.Name(true))
		o.L(`var x %s`, field.Type())
		o.L(`if err := dec.Decode(&x); err != nil {`)
		o.L("return fmt.Errorf(`failed to decode value for key %q: %%w`, err)", field.JSON())
		o.L(`}`)
		if IsIndirect(field) {
			o.L(`v.%s = x`, field.Name(false))
		} else {
			o.L(`v.%s = &x`, field.Name(false))
		}
	}
	o.L(`default:`)
	o.L(`var x interface{}`)
	o.L(`if rx, ok := registry.Get(tok); ok {`)
	o.L(`x = rx`)
	o.L(`if err := dec.Decode(x); err != nil {`)
	o.L("return fmt.Errorf(`failed to decode value for key %%q: %%w`, tok, err)")
	o.L(`}`)
	o.L(`} else {`)
	o.L(`if err := dec.Decode(&x); err != nil {`)
	o.L("return fmt.Errorf(`failed to decode value for key %%q: %%w`, tok, err)")
	o.L(`}`)
	o.L(`}`)
	o.L(`if privateParams == nil {`)
	o.L(`privateParams = make(map[string]interface{})`)
	o.L(`}`)
	o.L(`privateParams[tok] = x`)
	o.L(`}`)
	o.L(`}`)
	o.L(`}`)
	o.L(`if privateParams != nil {`)
	o.L(`v.privateParams = privateParams`)
	o.L(`}`)
	o.L(`return nil`)
	o.L(`}`)

	o.LL(`func (v *%s) AsMap(dst map[string]interface{}) error {`, object.Name(true))
	o.L(`for _, pair := range v.makePairs() {`)
	o.L(`dst[pair.Key] = pair.Value`)
	o.L(`}`)
	o.L(`return nil`)
	o.L(`}`)

	o.LL(`type %sBuilder struct {`, object.Name(true))
	o.L(`once sync.Once`)
	o.L(`mu sync.Mutex`)
	o.L(`err error`)
	o.L(`validator %sValidator`, object.Name(true))
	o.L(`object *%s`, object.Name(true))
	o.L(`}`)

	o.LL(`func (b *Builder) %[1]s() *%[1]sBuilder {`, object.Name(true))
	o.L(`return New%sBuilder()`, object.Name(true))
	o.L(`}`)

	o.LL(`func New%[1]sBuilder() *%[1]sBuilder {`, object.Name(true))
	o.L(`var b %sBuilder`, object.Name(true))
	o.L(`b.init()`)
	o.L(`return &b`)
	o.L(`}`)

	o.LL(`func (b *%[1]sBuilder) From(in *%[1]s) *%[1]sBuilder {`, object.Name(true))
	o.L(`b.once.Do(b.init)`)
	o.L(`b.object = in.Clone()`)
	o.L(`return b`)
	o.L(`}`)

	o.LL(`func (b *%sBuilder) init() {`, object.Name(true))
	o.L(`b.err = nil`)
	o.L(`b.validator = nil`)
	o.L(`b.object = &%s{}`, object.Name(true))
	if schema := object.String(`schema`); schema != "" {
		if !object.Bool(`skipCommonFields`) {
			o.LL(`b.object.schemas = make(schemas)`)
			o.L(`b.object.schemas.Add(%sSchemaURI)`, object.Name(true))
		}
	}
	o.L(`}`)

	for _, field := range object.Fields() {
		// If the argument is a slice, the parameter type should be varg
		if field.Name(false) == `schemas` {
			o.LL(`func (b *%[1]sBuilder) %[2]s(v ...string) *%[1]sBuilder {`, object.Name(true), field.Name(true))
		} else if IsSlice(field) {
			o.LL(`func (b *%[1]sBuilder) %[2]s(v ...%[3]s) *%[1]sBuilder {`, object.Name(true), field.Name(true), strings.TrimPrefix(field.Type(), `[]`))
		} else {
			o.LL(`func (b *%[1]sBuilder) %[2]s(v %[3]s) *%[1]sBuilder {`, object.Name(true), field.Name(true), field.Type())
		}
		o.L(`b.mu.Lock()`)
		o.L(`defer b.mu.Unlock()`)
		o.L(`b.once.Do(b.init)`)
		o.L(`if b.err != nil {`)
		o.L(`return b`)
		o.L(`}`)
		if field.Name(false) == `schemas` {
			o.L(`for _, schema := range v {`)
			o.L(`b.object.schemas.Add(schema)`)
			o.L(`}`)
		} else {
			o.L(`if err := b.object.Set(%q, v); err != nil {`, field.JSON())
			o.L(`b.err = err`)
			o.L(`}`)
		}
		o.L(`return b`)
		o.L(`}`)
	}

	if object.String(`schema`) != "" {
		o.LL(`func (b *%[1]sBuilder) Extension(uri string, value interface{}) *%[1]sBuilder {`, object.Name(true))
		o.L(`b.mu.Lock()`)
		o.L(`defer b.mu.Unlock()`)
		o.L(`b.once.Do(b.init)`)
		o.L(`if b.err != nil {`)
		o.L(`return b`)
		o.L(`}`)
		o.L(`b.object.schemas.Add(uri)`)
		o.L(`if err := b.object.Set(uri, value); err != nil {`)
		o.L(`b.err = err`)
		o.L(`}`)
		o.L(`return b`)
		o.L(`}`)
	}

	o.LL(`func (b *%[1]sBuilder) Validator(v %[1]sValidator) *%[1]sBuilder {`, object.Name(true))
	o.L(`b.mu.Lock()`)
	o.L(`defer b.mu.Unlock()`)
	o.L(`b.once.Do(b.init)`)
	o.L(`if b.err != nil {`)
	o.L(`return b`)
	o.L(`}`)
	o.L(`b.validator = v`)
	o.L(`return b`)
	o.L(`}`)

	o.LL(`func (b *%[1]sBuilder) Build() (*%[1]s, error) {`, object.Name(true))
	o.L(`b.mu.Lock()`)
	o.L(`defer b.mu.Unlock()`)
	o.L(`object := b.object`)
	o.L(`validator := b.validator`)
	o.L(`err := b.err`)
	o.L(`b.once = sync.Once{}`)
	o.L(`if err != nil {`)
	o.L(`return nil, err`)
	o.L(`}`)
	o.L(`if object == nil {`)
	o.L(`return nil, fmt.Errorf("resource.%sBuilder: object was not initialized")`, object.Name(true))
	o.L(`}`)
	o.L(`if validator == nil {`)
	o.L(`validator = Default%sValidator`, object.Name(true))
	o.L(`}`)
	// There's a rare circumstance in which somebody knowingly or unknowningly
	// set the validator to nil
	o.L(`if validator != nil {`)
	o.L(`if err := validator.Validate(object); err != nil {`)
	o.L(`return nil, err`)
	o.L(`}`)
	o.L(`}`)
	o.L(`return object, nil`)
	o.L(`}`)

	o.LL(`func (b *%[1]sBuilder) MustBuild() *%[1]s {`, object.Name(true))
	o.L(`object, err := b.Build()`)
	o.L(`if err != nil {`)
	o.L(`panic(err)`)
	o.L(`}`)
	o.L(`return object`)
	o.L(`}`)
	if err := o.WriteFile(xstrings.Snake(object.Name(false))+`_gen.go`, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}
		return fmt.Errorf(`failed to write to %s_gen.go: %w`, xstrings.Snake(object.Name(false)), err)
	}
	return nil
}

func generateEnt(object *codegen.Object) error {
	// for the time being, only generate for hardcoded objects.
	// later, move this definition to objects.yml
	switch object.Name(true) {
	case `User`, `Group`, `Email`, `Names`:
	default:
		return nil
	}

	fmt.Printf("  ⌛ Generating ent adapters for %s...\n", object.Name(true))

	if err := generateSchema(object); err != nil {
		return fmt.Errorf(`failed to generate schema: %w`, err)
	}

	if err := generateUtilities(object); err != nil {
		return fmt.Errorf(`failed to generate utilities: %w`, err)
	}
	return nil
}

func generateSchema(object *codegen.Object) error {
	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)

	o.L(`package schema`)

	o.LL(`func (%s) Fields() []ent.Field {`, object.Name(true))
	o.L(`return []ent.Field{`)
	for _, field := range object.Fields() {
		if field.Name(false) == "schemas" {
			continue
		}

		var ft = field.Type()
		if strings.HasPrefix(ft, `[]`) || strings.HasPrefix(ft, `*`) {
			continue
		}

		var entMethod = xstrings.Camel(ft)
		if v := field.String(`ent_build_method`); v != "" {
			entMethod = v
		}

		var entName = field.Name(false)
		if v := field.String(`ent_name`); v != "" {
			entName = v
		}
		var entType = field.String(`ent_type`)
		var entDefault = field.String(`ent_default`)

		if entType != "" {
			o.L(`field.%s(%q, %s)`, entMethod, entName, entType)
		} else {
			o.L(`field.%s(%q)`, entMethod, entName)
		}

		if entDefault != "" {
			o.R(`.Default(%s)`, entDefault)
		}
		if !field.IsRequired() {
			o.R(`.Optional()`)
		}

		if field.Bool(`ent_unique`) {
			o.R(`.Unique()`)
		}
		if field.Bool(`ent_notempty`) {
			o.R(`.NotEmpty()`)
		}
		o.R(`,`)
	}
	o.L(`}`)
	o.L(`}`)

	fn := filepath.Join(`..`, `sample`, `ent`, `schema`, xstrings.Snake(object.Name(false))+`_gen.go`)
	if err := o.WriteFile(fn, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}
		return fmt.Errorf(`failed to write to %s: %w`, fn, err)
	}
	return nil
}

func generateUtilities(object *codegen.Object) error {
	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)

	o.L(`package sample`)

	o.LL(`import (`)
	o.L(`"github.com/cybozu-go/scim/resource"`)
	o.L(`"github.com/cybozu-go/scim/sample/ent"`)
	o.L(`"github.com/cybozu-go/scim/sample/ent/%s"`, object.Name(false))
	o.L(`)`)

	if object.String(`schema`) != "" {
		o.LL(`func %sLoadEntFields(q *ent.%sQuery, fields []string) {`, object.Name(false), object.Name(true))
		o.L(`if len(fields) == 0 {`)
		o.L(`fields = []string{`)
		for i, field := range object.Fields() {
			switch field.Name(false) {
			case "schemas", "meta":
				continue
			}
			if field.Bool(`skipCommonFields`) {
				switch field.Name(false) {
				case "id", "externalID":
					continue
				}
			}

			if i > 0 {
				o.R(`,`)
			}
			o.R(`%q`, field.Name(false))
		}
		o.R(`}`)
		o.L(`}`)

		o.L(`selectNames := make([]string, 0, len(fields))`)
		o.L(`for _, f := range fields {`)
		o.L(`switch f {`)
		for _, field := range object.Fields() {
			if field.Name(false) == "schemas" {
				continue
			}
			if field.Bool(`skipCommonFields`) {
				switch field.Name(false) {
				case "id", "externalID", "meta":
					continue
				}
			}

			o.L(`case %q:`, field.Name(false))
			// Special case
			var ft = field.Type()
			if strings.HasPrefix(ft, `[]`) || strings.HasPrefix(ft, `*`) {
				// TODO: later
				switch field.Name(false) {
				case `emails`, `name`:
					o.L(`q.With%s()`, field.Name(true))
				}
				continue
			} else {
				// Otherwise, accumulate in the list of names
				o.L(`selectNames = append(selectNames, %s.Field%s)`, object.Name(false), field.Name(true))
			}
		}
		o.L(`}`)
		o.L(`}`)
		o.L(`q.Select(selectNames...)`)
		o.L(`}`)
	}

	o.LL(`func %[1]sResourceFromEnt(in *ent.%[1]s) (*resource.%[1]s, error) {`, object.Name(true))
	o.L(`var b resource.Builder`)

	o.LL(`builder := b.%s()`, object.Name(true))

	if !object.Bool(`skipCommonFields`) {
		o.LL(`meta, err := b.Meta().`)
		o.L(`ResourceType(%q).`, object.Name(true))
		o.L(`Location(%q+in.ID.String()).` /* TODO: FIXME */, fmt.Sprintf(`https://foobar.com/scim/v2/%s/`, object.Name(true)))
		o.L(`Build()`)
		o.L(`if err != nil {`)
		o.L(`return nil, fmt.Errorf("failed to build meta information for %s")`, object.Name(true))
		o.L(`}`)
		o.LL(`builder.`)
		o.L(`Meta(meta)`)
	}
	for _, field := range object.Fields() {
		if field.Name(false) == "schemas" {
			continue
		}

		switch field.Name(false) {
		case `emails`, `name`:
		default:
			continue
		}

		// TODO: include others

		rsname := strings.TrimSuffix(field.Name(true), "s")
		if rsname == "Name" {
			rsname = "Names"
		}
		o.LL(`if el := len(in.Edges.%s); el > 0 {`, field.Name(true))
		o.L(`list := make([]*resource.%s, 0, el)`, rsname)
		o.L(`for _, ine := range in.Edges.%s {`, field.Name(true))
		o.L(`r, err := %sResourceFromEnt(ine)`, rsname)
		o.L(`if err != nil {`)
		o.L(`return nil, fmt.Errorf("failed to build %s information for %s")`, field.Name(false), object.Name(true))
		o.L(`}`)
		o.L(`list = append(list, r)`)
		o.L(`}`)

		if strings.HasPrefix(field.Type(), "*") {
			o.L(`builder.%s(list[0])`, field.Name(true))
		} else {
			o.L(`builder.%s(list...)`, field.Name(true))
		}
		o.L(`}`)
	}

	for _, field := range object.Fields() {
		switch field.Name(true) {
		// FIXME: do't hard codethis
		case "Password":
			continue
		case "ID":
			o.L(`builder.%[1]s(in.%[1]s.String())`, field.Name(true))
		case "Schemas", "Meta", "Members", "Addresses", "Emails", "Entitlements", "IMS", "NickName", "Name", "Groups", "PhoneNumbers", "ProfileURL", "Title", "Roles", "X509Certificates":
			// TODO: FIXME
		default:
			o.L(`if !reflect.ValueOf(in.%s).IsZero() {`, field.Name(true))
			o.L(`builder.%[1]s(in.%[1]s)`, field.Name(true))
			o.L(`}`)
		}
	}
	o.L(`return builder.Build()`)
	o.L(`}`)

	o.LL(`func %sEntFileFromSCIM(s string) string {`, object.Name(true))
	o.L(`switch s {`)
	for _, field := range object.Fields() {
		if strings.HasPrefix(field.Type(), `[]`) || strings.HasPrefix(field.Type(), `*`) {
			continue
		}
		switch field.Name(false) {
		case `schemas`:
			continue
		default:
		}
		o.L(`case resource.%s%sKey:`, object.Name(true), field.Name(true))
		o.L(`return %s.Field%s`, object.Name(false), field.Name(true))
	}
	o.L(`default:`)
	o.L(`return s`)
	o.L(`}`)
	o.L(`}`)

	o.LL(`func %sPresencePredicate(scimField string) predicate.%s {`, object.Name(false), object.Name(true))
	o.L(`switch scimField {`)
	for _, field := range object.Fields() {
		switch field.Name(false) {
		case `schemas`:
			continue
		default:
		}
		if field.Type() != "string" {
			continue
		}
		if field.IsRequired() {
			continue
		}
		o.L(`case resource.%s%sKey:`, object.Name(true), field.Name(true))
		o.L(`return %[1]s.And(%[1]s.%[2]sNotNil(), %[1]s.%[2]sNEQ(""))`, object.Name(false), field.Name(true))
	}
	o.L(`default:`)
	o.L(`return nil`)
	o.L(`}`)
	o.L(`}`)

	fn := filepath.Join(`..`, `sample`, xstrings.Snake(object.Name(false))+`_gen.go`)
	if err := o.WriteFile(fn, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}
		return fmt.Errorf(`failed to write to %s: %w`, fn, err)
	}
	return nil
}
