package token

const (
	// token types
	Value  = `!!!value`
	Ident  = `!!!ident`
	Number = `!!!number`
	EOF    = `!!!eof`

	// keywords
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
	Dot                    = `.`
	LParen                 = `(`
	RParen                 = `)`
	LBracket               = `[`
	RBracket               = `]`
	True                   = `true`
	False                  = `false`
	Null                   = `null`
)
