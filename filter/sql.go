package filter

import (
	"fmt"
	"regexp"
	"strings"

	goqu "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type ColumnMapper interface {
	// Column returns the SCIM resource field specification
	// as a SQL table column. The column is expressed as exp.IdentifierExpression,
	// which can be qualified with a table name (i.e. goqu.C("foo").Table("bar"))
	//
	// If the table name is not the same as the main table name,
	// then the generator tries to create an INNER join
	Column(string) exp.IdentifierExpression
}

type ColumnMapperFunc func(string) exp.IdentifierExpression

func (fn ColumnMapperFunc) Column(n string) exp.IdentifierExpression {
	return fn(n)
}

var nilMapper = ColumnMapperFunc(defaultMapper)

func defaultMapper(n string) exp.IdentifierExpression {
	if !strings.ContainsRune(n, '.') {
		return goqu.C(n)
	}
	parts := strings.SplitN(n, ".", 2)
	return goqu.C(parts[1]).Table(parts[0])
}

type MapColumnMapper map[string]string

func (m MapColumnMapper) Column(n string) exp.IdentifierExpression {
	col, ok := m[n]
	if !ok {
		col = n
	}
	return defaultMapper(col)
}

// SQL generates a rather simplistic SQL from the given
// SCIM filter. The resulting SQL assumes that you have a very
// simplistic table normalization, and therefore this will
// most likely not scale well in your production system.
// Therefore, only use it as an example to implement something
// out fo the SCIM filter parser
func SQL(src, table string, options ...SQLOption) (string, []interface{}, error) {
	var mapper ColumnMapper
	var dialect string

	dialect = "default"
	mapper = nilMapper
	//nolint:forcetypeassert
	for _, option := range options {
		switch option.Ident() {
		case identColumnMapper{}:
			mapper = option.Value().(ColumnMapper)
		case identDialect{}:
			dialect = option.Value().(string)
		}
	}

	expr, err := Parse(src)
	if err != nil {
		return "", nil, fmt.Errorf(`failed to parse filter: %w`, err)
	}

	var sv = sqlVisitor{
		mapper: mapper,
	}

	if err := sv.visit(expr); err != nil {
		return "", nil, fmt.Errorf(`failed to generate SQL: %w`, err)
	}
	b := goqu.Dialect(dialect).From(table).Prepared(true)

	for _, extraTable := range sv.tables {
		// TODO: This currently just assumes that for any extra
		// tables that we use, there's always going to be a column
		// named "${extraTable}.${table}_id" which can be joined with
		// "${table}.id".
		b = b.Join(
			goqu.T(extraTable),
			goqu.On(goqu.C(fmt.Sprintf("%s_id", table)).Table(extraTable).Eq(goqu.C("id").Table(table))),
		)
	}

	return b.Where(sv.accum...).ToSQL()
}

type sqlVisitor struct {
	mapper ColumnMapper
	accum  []exp.Expression
	tables []string
}

func (sv *sqlVisitor) sqlValue(v interface{}) (interface{}, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case interface{ Lit() string }: // IdentifierExpr, AttrValueExpr
		return v.Lit(), nil
	case BoolExpr:
		return v.Lit(), nil
	case NumberExpr:
		return v.Lit(), nil
	default:
		return nil, fmt.Errorf(`unhandled type: %T`, v)
	}
}

func (sv *sqlVisitor) visit(v interface{}) error {
	switch v := v.(type) {
	case PresenceExpr:
		return sv.visitPresenceExpr(v)
	case CompareExpr:
		return sv.visitCompareExpr(v)
	case RegexExpr:
		return sv.visitRegexExpr(v)
	case LogExpr: // RENAME ME TO LogicalStatement
		return sv.visitLogExpr(v)
	case ParenExpr:
		return sv.visitParenExpr(v)
	case ValuePath:
		return sv.visitValuePath(v)
	default:
		return fmt.Errorf(`unhandled statement type: %T`, v)
	}
}

func (sv *sqlVisitor) visitLogExpr(v LogExpr) error {
	if err := sv.visit(v.LHE()); err != nil {
		return fmt.Errorf(`failed to parse left hand side of %q statement: %w`, v.Operator(), err)
	}
	if err := sv.visit(v.RHS()); err != nil {
		return fmt.Errorf(`failed to parse right hand side of %q statement: %w`, v.Operator(), err)
	}

	accum := sv.accum
	sv.accum = nil

	var l exp.ExpressionList
	switch v.Operator() {
	case "and":
		l = goqu.And(accum...)
	case "or":
		l = goqu.Or(accum...)
	default:
		return fmt.Errorf(`unhandled logical statement operator %q`, v.Operator())
	}
	sv.accum = append(sv.accum, l)
	return nil
}

func (sv *sqlVisitor) visitPresenceExpr(v PresenceExpr) error {
	attr, err := sv.sqlValue(v.Attr())
	sattr, ok := attr.(string)
	if err != nil || !ok {
		if err == nil && !ok {
			err = fmt.Errorf(`expected string, got %T`, attr)
		}
		return fmt.Errorf(`left hand side of PresenceExpr is not valid: %w`, err)
	}

	switch v.Operator() {
	case "pr":
		col := sv.mapper.Column(sattr)
		if t := col.GetTable(); t != "" {
			sv.tables = append(sv.tables, t)
		}
		sv.accum = append(sv.accum, col.IsNotNull())
		return nil
	default:
		return fmt.Errorf(`unhandled attr operator %q`, v.Operator())
	}
}

