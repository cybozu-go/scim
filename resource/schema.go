package resource

// AttributeByName fetches a schema attribute by its JSON field name.
// (i.e. you must use `$ref` instead of `Reference`, `name` instead of
// `Name`, etc)
//
// If an attribute with the given name does not exist,
// the second return value is false.
//
// Sub-attributes can also be specified by concatenating
// the field names with a dot ('.'), for example `members.value`
func (v *Schema) AttributeByName(name string) (*SchemaAttribute, bool) {
	for _, attr := range v.Attributes() {
		if attr.Name() == name {
			return attr, true
		}
	}

	return nil, false
}
