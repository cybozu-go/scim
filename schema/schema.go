//go:generate ../tools/cmd/genschema.sh

package schema

import "github.com/lestrrat-go/scim/resource"

var schemas = make(map[string]*resource.Schema)

func Register(s string, schema *resource.Schema) {
	schemas[s] = schema
}

func Get(s string) (*resource.Schema, bool) {
	schema, ok := schemas[s]
	return schema, ok
}
