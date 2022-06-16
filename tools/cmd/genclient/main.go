package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
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
	var resourcesFile = flag.String("resources", "resources.yml", "")
	flag.Parse()

	resourcesSrc, err := yaml2json(*resourcesFile)
	if err != nil {
		return err
	}

	callsSrc, err := yaml2json(*callsFile)
	if err != nil {
		return err
	}

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
	}

	var calls struct {
		Services []Service `json:"services"`
	}
	if err := json.NewDecoder(bytes.NewReader(callsSrc)).Decode(&calls); err != nil {
		return fmt.Errorf(`failed to decode %q: %w`, *callsFile, err)
	}

	for _, service := range calls.Services {
		if err := generateService(service, resources); err != nil {
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

func generateService(svc Service, resources map[string]*codegen.Object) error {
	fmt.Printf("  ⌛ Generating %s...\n", svc.Name)

	var buf bytes.Buffer
	o := codegen.NewOutput(&buf)
	o.L(`package client`)

	if desc := svc.Description; desc != "" {
		o.LL(`// %s`, strings.TrimSpace(desc))
	} else {
		o.R("\n")
	}
	o.L(`type %s struct {`, svc.Name)
	o.L(`client *Client`)
	o.L(`}`)

	o.LL(`func (client *Client) %s() *%s {`, strings.TrimSuffix(svc.Name, `Service`), svc.Name)
	o.L(`return &%s{`, svc.Name)
	o.L(`client: client,`)
	o.L(`}`)
	o.L(`}`)

	for _, call := range svc.Calls {
		fmt.Printf("    ⌛ Call %s...\n", call.Name(true))

		if err := generateCall(o, svc, call, resources); err != nil {
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

func generateCall(o *codegen.Output, svc Service, call *codegen.Object, resources map[string]*codegen.Object) error {
	rstype := call.String(`resource`)

	o.LL(`type %s struct {`, call.Name(true))
	if rstype != "" {
		o.L(`builder *resource.%sBuilder`, call.String(`resource`))
	}
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
		rs, ok := resources[rstype]
		if !ok {
			return fmt.Errorf(`resource %q not found`, rstype)
		}

		if resType == "" {
			resType = `resource.` + rs.Name(true)
		}

		var fields []codegen.Field

		rschema, ok := schema.GetByResourceType(rstype)
		if !ok {
			fields = append(rs.Fields(), optional...)
		} else {
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

			allowed := make(map[string]struct{})
			if len(mutabilities) > 0 {
				attrs := rschema.Attributes()
				// I found out RFC7643 talks about “externalId” field, but it’s not in any of the defined schemas :(
				eid := resource.NewSchemaAttributeBuilder().
					Name(`externalId`).
					Mutability(resource.MutWriteOnly).
					MustBuild()

				for _, attr := range append(attrs, eid) {
					mut := string(attr.Mutability())
					if _, ok := mutabilities[mut]; !ok {
						continue
					}

					allowed[attr.Name()] = struct{}{}
				}
			}

			for _, f := range append(rs.Fields(), optional...) {
				if len(mutabilities) > 0 {
					if _, ok := allowed[f.JSON()]; !ok {
						continue
					}
				}
				fields = append(fields, f)
			}
		}

		for _, field := range fields {
			var typ string
			var isSlice bool
			if IsSlice(field) {
				typ = strings.TrimPrefix(field.Type(), `[]`)
				isSlice = true
			} else {
				typ = field.Type()
			}

			var hasPtrPrefix bool
			if strings.HasPrefix(typ, `*`) {
				typ = strings.TrimPrefix(typ, `*`)
				hasPtrPrefix = true
			}

			if _, ok := resources[typ]; ok {
				typ = `resource.` + typ
			}
			if hasPtrPrefix {
				typ = `*` + typ
			}

			// If the argument is a slice, the parameter type should be varg
			if isSlice {
				o.LL(`func (call *%[1]s) %[2]s(v ...%[3]s) *%[1]s {`, call.Name(true), field.Name(true), typ)
				o.L(`call.builder.%s(v...)`, field.Name(true))
			} else {
				o.LL(`func (call *%[1]s) %[2]s(v %[3]s) *%[1]s {`, call.Name(true), field.Name(true), typ)
				o.L(`call.builder.%s(v)`, field.Name(true))
			}
			o.L(`return call`)
			o.L(`}`)
		}

		if jsonPayload {
			o.LL(`// Extension allows users to register an extension using the fully qualified URI`)
			o.L(`func (call *%[1]s) Extension(uri string, value interface{}) *%[1]s {`, call.Name(true))
			o.L(`call.builder.Extension(uri, value)`)
			o.L(`return call`)
			o.L(`}`)

			o.LL(`func (call *%[1]s) Validator(v resource.%[2]sValidator) *%[1]s {`, call.Name(true), rs.Name(true))
			o.L(`call.builder.Validator(v)`)
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

	if rstype != "" {
		o.L(`payload, err := call.builder.Build()`)
		o.L(`if err != nil {`)
		o.L("return %sfmt.Errorf(`failed to generate request payload for %s: %%w`, err)", errPrefix, call.Name(true))
		o.L(`}`)
		o.R("\n")
	}
	o.L(`trace := call.trace`)
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
	o.L(`if trace != nil {`)
	o.L(`buf, _ := httputil.DumpResponse(res, true)`)
	o.L(`fmt.Fprintf(trace, "%%s\n", buf)`)
	o.L(`}`)
	o.L(`if err != nil {`)
	o.L("return %sfmt.Errorf(`failed to send request to %%q: %%w`, u, err)", errPrefix)
	o.L(`}`)
	o.L(`defer res.Body.Close()`)

	successStatus := call.String(`successStatus`)
	if successStatus == "" {
		successStatus = `http.StatusOK`
	}
	o.LL(`if res.StatusCode != %s {`, successStatus)
	o.L("return %sfmt.Errorf(`expected call response %%d, got (%%d)`, %s, res.StatusCode)", errPrefix, successStatus)
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
