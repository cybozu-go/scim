package sample

import (
	"fmt"
	"reflect"

	"entgo.io/ent/dialect/sql"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/email"
	"github.com/cybozu-go/scim/sample/ent/predicate"
	"github.com/cybozu-go/scim/sample/ent/role"
	"github.com/cybozu-go/scim/sample/ent/user"
)

func userLoadEntFields(q *ent.UserQuery, fields []string) {
	if len(fields) == 0 {
		fields = []string{resource.UserActiveKey, resource.UserAddressesKey, resource.UserDisplayNameKey, resource.UserEmailsKey, resource.UserEntitlementsKey, resource.UserExternalIDKey, resource.UserGroupsKey, resource.UserIDKey, resource.UserIMSKey, resource.UserLocaleKey, resource.UserNameKey, resource.UserNickNameKey, resource.UserPasswordKey, resource.UserPhoneNumbersKey, resource.UserPreferredLanguageKey, resource.UserProfileURLKey, resource.UserRolesKey, resource.UserTimezoneKey, resource.UserTitleKey, resource.UserUserNameKey, resource.UserUserTypeKey, resource.UserX509CertificatesKey}
	}
	selectNames := make([]string, 0, len(fields))
	for _, f := range fields {
		switch f {
		case resource.UserActiveKey:
			selectNames = append(selectNames, user.FieldActive)
		case resource.UserAddressesKey:
		case resource.UserDisplayNameKey:
			selectNames = append(selectNames, user.FieldDisplayName)
		case resource.UserEmailsKey:
			q.WithEmails()
		case resource.UserEntitlementsKey:
		case resource.UserExternalIDKey:
			selectNames = append(selectNames, user.FieldExternalID)
		case resource.UserGroupsKey:
		case resource.UserIDKey:
			selectNames = append(selectNames, user.FieldID)
		case resource.UserIMSKey:
		case resource.UserLocaleKey:
			selectNames = append(selectNames, user.FieldLocale)
		case resource.UserMetaKey:
		case resource.UserNameKey:
			q.WithName()
		case resource.UserNickNameKey:
			selectNames = append(selectNames, user.FieldNickName)
		case resource.UserPasswordKey:
			selectNames = append(selectNames, user.FieldPassword)
		case resource.UserPhoneNumbersKey:
		case resource.UserPreferredLanguageKey:
			selectNames = append(selectNames, user.FieldPreferredLanguage)
		case resource.UserProfileURLKey:
			selectNames = append(selectNames, user.FieldProfileURL)
		case resource.UserRolesKey:
		case resource.UserTimezoneKey:
			selectNames = append(selectNames, user.FieldTimezone)
		case resource.UserTitleKey:
			selectNames = append(selectNames, user.FieldTitle)
		case resource.UserUserNameKey:
			selectNames = append(selectNames, user.FieldUserName)
		case resource.UserUserTypeKey:
			selectNames = append(selectNames, user.FieldUserType)
		case resource.UserX509CertificatesKey:
		}
	}
	q.Select(selectNames...)
}

func userLocation(id string) string {
	return "https://foobar.com/scim/v2/Users/" + id
}

