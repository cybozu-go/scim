// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/server/sample/ent/user"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// ExternalID holds the value of the "externalID" field.
	ExternalID string `json:"externalID,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// PreferredLanguage holds the value of the "preferredLanguage" field.
	PreferredLanguage string `json:"preferredLanguage,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale string `json:"locale,omitempty"`
	// Timezone holds the value of the "timezone" field.
	Timezone string `json:"timezone,omitempty"`
	// UserType holds the value of the "userType" field.
	UserType string `json:"userType,omitempty"`
	// UserName holds the value of the "userName" field.
	UserName string `json:"userName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges       UserEdges `json:"edges"`
	group_users *uuid.UUID
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Emails holds the value of the emails edge.
	Emails []*Email `json:"emails,omitempty"`
	// Names holds the value of the names edge.
	Names []*Name `json:"names,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// EmailsOrErr returns the Emails value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EmailsOrErr() ([]*Email, error) {
	if e.loadedTypes[1] {
		return e.Emails, nil
	}
	return nil, &NotLoadedError{edge: "emails"}
}

// NamesOrErr returns the Names value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NamesOrErr() ([]*Name, error) {
	if e.loadedTypes[2] {
		return e.Names, nil
	}
	return nil, &NotLoadedError{edge: "names"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldActive:
			values[i] = new(sql.NullBool)
		case user.FieldExternalID, user.FieldPassword, user.FieldPreferredLanguage, user.FieldLocale, user.FieldTimezone, user.FieldUserType, user.FieldUserName:
			values[i] = new(sql.NullString)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		case user.ForeignKeys[0]: // group_users
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				u.Active = value.Bool
			}
		case user.FieldExternalID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field externalID", values[i])
			} else if value.Valid {
				u.ExternalID = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldPreferredLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preferredLanguage", values[i])
			} else if value.Valid {
				u.PreferredLanguage = value.String
			}
		case user.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				u.Locale = value.String
			}
		case user.FieldTimezone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timezone", values[i])
			} else if value.Valid {
				u.Timezone = value.String
			}
		case user.FieldUserType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userType", values[i])
			} else if value.Valid {
				u.UserType = value.String
			}
		case user.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userName", values[i])
			} else if value.Valid {
				u.UserName = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_users", values[i])
			} else if value.Valid {
				u.group_users = new(uuid.UUID)
				*u.group_users = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return (&UserClient{config: u.config}).QueryGroups(u)
}

// QueryEmails queries the "emails" edge of the User entity.
func (u *User) QueryEmails() *EmailQuery {
	return (&UserClient{config: u.config}).QueryEmails(u)
}

// QueryNames queries the "names" edge of the User entity.
func (u *User) QueryNames() *NameQuery {
	return (&UserClient{config: u.config}).QueryNames(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", u.Active))
	builder.WriteString(", externalID=")
	builder.WriteString(u.ExternalID)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", preferredLanguage=")
	builder.WriteString(u.PreferredLanguage)
	builder.WriteString(", locale=")
	builder.WriteString(u.Locale)
	builder.WriteString(", timezone=")
	builder.WriteString(u.Timezone)
	builder.WriteString(", userType=")
	builder.WriteString(u.UserType)
	builder.WriteString(", userName=")
	builder.WriteString(u.UserName)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
