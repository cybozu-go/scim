package resource

import "fmt"

// Error() returns the stringified version of the SCIM error
func (e *Error) Error() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return fmt.Sprintf(`scim error: status="%d", detail=%q (%s)`, e.Status(), e.Detail(), e.ScimType())
}
