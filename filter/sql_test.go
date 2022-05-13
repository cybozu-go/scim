package filter_test

import (
	"testing"

	"github.com/lestrrat-go/scim/filter"
	"github.com/stretchr/testify/require"
)

func TestSQL(t *testing.T) {
	testcases := []struct {
		Filter       string
		ExpectedSQL  string
		ExpectedArgs []interface{}
		Error        bool
		SQLOptions   []filter.SQLOption
	}{
		{
			Filter:       `userName eq "bjensen"`,
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("userName" = ?)`,
			ExpectedArgs: []interface{}{"bjensen"},
		},
		{
			Filter: `name.familyName co "O'Malley"`,
			SQLOptions: []filter.SQLOption{
				filter.WithColumnMapper(filter.MapColumnMapper(map[string]string{
					`name.familyName`: `familyName`,
				})),
			},
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("familyName" ~ ?)`,
			ExpectedArgs: []interface{}{`O'Malley`},
		},
		{
			Filter:       `userName sw "J"`,
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("userName" ~ ?)`,
			ExpectedArgs: []interface{}{`^J`},
		},
		{
			Filter:       `title pr`,
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("title" IS NOT NULL)`,
			ExpectedArgs: []interface{}{},
		},
		{
			Filter: `meta.lastModified gt "2011-05-13T04:42:34Z"`,
			SQLOptions: []filter.SQLOption{
				filter.WithColumnMapper(filter.MapColumnMapper(map[string]string{
					`meta.lastModified`: `lastModified`,
				})),
			},
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("lastModified" > ?)`,
			ExpectedArgs: []interface{}{"2011-05-13T04:42:34Z"},
		},
		{
			Filter: `meta.lastModified ge "2011-05-13T04:42:34Z"`,
			SQLOptions: []filter.SQLOption{
				filter.WithColumnMapper(filter.MapColumnMapper(map[string]string{
					`meta.lastModified`: `lastModified`,
				})),
			},
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("lastModified" >= ?)`,
			ExpectedArgs: []interface{}{"2011-05-13T04:42:34Z"},
		},
		{
			Filter: `meta.lastModified lt "2011-05-13T04:42:34Z"`,
			SQLOptions: []filter.SQLOption{
				filter.WithColumnMapper(filter.MapColumnMapper(map[string]string{
					`meta.lastModified`: `lastModified`,
				})),
			},
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("lastModified" < ?)`,
			ExpectedArgs: []interface{}{"2011-05-13T04:42:34Z"},
		},
		{
			Filter: `meta.lastModified le "2011-05-13T04:42:34Z"`,
			SQLOptions: []filter.SQLOption{
				filter.WithColumnMapper(filter.MapColumnMapper(map[string]string{
					`meta.lastModified`: `lastModified`,
				})),
			},
			ExpectedSQL:  `SELECT * FROM "users" WHERE ("lastModified" <= ?)`,
			ExpectedArgs: []interface{}{"2011-05-13T04:42:34Z"},
		},
		{
			Filter:       `title pr and userType eq "Employee"`,
			ExpectedSQL:  `SELECT * FROM "users" WHERE (("title" IS NOT NULL) AND ("userType" = ?))`,
			ExpectedArgs: []interface{}{"Employee"},
		},
		{
			Filter:       `title pr or userType eq "Intern"`,
			ExpectedSQL:  `SELECT * FROM "users" WHERE (("title" IS NOT NULL) OR ("userType" = ?))`,
			ExpectedArgs: []interface{}{"Intern"},
		},
		// Currently not supported:
		// filter=
		//   schemas eq "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"
		// filter=urn:ietf:params:scim:schemas:core:2.0:User:userName sw "J"
		{
			Filter:       `userType eq "Employee" and (emails.type eq "work")`,
			ExpectedSQL:  `SELECT * FROM "users" INNER JOIN "emails" ON ("emails"."users_id" = "users"."id") WHERE (("userType" = ?) AND ("emails"."type" = ?))`,
			ExpectedArgs: []interface{}{"Employee", "work"},
		},
		{
			Filter:       `userType eq "Employee" and emails[type eq "work" and value co "@example.com"]`,
			ExpectedSQL:  `SELECT * FROM "users" INNER JOIN "emails" ON ("emails"."users_id" = "users"."id") WHERE (("userType" = ?) AND (("emails"."type" = ?) AND ("emails"."value" ~ ?)))`,
			ExpectedArgs: []interface{}{"Employee", "work", "@example\\.com"},
		},
		{
			// Filter: `userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
			Filter:       `userType ne "Employee" and not (emails.value co "example.com" or emails.value co "example.org")`,
			ExpectedSQL:  `SELECT * FROM "users" INNER JOIN "emails" ON ("emails"."users_id" = "users"."id") INNER JOIN "emails" ON ("emails"."users_id" = "users"."id") WHERE (("userType" != ?) AND NOT(((("emails"."value" ~ ?) OR ("emails"."value" ~ ?)))))`,
			ExpectedArgs: []interface{}{"Employee", "example\\.com", "example\\.org"},
		},
		{
			Filter:       `emails[type eq "work" and value co "@example.com"] or ims[type eq "xmpp" and value co "@foo.com"]`,
			ExpectedSQL:  `SELECT * FROM "users" INNER JOIN "emails" ON ("emails"."users_id" = "users"."id") INNER JOIN "ims" ON ("ims"."users_id" = "users"."id") WHERE ((("emails"."type" = ?) AND ("emails"."value" ~ ?)) OR (("ims"."type" = ?) AND ("ims"."value" ~ ?)))`,
			ExpectedArgs: []interface{}{"work", "@example\\.com", "xmpp", "@foo\\.com"},
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.Filter, func(t *testing.T) {
			sqlstmt, args, err := filter.SQL(tc.Filter, "users", tc.SQLOptions...)
			if tc.Error {
				require.Error(t, err, `filter.SQL should fail`)
			} else {
				require.NoError(t, err, `filter.SQL should succeed`)
				require.Equal(t, tc.ExpectedSQL, sqlstmt)
				require.Equal(t, tc.ExpectedArgs, args)
			}
		})
	}
}
