// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/sample/ent/name"
	"github.com/cybozu-go/scim/sample/ent/user"
	"github.com/google/uuid"
)

// Name is the model entity for the Name schema.
type Name struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FamilyName holds the value of the "familyName" field.
	FamilyName string `json:"familyName,omitempty"`
	// Formatted holds the value of the "formatted" field.
	Formatted string `json:"formatted,omitempty"`
	// GivenName holds the value of the "givenName" field.
	GivenName string `json:"givenName,omitempty"`
	// HonorificPrefix holds the value of the "honorificPrefix" field.
	HonorificPrefix string `json:"honorificPrefix,omitempty"`
	// HonorificSuffix holds the value of the "honorificSuffix" field.
	HonorificSuffix string `json:"honorificSuffix,omitempty"`
	// MiddleName holds the value of the "middleName" field.
	MiddleName string `json:"middleName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NameQuery when eager-loading is set.
	Edges      NameEdges `json:"edges"`
	user_names *uuid.UUID
}

// NameEdges holds the relations/edges for other nodes in the graph.
type NameEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NameEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// The edge users was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Name) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case name.FieldID:
			values[i] = new(sql.NullInt64)
		case name.FieldFamilyName, name.FieldFormatted, name.FieldGivenName, name.FieldHonorificPrefix, name.FieldHonorificSuffix, name.FieldMiddleName:
			values[i] = new(sql.NullString)
		case name.ForeignKeys[0]: // user_names
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Name", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Name fields.
func (n *Name) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case name.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case name.FieldFamilyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field familyName", values[i])
			} else if value.Valid {
				n.FamilyName = value.String
			}
		case name.FieldFormatted:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field formatted", values[i])
			} else if value.Valid {
				n.Formatted = value.String
			}
		case name.FieldGivenName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field givenName", values[i])
			} else if value.Valid {
				n.GivenName = value.String
			}
		case name.FieldHonorificPrefix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field honorificPrefix", values[i])
			} else if value.Valid {
				n.HonorificPrefix = value.String
			}
		case name.FieldHonorificSuffix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field honorificSuffix", values[i])
			} else if value.Valid {
				n.HonorificSuffix = value.String
			}
		case name.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middleName", values[i])
			} else if value.Valid {
				n.MiddleName = value.String
			}
		case name.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_names", values[i])
			} else if value.Valid {
				n.user_names = new(uuid.UUID)
				*n.user_names = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Name entity.
func (n *Name) QueryUsers() *UserQuery {
	return (&NameClient{config: n.config}).QueryUsers(n)
}

// Update returns a builder for updating this Name.
// Note that you need to call Name.Unwrap() before calling this method if this Name
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Name) Update() *NameUpdateOne {
	return (&NameClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Name entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Name) Unwrap() *Name {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Name is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Name) String() string {
	var builder strings.Builder
	builder.WriteString("Name(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", familyName=")
	builder.WriteString(n.FamilyName)
	builder.WriteString(", formatted=")
	builder.WriteString(n.Formatted)
	builder.WriteString(", givenName=")
	builder.WriteString(n.GivenName)
	builder.WriteString(", honorificPrefix=")
	builder.WriteString(n.HonorificPrefix)
	builder.WriteString(", honorificSuffix=")
	builder.WriteString(n.HonorificSuffix)
	builder.WriteString(", middleName=")
	builder.WriteString(n.MiddleName)
	builder.WriteByte(')')
	return builder.String()
}

// Names is a parsable slice of Name.
type Names []*Name

func (n Names) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}