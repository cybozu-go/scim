// Code generated by entc, DO NOT EDIT.

package phonenumber

const (
	// Label holds the string label denoting the phonenumber type in the database.
	Label = "phone_number"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldDisplay holds the string denoting the display field in the database.
	FieldDisplay = "display"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPrimary holds the string denoting the primary field in the database.
	FieldPrimary = "primary"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the phonenumber in the database.
	Table = "phone_numbers"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "phone_numbers"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "phone_number_user"
)

// Columns holds all SQL columns for phonenumber fields.
var Columns = []string{
	FieldID,
	FieldValue,
	FieldDisplay,
	FieldType,
	FieldPrimary,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "phone_numbers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"phone_number_user",
	"user_phone_numbers",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}