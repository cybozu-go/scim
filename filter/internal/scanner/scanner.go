// This file is based on https://github.com/hiyosi/filter/blob/master/lexer.go
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
package scanner

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/cybozu-go/scim/filter/internal/expr"
	"github.com/cybozu-go/scim/filter/internal/token"
)

var caseInsensitiveKeywords map[string]struct{}

func init() {
	caseInsensitiveKeywords = make(map[string]struct{})
	for _, k := range []string{token.NotOp, token.AndOp, token.OrOp, token.PresenceOp, token.EqualOp, token.NotEqualOp, token.ContainsOp, token.StartsWithOp, token.EndsWithOp, token.GreaterThanOp, token.GreaterThanOrEqualToOp, token.LessThanOp, token.LessThanOrEqualToOp} {
		caseInsensitiveKeywords[strings.ToLower(k)] = struct{}{}
	}
}

type Scanner interface {
	Scan() (int, interface{}, expr.Position, error)
}

type scanner struct {
	dialect  Dialect
	src      []rune
	offset   int
	lineHead int
	line     int
}

type Dialect interface {
	// return token type for yacc parser
	// return -1 if not found
	TokenType(string) int

	// return lower-cased version if case-insensitive
	Normalize(string) string

	// returns true if this is a keyword
	IsKeyword(string) bool
}

func New(src string, d Dialect) Scanner {
	return &scanner{
		dialect: d,
		src:     []rune(src),
	}
}

//nolint:nonamedreturns
func (s *scanner) Scan() (tok int, lit interface{}, pos expr.Position, err error) {
	s.skipWhiteSpace()
	pos = s.position()

	switch ch := s.peek(); {
	case unicode.IsLetter(ch) || ch == '"':
		if ch == '"' {
			tok, lit = s.dialect.TokenType(token.Value), s.scanAttrValue()
		} else {
			ident := s.scanIdentifier()
			normalized := s.dialect.Normalize(ident)
			if s.dialect.IsKeyword(normalized) {
				tok = s.dialect.TokenType(normalized)
				lit = normalized
			} else {
				tok = s.dialect.TokenType(token.Ident)
				lit = ident
			}
		}
	case isDigit(ch):
		i, _ := strconv.Atoi(s.scanNumber())
		tok = s.dialect.TokenType(token.Number)
		lit = i
	case ch == '(':
		tok = s.dialect.TokenType(token.LParen)
		lit = "("
		s.next()
	case ch == ')':
		tok = s.dialect.TokenType(token.RParen)
		lit = ")"
		s.next()
	case ch == '[':
		tok = s.dialect.TokenType(token.LBracket)
		lit = "["
		s.next()
	case ch == ']':
		tok = s.dialect.TokenType(token.RBracket)
		lit = "]"
		s.next()
	case ch == '.':
		tok = s.dialect.TokenType(token.Dot)
		lit = "."
		s.next()
	case ch == -1:
		tok = s.dialect.TokenType(token.EOF)
		s.next()
	default:
		err = fmt.Errorf(`scan error: line %d, column %d: %q`, pos.Line, pos.Column, ch)
	}
	return
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsPunct(ch) || unicode.IsSymbol(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	}
	return -1
}

func (s *scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *scanner) position() expr.Position {
	return expr.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *scanner) skipWhiteSpace() {
	for isWhiteSpace(s.peek()) {
		s.next()
	}
}

func (s *scanner) scanAttrValue() string {
	var ret []rune
	s.next()
	for isLetter(s.peek()) || isDigit(s.peek()) || s.peek() == ' ' {
		ret = append(ret, s.peek())
		s.next()
		if ret[len(ret)-1] != '\\' && s.peek() == '"' {
			s.next()
			break
		}
	}
	return string(ret)
}

func (s *scanner) scanIdentifier() string {
	var ret []rune
	begin := s.peek()

	switch begin {
	default:
		for unicode.IsLetter(s.peek()) || isDigit(s.peek()) || s.peek() == '-' || s.peek() == '_' || s.peek() == '.' || s.peek() == ':' || s.peek() == '/' {
			ret = append(ret, s.peek())
			s.next()
		}
	}
	return string(ret)
}

func (s *scanner) scanNumber() string {
	var ret []rune
	for isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}
