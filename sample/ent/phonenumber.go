// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/sample/ent/phonenumber"
	"github.com/cybozu-go/scim/sample/ent/user"
	"github.com/google/uuid"
)

// PhoneNumber is the model entity for the PhoneNumber schema.
type PhoneNumber struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Display holds the value of the "display" field.
	Display string `json:"display,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Primary holds the value of the "primary" field.
	Primary bool `json:"primary,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PhoneNumberQuery when eager-loading is set.
	Edges              PhoneNumberEdges `json:"edges"`
	phone_number_user  *uuid.UUID
	user_phone_numbers *uuid.UUID
}

// PhoneNumberEdges holds the relations/edges for other nodes in the graph.
type PhoneNumberEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhoneNumberEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PhoneNumber) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case phonenumber.FieldPrimary:
			values[i] = new(sql.NullBool)
		case phonenumber.FieldID:
			values[i] = new(sql.NullInt64)
		case phonenumber.FieldValue, phonenumber.FieldDisplay, phonenumber.FieldType:
			values[i] = new(sql.NullString)
		case phonenumber.ForeignKeys[0]: // phone_number_user
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case phonenumber.ForeignKeys[1]: // user_phone_numbers
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type PhoneNumber", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PhoneNumber fields.
func (pn *PhoneNumber) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case phonenumber.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pn.ID = int(value.Int64)
		case phonenumber.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				pn.Value = value.String
			}
		case phonenumber.FieldDisplay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display", values[i])
			} else if value.Valid {
				pn.Display = value.String
			}
		case phonenumber.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pn.Type = value.String
			}
		case phonenumber.FieldPrimary:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field primary", values[i])
			} else if value.Valid {
				pn.Primary = value.Bool
			}
		case phonenumber.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number_user", values[i])
			} else if value.Valid {
				pn.phone_number_user = new(uuid.UUID)
				*pn.phone_number_user = *value.S.(*uuid.UUID)
			}
		case phonenumber.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_phone_numbers", values[i])
			} else if value.Valid {
				pn.user_phone_numbers = new(uuid.UUID)
				*pn.user_phone_numbers = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the PhoneNumber entity.
func (pn *PhoneNumber) QueryUser() *UserQuery {
	return (&PhoneNumberClient{config: pn.config}).QueryUser(pn)
}

// Update returns a builder for updating this PhoneNumber.
// Note that you need to call PhoneNumber.Unwrap() before calling this method if this PhoneNumber
// was returned from a transaction, and the transaction was committed or rolled back.
func (pn *PhoneNumber) Update() *PhoneNumberUpdateOne {
	return (&PhoneNumberClient{config: pn.config}).UpdateOne(pn)
}

// Unwrap unwraps the PhoneNumber entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pn *PhoneNumber) Unwrap() *PhoneNumber {
	tx, ok := pn.config.driver.(*txDriver)
	if !ok {
		panic("ent: PhoneNumber is not a transactional entity")
	}
	pn.config.driver = tx.drv
	return pn
}

// String implements the fmt.Stringer.
func (pn *PhoneNumber) String() string {
	var builder strings.Builder
	builder.WriteString("PhoneNumber(")
	builder.WriteString(fmt.Sprintf("id=%v", pn.ID))
	builder.WriteString(", value=")
	builder.WriteString(pn.Value)
	builder.WriteString(", display=")
	builder.WriteString(pn.Display)
	builder.WriteString(", type=")
	builder.WriteString(pn.Type)
	builder.WriteString(", primary=")
	builder.WriteString(fmt.Sprintf("%v", pn.Primary))
	builder.WriteByte(')')
	return builder.String()
}

// PhoneNumbers is a parsable slice of PhoneNumber.
type PhoneNumbers []*PhoneNumber

func (pn PhoneNumbers) config(cfg config) {
	for _i := range pn {
		pn[_i].config = cfg
	}
}