package schema_test

import (
	"testing"

	"github.com/cybozu-go/scim/schema"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {
	t.Run(`Unique`, func(t *testing.T) {
		seen := make(map[string]int)
		for i, r := range schema.All() {
			prev, ok := seen[r.Name()]
			require.False(t, ok, `resource %q should be unique (duplicate at item %d, first seen at item %d)`, r.Name(), i, prev)

			seen[r.Name()] = i
		}

		for n := range seen {
			t.Logf("schema name %q", n)
		}
	})
}
