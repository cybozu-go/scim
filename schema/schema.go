//go:generate ../tools/cmd/genschema.sh

package schema

import "github.com/cybozu-go/scim/resource"

var schemas = make(map[string]*resource.Schema)

func Register(s string, schema *resource.Schema) {
	schemas[s] = schema
}

func Get(s string) (*resource.Schema, bool) {
	schema, ok := schemas[s]
	return schema, ok
}

func All() []*resource.Schema {
	list := make([]*resource.Schema, 0, len(schemas))
	for _, s := range schemas {
		list = append(list, s)
	}
	return list
}
