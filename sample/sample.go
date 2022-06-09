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
	"github.com/lestrrat-go/rungroup"
	"golang.org/x/text/secure/precis"

	// default driver
	_ "github.com/mattn/go-sqlite3"
)

//FIXME
var _ = emailPresencePredicate
var _ = groupPresencePredicate
var _ = namesPresencePredicate
var _ = emailStartsWithPredicate
var _ = namesStartsWithPredicate

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

func (b *Backend) createEmails(in *resource.User) ([]*ent.Email, error) {
	emails := make([]*ent.Email, len(in.Emails()))
	var hasPrimary bool
	for i, v := range in.Emails() {
		emailCreateCall := b.db.Email.Create()
		emailCreateCall.SetValue(v.Value())

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
			emailCreateCall.SetPrimary(true)
			hasPrimary = true
		}

		email, err := emailCreateCall.Save(context.TODO())
		if err != nil {
			return nil, fmt.Errorf(`failed to save email %d: %w`, i, err)
		}

		emails[i] = email
	}
	return emails, nil
}

func (b *Backend) createName(v *resource.Names) (*ent.Names, error) {
	nameCreateCall := b.db.Names.Create()
	if v.HasFamilyName() {
		nameCreateCall.SetFamilyName(v.FamilyName())
	}
	if v.HasFormatted() {
		nameCreateCall.SetFormatted(v.Formatted())
	}
	if v.HasGivenName() {
		nameCreateCall.SetGivenName(v.GivenName())
	}
	if v.HasHonorificPrefix() {
		nameCreateCall.SetHonorificPrefix(v.HonorificPrefix())
	}
	if v.HasHonorificSuffix() {
		nameCreateCall.SetHonorificSuffix(v.HonorificSuffix())
	}
	if v.HasMiddleName() {
		nameCreateCall.SetMiddleName(v.MiddleName())
	}

	name, err := nameCreateCall.Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to save name: %w`, err)
	}

	return name, nil
}

func (b *Backend) CreateUser(in *resource.User) (*resource.User, error) {
	// TODO generate ETag
	// Note: use W/"...." Etags is probably better

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

	var emails []*ent.Email
	if in.HasEmails() {
		created, err := b.createEmails(in)
		if err != nil {
			return nil, fmt.Errorf(`failed to create emails: %w`, err)
		}
		createUserCall.AddEmails(created...)
		emails = created
	}

	var name []*ent.Names
	if in.HasName() {
		created, err := b.createName(in.Name())
		if err != nil {
			return nil, fmt.Errorf(`failed to create name: %w`, err)
		}
		createUserCall.AddName(created)
		name = []*ent.Names{created}
	}

	// now save the data
	u, err := createUserCall.
		Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to save user: %w`, err)
	}

	// u is the generated *ent.User, but it has no edges because
	// it is only populated during querying.
	// We could either populate u.Edges ourselves or re-fetch the
	// user via a query+eager-loading.
	// For the time being, we're just going to populate it ourselves
	u.Edges.Emails = emails
	u.Edges.Name = name

	return UserResourceFromEnt(u)
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
		ClearEmails().
		ClearName()

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

	var emails []*ent.Email
	if in.HasEmails() {
		created, err := b.createEmails(in)
		if err != nil {
			return nil, fmt.Errorf(`failed to create emails: %w`, err)
		}
		replaceUserCall.AddEmails(created...)
		emails = created
	}

	var name []*ent.Names
	if in.HasName() {
		created, err := b.createName(in.Name())
		if err != nil {
			return nil, fmt.Errorf(`failed to create name: %w`, err)
		}
		replaceUserCall.AddName(created)
		name = []*ent.Names{created}
	}

	if _, err := replaceUserCall.Save(context.TODO()); err != nil {
		return nil, fmt.Errorf(`failed to update user: %w`, err)
	}

	// u is the generated *ent.User, but it has no edges because
	// it is only populated during querying.
	// We could either populate u.Edges ourselves or re-fetch the
	// user via a query+eager-loading.
	// For the time being, we're just going to populate it ourselves
	u.Edges.Emails = emails
	u.Edges.Name = name

	return UserResourceFromEnt(u)
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

func (b *Backend) memberIDs(members []*resource.GroupMember) ([]uuid.UUID, []uuid.UUID, error) {
	var userMembers []uuid.UUID
	var groupMembers []uuid.UUID

	for _, member := range members {
		asUUID, err := uuid.Parse(member.Value())
		if err != nil {
			return nil, nil, fmt.Errorf(`expected "value" to contain a valid UUID: %w`, err)
		}

		if strings.Contains(member.Ref(), `/Users/`) {
			userMembers = append(userMembers, asUUID)
		} else if strings.Contains(member.Ref(), `/Groups/`) {
			groupMembers = append(groupMembers, asUUID)
		} else {
			return nil, nil, fmt.Errorf(`$ref is required in group "members" attribute when creating Groups`)
		}
	}

	return userMembers, groupMembers, nil
}

