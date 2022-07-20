package expr

type ExprStatement interface {
	Interface
	Expr() Interface
}

type exprStatement struct {
	expr Interface
}

func (*exprStatement) expression() {}

func NewExprStatement(expr Interface) ExprStatement {
	return &exprStatement{
		expr: expr,
	}
}

func (e *exprStatement) Expr() Interface {
	return e.expr
}

type PresenceExpr interface {
	Interface
	Attr() Interface
	Operator() string
}

type presenceExpr struct {
	attr     Interface
	operator string
}

func (*presenceExpr) expression() {}

func NewPresenceExpr(attr Interface, operator string) PresenceExpr {
	return &presenceExpr{
		attr:     attr,
		operator: operator,
	}
}

func (e *presenceExpr) Attr() Interface {
	return e.attr
}

func (e *presenceExpr) Operator() string {
	return e.operator
}

type CompareExpr interface {
	Interface
	LHE() Interface
	Operator() string
	RHE() Interface
}

type compareExpr struct {
	lhe      Interface
	operator string
	rhe      Interface
}

func (*compareExpr) expression() {}

func NewCompareExpr(lhe Interface, operator string, rhe Interface) CompareExpr {
	return &compareExpr{
		lhe:      lhe,
		operator: operator,
		rhe:      rhe,
	}
}

func (e *compareExpr) LHE() Interface {
	return e.lhe
}

func (e *compareExpr) Operator() string {
	return e.operator
}

func (e *compareExpr) RHE() Interface {
	return e.rhe
}

type RegexExpr interface {
	Interface
	LHE() Interface
	Operator() string
	Value() interface{}
}

type regexExpr struct {
	lhe      Interface
	operator string
	value    interface{}
}

func (*regexExpr) expression() {}

func NewRegexExpr(lhe Interface, operator string, value interface{}) RegexExpr {
	return &regexExpr{
		lhe:      lhe,
		operator: operator,
		value:    value,
	}
}

func (e *regexExpr) LHE() Interface {
	return e.lhe
}

func (e *regexExpr) Operator() string {
	return e.operator
}

func (e *regexExpr) Value() interface{} {
	return e.value
}

type ParenExpr interface {
	Interface
	Operator() string
	SubExpr() Interface
}

type parenExpr struct {
	operator string
	subExpr  Interface
}

func (*parenExpr) expression() {}

func NewParenExpr(operator string, subExpr Interface) ParenExpr {
	return &parenExpr{
		operator: operator,
		subExpr:  subExpr,
	}
}

func (e *parenExpr) Operator() string {
	return e.operator
}

func (e *parenExpr) SubExpr() Interface {
	return e.subExpr
}

type LogExpr interface {
	Interface
	LHE() Interface
	Operator() string
	RHS() Interface
}

type logExpr struct {
	lhe      Interface
	operator string
	rhS      Interface
}

func (*logExpr) expression() {}

func NewLogExpr(lhe Interface, operator string, rhS Interface) LogExpr {
	return &logExpr{
		lhe:      lhe,
		operator: operator,
		rhS:      rhS,
	}
}

func (e *logExpr) LHE() Interface {
	return e.lhe
}

func (e *logExpr) Operator() string {
	return e.operator
}

func (e *logExpr) RHS() Interface {
	return e.rhS
}

type ValuePath interface {
	Interface
	ParentAttr() Interface
	SubAttr() Interface
	SubExpr() Interface
}

type valuePath struct {
	parentAttr Interface
	subAttr    Interface
	subExpr    Interface
}

func (*valuePath) expression() {}

func NewValuePath(parentAttr Interface, subAttr Interface, subExpr Interface) ValuePath {
	return &valuePath{
		parentAttr: parentAttr,
		subAttr:    subAttr,
		subExpr:    subExpr,
	}
}

func (e *valuePath) ParentAttr() Interface {
	return e.parentAttr
}

func (e *valuePath) SubAttr() Interface {
	return e.subAttr
}

func (e *valuePath) SubExpr() Interface {
	return e.subExpr
}

type NumberExpr interface {
	Interface
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
	Interface
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
	Interface
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
	Interface
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
