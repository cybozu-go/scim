package resource_test

import (
	"testing"

	"github.com/cybozu-go/scim/schema"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {
	s, ok := schema.GetByResourceType(`User`)
	require.True(t, ok, `schema.Get should succeed`)

	t.Run(`attribute mutability`, func(t *testing.T) {
		testcases := []struct {
			Name  string
			Read  bool
			Write bool
		}{
			{
				Name: `groups`,
				Read: true,
			},
			{
				Name: `id`,
				Read: true,
			},
			{
				Name:  `password`,
				Write: true,
			},
			{
				Name:  `userName`,
				Read:  true,
				Write: true,
			},
		}
		for _, tc := range testcases {
			tc := tc
			t.Run(tc.Name, func(t *testing.T) {
				attr, ok := s.AttributeByName(tc.Name)
				require.True(t, ok, `s.AttributeByName should succeed`)

				if tc.Read {
					require.True(t, attr.CanRead(), `attribute %q should be readable`, tc.Name)
				} else {
					require.False(t, attr.CanRead(), `attribute %q should NOT be readable`, tc.Name)
				}

				if tc.Write {
					require.True(t, attr.CanWrite(), `attribute %q should be writable`, tc.Name)
				} else {
					require.False(t, attr.CanWrite(), `attribute %q should NOT be writable`, tc.Name)
				}
			})
		}
	})
}
