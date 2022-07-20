package filter

import (
	"github.com/cybozu-go/scim/filter/internal/expr"
	"github.com/cybozu-go/scim/filter/internal/token"
)

const (
	NotOp                  = token.NotOp
	AndOp                  = token.AndOp
	OrOp                   = token.OrOp
	PresenceOp             = token.PresenceOp
	EqualOp                = token.EqualOp
	NotEqualOp             = token.NotEqualOp
	GreaterThanOp          = token.GreaterThanOp
	GreaterThanOrEqualToOp = token.GreaterThanOrEqualToOp
	LessThanOp             = token.LessThanOp
	LessThanOrEqualToOp    = token.LessThanOrEqualToOp
	ContainsOp             = token.ContainsOp
	StartsWithOp           = token.StartsWithOp
	EndsWithOp             = token.EndsWithOp
	Dot                    = token.Dot
	LParen                 = token.LParen
	RParen                 = token.RParen
	LBracket               = token.LBracket
	RBracket               = token.RBracket
	True                   = token.True
	False                  = token.False
	Null                   = token.Null
)

type ExprStatement = expr.ExprStatement

func NewExprStatement(e Expr) ExprStatement {
	return expr.NewExprStatement(e)
}

type PresenceExpr = expr.PresenceExpr

func NewPresenceExpr(attr Expr, operator string) PresenceExpr {
	return expr.NewPresenceExpr(attr, operator)
}

type CompareExpr = expr.CompareExpr

func NewCompareExpr(lhe Expr, operator string, rhe Expr) CompareExpr {
	return expr.NewCompareExpr(lhe, operator, rhe)
}

type RegexExpr = expr.RegexExpr

func NewRegexExpr(lhe Expr, operator string, value interface{}) RegexExpr {
	return expr.NewRegexExpr(lhe, operator, value)
}

type ParenExpr = expr.ParenExpr

func NewParenExpr(operator string, subExpr Expr) ParenExpr {
	return expr.NewParenExpr(operator, subExpr)
}

type LogExpr = expr.LogExpr

func NewLogExpr(lhe Expr, operator string, rhS Expr) LogExpr {
	return expr.NewLogExpr(lhe, operator, rhS)
}

type ValuePath = expr.ValuePath

func NewValuePath(parentAttr Expr, subAttr Expr, subExpr Expr) ValuePath {
	return expr.NewValuePath(parentAttr, subAttr, subExpr)
}

type NumberExpr = expr.NumberExpr

func NewNumberExpr(lit int) NumberExpr {
	return expr.NewNumberExpr(lit)
}

type IdentifierExpr = expr.IdentifierExpr

func NewIdentifierExpr(lit string) IdentifierExpr {
	return expr.NewIdentifierExpr(lit)
}

type AttrValueExpr = expr.AttrValueExpr

func NewAttrValueExpr(lit string) AttrValueExpr {
	return expr.NewAttrValueExpr(lit)
}

type BoolExpr = expr.BoolExpr

func NewBoolExpr(lit bool) BoolExpr {
	return expr.NewBoolExpr(lit)
}
