package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
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

		object.Organize()
		if err := generateObject(object); err != nil {
			return fmt.Errorf(`failed to generate object: %s`, err)
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
	return strings.HasPrefix(s, `*`) || strings.HasPrefix(s, `[]`) || strings.HasSuffix(s, `List`)
}

func generateObject(object *codegen.Object) error {
	fmt.Printf("  ⌛ Generating %s...\n", object.Name(true))

	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)

	o.L(`package resource`)

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

	o.LL(`var Default%[1]sValidator %[1]sValidator`, object.Name(true))

	for _, field := range object.Fields() {
		o.LL(`func (v *%s) %s() %s {`, object.Name(true), field.Name(true), field.Type())
		o.L(`v.mu.RLock()`)
		o.L(`defer v.mu.RUnlock()`)
		if IsIndirect(field) {
			o.L(`return v.%s`, field.Name(false))
		} else {
			o.L(`if v.%s == nil {`, field.Name(false))
			o.L(`return %s`, codegen.ZeroVal(field.Type()))
			o.L(`}`)
			o.L(`return *(v.%s)`, field.Name(false))
		}
		o.L(`}`)
	}

	o.LL(`func (v *%s) MarshalJSON() ([]byte, error) {`, object.Name(true))
	o.L(`type pair struct {`)
	o.L(`Key string`)
	o.L(`Value interface{}`)
	o.L(`}`)
	o.L(`var pairs []pair`)
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
	o.L(`for _, option := range options {`)
	o.L(`switch option.Ident() {`)
	o.L(`case identExtension{}:`)
	o.L(`ext = option.Value().(string)`)
	o.L(`}`)
	o.L(`}`)

	o.L(`switch name {`)
	for _, field := range object.Fields() {
		o.L(`case %q:`, field.JSON())
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
		o.L(`case %q:`, field.JSON())
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
		o.L(`case %q:`, field.JSON())
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
	o.L(`if rt, ok := registry.Get(tok); ok {`)
	o.L(`x = reflect.New(rt).Interface()`)
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

	o.LL(`type %sBuilder struct {`, object.Name(true))
	o.L(`mu sync.Mutex`)
	o.L(`err error`)
	o.L(`validator %sValidator`, object.Name(true))
	o.L(`object *%s`, object.Name(true))
	o.L(`}`)

	o.LL(`func (b *Builder) %[1]s() *%[1]sBuilder {`, object.Name(true))
	o.L(`return &%sBuilder{}`, object.Name(true))
	o.L(`}`)

	for _, field := range object.Fields() {
		// If the argument is a slice, the parameter type should be varg
		if IsSlice(field) {
			o.LL(`func (b *%[1]sBuilder) %[2]s(v ...%[3]s) *%[1]sBuilder {`, object.Name(true), field.Name(true), strings.TrimPrefix(field.Type(), `[]`))
		} else {
			o.LL(`func (b *%[1]sBuilder) %[2]s(v %[3]s) *%[1]sBuilder {`, object.Name(true), field.Name(true), field.Type())
		}
		o.L(`b.mu.Lock()`)
		o.L(`defer b.mu.Unlock()`)
		o.L(`if b.err != nil {`)
		o.L(`return b`)
		o.L(`}`)
		o.L(`if b.object == nil {`)
		o.L(`b.object = &%s{}`, object.Name(true))
		o.L(`}`)
		o.L(`if err := b.object.Set(%q, v); err != nil {`, field.JSON())
		o.L(`b.err = err`)
		o.L(`}`)
		o.L(`return b`)
		o.L(`}`)
	}

	o.LL(`func (b *%[1]sBuilder) Extension(uri string, value interface{}) *%[1]sBuilder {`, object.Name(true))
	o.L(`b.mu.Lock()`)
	o.L(`defer b.mu.Unlock()`)
	o.L(`if b.err != nil {`)
	o.L(`return b`)
	o.L(`}`)
	o.L(`if b.object == nil {`)
	o.L(`b.object = &%s{}`, object.Name(true))
	o.L(`}`)
	o.L(`if err := b.object.Set(uri, value); err != nil {`)
	o.L(`b.err = err`)
	o.L(`}`)
	o.L(`return b`)
	o.L(`}`)

	o.LL(`func (b *%[1]sBuilder) Validator(v %[1]sValidator) *%[1]sBuilder {`, object.Name(true))
	o.L(`b.mu.Lock()`)
	o.L(`defer b.mu.Unlock()`)
	o.L(`if b.err != nil {`)
	o.L(`return b`)
	o.L(`}`)
	o.L(`b.validator = v`)
	o.L(`return b`)
	o.L(`}`)

	o.LL(`func (b *%[1]sBuilder) Build() (*%[1]s, error) {`, object.Name(true))
	o.L(`object := b.object`)
	o.L(`validator := b.validator`)
	o.L(`b.object = nil`)
	o.L(`b.validator = nil`)
	o.L(`if object == nil {`)
	o.L(`return nil, fmt.Errorf("resource.%sBuilder: object was not initialized")`, object.Name(true))
	o.L(`}`)
	o.L(`if err := b.err; err != nil {`)
	o.L(`return nil, err`)
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
