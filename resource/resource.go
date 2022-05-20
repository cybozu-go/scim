//go:generate ../tools/cmd/genresources.sh

package resource

func init() {
	DefaultUserValidator = UserValidateFunc(defaultUserValidate)
}

type pair struct {
	Key   string
	Value interface{}
}

func defaultUserValidate(v *User) error {
	return nil
}

type Builder struct{}

type AuthenticationSchemeType string

const (
	InvalidAuthenticationScheme AuthenticationSchemeType = ""
	OAuth                       AuthenticationSchemeType = "oauth"
	OAuth2                      AuthenticationSchemeType = "oauth2"
	OAuthBearerToken            AuthenticationSchemeType = "oauthbearertoken"
	HTTPBasic                   AuthenticationSchemeType = "httpbasic"
	HTTPDigest                  AuthenticationSchemeType = "httpdigest"
)

type DataType string

const (
	String    DataType = "string"
	Boolean   DataType = "boolean"
	Decimal   DataType = "decimal"
	Integer   DataType = "integer"
	DateTime  DataType = "dateTime"
	Reference DataType = "reference"
	Complex   DataType = "complex"
)

type Mutability string

const (
	MutReadOnly  Mutability = `readOnly`
	MutReadWrite Mutability = `readWrite`
	MutImmutable Mutability = `immutable`
	MutWriteOnly Mutability = `writeOnly`
)

type Uniqueness string

const (
	UniqNone   Uniqueness = `none`
	UniqServer Uniqueness = `server`
	UniqGlobal Uniqueness = `global`
)

type Returned string

const (
	ReturnedAlways  Returned = "always"
	ReturnedNever   Returned = "never"
	ReturnedDefault Returned = "default"
	ReturnedRequest Returned = "request"
)
