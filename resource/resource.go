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
