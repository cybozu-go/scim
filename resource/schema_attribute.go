package resource

// CanWrite returns true if the mutability is either readWrite or writeOnly.
//
// The result is evaluated in context of the SCIM server, from the PoV of the
// SCIM client.
func (v *SchemaAttribute) CanWrite() bool {
	switch v.Mutability() {
	case MutReadWrite, MutWriteOnly:
		return true
	default:
		return false
	}
}

// CanRead returns true if the mutability is either readOnly, readWrite, immutable.
//
// The result is evaluated in context of the SCIM server, from the PoV of the
// SCIM client.
func (v *SchemaAttribute) CanRead() bool {
	switch v.Mutability() {
	case MutReadOnly, MutReadWrite, MutImmutable:
		return true
	default:
		return false
	}
}
