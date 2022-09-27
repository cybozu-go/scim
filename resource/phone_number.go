package resource

import (
	"fmt"
	"net/url"
)

type PhoneNumberValue string

func (v *PhoneNumberValue) Get() string {
	return string(*v)
}

func (v *PhoneNumberValue) Accept(in interface{}) error {
	switch in := in.(type) {
	case string:
		// This is a very simplified version of validating RFC3966.
		// We just make sure that it's a valid URI, and that the schema
		// is "tel"
		u, err := url.Parse(in)
		if err != nil {
			return fmt.Errorf(`failed to parse phone number: %w`, err)
		}
		if u.Scheme != `tel` {
			return fmt.Errorf(`phone number scheme must be "tel", got %s`, u.Scheme)
		}
		*v = PhoneNumberValue(in)
		return nil
	default:
		return fmt.Errorf(`phoneNumber.value must be of string type (got %T)`, in)
	}
}
