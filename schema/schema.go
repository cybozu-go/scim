//go:generate ../tools/cmd/genschema.sh

// Package schema contains the schema for SCIM resources.
// You can only query top-level schema, but you can drill down
// into each sub-attributes as necessary.
package schema

import (
	"github.com/cybozu-go/scim/resource"
)

var schemaByType = make(map[string]*resource.Schema)
var schemaByURI = make(map[string]*resource.Schema)

// Registers a system schema so that it can be queried by clients
func Register(schema *resource.Schema) {
	schemaByType[schema.Name()] = schema
	if uri := schema.ID(); uri != "" {
		schemaByURI[uri] = schema
	}
}

// Get returns a schema by its schema URI
func Get(s string) (*resource.Schema, bool) {
	schema, ok := schemaByURI[s]
	return schema, ok
}

// GetByResourceType returns a schema by the associated type name (e.g. `User`, `Group`, `EnterpriseUser`, etc)
func GetByResourceType(s string) (*resource.Schema, bool) {
	schema, ok := schemaByType[s]
	return schema, ok
}

// All returns a list of all schemas that are registered
func All() []*resource.Schema {
	list := make([]*resource.Schema, 0, len(schemaByType))
	for _, s := range schemaByType {
		list = append(list, s)
	}
	return list
}
