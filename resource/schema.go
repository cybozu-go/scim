package resource

// AttributeByName fetches a schema attribute by name.
//
// If an attribute with the given name does not exist,
// the second return value is false.
//
// Sub-attributes can also be specified by concatenating
// the field names with a dot ('.'), for example `members.value`
func (v *Schema) AttributeByName(name string) (*SchemaAttribute, bool) {
	// resources are basically immutable, so we can safely cache this result
	v.attrByNameInitOnce.Do(v.populateAttrByName)

	attr, ok := v.attrByName[name]
	return attr, ok
}

func (v *Schema) populateAttrByName() {
	v.attrByName = make(map[string]*SchemaAttribute)
	for _, attr := range v.Attributes() {
		v.attrByName[attr.Name()] = attr
	}
}
