//go:generate go generate ./ent
//go:generate perl autofix.pl

package sample

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"entgo.io/ent/dialect"
	"github.com/cybozu-go/scim/filter"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/sample/ent"
	"github.com/cybozu-go/scim/sample/ent/group"
	"github.com/cybozu-go/scim/sample/ent/predicate"
	"github.com/cybozu-go/scim/sample/ent/user"
	"github.com/google/uuid"
	"golang.org/x/text/secure/precis"

	// default driver
	_ "github.com/mattn/go-sqlite3"
)

//FIXME
var _ = emailPresencePredicate
var _ = groupPresencePredicate
var _ = namesPresencePredicate

type Backend struct {
	db *ent.Client
}

func New(connspec string) (*Backend, error) {
	client, err := ent.Open(dialect.SQLite, connspec)
	if err != nil {
		return nil, fmt.Errorf(`failed to open database: %w`, err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf(`failed to create schema resources: %w`, err)
	}

	return &Backend{
		db: client,
	}, nil
}

func (b *Backend) Close() error {
	return b.db.Close()
}

var chars []byte
var maxchars *big.Int

func init() {
	charmap := make(map[byte]struct{})
	for i := 0x21; i < 0x7E; i++ {
		charmap[byte(i)] = struct{}{}
	}
	delete(charmap, 'I')
	delete(charmap, 'O')
	delete(charmap, '\\')
	delete(charmap, 'l')
	delete(charmap, 'o')
	for c := range charmap {
		chars = append(chars, c)
	}
	maxchars = big.NewInt(int64(len(chars)))
}

func randomString(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		bn, err := rand.Int(rand.Reader, maxchars)
		if err != nil {
			panic(err)
		}
		b.WriteByte(chars[int(bn.Int64())])
	}
	return b.String()
}

func (b *Backend) createEmails(in *resource.User, modifyUserCall interface{}) ([]*resource.Email, error) {
	emails := make([]*resource.Email, 0, len(in.Emails()))
	var hasPrimary bool
	for _, v := range in.Emails() {
		emailCreateCall := b.db.Email.Create().
			SetValue(v.Value())
		if v.HasDisplay() {
			emailCreateCall.SetDisplay(v.Display())
		}
		if v.HasType() {
			emailCreateCall.SetType(v.Type())
		}
		if sv := v.Primary(); sv {
			if hasPrimary {
				return nil, fmt.Errorf(`invalid user.emails: multiple emails have been set to primary`)
			}
			emailCreateCall.SetPrimary(sv)
		}
		email, err := emailCreateCall.Save(context.TODO())
		if err != nil {
			return nil, fmt.Errorf(`failed to create user: failed to create email: %w`, err)
		}

		switch modifyUserCall := modifyUserCall.(type) {
		case *ent.UserCreate:
			modifyUserCall.AddEmails(email)
		case *ent.UserUpdate:
			modifyUserCall.AddEmails(email)
		case *ent.UserUpdateOne:
			modifyUserCall.AddEmails(email)
		default:
			return nil, fmt.Errorf(`invalid value passed as modifyUserCall: %T`, modifyUserCall)
		}

		// TODO: delete in case userCreate fails?

		emails = append(emails, v)
	}
	return emails, nil
}

