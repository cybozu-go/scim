package pparser

import (
	"log"
	"strings"

	"github.com/cybozu-go/scim/filter/internal/token"
)

const (
	tEOF     = -1
	tUNKNOWN = 0
)

var _ = tUNKNOWN // avoid unused error

var keywords = map[string]int{
	token.NotOp:                  tNOT,
	token.AndOp:                  tAND,
	token.OrOp:                   tOR,
	token.PresenceOp:             tPR,
	token.EqualOp:                tEQ,
	token.NotEqualOp:             tNE,
	token.ContainsOp:             tCO,
	token.StartsWithOp:           tSW,
	token.EndsWithOp:             tEW,
	token.GreaterThanOp:          tGT,
	token.GreaterThanOrEqualToOp: tGE,
	token.LessThanOp:             tLT,
	token.LessThanOrEqualToOp:    tLE,
	token.True:                   tTRUE,
	token.False:                  tFALSE,
	token.Null:                   tNULL,
}

var families = map[string]int{
	token.Value:  tVALUE,
	token.Number: tNUMBER,
	token.Ident:  tIDENT,
}

var miscellaneous = map[string]int{
	token.Dot:      tDOT,
	token.LParen:   tLPAREN,
	token.RParen:   tRPAREN,
	token.LBracket: tLBOXP,
	token.RBracket: tRBOXP,
	token.EOF:      tEOF,
}

type Dialect struct{}

func (Dialect) IsKeyword(s string) bool {
	_, ok := keywords[s]
	return ok
}

func (Dialect) Normalize(s string) string {
	l := strings.ToLower(s)
	_, ok := keywords[l]
	if !ok {
		return s
	}
	return l
}

func (Dialect) TokenType(s string) int {
	if v, ok := keywords[s]; ok {
		return v
	}

	if v, ok := families[s]; ok {
		return v
	}

	if v, ok := miscellaneous[s]; ok {
		return v
	}

	log.Printf("Could not find token for %q", s)
	return tUNKNOWN
}
