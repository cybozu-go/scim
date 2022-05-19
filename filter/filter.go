//go:generate goyacc -l -o parser.go -v parser.output parser.go.y
//go:generate rm parser.output
//go:generate ../tools/cmd/genfilter.sh
package filter

const (
	NotOp                  = `not`
	AndOp                  = `and`
	OrOp                   = `or`
	PresenceOp             = `pr`
	EqualOp                = `eq`
	NotEqualOp             = `ne`
	GreaterThanOp          = `gt`
	GreaterThanOrEqualToOp = `ge`
	LessThanOp             = `lt`
	LessThanOrEqualToOp    = `le`
	ContainsOp             = `co`
	StartsWithOp           = `sw`
	EndsWithOp             = `ew`
)

// Expr is an interface to group AST nodes that are expressions
type Expr interface {
	expression()
}

// Parse takes a string input and converts it into an expression
func Parse(src string) (Expr, error) {
	s := newScanner(src)
	return parse(s)
}
