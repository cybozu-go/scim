package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/schema"
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

type Service struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Calls       []*codegen.Object `json:"calls"`
}

func _main() error {
	codegen.RegisterZeroVal(`AuthenticationSchemeType`, `InvalidAuthenticationScheme`)

	// Contains the "call" objects
	var callsFile = flag.String("calls", "calls.yml", "")

	// Contains the resource objects that the calls use
	// var resourcesFile = flag.String("resources", "resources.yml", "")
	flag.Parse()

	//	resourcesSrc, err := yaml2json(*resourcesFile)
	//	if err != nil {
	//		return err
	//	}

	callsSrc, err := yaml2json(*callsFile)
	if err != nil {
		return err
	}

	/*
		var resourceDefs struct {
			Common  codegen.FieldList
			Objects []*codegen.Object `json:"objects"`
		}
		if err := json.NewDecoder(bytes.NewReader(resourcesSrc)).Decode(&resourceDefs); err != nil {
			return fmt.Errorf(`failed to decode %q: %w`, *resourcesFile, err)
		}

		resources := make(map[string]*codegen.Object)
		for _, object := range resourceDefs.Objects {
			// Each object needs a common set of fields.
			if !object.Bool(`skipCommonFields`) {
				for _, commonField := range resourceDefs.Common {
					object.AddField(commonField)
				}
			}

			object.Organize()
			resources[object.Name(true)] = object
		}*/

	var calls struct {
		Services []Service `json:"services"`
	}
	if err := json.NewDecoder(bytes.NewReader(callsSrc)).Decode(&calls); err != nil {
		return fmt.Errorf(`failed to decode %q: %w`, *callsFile, err)
	}

	for _, service := range calls.Services {
		if err := generateService(service); err != nil {
			return fmt.Errorf(`failed to generate service: %s`, err)
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

func generateService(svc Service) error {
	fmt.Printf("  ⌛ Generating %s...\n", svc.Name)

	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)
	o.L(`package client`)

	o.LL(`import (`)
	for _, pkg := range []string{`github.com/goccy/go-yaml`} {
		o.L(`%q`, pkg)
	}
	o.L(`)`)

	o.Comment(svc.Description)
	o.L(`type %s struct {`, svc.Name)
	o.L(`client *Client`)
	o.L(`}`)

	if svc.Name != "SearchService" {
		o.Comment(fmt.Sprintf(`%s creates a new Service object to perform an operation`, strings.TrimSuffix(svc.Name, `Service`)))
		o.L(`func (client *Client) %s() *%s {`, strings.TrimSuffix(svc.Name, `Service`), svc.Name)
		o.L(`return &%s{`, svc.Name)
		o.L(`client: client,`)
		o.L(`}`)
		o.L(`}`)
	}

	for _, call := range svc.Calls {
		fmt.Printf("    ⌛ Call %s...\n", call.Name(true))

		if err := generateCall(o, svc, call); err != nil {
			return fmt.Errorf(`failed to generate call %s: %w`, call.Name(true), err)
		}
	}

	fn := xstrings.Snake(svc.Name) + `_gen.go`
	if err := o.WriteFile(fn, codegen.WithFormatCode(true)); err != nil {
		if cfe, ok := err.(codegen.CodeFormatError); ok {
			fmt.Fprint(os.Stderr, cfe.Source())
		}
		return fmt.Errorf(`failed to write to %s_gen.go: %w`, fn, err)
	}
	return nil
}

func generateCall(o *codegen.Output, svc Service, call *codegen.Object) error {
	rstype := call.String(`resource`)

	o.Comment(fmt.Sprintf(`%s is an encapsulation of a SCIM operation.`, call.Name(true)))
	o.L(`type %s struct {`, call.Name(true))
	if rstype != "" {
		o.L(`builder *resource.%sBuilder`, rstype)
		o.L(`object *resource.%s`, rstype)
	}
	o.L(`err error`)
	o.L(`client *Client`)
	o.L(`trace io.Writer`)

	call.Organize()

	var required []codegen.Field
	var optional []codegen.Field
	for _, f := range call.Fields() {
		// TODO: not handling complex types here yet
		o.L(`%s %s`, f.Name(false), f.Type())
		if f.IsRequired() {
			required = append(required, f)
		} else {
			optional = append(optional, f)
		}
	}
	o.L(`}`)

	methodName := strings.TrimSuffix(call.Name(true), `Call`)
	if v := call.String(`method_name`); v != "" {
		methodName = v
	}

	if rstype != "" {
		o.LL(`func (call *%s) payload() (*resource.%s, error) {`, call.Name(true), rstype)
		o.L(`if object := call.object; object != nil {`)
		o.L(`return object, nil`)
		o.L(`}`)
		o.L(`return call.builder.Build()`)
		o.L(`}`)

		o.LL(`func (call *%[1]s) FromJSON(data []byte) *%[1]s {`, call.Name(true))
		o.L(`if call.err != nil {`)
		o.L(`return call`)
		o.L(`}`)
		o.L(`var in resource.%s`, rstype)
		o.L(`if err := json.Unmarshal(data, &in); err != nil {`)
		o.L(`call.err = fmt.Errorf("failed to decode data: %%w", err)`)
		o.L(`return call`)
		o.L(`}`)
		o.L(`call.object = &in`)
		o.L(`return call`)
		o.L(`}`)
	}

	o.R("\n")
	if desc := call.String(`description`); desc != "" {
		scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(desc)))
		for scanner.Scan() {
			o.L(`// %s`, scanner.Text())
		}
	}

	o.L(`func (svc *%s) %s(`, svc.Name, methodName)
	for i, f := range required {
		if i > 0 {
			o.R(`, `)
		}
		o.R(`%s %s`, f.Name(false), f.Type())
	}
	o.R(`) *%s {`, call.Name(true))
	o.L(`return &%s{`, call.Name(true))
	if rstype != "" {
		o.L(`builder: resource.New%sBuilder(),`, call.String(`resource`))
	}
	o.L(`client: svc.client,`)
	for _, f := range required {
		o.L(`%[1]s: %[1]s,`, f.Name(false))
	}
	o.L(`}`)
	o.L(`}`)

	resType := call.String(`response_type`)
	jsonPayload := call.Bool(`jsonPayload`)

	// if we need to generate a resource to be sent to the server, we look it up
	// an empty "resource" means no resource generation is necessary, and we can
	// just send an "empty" request
	if rstype != "" {
		allowed := make(map[string]struct{})
		s, ok := schema.GetByResourceType(rstype)
		if ok {
			mutabilities := make(map[string]struct{})
			if iface, ok := call.Extra(`allowedMutability`); ok {
				if vam, ok := iface.([]interface{}); ok {
					for _, v := range vam {
						if sv, ok := v.(string); ok {
							mutabilities[sv] = struct{}{}
						}
					}
				}
			}

			if len(mutabilities) > 0 {
				for _, attr := range s.Attributes() {
					mut := string(attr.Mutability())
					if _, ok := mutabilities[mut]; !ok {
						continue
					}

					allowed[attr.GoAccessorName()] = struct{}{}
				}
			}
		}

		builder, ok := resource.LookupBuilderByName(rstype)
		if !ok {
			return fmt.Errorf(`could not find builder for resource type %q`, rstype)
		}

		var tpHelper func(*strings.Builder, reflect.Type)
		tpHelper = func(sb *strings.Builder, t reflect.Type) {
			switch t.Kind() {
			case reflect.Slice:
				sb.WriteString(`[]`)
				tpHelper(sb, t.Elem())
			case reflect.Ptr:
				sb.WriteString(`*`)
				tpHelper(sb, t.Elem())
			case reflect.Interface:
				if t.Name() == "" {
					sb.WriteString(`interface{}`)
				}
				fallthrough
			default:
				if pkg := t.PkgPath(); pkg != "" {
					i := strings.LastIndex(pkg, "/")
					if i > 0 {
						pkg = pkg[i+1:]
					}
					sb.WriteString(pkg)
					sb.WriteRune('.')
				}
				sb.WriteString(t.Name())
			}
		}

		typeName := func(t reflect.Type) string {
			var sb strings.Builder
			tpHelper(&sb, t)
			return sb.String()
		}

		rt := reflect.TypeOf(builder)
		for i := 0; i < rt.NumMethod(); i++ {
			m := rt.Method(i)
			switch m.Name {
			case `Build`, `MustBuild`, `From`, `Extension`, `SetField`:
				continue
			}
			if len(allowed) > 0 {
				if _, ok := allowed[m.Name]; !ok {
					continue
				}
			}

			o.LL(`func(call *%[1]s) %[2]s(`, call.Name(true), m.Name)
			mt := m.Type

			// There should be only one argument, so we use 1
			if mt.IsVariadic() {
				o.R(`in ...%s`, mt.In(1).Elem())
			} else {
				o.R(`in %s`, typeName(mt.In(1)))
			}
			o.R(`) *%s {`, call.Name(true))

			o.R(`call.builder.%s(in`, m.Name)
			// variadic args need expanding
			if mt.IsVariadic() {
				o.R(`...`)
			}
			o.R(`)`)
			o.L(`return call`)
			o.L(`}`)
		}

		if jsonPayload {
			o.LL(`// Extension allows users to register an extension using the fully qualified URI`)
			o.L(`func (call *%[1]s) Extension(uri string, value interface{}) *%[1]s {`, call.Name(true))
			o.L(`call.builder.Extension(uri, value)`)
			o.L(`return call`)
			o.L(`}`)
		}
	}

	o.LL(`func (call *%[1]s) Trace(w io.Writer) *%[1]s {`, call.Name(true))
	o.L(`call.trace = w`)
	o.L(`return call`)
	o.L(`}`)

	if tmplSrc := call.String(`makeURLTemplate`); tmplSrc != "" {
		var vars = struct {
			ObjectName string
			Path       string
		}{
			ObjectName: call.Name(true),
			Path:       call.String(`path`),
		}
		tmpl, err := template.New("makeURL").Parse(tmplSrc)
		if err != nil {
			return fmt.Errorf(`failed to parse makeURLTemplate for %s: %w`, call.Name(true), err)
		}
		var makeURLOut bytes.Buffer
		if err := tmpl.Execute(&makeURLOut, vars); err != nil {
			return fmt.Errorf(`failed to execute template: %w`, err)
		}
		o.LL(makeURLOut.String())
	} else {
		o.LL(`func (call *%[1]s) makeURL() string {`, call.Name(true))
		o.L(`return	call.client.baseURL + %q`, call.String(`path`))
		o.L(`}`)
	}

	if resType == "" {
		return fmt.Errorf(`response_type is not specified for %q`, call.Name(true))
	}

	var errPrefix string
	if resType == "none" {
		o.LL(`func (call *%s) Do(ctx context.Context) error {`, call.Name(true))
	} else {
		errPrefix = "nil, "
		o.LL(`func (call *%s) Do(ctx context.Context) (*%s, error) {`, call.Name(true), resType)
	}

	o.L(`if err := call.err; err != nil {`)
	o.L(`return %sfmt.Errorf("failed to build request: %%w", err)`, errPrefix)
	o.L(`}`)

	if rstype != "" {
		o.L(`payload, err := call.payload()`)
		o.L(`if err != nil {`)
		o.L("return %sfmt.Errorf(`failed to generate request payload for %s: %%w`, err)", errPrefix, call.Name(true))
		o.L(`}`)
		o.R("\n")
	}
	o.L(`trace := call.trace`)
	o.L(`if trace == nil {`)
	o.L(`trace = call.client.trace`)
	o.L(`}`)
	o.L(`u := call.makeURL()`)
	o.L(`if trace != nil {`)
	o.L(`fmt.Fprintf(trace, "trace: client sending call request to %%q\n", u)`)
	o.L(`}`)

	if rstype == "" {
		o.LL(`req, err := http.NewRequestWithContext(ctx, %s, u, nil)`, call.String(`http_method`))
	} else {
		if jsonPayload {
			o.LL(`var body bytes.Buffer`)
			o.L(`if err := json.NewEncoder(&body).Encode(payload); err != nil {`)
			o.L("return %sfmt.Errorf(`failed to encode call request: %%w`, err)", errPrefix)
			o.L(`}`)
			o.LL(`req, err := http.NewRequestWithContext(ctx, %s, u, &body)`, call.String(`http_method`))
		} else {
			o.LL(`var vals url.Values`)
			o.L(`m := make(map[string]interface{})`)
			o.L(`if err := payload.AsMap(m); err != nil {`)
			o.L("return %sfmt.Errorf(`failed to convert resource into map: %%w`, err)", errPrefix)
			o.L(`}`)
			o.L(`if len(m) > 0 {`)
			o.L(`vals = make(url.Values)`)
			o.L(`for key, value := range m {`)
			// HACK: this needs to be fixed
			o.L(`switch value := value.(type) {`)
			o.L(`case []string:`)
			o.L(`vals.Add(key, strings.Join(value, ","))`)
			o.L(`default:`)
			// TODO: this is over simplified
			o.L("vals.Add(key, fmt.Sprintf(`%%s`, value))")
			o.L(`}`)
			o.L(`}`)
			o.L(`}`)

			o.L(`if enc := vals.Encode(); len(enc) > 0 {`)
			o.L(`u = u + "?"+ vals.Encode()`)
			o.L(`}`)
			o.L(`req, err := http.NewRequestWithContext(ctx, %s, u, nil)`, call.String(`http_method`))
		}
	}

	o.L(`if err != nil {`)
	o.L("return %sfmt.Errorf(`failed to create new HTTP request: %%w`, err)", errPrefix)
	o.L(`}`)
	if jsonPayload {
		o.LL("req.Header.Set(`Content-Type`, `application/scim+json`)")
	}
	o.L("req.Header.Set(`Accept`, `application/scim+json`)")

	o.LL(`if trace != nil {`)
	o.L(`buf, _ := httputil.DumpRequestOut(req, true)`)
	o.L(`fmt.Fprintf(trace, "%%s\n", buf)`)
	o.L(`}`)

	o.LL(`res, err := call.client.httpcl.Do(req)`)
	o.L(`if err != nil {`)
	o.L("return %sfmt.Errorf(`failed to send request to %%q: %%w`, u, err)", errPrefix)
	o.L(`}`)
	o.L(`if trace != nil {`)
	o.L(`buf, _ := httputil.DumpResponse(res, true)`)
	o.L(`fmt.Fprintf(trace, "%%s\n", buf)`)
	o.L(`}`)
	o.L(`defer res.Body.Close()`)

	// This is fugly, but for special cases we don't parse the result
	if strings.HasPrefix(call.Name(true), "Patch") {
		o.LL(`if res.StatusCode == http.StatusNoContent {`)
		o.L(`//nolint:nilnil`)
		o.L(`return %snil`, errPrefix) // we end up retrieving nil, nil
		o.L(`}`)
	}

	successStatus := call.String(`successStatus`)
	if successStatus == "" {
		successStatus = `http.StatusOK`
	}
	o.LL(`if res.StatusCode != %s {`, successStatus)
	// First, attempt to decode the body into a resource.Error type
	o.L(`var serr resource.Error`)
	o.L(`var resBody bytes.Buffer`)
	o.L(`if err := json.NewDecoder(io.TeeReader(res.Body, &resBody)).Decode(&serr); err != nil {`)
	// If unmarshaling into an error type is not possible, then
	// ...well, return the payload as part of the error so the
	// user gets some hints
	o.L(`return %sfmt.Errorf("expected %%d (got %%d): %%s", %s, res.StatusCode, resBody.String())`, errPrefix, successStatus)
	o.L(`}`)
	o.L("return %s&serr", errPrefix)
	o.L(`}`)

	if resType == "none" {
		o.LL(`return nil`)
	} else {
		o.LL(`var respayload %s`, resType)
		o.L(`if err := json.NewDecoder(res.Body).Decode(&respayload); err != nil {`)
		o.L("return nil, fmt.Errorf(`failed to decode call response: %%w`, err)")
		o.L(`}`)
		o.LL(`return &respayload, nil`)
	}
	o.L(`}`)
	return nil
}