func (b *Backend) CreateUser(in *resource.User) (*resource.User, error) {
	// TODO generate ETag
	// Note: use W/"...." Etags is probably better

	var builder resource.Builder

	// Generate a random password if none is given
	password := in.Password()
	if password == "" {
		password = randomString(25)
	} else {
		norm, err := precis.OpaqueString.String(password)
		if err != nil {
			return nil, fmt.Errorf(`failed to normalize password: %w`, err)
		}
		password = norm
	}

	createUserCall := b.db.User.Create().
		SetUserName(in.UserName()).
		SetPassword(password)

	// optional fields
	if in.HasDisplayName() {
		createUserCall.SetDisplayName(in.DisplayName())
	}

	if in.HasExternalID() {
		createUserCall.SetExternalID(in.ExternalID())
	}

	if in.HasUserType() {
		createUserCall.SetUserType(in.UserType())
	}

	if in.HasPreferredLanguage() {
		createUserCall.SetPreferredLanguage(in.PreferredLanguage())
	}

	if in.HasLocale() {
		createUserCall.SetLocale(in.Locale())
	}

	if in.HasTimezone() {
		createUserCall.SetTimezone(in.Timezone())
	}

	emails, err := b.createEmails(in, createUserCall)
	if err != nil {
		return nil, fmt.Errorf(`failed to create emails: %w`, err)
	}

	var name *resource.Names
	if v := in.Name(); v != nil {
		nameCreateCall := b.db.Names.Create()
		if sv := v.FamilyName(); sv != "" {
			nameCreateCall.SetFamilyName(sv)
		}
		if sv := v.Formatted(); sv != "" {
			nameCreateCall.SetFormatted(sv)
		}
		if sv := v.GivenName(); sv != "" {
			nameCreateCall.SetGivenName(sv)
		}
		if sv := v.HonorificPrefix(); sv != "" {
			nameCreateCall.SetHonorificPrefix(sv)
		}
		if sv := v.HonorificSuffix(); sv != "" {
			nameCreateCall.SetHonorificSuffix(sv)
		}
		if sv := v.MiddleName(); sv != "" {
			nameCreateCall.SetMiddleName(sv)
		}
		created, err := nameCreateCall.Save(context.TODO())
		if err != nil {
			return nil, fmt.Errorf(`failed to create user: failed to create name: %w`, err)
		}
		createUserCall.AddName(created)
		// TODO: delete in case userCreate fails?

		name = v
	}

	// now save the data
	u, err := createUserCall.
		Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to save data: %w`, err)
	}

	userBuilder := builder.User().
		ID(u.ID.String()).
		UserName(u.UserName).
		Meta(
			builder.Meta().
				ResourceType(`User`).
				Location(`https://foobar.com/scim/v2/Users/` + u.ID.String()).
				MustBuild(),
		)

	if len(emails) > 0 {
		userBuilder.Emails(emails...)
	}
	if name != nil {
		userBuilder.Name(name)
	}

	return userBuilder.Build()
}

func (b *Backend) RetrieveUser(id string, fields ...string) (*resource.User, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse ID: %w`, err)
	}

	userQuery := b.db.User.Query().
		Where(user.IDEQ(parsedUUID))

	userLoadEntFields(userQuery, fields)

	u, err := userQuery.
		Only(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to retrieve user: %w`, err)
	}

	return UserResourceFromEnt(u)
}

func (b *Backend) ReplaceUser(id string, in *resource.User) (*resource.User, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse ID: %w`, err)
	}

	// TODO: is it possible to just grab the ID or check existence?
	u, err := b.db.User.Query().
		Where(user.IDEQ(parsedUUID)).
		Only(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to retrieve user: %w`, err)
	}

	replaceUserCall := u.Update().
		ClearEmails()

	// optional fields
	if in.HasExternalID() {
		replaceUserCall.SetExternalID(in.ExternalID())
	}

	if in.HasUserType() {
		replaceUserCall.SetUserType(in.UserType())
	}

	if in.HasPreferredLanguage() {
		replaceUserCall.SetPreferredLanguage(in.PreferredLanguage())
	}

	if in.HasLocale() {
		replaceUserCall.SetLocale(in.Locale())
	}

	if in.HasTimezone() {
		replaceUserCall.SetTimezone(in.Timezone())
	}

	var emails []*resource.Email
	if in.HasEmails() {
		v, err := b.createEmails(in, replaceUserCall)
		if err != nil {
			return nil, fmt.Errorf(`failed to create emails: %w`, err)
		}
		emails = v
	}

	var builder resource.Builder
	userBuilder := builder.User().
		ID(id).
		ExternalID(u.ExternalID).
		UserName(u.UserName).
		Meta(
			builder.Meta().
				ResourceType(`User`).
				Location(`https://foobar.com/scim/v2/Users/` + u.ID.String()).
				MustBuild(),
		)

	if len(emails) > 0 {
		userBuilder.Emails(emails...)
	}

	return userBuilder.Build()
}

