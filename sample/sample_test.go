package sample_test

import (
	"testing"

	"github.com/cybozu-go/scim/sample"
	"github.com/cybozu-go/scim/test"
	"github.com/stretchr/testify/require"
)

func TestSample(t *testing.T) {
	s, err := sample.New("file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err, `sample.New should succeed`)

	test.RunConformanceTests(t, "Sample backend", s)
}
