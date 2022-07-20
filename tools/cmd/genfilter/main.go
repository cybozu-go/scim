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
		Objects []*codegen.Object `json:"objects"`
	}
	if err := json.NewDecoder(bytes.NewReader(jsonSrc)).Decode(&def); err != nil {
		return fmt.Errorf(`failed to decode %q: %w`, *objectsFile, err)
	}

	for _, object := range def.Objects {
		object.Organize()
	}

	if err := generateExpr(def.Objects); err != nil {
		return fmt.Errorf(`failed to generate filter expressions: %w`, err)
	}

	if err := generateAliases(def.Objects); err != nil {
		return fmt.Errorf(`failed to generate filter aliases: %w`, err)
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

func generateExpr(objects []*codegen.Object) error {
	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)

	o.L(`package expr`)

	for _, object := range objects {
		o.LL(`type %s interface {`, object.Name(true))
		o.L(`Interface`)
		for _, field := range object.Fields() {
			o.L(`%s() %s`, field.Name(true), field.Type())
		}
		o.L(`}`)

		o.LL(`type %s struct {`, object.Name(false))
		for _, field := range object.Fields() {
			o.L(`%s %s`, field.Name(false), field.Type())
		}
		o.L(`}`)

		o.LL(`func (*%s) expression() {}`, object.Name(false))

		o.LL(`func New%s(`, object.Name(true))
		for i, field := range object.Fields() {
			if i > 0 {
				o.R(`,`)
			}
			o.R(`%s %s`, field.Name(false), field.Type())
		}
		o.R(`) %s {`, object.Name(true))
		o.L(`return &%s{`, object.Name(false))
		for _, field := range object.Fields() {
			o.L(`%[1]s: %[1]s,`, field.Name(false))
		}
		o.L(`}`)
		o.L(`}`)

		for _, field := range object.Fields() {
			o.LL(`func (e *%s) %s() %s {`, object.Name(false), field.Name(true), field.Type())
			o.L(`return e.%s`, field.Name(false))
			o.L(`}`)
		}
	}
	const fn = `internal/expr/expr_gen.go`
	if err := o.WriteFile(fn, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}
		return fmt.Errorf(`failed to write to %s: %w`, fn, err)
	}
	return nil
}

func generateAliases(objects []*codegen.Object) error {
	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)

	o.L(`package filter`)

	o.LL(`import (`)
	o.L(`"github.com/cybozu-go/scim/filter/internal/expr"`)
	o.L(`"github.com/cybozu-go/scim/filter/internal/token"`)
	o.L(`)`)
	tokens := []string{
		`NotOp`,
		`AndOp`,
		`OrOp`,
		`PresenceOp`,
		`EqualOp`,
		`NotEqualOp`,
		`GreaterThanOp`,
		`GreaterThanOrEqualToOp`,
		`LessThanOp`,
		`LessThanOrEqualToOp`,
		`ContainsOp`,
		`StartsWithOp`,
		`EndsWithOp`,
		`Dot`,
		`LParen`,
		`RParen`,
		`LBracket`,
		`RBracket`,
		`True`,
		`False`,
		`Null`,
	}

	o.LL(`const (`)
	for _, token := range tokens {
		o.L(`%[1]s = token.%[1]s`, token)
	}
	o.L(`)`)

	for _, object := range objects {
		o.LL(`type %[1]s = expr.%[1]s`, object.Name(true))
		o.LL(`func New%s(`, object.Name(true))
		for i, field := range object.Fields() {
			if i > 0 {
				o.R(`,`)
			}
			arg := field.Name(false)
			if arg == `expr` {
				arg = `e`
			}
			typ := field.Type()
			if typ == "Interface" {
				typ = "Expr"
			}
			o.R(`%s %s`, arg, typ)
		}
		o.R(`) %s {`, object.Name(true))
		o.L(`return expr.New%s(`, object.Name(true))
		for i, field := range object.Fields() {
			if i > 0 {
				o.R(`,`)
			}
			arg := field.Name(false)
			if arg == `expr` {
				arg = `e`
			}
			o.R(`%s`, arg)
		}
		o.R(`)`)
		o.L(`}`)
	}
	const fn = `expr_gen.go`
	if err := o.WriteFile(fn, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}
		return fmt.Errorf(`failed to write to %s: %w`, fn, err)
	}
	return nil
}
