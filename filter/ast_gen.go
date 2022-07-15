package filter

type ExprStatement interface {
	Expr
	Expr() Expr
}

type exprStatement struct {
	expr Expr
}

func (*exprStatement) expression() {}

func NewExprStatement(expr Expr) ExprStatement {
	return &exprStatement{
		expr: expr,
	}
}

func (e *exprStatement) Expr() Expr {
	return e.expr
}

type PresenceExpr interface {
	Expr
	Attr() Expr
	Operator() string
}

type presenceExpr struct {
	attr     Expr
	operator string
}

func (*presenceExpr) expression() {}

func NewPresenceExpr(attr Expr, operator string) PresenceExpr {
	return &presenceExpr{
		attr:     attr,
		operator: operator,
	}
}

func (e *presenceExpr) Attr() Expr {
	return e.attr
}

func (e *presenceExpr) Operator() string {
	return e.operator
}

type CompareExpr interface {
	Expr
	LHE() Expr
	Operator() string
	RHE() Expr
}

type compareExpr struct {
	lhe      Expr
	operator string
	rhe      Expr
}

func (*compareExpr) expression() {}

func NewCompareExpr(lhe Expr, operator string, rhe Expr) CompareExpr {
	return &compareExpr{
		lhe:      lhe,
		operator: operator,
		rhe:      rhe,
	}
}

func (e *compareExpr) LHE() Expr {
	return e.lhe
}

func (e *compareExpr) Operator() string {
	return e.operator
}

func (e *compareExpr) RHE() Expr {
	return e.rhe
}

type RegexExpr interface {
	Expr
	LHE() Expr
	Operator() string
	Value() interface{}
}

type regexExpr struct {
	lhe      Expr
	operator string
	value    interface{}
}

func (*regexExpr) expression() {}

func NewRegexExpr(lhe Expr, operator string, value interface{}) RegexExpr {
	return &regexExpr{
		lhe:      lhe,
		operator: operator,
		value:    value,
	}
}

func (e *regexExpr) LHE() Expr {
	return e.lhe
}

func (e *regexExpr) Operator() string {
	return e.operator
}

func (e *regexExpr) Value() interface{} {
	return e.value
}

type ParenExpr interface {
	Expr
	Operator() string
	SubExpr() Expr
}

type parenExpr struct {
	operator string
	subExpr  Expr
}

func (*parenExpr) expression() {}

func NewParenExpr(operator string, subExpr Expr) ParenExpr {
	return &parenExpr{
		operator: operator,
		subExpr:  subExpr,
	}
}

func (e *parenExpr) Operator() string {
	return e.operator
}

func (e *parenExpr) SubExpr() Expr {
	return e.subExpr
}

type LogExpr interface {
	Expr
	LHE() Expr
	Operator() string
	RHS() Expr
}

type logExpr struct {
	lhe      Expr
	operator string
	rhS      Expr
}

func (*logExpr) expression() {}

func NewLogExpr(lhe Expr, operator string, rhS Expr) LogExpr {
	return &logExpr{
		lhe:      lhe,
		operator: operator,
		rhS:      rhS,
	}
}

func (e *logExpr) LHE() Expr {
	return e.lhe
}

func (e *logExpr) Operator() string {
	return e.operator
}

func (e *logExpr) RHS() Expr {
	return e.rhS
}

type ValuePath interface {
	Expr
	ParentAttr() Expr
	SubAttr() Expr
	SubExpr() Expr
}

type valuePath struct {
	parentAttr Expr
	subAttr    Expr
	subExpr    Expr
}

func (*valuePath) expression() {}

func NewValuePath(parentAttr Expr, subAttr Expr, subExpr Expr) ValuePath {
	return &valuePath{
		parentAttr: parentAttr,
		subAttr:    subAttr,
		subExpr:    subExpr,
	}
}

func (e *valuePath) ParentAttr() Expr {
	return e.parentAttr
}

func (e *valuePath) SubAttr() Expr {
	return e.subAttr
}

func (e *valuePath) SubExpr() Expr {
	return e.subExpr
}

type NumberExpr interface {
	Expr
	Lit() int
}

type numberExpr struct {
	lit int
}

func (*numberExpr) expression() {}

func NewNumberExpr(lit int) NumberExpr {
	return &numberExpr{
		lit: lit,
	}
}

func (e *numberExpr) Lit() int {
	return e.lit
}

type IdentifierExpr interface {
	Expr
	Lit() string
}

type identifierExpr struct {
	lit string
}

func (*identifierExpr) expression() {}

func NewIdentifierExpr(lit string) IdentifierExpr {
	return &identifierExpr{
		lit: lit,
	}
}

func (e *identifierExpr) Lit() string {
	return e.lit
}

type AttrValueExpr interface {
	Expr
	Lit() string
}

type attrValueExpr struct {
	lit string
}

func (*attrValueExpr) expression() {}

func NewAttrValueExpr(lit string) AttrValueExpr {
	return &attrValueExpr{
		lit: lit,
	}
}

func (e *attrValueExpr) Lit() string {
	return e.lit
}

type BoolExpr interface {
	Expr
	Lit() bool
}

type boolExpr struct {
	lit bool
}

func (*boolExpr) expression() {}

func NewBoolExpr(lit bool) BoolExpr {
	return &boolExpr{
		lit: lit,
	}
}

func (e *boolExpr) Lit() bool {
	return e.lit
}
