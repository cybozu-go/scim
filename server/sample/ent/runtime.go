// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/cybozu-go/scim/server/sample/ent/email"
	"github.com/cybozu-go/scim/server/sample/ent/group"
	"github.com/cybozu-go/scim/server/sample/ent/schema"
	"github.com/cybozu-go/scim/server/sample/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	emailFields := schema.Email{}.Fields()
	_ = emailFields
	// emailDescPrimary is the schema descriptor for primary field.
	emailDescPrimary := emailFields[1].Descriptor()
	// email.DefaultPrimary holds the default value on creation for the primary field.
	email.DefaultPrimary = emailDescPrimary.Default.(bool)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescDisplayName is the schema descriptor for displayName field.
	groupDescDisplayName := groupFields[1].Descriptor()
	// group.DisplayNameValidator is a validator for the "displayName" field. It is called by the builders before save.
	group.DisplayNameValidator = groupDescDisplayName.Validators[0].(func(string) error)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[0].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescUserName is the schema descriptor for userName field.
	userDescUserName := userFields[8].Descriptor()
	// user.UserNameValidator is a validator for the "userName" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[2].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
