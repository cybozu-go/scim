//go:generate ../tools/cmd/genresources.sh

package resource

func init() {
	DefaultUserValidator = UserValidateFunc(defaultUserValidate)
}

func defaultUserValidate(v *User) error {
	return nil
}

type Builder struct{}
