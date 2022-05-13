// This file is based on https://github.com/hiyosi/filter/blob/master/parser.go.y
//
// Licensed under MIT license
// 
// The MIT License (MIT)
// 
// Copyright (c) 2015 Tomoya Usami
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
%{
package filter

import "fmt"

type token struct {
	tok int
	lit interface{}
	pos position
}

%}

%union{
	filters Expr
	expr	Expr
	tok	token
}

%type<filters> filters
%type<expr> expr
%type<expr> attrName
%type<expr> attrValue

%token<tok> tIDENT tTRUE tFALSE tNULL tVALUE tNUMBER tPR tEQ tNE tCO tSW tEW tGT tGE tLT tLE tAND tOR tNOT tLPAREN tRPAREN tLBOXP tRBOXP tSP

%left  tAND
%left  tOR
%right tNOT

%%

filters
	: expr
	{
		$$ = $1
		if l, ok := yylex.(*lexer); ok {
			l.expr = $$
		}
	}

expr //TODO: gt,ge,lt,le operators should take string and boolean
        : attrName tPR
        {
                $$ = NewPresenceExpr($1, $2.lit.(string))
        }
        | attrName tEQ attrValue
        {
                $$ = NewCompareExpr($1, $2.lit.(string), $3)
        }
        | attrName tNE attrValue
        {
                $$ = NewCompareExpr($1, $2.lit.(string), $3)
        }
        | attrName tCO tVALUE
        {
                $$ = NewRegexExpr($1, $2.lit.(string), $3.lit)
        }
        | attrName tSW tVALUE
        {
                $$ = NewRegexExpr($1, $2.lit.(string), $3.lit)
        }
        | attrName tEW tVALUE
        {
                $$ = NewRegexExpr($1, $2.lit.(string), $3.lit)
        }
        | attrName tGT attrValue
        {
                $$ = NewCompareExpr($1, $2.lit.(string), $3)
        }
        | attrName tGE attrValue
        {
                $$ = NewCompareExpr($1, $2.lit.(string), $3)
        }
        | attrName tLT attrValue
        {
                $$ = NewCompareExpr($1, $2.lit.(string), $3)
        }
        | attrName tLE attrValue
        {
                $$ = NewCompareExpr($1, $2.lit.(string), $3)
        }
        | expr tAND expr
        {
                $$ = NewLogExpr($1, $2.lit.(string), $3)
        }
        | expr tOR expr
        {
                $$ = NewLogExpr($1, $2.lit.(string), $3)
        }
        | tLPAREN expr tRPAREN
        {
                $$ = NewParenExpr("", $2)
        }
        | tNOT tLPAREN expr tRPAREN
        {
                $$ = NewParenExpr($1.lit.(string), $3)
        }
        | attrName tLBOXP expr tRBOXP
        {
                $$ = NewValuePath($1, $3)
        }

attrName
        : tIDENT 
	{
		$$ = NewIdentifierExpr($1.lit.(string))
        }

attrValue
        : tNUMBER
	{
		$$ = NewNumberExpr($1.lit.(int))
	}
        | tTRUE
        {
       		$$ = NewBoolExpr(true)
        }
        | tFALSE
        {
       		$$ = NewBoolExpr(false)
        }
        | tNULL
        {
       		$$ = NewIdentifierExpr($1.lit.(string))
        }        
        | tVALUE
	{
		$$ = NewAttrValueExpr($1.lit.(string))
        }        

%%

type lexer struct {
	s          *scanner
	recentLit  interface{}
	recentPos  position
	expr       Expr
	err        chan error
}

func (l *lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := l.s.Scan()
	if err != nil {
		l.emitError(err)
		return -1
	}
	if tok == tEOF {
		return 0
	}
	lval.tok = token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *lexer) makeError(e interface{}) error {
	switch e := e.(type) {
	case error:
		return fmt.Errorf(`parse error: line %d, column %d: %q: %w`, l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
	default:
		return fmt.Errorf(`parse error: line %d, column %d: %q: %s`, l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
	}
}

// implements yylexer, so it must stay
func (l *lexer) Error(e string) {
	l.emitError(l.makeError(e))
}

func (l *lexer) emitError(err error) {
	select {
	case l.err <- err:
	default:
	}
}

func parse(s *scanner) (Expr, error) {
	l := lexer{s: s, err: make(chan error, 1)}
	if yyParse(&l) != 0 {
		return nil, <-l.err
	}
	return l.expr, nil
}
