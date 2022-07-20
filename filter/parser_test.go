// This file is based on https://github.com/hiyosi/filter/blob/master/parser_test.go
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

package filter_test

import (
	"testing"

	"github.com/cybozu-go/scim/filter"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	testcases := []struct {
		Filter  string
		Expr    filter.Expr
		Error   bool
		Options []filter.ParseOption
	}{
		{
			Filter: "foo bar baz",
			Error:  true,
		},
		{
			Filter: `ham eq spam`,
			Error:  true,
		},
		{
			Filter: `123 eq "spam"`,
			Error:  true,
		},
		{
			Filter: `123ham eq "spam"`,
			Error:  true,
		},
		{
			Filter: `ham@123 eq "spam"`,
			Error:  true,
		},
		{
			Filter: `ham co 123`,
			Error:  true,
		},
		{
			Filter: `ham co true`,
			Error:  true,
		},
		{
			Filter: `ham co false`,
			Error:  true,
		},
		{
			Filter: `123 co 123`,
			Error:  true,
		},
		{
			Filter: `ham sw 123`,
			Error:  true,
		},
		{
			Filter: `ham sw true`,
			Error:  true,
		},
		{
			Filter: `ham sw false`,
			Error:  true,
		},
		{
			Filter: `123 sw 123`,
			Error:  true,
		},
		{
			Filter: `ham ew 123`,
			Error:  true,
		},
		{
			Filter: `ham ew true`,
			Error:  true,
		},
		{
			Filter: `ham ew false`,
			Error:  true,
		},
		{
			Filter: `123 ew 123`,
			Error:  true,
		},
		{
			Filter: `123 gt 123`,
			Error:  true,
		},
		{
			Filter: `123 ge 123`,
			Error:  true,
		},
		{
			Filter: `123 lt 123`,
			Error:  true,
		},
		{
			Filter: `123 le 123`,
			Error:  true,
		},
		{
			Filter: `ham and 123`,
			Error:  true,
		},
		{
			Filter: `ham and (foo eq "spam")`,
			Error:  true,
		},
		{
			Filter: `ham or 123`,
			Error:  true,
		},
		{
			Filter: `ham or (foo eq "spam")`,
			Error:  true,
		},
		{
			Filter: `(foo)`,
			Error:  true,
		},
		{
			Filter: `(foo bar baz)`,
			Error:  true,
		},
		{
			Filter: `(foo bar baz]`,
			Error:  true,
		},
		{
			Filter: `[foo bar baz)`,
			Error:  true,
		},
		{
			Filter: `ham eq "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham"),
				filter.EqualOp,
				filter.NewAttrValueExpr("spam"),
			),
		},
		{
			Filter: `ham eq "spam@example.com"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham"),
				filter.EqualOp,
				filter.NewAttrValueExpr("spam@example.com"),
			),
		},
		{
			Filter: `ham eq "山田 太郎"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham"),
				filter.EqualOp,
				filter.NewAttrValueExpr("山田 太郎"),
			),
		},
		{
			Filter: `ham eq null`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham"),
				filter.EqualOp,
				filter.NewIdentifierExpr("null"),
			),
		},
		{
			Filter: `ham123 eq "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham123"),
				filter.EqualOp,
				filter.NewAttrValueExpr("spam"),
			),
		},
		{
			Filter: `ham-123 eq "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham-123"),
				filter.EqualOp,
				filter.NewAttrValueExpr("spam"),
			),
		},
		{
			Filter: `ham_123 eq "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham_123"),
				filter.EqualOp,
				filter.NewAttrValueExpr("spam"),
			),
		},
		{
			Filter: `ham eq 123`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham"),
				filter.EqualOp,
				filter.NewNumberExpr(123),
			),
		},
		{
			Filter: "ham eq true",
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr("ham"),
				filter.EqualOp,
				filter.NewBoolExpr(true),
			),
		},
		{
			Filter: `ham eq false`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.EqualOp,
				filter.NewBoolExpr(false),
			),
		},
		{
			Filter: `ham ne "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.NotEqualOp,
				filter.NewAttrValueExpr(`spam`),
			),
		},
		{
			Filter: `ham ne 123`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.NotEqualOp,
				filter.NewNumberExpr(123),
			),
		},
		{
			Filter: `ham ne true`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.NotEqualOp,
				filter.NewBoolExpr(true),
			),
		},
		{
			Filter: `ham ne false`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.NotEqualOp,
				filter.NewBoolExpr(false),
			),
		},
		{
			Filter: `ham ne null`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.NotEqualOp,
				filter.NewIdentifierExpr(`null`),
			),
		},
		// contains
		{
			Filter: `ham co "spam"`,
			Expr: filter.NewRegexExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.ContainsOp,
				`spam`,
			),
		},
		// starts with
		{
			Filter: `ham sw "spam"`,
			Expr: filter.NewRegexExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.StartsWithOp,
				`spam`,
			),
		},
		// ends with
		{
			Filter: `ham ew "spam"`,
			Expr: filter.NewRegexExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.EndsWithOp,
				`spam`,
			),
		},
		// presence (has value)
		{
			Filter: `ham pr`,
			Expr: filter.NewPresenceExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.PresenceOp,
			),
		},
		// greater than
		{
			Filter: `ham gt "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOp,
				filter.NewAttrValueExpr(`spam`),
			),
		},

		{
			Filter: `ham gt 123`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOp,
				filter.NewNumberExpr(123),
			),
		},
		{
			Filter: `ham gt true`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOp,
				filter.NewBoolExpr(true),
			),
		},
		{
			Filter: `ham gt false`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOp,
				filter.NewBoolExpr(false),
			),
		},
		{
			Filter: `ham gt null`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOp,
				filter.NewIdentifierExpr(`null`),
			),
		},
		// greater than or equal
		{
			Filter: `ham ge "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOrEqualToOp,
				filter.NewAttrValueExpr(`spam`),
			),
		},
		{
			Filter: `ham ge 123`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOrEqualToOp,
				filter.NewNumberExpr(123),
			),
		},
		{
			Filter: `ham ge true`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOrEqualToOp,
				filter.NewBoolExpr(true),
			),
		},
		{
			Filter: `ham ge false`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOrEqualToOp,
				filter.NewBoolExpr(false),
			),
		},
		{
			Filter: `ham ge null`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.GreaterThanOrEqualToOp,
				filter.NewIdentifierExpr(`null`),
			),
		},

		// less than
		{
			Filter: `ham lt "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOp,
				filter.NewAttrValueExpr(`spam`),
			),
		},
		{
			Filter: `ham lt 123`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOp,
				filter.NewNumberExpr(123),
			),
		},
		{
			Filter: `ham lt true`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOp,
				filter.NewBoolExpr(true),
			),
		},
		{
			Filter: `ham lt false`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOp,
				filter.NewBoolExpr(false),
			),
		},
		{
			Filter: `ham lt null`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOp,
				filter.NewIdentifierExpr(`null`),
			),
		},
		// less than equal
		{
			Filter: `ham le "spam"`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOrEqualToOp,
				filter.NewAttrValueExpr(`spam`),
			),
		},
		{
			Filter: `ham le 123`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOrEqualToOp,
				filter.NewNumberExpr(123),
			),
		},
		{
			Filter: `ham le true`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOrEqualToOp,
				filter.NewBoolExpr(true),
			),
		},
		{
			Filter: `ham le false`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOrEqualToOp,
				filter.NewBoolExpr(false),
			),
		},
		{
			Filter: `ham le null`,
			Expr: filter.NewCompareExpr(
				filter.NewIdentifierExpr(`ham`),
				filter.LessThanOrEqualToOp,
				filter.NewIdentifierExpr(`null`),
			),
		},

		// Logical And
		{
			Filter: `ham pr and ham eq "spam"`,
			Expr: filter.NewLogExpr(
				filter.NewPresenceExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.PresenceOp,
				),
				filter.AndOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`spam`),
				),
			),
		},
		{
			Filter: `ham eq "spam" and foo eq "bar"`,
			Expr: filter.NewLogExpr(
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`spam`),
				),
				filter.AndOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`foo`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`bar`),
				),
			),
		},
		{
			Filter: `ham co "spam" and foo eq "bar"`,
			Expr: filter.NewLogExpr(
				filter.NewRegexExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.ContainsOp,
					`spam`,
				),
				filter.AndOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`foo`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`bar`),
				),
			),
		},
		// Logical or
		{
			Filter: `ham pr or ham eq "spam"`,
			Expr: filter.NewLogExpr(
				filter.NewPresenceExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.PresenceOp,
				),
				filter.OrOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`spam`),
				),
			),
		},
		{
			Filter: `ham eq "spam" or foo eq "bar"`,
			Expr: filter.NewLogExpr(
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`spam`),
				),
				filter.OrOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`foo`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`bar`),
				),
			),
		},
		{
			Filter: `ham co "spam" or foo eq "bar"`,
			Expr: filter.NewLogExpr(
				filter.NewRegexExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.ContainsOp,
					`spam`,
				),
				filter.OrOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`foo`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`bar`),
				),
			),
		},

		// Precedence grouping
		{
			Filter: `(ham eq "spam")`,
			Expr: filter.NewParenExpr(
				``,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`spam`),
				),
			),
		},
		{
			Filter: `(ham co "spam")`,
			Expr: filter.NewParenExpr(
				``,
				filter.NewRegexExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.ContainsOp,
					`spam`,
				),
			),
		},
		{
			Filter: `(ham eq "spam") and (foo co "bar")`,
			Expr: filter.NewLogExpr(
				filter.NewParenExpr(
					``,
					filter.NewCompareExpr(
						filter.NewIdentifierExpr(`ham`),
						filter.EqualOp,
						filter.NewAttrValueExpr(`spam`),
					),
				),
				filter.AndOp,
				filter.NewParenExpr(
					``,
					filter.NewRegexExpr(
						filter.NewIdentifierExpr(`foo`),
						filter.ContainsOp,
						`bar`,
					),
				),
			),
		},

		// Complex  attribute filter grouping
		{
			Filter: `foo[ham eq "spam"]`,
			Expr: filter.NewValuePath(
				filter.NewIdentifierExpr(`foo`),
				nil,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`spam`),
				),
			),
		},
		// not function
		{
			Filter: `not (ham eq "spam")`,
			Expr: filter.NewParenExpr(
				filter.NotOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`spam`),
				),
			),
		},
		{
			Filter: `not (ham co "spam")`,
			Expr: filter.NewParenExpr(
				filter.NotOp,
				filter.NewRegexExpr(
					filter.NewIdentifierExpr(`ham`),
					filter.ContainsOp,
					`spam`,
				),
			),
		},
		{
			Filter: `not ( foo[ham eq "spam"] )`,
			Expr: filter.NewParenExpr(
				filter.NotOp,
				filter.NewValuePath(
					filter.NewIdentifierExpr(`foo`),
					nil,
					filter.NewCompareExpr(
						filter.NewIdentifierExpr(`ham`),
						filter.EqualOp,
						filter.NewAttrValueExpr(`spam`),
					),
				),
			),
		},
		// associative
		{
			Filter: `ham eq "spam" and foo eq "bar" and baz eq "qux"`,
			Expr: filter.NewLogExpr(
				filter.NewLogExpr(
					filter.NewCompareExpr(
						filter.NewIdentifierExpr(`ham`),
						filter.EqualOp,
						filter.NewAttrValueExpr(`spam`),
					),
					filter.AndOp,
					filter.NewCompareExpr(
						filter.NewIdentifierExpr(`foo`),
						filter.EqualOp,
						filter.NewAttrValueExpr(`bar`),
					),
				),
				filter.AndOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`baz`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`qux`),
				),
			),
		},
		{
			Filter: `ham eq "spam" or foo eq "bar" and baz eq "qux"`,
			Expr: filter.NewLogExpr(
				filter.NewLogExpr(
					filter.NewCompareExpr(
						filter.NewIdentifierExpr(`ham`),
						filter.EqualOp,
						filter.NewAttrValueExpr(`spam`),
					),
					filter.OrOp,
					filter.NewCompareExpr(
						filter.NewIdentifierExpr(`foo`),
						filter.EqualOp,
						filter.NewAttrValueExpr(`bar`),
					),
				),
				filter.AndOp,
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`baz`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`qux`),
				),
			),
		},
		{
			Filter: `members[value eq "2819c223-7f76-453a-919d-413861904646"].displayName`,
			Options: []filter.ParseOption{
				filter.WithPatchExpression(true),
			},
			Expr: filter.NewValuePath(
				filter.NewIdentifierExpr(`members`),
				filter.NewIdentifierExpr(`displayName`),
				filter.NewCompareExpr(
					filter.NewIdentifierExpr(`value`),
					filter.EqualOp,
					filter.NewAttrValueExpr(`2819c223-7f76-453a-919d-413861904646`),
				),
			),
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.Filter, func(t *testing.T) {
			expr, err := filter.Parse(tc.Filter, tc.Options...)
			if tc.Error {
				require.Error(t, err, `filter.Parse should fail`)
			} else {
				require.NoError(t, err, `filter.Parse should succeed`)
				require.Equal(t, tc.Expr, expr, `values should match`)
			}
		})
	}
}
