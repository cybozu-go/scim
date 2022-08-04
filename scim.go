//go:generate ./tools/cmd/genoptions.sh
//go:generate go generate ./filter
//go:generate go generate ./resource
//go:generate go generate ./sample
//go:generate go generate ./client

package scim

type ResourceNotFoundError struct{}

func (ResourceNotFoundError) Error() string {
	return "resource not found"
}
