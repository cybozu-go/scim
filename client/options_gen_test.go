// This file is auto-generated by tools/cmd/genoptions/main.go. DO NOT EDIT

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionIdent(t *testing.T) {
	require.Equal(t, "WithClient", identClient{}.String())
	require.Equal(t, "WithTrace", identTrace{}.String())
}
