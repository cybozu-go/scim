//go:generate goyacc -l -o internal/fparser/parser.go -v filter_parser.output internal/fparser/parser.go.y
//go:generate goyacc -l -o internal/pparser/parser.go -v patch_parser.output internal/pparser/parser.go.y
//go:generate rm filter_parser.output patch_parser.output
//go:generate ../tools/cmd/genfilter.sh

package filter

import (
	"github.com/cybozu-go/scim/filter/internal/expr"
	"github.com/cybozu-go/scim/filter/internal/fparser"
	"github.com/cybozu-go/scim/filter/internal/pparser"
)

type Expr = expr.Interface

// Parse takes a string input and converts it into an expression.
// The `options` parameter can be specified to toggle specific behavior.
//
// By default the the parser expects an expression that is used for SCIM filters.
// But by specifying filter.WithPatchExpression(true), it adds support for
// allowing a single "valuePath" element to be present.
func Parse(src string, options ...ParseOption) (Expr, error) {
	parseFn := fparser.Parse
	for _, option := range options {
		switch option.Ident() {
		case identPatchExpression{}:
			parseFn = pparser.Parse
		}
	}
	return parseFn(src)
}
