// This file is auto-generated by tools/cmd/genoptions/main.go. DO NOT EDIT

package sql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionIdent(t *testing.T) {
	require.Equal(t, "WithColumnMapper", identColumnMapper{}.String())
	require.Equal(t, "WithDialect", identDialect{}.String())
}