func (sv *sqlVisitor) visitCompareExpr(v CompareExpr) error {
	lhe, err := sv.sqlValue(v.LHE())
	slhe, ok := lhe.(string)
	if err != nil || !ok {
		if err == nil && !ok {
			err = fmt.Errorf(`expected string, got %T`, lhe)
		}
		return fmt.Errorf(`left hand side of CompareExpr is not valid: %w`, err)
	}

	rhe, err := sv.sqlValue(v.RHE())
	if err != nil {
		return fmt.Errorf(`right hand side of CompareExpr is not valid: %w`, err)
	}

	var expr exp.Expression
	col := sv.mapper.Column(slhe)
	switch v.Operator() {
	case EqualOp:
		expr = col.Eq(rhe)
	case NotEqualOp:
		expr = col.Neq(rhe)
	case "gt":
		expr = col.Gt(rhe)
	case "ge":
		expr = col.Gte(rhe)
	case "lt":
		expr = col.Lt(rhe)
	case "le":
		expr = col.Lte(rhe)
	default:
		return fmt.Errorf(`unhandled comparison operator %q`, v.Operator())
	}
	if t := col.GetTable(); t != "" {
		sv.tables = append(sv.tables, t)
	}
	sv.accum = append(sv.accum, expr)
	return nil
}

func (sv *sqlVisitor) visitRegexExpr(v RegexExpr) error {
	lhe, err := sv.sqlValue(v.LHE())
	slhe, ok := lhe.(string)
	if err != nil || !ok {
		return fmt.Errorf(`left hand side of RegexExpr is not valid`)
	}

	rhe, err := sv.sqlValue(v.Value())
	if err != nil {
		return fmt.Errorf(`right hand side of RegexExpr is not valid: %w`, err)
	}
	// convert rhe to string so it can be passed to regexp.QuoteMeta
	srhe := fmt.Sprintf(`%v`, rhe)

	var expr exp.Expression
	col := sv.mapper.Column(slhe)
	switch v.Operator() {
	case "co":
		expr = col.RegexpLike(regexp.QuoteMeta(srhe))
	case "sw":
		expr = col.RegexpLike("^" + regexp.QuoteMeta(srhe))
	case "ew":
		expr = col.RegexpLike(regexp.QuoteMeta(srhe) + "$")
	default:
		return fmt.Errorf(`unhandled regexp operator %q`, v.Operator())
	}
	if t := col.GetTable(); t != "" {
		sv.tables = append(sv.tables, t)
	}
	sv.accum = append(sv.accum, expr)
	return nil
}

func (sv *sqlVisitor) visitParenExpr(v ParenExpr) error {
	prev := sv.accum
	sv.accum = nil
	if err := sv.visit(v.SubExpr()); err != nil {
		return fmt.Errorf(`failed to visit statements inside (): %w`, err)
	}
	accum := sv.accum
	sv.accum = prev

	if v.Operator() == "not" {
		accum = []exp.Expression{goqu.Func("NOT", accum)}
	}

	sv.accum = append(sv.accum, accum...)
	return nil
}

func (sv *sqlVisitor) visitValuePath(v ValuePath) error {
	parent, ok := v.ParentAttr().(IdentifierExpr)
	if !ok {
		return fmt.Errorf(`expected identifier expression in parent for group statement, got %T`, v.ParentAttr)
	}

	prev := sv.accum
	sv.accum = nil
	if err := sv.visit(v.SubExpr()); err != nil {
		return fmt.Errorf(`failed to parse sub statement for group statement: %w`, err)
	}
	accum := sv.accum
	sv.accum = nil

	exprs := make([]exp.Expression, len(accum))
	for _, expr := range accum {
		exprs = append(exprs, qualifyTable(expr, parent.Lit()))
	}
	sv.accum = prev
	sv.accum = append(sv.accum, exprs...)
	sv.tables = append(sv.tables, parent.Lit())
	return nil
}

func qualifyTable(expr exp.Expression, parent string) exp.Expression {
	switch expr := expr.(type) {
	case exp.ExpressionList:
		// AND/OR expressions.
		var subexprs []exp.Expression
		for _, subexpr := range expr.Expressions() {
			subexprs = append(subexprs, qualifyTable(subexpr, parent))
		}
		return exp.NewExpressionList(expr.Type(), subexprs...)
	case exp.BooleanExpression:
		lhs := qualifyTable(expr.LHS(), parent)
		rhs := expr.RHS()
		if erhs, ok := rhs.(exp.Expression); ok {
			rhs = qualifyTable(erhs, parent)
		}
		return exp.NewBooleanExpression(expr.Op(), lhs, rhs)
	case exp.IdentifierExpression:
		if expr.GetTable() != "" {
			return expr
		}
		return expr.Table(parent)
	default:
		panic(fmt.Sprintf("unhandled expression type: %T", expr))
	}
}
