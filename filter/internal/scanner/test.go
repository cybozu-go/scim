package scanner

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cybozu-go/scim/filter/internal/token"
)

func runSingleTest(t *testing.T, src string, expectTok int, dialect Dialect) {
	s := New(src, dialect)
	tok, lit, _, err := s.Scan()
	if err != nil {
		t.Errorf(`error while scanning: %s`, err)
	}

	if tok != expectTok {
		t.Errorf("Expect Scanner{%q}.Scan() = %#v, _ want %#v", src, tok, expectTok)
	}

	switch lit.(type) {
	case string:
		if lit != strings.Replace(src, "\"", "", -1) {
			t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
		}
	case int:
		s, _ := strconv.Atoi(src)
		if lit != s {
			t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
		}
	case bool:
		s, _ := strconv.ParseBool(src)
		if lit != s {
			t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
		}
	default:
		t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
	}

	tok, _, _, err = s.Scan()
	if err != nil {
		t.Errorf(`error scanning: %s`, err)
	}
	if expected := dialect.TokenType(token.EOF); tok != expected {
		t.Errorf("Expect Scanner{%q}.Scan() = %#v, _ want %#v", src, tok, expected)
	}
}

func TestScanner(t *testing.T, dialect Dialect) {
	t.Helper()

	identTok := dialect.TokenType(token.Ident)
	if identTok == -1 {
		t.Errorf("could not find token type for %s", token.Ident)
	}

	valueTok := dialect.TokenType(token.Value)
	if valueTok == -1 {
		t.Errorf("could not find token type for %s", token.Value)
	}

	numberTok := dialect.TokenType(token.Number)
	if numberTok == -1 {
		t.Errorf("could not find token type for %s", token.Number)
	}

	trueTok := dialect.TokenType(token.True)
	if trueTok == -1 {
		t.Errorf("could not find token type for %s", token.True)
	}

	falseTok := dialect.TokenType(token.False)
	if falseTok == -1 {
		t.Errorf("could not find token type for %s", token.False)
	}

	lparenTok := dialect.TokenType(token.LParen)
	if lparenTok == -1 {
		t.Errorf("could not find token type for %s", token.LParen)
	}

	rparenTok := dialect.TokenType(token.RParen)
	if rparenTok == -1 {
		t.Errorf("could not find token type for %s", token.RParen)
	}

	lbracketTok := dialect.TokenType(token.LBracket)
	if lbracketTok == -1 {
		t.Errorf("could not find token type for %s", token.LBracket)
	}

	rbracketTok := dialect.TokenType(token.RBracket)
	if rbracketTok == -1 {
		t.Errorf("could not find token type for %s", token.RBracket)
	}

	testcases := []struct {
		Input string
		Token int
	}{
		{
			Input: "ham",
			Token: identTok,
		},
		{
			Input: "ham123",
			Token: identTok,
		},
		{
			Input: "ham-123",
			Token: identTok,
		},
		{
			Input: "ham_123",
			Token: identTok,
		},
		{
			Input: "ham.123",
			Token: identTok,
		},
		{
			Input: "ham:123",
			Token: identTok,
		},
		{
			Input: "ham//123",
			Token: identTok,
		},
		{
			Input: `"ham"`,
			Token: valueTok,
		},
		{
			Input: `"ham spam"`,
			Token: valueTok,
		},
		{
			Input: `"ham@spam"`,
			Token: valueTok,
		},
		{
			Input: `"ham (spam)"`,
			Token: valueTok,
		},
		{
			Input: `123`,
			Token: numberTok,
		},
		{
			Input: `true`,
			Token: trueTok,
		},
		{
			Input: `false`,
			Token: falseTok,
		},
		{
			Input: `(`,
			Token: lparenTok,
		},
		{
			Input: `)`,
			Token: rparenTok,
		},
		{
			Input: `[`,
			Token: lbracketTok,
		},
		{
			Input: `]`,
			Token: rbracketTok,
		},
	}

	for _, tc := range testcases {
		runSingleTest(t, tc.Input, tc.Token, dialect)
	}
}