func (b *Backend) DeleteUser(id string) error {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf(`failed to parse ID: %w`, err)
	}

	if err := b.db.User.DeleteOneID(parsedUUID).Exec(context.TODO()); err != nil {
		return fmt.Errorf(`failed to delete user: %w`, err)
	}

	return nil
}

func (b *Backend) CreateGroup(in *resource.Group) (*resource.Group, error) {
	var userMembers []uuid.UUID
	var groupMembers []uuid.UUID

	for _, member := range in.Members() {
		asUUID, err := uuid.Parse(member.Value())
		if err != nil {
			return nil, fmt.Errorf(`expected "value" to contain a valid UUID: %w`, err)
		}
		if strings.Contains(member.Ref(), `/v2/Users`) {
			userMembers = append(userMembers, asUUID)
		} else if strings.Contains(member.Ref(), `/v2/Groups`) {
			groupMembers = append(groupMembers, asUUID)
		} else {
			return nil, fmt.Errorf(`$ref is required in group "members" attribute when creating Groups`)
		}
	}

	createGroupCall := b.db.Group.Create().
		SetDisplayName(in.DisplayName())
	if len(userMembers) > 0 {
		createGroupCall.AddUserIDs(userMembers...)
	}

	if len(groupMembers) > 0 {
		createGroupCall.AddChildIDs(groupMembers...)
	}

	g, err := createGroupCall.
		Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to save data: %w`, err)
	}

	var builder resource.Builder

	return builder.Group().
		ID(g.ID.String()).
		DisplayName(in.DisplayName()).
		Members(in.Members()...).
		Meta(
			builder.Meta().
				ResourceType(`Group`).
				Location(`https://foobar.com/scim/v2/Groups/` + g.ID.String()).
				MustBuild(),
		).
		Build()
}

// XXX passing these boolean variables is so ugly
func buildWhere(src string, buildUsers, buildGroups bool) ([]predicate.User, []predicate.Group, error) {
	expr, err := filter.Parse(src)
	if err != nil {
		return nil, nil, fmt.Errorf(`failed to parse filter: %w`, err)
	}

	var v filterVisitor

	// XXX while /.search (at the root level) allows querying for
	// all resources (well, User and Group only, really), /Users/.search
	// and /Group/.search restrict the search domain to either User or Group
	// only. In this case we need to limit the predicates that we generate

	// we do this by explicitly initializing the storage space
	// (the []predicate.* fields) with a non-nill value
	if buildUsers {
		v.users = []predicate.User{}
	}
	if buildGroups {
		v.groups = []predicate.Group{}
	}

	if err := v.visit(expr); err != nil {
		return nil, nil, fmt.Errorf(`failed to parse filter expression: %w`, err)
	}

	return v.users, v.groups, nil
}

type filterVisitor struct {
	users  []predicate.User
	groups []predicate.Group
}

func (v *filterVisitor) visit(expr filter.Expr) error {
	switch expr := expr.(type) {
	case filter.PresenceExpr:
		return v.visitPresenceExpr(expr)
	case filter.CompareExpr:
		return v.visitCompareExpr(expr)
	case filter.RegexExpr:
		return v.visitRegexExpr(expr)
	case filter.LogExpr: // RENAME ME TO LogicalStatement
		return v.visitLogExpr(expr)
	case filter.ParenExpr:
		return v.visitParenExpr(expr)
	case filter.ValuePath:
		return v.visitValuePath(expr)
	default:
		return fmt.Errorf(`unhandled statement type: %T`, expr)
	}
}

