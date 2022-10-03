package resource

import "fmt"

// Error() returns the stringified version of the SCIM error
func (v *Error) Error() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return fmt.Sprintf(`scim error: status="%d", detail=%q (%s)`, v.Status(), v.Detail(), v.SCIMType())
}
