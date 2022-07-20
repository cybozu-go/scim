package fparser_test

import (
	"testing"

	"github.com/cybozu-go/scim/filter/internal/fparser"
	"github.com/cybozu-go/scim/filter/internal/scanner"
)

func TestScanner(t *testing.T) {
	scanner.TestScanner(t, fparser.Dialect{})
}