func exprAttr(expr interface{}) (interface{}, error) {
	switch v := expr.(type) {
	case string:
		return v, nil
	case interface{ Lit() string }: // IdentifierExpr, AttrValueExpr
		return v.Lit(), nil
	case filter.BoolExpr:
		return v.Lit(), nil
	case filter.NumberExpr:
		return v.Lit(), nil
	default:
		return nil, fmt.Errorf(`unhandled type: %T`, v)
	}
}

func (v *filterVisitor) visitPresenceExpr(expr filter.PresenceExpr) error {
	attr, err := exprAttr(expr.Attr())
	sattr, ok := attr.(string)
	if err != nil || !ok {
		if err == nil && !ok {
			err = fmt.Errorf(`expected string, got %T`, attr)
		}
		return fmt.Errorf(`left hand side of PresenceExpr is not valid: %w`, err)
	}

	switch expr.Operator() {
	case "pr":
		if v.users != nil {
			if pred := userPresencePredicate(sattr); pred != nil {
				v.users = append(v.users, pred)
			}
		}
		return nil
	default:
		return fmt.Errorf(`unhandled attr operator %q`, expr.Operator())
	}
}

func (v *filterVisitor) visitCompareExpr(expr filter.Expr) error {
	return fmt.Errorf(`unimplemented`)
}
func (v *filterVisitor) visitRegexExpr(expr filter.Expr) error {
	return fmt.Errorf(`unimplemented`)
}
func (v *filterVisitor) visitLogExpr(expr filter.Expr) error {
	return fmt.Errorf(`unimplemented`)
}
func (v *filterVisitor) visitParenExpr(expr filter.Expr) error {
	return fmt.Errorf(`unimplemented`)
}
func (v *filterVisitor) visitValuePath(expr filter.Expr) error {
	return fmt.Errorf(`unimplemented`)
}

func (b *Backend) Search(in *resource.SearchRequest) (*resource.ListResponse, error) {
	userWhere, groupWhere, err := buildWhere(in.Filter(), true, true)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse filter: %w`, err)
	}

	var list []interface{}

	if len(userWhere) > 0 {
		users, err := b.db.User.Query().Where(userWhere...).
			All(context.TODO())
		if err != nil {
			return nil, fmt.Errorf(`failed to execute query: %w`, err)
		}

		for _, user := range users {
			r, err := UserResourceFromEnt(user)
			if err != nil {
				return nil, fmt.Errorf(`failed to convert internal data to SCIM resource: %w`, err)
			}
			list = append(list, r)
		}
	}

	_ = groupWhere

	/*
			if attrs := in.Attributes(); len(attrs) > 0 {
				// TODO: need to generate SCIM name to ent name converter?
				snakeAttrs := make([]string, len(attrs))
				for i, attr := range attrs {
					snakeAttrs[i] = xstrings.Snake(attr)
				}
				q.Select(snakeAttrs...)
			}

			q.Where(
				user.DisplayNameHasPrefix("smith"),
			)

		users, err := q.All(context.TODO())
		if err != nil {
			return nil, fmt.Errorf(`failed to execute query: %w`, err)
		}

		list := make([]interface{}, len(users))
		for i, user := range users {
			created, err := UserResourceFromEnt(user)
			if err != nil {
				return nil, fmt.Errorf(`failed to convert internal data to SCIM resource: %w`, err)
			}
			list[i] = created
		}
	*/

	var builder resource.Builder
	return builder.ListResponse().
		TotalResults(len(list)).
		Resources(list...).
		Build()
}

func (b *Backend) RetrieveGroup(id string, fields ...string) (*resource.Group, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse ID: %w`, err)
	}

	groupQuery := b.db.Group.Query().
		Where(group.IDEQ(parsedUUID))

	groupLoadEntFields(groupQuery, fields)

	u, err := groupQuery.
		Only(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to retrieve group: %w`, err)
	}

	return GroupResourceFromEnt(u)
}
