package examples_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cybozu-go/scim/resource"
)

func Example_ResourceExtension() {
	var b resource.Builder

	user, err := b.User().
		ID("foo").
		UserName("foo").
		Extension(
			resource.EnterpriseUserSchemaURI,
			b.EnterpriseUser().
				CostCenter("foo").
				Department("foo").
				Division("foo").
				EmployeeNumber("foo").
				Organization("foo").
				MustBuild(),
		).
		Build()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	json.NewEncoder(os.Stdout).Encode(user)
	// OUTPUT:
	// {"id":"foo","schemas":["urn:ietf:params:scim:schemas:core:2.0:User","urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"],"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User":{"costCenter":"foo","department":"foo","division":"foo","employeeNumber":"foo","organization":"foo"},"userName":"foo"}
}
