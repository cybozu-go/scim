package expr

type Position struct {
	Line   int
	Column int
}

// Expr is an interface to group AST nodes that are expressions
type Interface interface {
	expression()
}
