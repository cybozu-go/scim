// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EmailsColumns holds the columns for the "emails" table.
	EmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "display", Type: field.TypeString, Nullable: true},
		{Name: "primary", Type: field.TypeBool, Default: false},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "value", Type: field.TypeString},
		{Name: "user_emails", Type: field.TypeUUID, Nullable: true},
	}
	// EmailsTable holds the schema information for the "emails" table.
	EmailsTable = &schema.Table{
		Name:       "emails",
		Columns:    EmailsColumns,
		PrimaryKey: []*schema.Column{EmailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "emails_users_emails",
				Columns:    []*schema.Column{EmailsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "display_name", Type: field.TypeString},
		{Name: "external_id", Type: field.TypeString},
		{Name: "group_children", Type: field.TypeUUID, Nullable: true},
		{Name: "user_groups", Type: field.TypeUUID, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_groups_children",
				Columns:    []*schema.Column{GroupsColumns[3]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "groups_users_groups",
				Columns:    []*schema.Column{GroupsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NamesColumns holds the columns for the "names" table.
	NamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "family_name", Type: field.TypeString, Nullable: true},
		{Name: "formatted", Type: field.TypeString, Nullable: true},
		{Name: "given_name", Type: field.TypeString, Nullable: true},
		{Name: "honorific_prefix", Type: field.TypeString, Nullable: true},
		{Name: "honorific_suffix", Type: field.TypeString, Nullable: true},
		{Name: "middle_name", Type: field.TypeString, Nullable: true},
		{Name: "user_names", Type: field.TypeUUID, Nullable: true},
	}
	// NamesTable holds the schema information for the "names" table.
	NamesTable = &schema.Table{
		Name:       "names",
		Columns:    NamesColumns,
		PrimaryKey: []*schema.Column{NamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "names_users_names",
				Columns:    []*schema.Column{NamesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "display_name", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "external_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "password", Type: field.TypeString, Size: 2147483647},
		{Name: "preferred_language", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "locale", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "timezone", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "user_type", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "user_name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "group_users", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_groups_users",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EmailsTable,
		GroupsTable,
		NamesTable,
		UsersTable,
	}
)

func init() {
	EmailsTable.ForeignKeys[0].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupsTable.ForeignKeys[1].RefTable = UsersTable
	NamesTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
}