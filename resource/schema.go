package resource

// AttributeByName fetches a schema attribute by name.
//
// If an attribute with the given name does not exist,
// the second return value is false.
//
// Sub-attributes can also be specified by concatenating
// the field names with a dot ('.'), for example `members.value`
func (s *Schema) AttributeByName(name string) (*SchemaAttribute, bool) {
	// resources are basically immutable, so we can safely cache this result
	s.attrByNameInitOnce.Do(s.populateAttrByName)

	attr, ok := s.attrByName[name]
	return attr, ok
}

func (s *Schema) populateAttrByName() {
	s.attrByName = make(map[string]*SchemaAttribute)
	for _, attr := range s.Attributes() {
		s.attrByName[attr.Name()] = attr
	}
}
