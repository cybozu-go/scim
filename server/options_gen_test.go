// This file is auto-generated by tools/cmd/genoptions/main.go. DO NOT EDIT

package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionIdent(t *testing.T) {
	require.Equal(t, "WithPath", identPath{}.String())
}