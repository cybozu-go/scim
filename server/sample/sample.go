//go:generate go generate ./ent
//go:generate perl autofix.pl

package sample

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"entgo.io/ent/dialect"
	"github.com/cybozu-go/scim/resource"
	"github.com/cybozu-go/scim/server/sample/ent"
	"github.com/cybozu-go/scim/server/sample/ent/user"
	"github.com/google/uuid"
	"github.com/lestrrat-go/xstrings"
	"golang.org/x/text/secure/precis"

	_ "github.com/mattn/go-sqlite3"
)

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

func (b *Backend) CreateUser(in *resource.User) (*resource.User, error) {
	// TODO generate ETag

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
	if v := in.ExternalID(); v != "" {
		createUserCall.SetExternalID(v)
	}

	if v := in.UserType(); v != "" {
		createUserCall.SetUserType(v)
	}

	if v := in.PreferredLanguage(); v != "" {
		createUserCall.SetPreferredLanguage(v)
	}

	if v := in.Locale(); v != "" {
		createUserCall.SetLocale(v)
	}

	if v := in.Timezone(); v != "" {
		createUserCall.SetTimezone(v)
	}

	var emails []*resource.Email
	var hasPrimary bool
	for _, v := range in.Emails() {
		emailCreateCall := b.db.Email.Create().
			SetValue(v.Value())
		if sv := v.Display(); sv != "" {
			emailCreateCall.SetDisplay(sv)
		}
		if sv := v.Type(); sv != "" {
			emailCreateCall.SetType(sv)
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

		createUserCall.AddEmails(email)
		// TODO: delete in case userCreate fails?

		emails = append(emails, v)
	}

	var name *resource.Names
	if v := in.Name(); v != nil {
		nameCreateCall := b.db.Name.Create()
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
		createUserCall.AddNames(created)
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

	userBuilder.Emails(emails...)
	if name != nil {
		userBuilder.Name(name)
	}

	return userBuilder.Build()
}

func (b *Backend) RetrieveUser(id string) (*resource.User, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse ID: %w`, err)
	}

	user, err := b.db.User.Query().
		Where(user.IDEQ(parsedUUID)).
		First(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to retrieve user: %w`, err)
	}

	return UserResourceFromEnt(user)
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
		json.NewEncoder(os.Stdout).Encode(member)
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

func (b *Backend) Search(in *resource.SearchRequest) (*resource.ListResponse, error) {
	q := b.db.User.Query()

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

	var builder resource.Builder
	return builder.ListResponse().
		TotalResults(len(list)).
		Resources(list...).
		Build()
}