func UserResourceFromEnt(in *ent.User) (*resource.User, error) {
	var b resource.Builder

	builder := b.User()

	meta, err := b.Meta().
		ResourceType("User").
		Location(userLocation(in.ID.String())).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build meta information for User")
	}

	builder.
		Meta(meta)

	if el := len(in.Edges.Emails); el > 0 {
		list := make([]*resource.Email, 0, el)
		for _, ine := range in.Edges.Emails {
			r, err := EmailResourceFromEnt(ine)
			if err != nil {
				return nil, fmt.Errorf("failed to build emails information for User")
			}
			list = append(list, r)
		}
		builder.Emails(list...)
	}

	if el := len(in.Edges.Name); el > 0 {
		list := make([]*resource.Names, 0, el)
		for _, ine := range in.Edges.Name {
			r, err := NamesResourceFromEnt(ine)
			if err != nil {
				return nil, fmt.Errorf("failed to build name information for User")
			}
			list = append(list, r)
		}
		builder.Name(list[0])
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

func UserEntFieldFromSCIM(s string) string {
	switch s {
	case resource.UserActiveKey:
		return user.FieldActive
	case resource.UserDisplayNameKey:
		return user.FieldDisplayName
	case resource.UserExternalIDKey:
		return user.FieldExternalID
	case resource.UserIDKey:
		return user.FieldID
	case resource.UserLocaleKey:
		return user.FieldLocale
	case resource.UserNickNameKey:
		return user.FieldNickName
	case resource.UserPasswordKey:
		return user.FieldPassword
	case resource.UserPreferredLanguageKey:
		return user.FieldPreferredLanguage
	case resource.UserProfileURLKey:
		return user.FieldProfileURL
	case resource.UserTimezoneKey:
		return user.FieldTimezone
	case resource.UserTitleKey:
		return user.FieldTitle
	case resource.UserUserNameKey:
		return user.FieldUserName
	case resource.UserUserTypeKey:
		return user.FieldUserType
	default:
		return s
	}
}

func userStartsWithPredicate(q *ent.UserQuery, scimField string, val interface{}) (predicate.User, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.UserDisplayNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserEmailsKey:
		switch subfield {
		case resource.EmailDisplayKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.DisplayEQ(val.(string))), nil
		case resource.EmailPrimaryKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.PrimaryEQ(val.(bool))), nil
		case resource.EmailTypeKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.TypeEQ(val.(string))), nil
		case resource.EmailValueKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserExternalIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserLocaleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserNickNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPasswordKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPreferredLanguageKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserProfileURLKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserRolesKey:
		switch subfield {
		case resource.RoleDisplayKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.DisplayEQ(val.(string))), nil
		case resource.RolePrimaryKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.PrimaryEQ(val.(bool))), nil
		case resource.RoleTypeKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.TypeEQ(val.(string))), nil
		case resource.RoleValueKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserTimezoneKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserTitleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserTypeKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasPrefix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func userEndsWithPredicate(q *ent.UserQuery, scimField string, val interface{}) (predicate.User, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.UserDisplayNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserEmailsKey:
		switch subfield {
		case resource.EmailDisplayKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.DisplayEQ(val.(string))), nil
		case resource.EmailPrimaryKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.PrimaryEQ(val.(bool))), nil
		case resource.EmailTypeKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.TypeEQ(val.(string))), nil
		case resource.EmailValueKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserExternalIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserLocaleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserNickNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPasswordKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPreferredLanguageKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserProfileURLKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserRolesKey:
		switch subfield {
		case resource.RoleDisplayKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.DisplayEQ(val.(string))), nil
		case resource.RolePrimaryKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.PrimaryEQ(val.(bool))), nil
		case resource.RoleTypeKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.TypeEQ(val.(string))), nil
		case resource.RoleValueKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserTimezoneKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserTitleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserTypeKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.HasSuffix(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func userContainsPredicate(q *ent.UserQuery, scimField string, val interface{}) (predicate.User, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.UserDisplayNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserEmailsKey:
		switch subfield {
		case resource.EmailDisplayKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.DisplayEQ(val.(string))), nil
		case resource.EmailPrimaryKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.PrimaryEQ(val.(bool))), nil
		case resource.EmailTypeKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.TypeEQ(val.(string))), nil
		case resource.EmailValueKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserExternalIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserLocaleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserNickNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPasswordKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPreferredLanguageKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserProfileURLKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserRolesKey:
		switch subfield {
		case resource.RoleDisplayKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.DisplayEQ(val.(string))), nil
		case resource.RolePrimaryKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.PrimaryEQ(val.(bool))), nil
		case resource.RoleTypeKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.TypeEQ(val.(string))), nil
		case resource.RoleValueKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserTimezoneKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserTitleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserTypeKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.Contains(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func userEqualsPredicate(q *ent.UserQuery, scimField string, val interface{}) (predicate.User, error) {
	_ = q
	field, subfield, err := splitScimField(scimField)
	if err != nil {
		return nil, err
	}
	_ = subfield // TODO: remove later
	switch field {
	case resource.UserDisplayNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserEmailsKey:
		switch subfield {
		case resource.EmailDisplayKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.DisplayEQ(val.(string))), nil
		case resource.EmailPrimaryKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.PrimaryEQ(val.(bool))), nil
		case resource.EmailTypeKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.TypeEQ(val.(string))), nil
		case resource.EmailValueKey:
			//nolint:forcetypeassert
			return user.HasEmailsWith(email.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserExternalIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserIDKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserLocaleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserNickNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPasswordKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserPreferredLanguageKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserProfileURLKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserRolesKey:
		switch subfield {
		case resource.RoleDisplayKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.DisplayEQ(val.(string))), nil
		case resource.RolePrimaryKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.PrimaryEQ(val.(bool))), nil
		case resource.RoleTypeKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.TypeEQ(val.(string))), nil
		case resource.RoleValueKey:
			//nolint:forcetypeassert
			return user.HasRolesWith(role.ValueEQ(val.(string))), nil
		default:
			return nil, fmt.Errorf("invalid filter specification: invalid subfield for %q", field)
		}
	case resource.UserTimezoneKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserTitleKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserNameKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	case resource.UserUserTypeKey:
		entFieldName := UserEntFieldFromSCIM(scimField)
		return predicate.User(func(s *sql.Selector) {
			//nolint:forcetypeassert
			s.Where(sql.EQ(s.C(entFieldName), val.(string)))
		}), nil
	default:
		return nil, fmt.Errorf("invalid filter field specification")
	}
}

func userPresencePredicate(scimField string) predicate.User {
	switch scimField {
	case resource.UserDisplayNameKey:
		return user.And(user.DisplayNameNotNil(), user.DisplayNameNEQ(""))
	case resource.UserExternalIDKey:
		return user.And(user.ExternalIDNotNil(), user.ExternalIDNEQ(""))
	case resource.UserLocaleKey:
		return user.And(user.LocaleNotNil(), user.LocaleNEQ(""))
	case resource.UserNickNameKey:
		return user.And(user.NickNameNotNil(), user.NickNameNEQ(""))
	case resource.UserPasswordKey:
		return user.And(user.PasswordNotNil(), user.PasswordNEQ(""))
	case resource.UserPreferredLanguageKey:
		return user.And(user.PreferredLanguageNotNil(), user.PreferredLanguageNEQ(""))
	case resource.UserProfileURLKey:
		return user.And(user.ProfileURLNotNil(), user.ProfileURLNEQ(""))
	case resource.UserTimezoneKey:
		return user.And(user.TimezoneNotNil(), user.TimezoneNEQ(""))
	case resource.UserTitleKey:
		return user.And(user.TitleNotNil(), user.TitleNEQ(""))
	case resource.UserUserTypeKey:
		return user.And(user.UserTypeNotNil(), user.UserTypeNEQ(""))
	default:
		return nil
	}
}
