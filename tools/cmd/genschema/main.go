package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/cybozu-go/scim/resource"
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
	var schemaFile = flag.String("schema", "schema.yml", "")
	flag.Parse()
	jsonSrc, err := yaml2json(*schemaFile)
	if err != nil {
		return err
	}

	var def struct {
		Common  []*resource.SchemaAttribute `json:"common"`
		Schemas []*resource.Schema          `json:"schemas"`
	}
	if err := json.Unmarshal(jsonSrc, &def); err != nil {
		return err
	}

	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)
	o.L(`package schema`)

	o.LL(`func init() {`)
	for i, r := range def.Schemas {
		if i > 0 {
			o.LL(`{`)
		} else {
			o.L(`{`)
		}
		o.L(`s := resource.NewSchemaBuilder().`)
		o.L(`ID(%q).`, r.ID())
		o.L(`Name(%q).`, r.Name())
		o.L(`Description(%q).`, r.Description())
		o.L(`Attributes(`)

		attrs := r.Attributes()
		attrs = append(attrs, def.Common...)

		sort.Slice(attrs, func(i, j int) bool {
			return attrs[i].Name() < attrs[j].Name()
		})

		for _, attr := range attrs {
			if err := generateAttr(o, attr); err != nil {
				return fmt.Errorf(`failed to generate attribute: %w`, err)
			}
			o.R(`,`)
		}
		o.L(`).`)
		o.L(`Name(%q).`, r.Name())
		o.L(`MustBuild()`)

		o.L(`Register(%q, s)`, r.Name())
		o.L(`Register(resource.%sSchemaURI, s)`, r.Name())
		o.L(`}`)
	}
	o.L(`}`)

	filename := `schema_gen.go`
	if err := o.WriteFile(filename, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}
		return fmt.Errorf(`failed to write to %s: %w`, filename, err)
	}
	return nil
}

func generateAttr(o *codegen.Output, attr *resource.SchemaAttribute) error {
	o.L(`resource.NewSchemaAttributeBuilder().`)
	o.L(`Name(%q).`, attr.Name())
	o.L(`Type(%q).`, attr.Type())
	o.L(`MultiValued(%t).`, attr.MultiValued())
	o.L(`Description(%q).`, attr.Description())
	o.L(`Required(%t).`, attr.Required())
	o.L(`CaseExact(%t).`, attr.CaseExact())

	var mut string
	switch attr.Mutability() {
	case resource.MutReadOnly:
		mut = `resource.MutReadOnly`
	case resource.MutReadWrite:
		mut = `resource.MutReadWrite`
	case resource.MutImmutable:
		mut = `resource.MutImmutable`
	case resource.MutWriteOnly:
		mut = `resource.MutWriteOnly`
	default:
		return fmt.Errorf(`unknown mutability: %q`, attr.Mutability())
	}
	o.L(`Mutability(%s).`, mut)

	var returned string
	switch attr.Returned() {
	case resource.ReturnedAlways:
		returned = `resource.ReturnedAlways`
	case resource.ReturnedNever:
		returned = `resource.ReturnedNever`
	case resource.ReturnedDefault:
		returned = `resource.ReturnedDefault`
	case resource.ReturnedRequest:
		returned = `resource.ReturnedRequest`
	default:
		return fmt.Errorf(`unknown returned: %q`, attr.Returned())
	}
	o.L(`Returned(%s).`, returned)

	var uniq string
	switch attr.Uniqueness() {
	case resource.UniqNone:
		uniq = `resource.UniqNone`
	case resource.UniqServer:
		uniq = `resource.UniqServer`
	case resource.UniqGlobal:
		uniq = `resource.UniqGlobal`
	default:
		return fmt.Errorf(`unknown uniqueness: %q`, attr.Uniqueness())
	}
	o.L(`Uniqueness(%s).`, uniq)

	subattrs := attr.Attributes()
	if len(subattrs) > 0 {
		o.L(`Attributes(`)
		for _, subattr := range subattrs {
			if err := generateAttr(o, subattr); err != nil {
				return err
			}
			o.R(`,`)
		}
		o.L(`).`)
	}

	o.L(`GoAccessorName(%q).`, goAccessorName(attr))

	o.L(`MustBuild()`)
	return nil
}

func goAccessorName(attr *resource.SchemaAttribute) string {
	s := attr.Name()
	if s == `$ref` {
		s = `reference`
	}
	s = xstrings.Camel(s)
	if s == `Ims` {
		s = `IMS`
	}
	if strings.HasSuffix(s, `Id`) {
		s = strings.TrimSuffix(s, `Id`) + `ID`
	}
	if strings.HasSuffix(s, `Url`) {
		s = strings.TrimSuffix(s, `Url`) + `URL`
	}
	return s
}
