package sample

import (
	"fmt"
	"reflect"

	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/user"
)

func userLoadEntFields(q *ent.UserQuery, fields []string) {
	selectNames := make([]string, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "active":
			selectNames = append(selectNames, user.FieldActive)
		case "addresses":
		case "displayName":
			selectNames = append(selectNames, user.FieldDisplayName)
		case "emails":
			q.WithEmails()
		case "entitlements":
		case "externalID":
			selectNames = append(selectNames, user.FieldExternalID)
		case "groups":
		case "id":
			selectNames = append(selectNames, user.FieldID)
		case "ims":
		case "locale":
			selectNames = append(selectNames, user.FieldLocale)
		case "meta":
		case "name":
		case "nickName":
			selectNames = append(selectNames, user.FieldNickName)
		case "password":
			selectNames = append(selectNames, user.FieldPassword)
		case "phoneNumbers":
		case "preferredLanguage":
			selectNames = append(selectNames, user.FieldPreferredLanguage)
		case "profileURL":
			selectNames = append(selectNames, user.FieldProfileURL)
		case "roles":
		case "timezone":
			selectNames = append(selectNames, user.FieldTimezone)
		case "title":
			selectNames = append(selectNames, user.FieldTitle)
		case "userName":
			selectNames = append(selectNames, user.FieldUserName)
		case "userType":
			selectNames = append(selectNames, user.FieldUserType)
		case "x509Certificates":
		}
	}
	q.Select(selectNames...)
}

func UserResourceFromEnt(in *ent.User) (*resource.User, error) {
	var b resource.Builder

	builder := b.User()

	meta, err := b.Meta().
		ResourceType("User").
		Location("https://foobar.com/scim/v2/User/" + in.ID.String()).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build meta information for User")
	}

	builder.
		Meta(meta)

	if el := len(in.Edges.Emails); el > 0 {
		emails := make([]*resource.Email, 0, el)
		for _, ine := range in.Edges.Emails {
			email, err := EmailResourceFromEnt(ine)
			if err != nil {
				return nil, fmt.Errorf("failed to build email information for User")
			}
			emails = append(emails, email)
		}
		builder.Emails(emails...)
	}
	if !reflect.ValueOf(in.Active).IsZero() {
		builder.Active(in.Active)
	}
	if !reflect.ValueOf(in.DisplayName).IsZero() {
		builder.DisplayName(in.DisplayName)
	}
	if !reflect.ValueOf(in.ExternalID).IsZero() {
		builder.ExternalID(in.ExternalID)
	}
	builder.ID(in.ID.String())
	if !reflect.ValueOf(in.Locale).IsZero() {
		builder.Locale(in.Locale)
	}
	if !reflect.ValueOf(in.Password).IsZero() {
		builder.Password(in.Password)
	}
	if !reflect.ValueOf(in.PreferredLanguage).IsZero() {
		builder.PreferredLanguage(in.PreferredLanguage)
	}
	if !reflect.ValueOf(in.Timezone).IsZero() {
		builder.Timezone(in.Timezone)
	}
	if !reflect.ValueOf(in.UserName).IsZero() {
		builder.UserName(in.UserName)
	}
	if !reflect.ValueOf(in.UserType).IsZero() {
		builder.UserType(in.UserType)
	}
	return builder.Build()
}
