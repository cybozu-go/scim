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
package filter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	tEOF     = -1
	tUNKNOWN = 0
)

var _ = tUNKNOWN // avoid unused error

var keywords = map[string]int{
	NotOp:                  tNOT,
	AndOp:                  tAND,
	OrOp:                   tOR,
	PresenceOp:             tPR,
	EqualOp:                tEQ,
	NotEqualOp:             tNE,
	ContainsOp:             tCO,
	StartsWithOp:           tSW,
	EndsWithOp:             tEW,
	GreaterThanOp:          tGT,
	GreaterThanOrEqualToOp: tGE,
	LessThanOp:             tLT,
	LessThanOrEqualToOp:    tLE,
	".":                    tDOT,
	"(":                    tLPAREN,
	")":                    tRPAREN,
	"[":                    tLBOXP,
	"]":                    tRBOXP,
	"true":                 tTRUE,
	"false":                tFALSE,
	"null":                 tNULL,
}

var caseInsensitiveKeywords map[string]struct{}

func init() {
	caseInsensitiveKeywords = make(map[string]struct{})
	for _, k := range []string{NotOp, AndOp, OrOp, PresenceOp, EqualOp, NotEqualOp, ContainsOp, StartsWithOp, EndsWithOp, GreaterThanOp, GreaterThanOrEqualToOp, LessThanOp, LessThanOrEqualToOp} {
		caseInsensitiveKeywords[strings.ToLower(k)] = struct{}{}
	}
}

type position struct {
	Line   int
	Column int
}

type scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

func newScanner(src string) *scanner {
	return &scanner{src: []rune(src)}
}

//nolint:nonamedreturns
func (s *scanner) Scan() (tok int, lit interface{}, pos position, err error) {
	s.skipWhiteSpace()
	pos = s.position()

	switch ch := s.peek(); {
	case unicode.IsLetter(ch) || ch == '"':
		if ch == '"' {
			tok, lit = tVALUE, s.scanAttrValue()
		} else {
			ident := s.scanIdentifier()
			// some operators need to be lower-cased
			lcident := strings.ToLower(ident)
			if _, ok := caseInsensitiveKeywords[lcident]; ok {
				ident = lcident
			}
			lit = ident

			if keyword, ok := keywords[ident]; ok {
				tok = keyword
			} else {
				tok = tIDENT
			}
		}
	case isDigit(ch):
		i, _ := strconv.Atoi(s.scanNumber())
		tok, lit = tNUMBER, i
	case ch == '(':
		tok = tLPAREN
		lit = "("
		s.next()
	case ch == ')':
		tok = tRPAREN
		lit = ")"
		s.next()
	case ch == '[':
		tok = tLBOXP
		lit = "["
		s.next()
	case ch == ']':
		tok = tRBOXP
		lit = "]"
		s.next()
	case ch == '.':
		tok = tDOT
		lit = "."
		s.next()
	case ch == -1:
		tok = tEOF
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

func (s *scanner) position() position {
	return position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
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