func (b *Backend) CreateGroup(in *resource.Group) (*resource.Group, error) {
	userMembers, groupMembers, err := b.memberIDs(in.Members())
	if err != nil {
		return nil, err
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
	case filter.PresenceOp:
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

func (v *filterVisitor) visitRegexExpr(expr filter.RegexExpr) error {
	lhe, err := exprAttr(expr.LHE())
	slhe, ok := lhe.(string)
	if err != nil || !ok {
		return fmt.Errorf(`left hand side of RegexExpr is not valid`)
	}

	rhe, err := exprAttr(expr.Value())
	if err != nil {
		return fmt.Errorf(`right hand side of RegexExpr is not valid: %w`, err)
	}
	// convert rhe to string so it can be passed to regexp.QuoteMeta
	srhe := fmt.Sprintf(`%v`, rhe)

	switch expr.Operator() {
	case filter.ContainsOp:
		return fmt.Errorf("unimplemented")
	case filter.StartsWithOp:
		if v.users != nil {
			if pred := userStartsWithPredicate(slhe, srhe); pred != nil {
				v.users = append(v.users, pred)
			}
		}
		if v.groups != nil {
			if pred := groupStartsWithPredicate(slhe, srhe); pred != nil {
				v.groups = append(v.groups, pred)
			}
		}
		return nil
	case filter.EndsWithOp:
		return fmt.Errorf("unimplemented")
	default:
		return fmt.Errorf(`unhandled regexp operator %q`, expr.Operator())
	}
}

func (v *filterVisitor) visitCompareExpr(expr filter.Expr) error {
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
	return b.search(in, true, true)
}

func (b *Backend) SearchUser(in *resource.SearchRequest) (*resource.ListResponse, error) {
	return b.search(in, true, false)
}

func (b *Backend) SearchGroup(in *resource.SearchRequest) (*resource.ListResponse, error) {
	return b.search(in, false, true)
}

func (b *Backend) search(in *resource.SearchRequest, searchUser, searchGroup bool) (*resource.ListResponse, error) {
	userWhere, groupWhere, err := buildWhere(in.Filter(), searchUser, searchGroup)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse filter: %w`, err)
	}

	var list []interface{}

	var g rungroup.Group
	if searchUser {
		_ = g.Add(rungroup.ActorFunc(func(ctx context.Context) error {
			users, err := b.db.User.Query().Where(userWhere...).
				All(ctx)
			if err != nil {
				return fmt.Errorf(`failed to execute query: %w`, err)
			}

			for _, user := range users {
				r, err := UserResourceFromEnt(user)
				if err != nil {
					return fmt.Errorf(`failed to convert internal data to SCIM resource: %w`, err)
				}
				list = append(list, r)
			}
			return nil
		}))
	}

	if searchGroup {
		_ = g.Add(rungroup.ActorFunc(func(ctx context.Context) error {
			groups, err := b.db.Group.Query().Where(groupWhere...).
				All(ctx)
			if err != nil {
				return fmt.Errorf(`failed to execute query: %w`, err)
			}

			for _, group := range groups {
				r, err := GroupResourceFromEnt(group)
				if err != nil {
					return fmt.Errorf(`failed to convert internal data to SCIM resource: %w`, err)
				}
				list = append(list, r)
			}
			return nil
		}))
	}

	if err := <-g.Run(context.TODO()); err != nil {
		return nil, err
	}

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
		WithUsers().
		WithChildren().
		Where(group.IDEQ(parsedUUID))

	groupLoadEntFields(groupQuery, fields)

	g, err := groupQuery.
		Only(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to retrieve group %s: %w`, id, err)
	}

	return GroupResourceFromEnt(g)
}

func (b *Backend) ReplaceGroup(id string, in *resource.Group) (*resource.Group, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse ID: %w`, err)
	}

	// TODO: is it possible to just grab the ID or check existence?
	g, err := b.db.Group.Query().
		Where(group.IDEQ(parsedUUID)).
		Only(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to retrieve group for replace: %w`, err)
	}

	replaceGroupCall := g.Update().
		ClearUsers().
		ClearChildren()

	// optional fields
	if in.HasDisplayName() {
		replaceGroupCall.SetDisplayName(in.DisplayName())
	}

	userMembers, groupMembers, err := b.memberIDs(in.Members())
	if err != nil {
		return nil, err
	}

	if len(userMembers) > 0 {
		replaceGroupCall.AddUserIDs(userMembers...)
	}

	if len(groupMembers) > 0 {
		replaceGroupCall.AddChildIDs(groupMembers...)
	}

	if _, err := replaceGroupCall.Save(context.TODO()); err != nil {
		return nil, fmt.Errorf(`failed to update group: %w`, err)
	}

	// Okay, I'm sure we can get the edges (users+children -> members)
	// somehow without re-fetching the data, but we're going to punt this
	return b.RetrieveGroup(id)
}

func (b *Backend) DeleteGroup(id string) error {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf(`failed to parse ID: %w`, err)
	}

	if err := b.db.Group.DeleteOneID(parsedUUID).Exec(context.TODO()); err != nil {
		return fmt.Errorf(`failed to delete group: %w`, err)
	}

	return nil
}
