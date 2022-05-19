package filter

import (
	"strconv"
	"strings"
	"testing"
)

func testScanner(t *testing.T, src string, expectTok int) {
	s := newScanner(src)
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
	if tok != tEOF {
		t.Errorf("Expect Scanner{%q}.Scan() = %#v, _ want %#v", src, tok, tEOF)
	}
}

func TestScanner(t *testing.T) {
	testScanner(t, "ham", tIDENT)
	testScanner(t, "ham123", tIDENT)
	testScanner(t, "ham-123", tIDENT)
	testScanner(t, "ham_123", tIDENT)
	testScanner(t, "ham.123", tIDENT)
	testScanner(t, "ham:123", tIDENT)
	testScanner(t, "ham//123", tIDENT)
	testScanner(t, "\"ham\"", tVALUE)
	testScanner(t, "\"ham spam\"", tVALUE)
	testScanner(t, "\"ham@spam\"", tVALUE)
	testScanner(t, "\"ham (spam)\"", tVALUE)
	testScanner(t, "123", tNUMBER)
	testScanner(t, "123", tNUMBER)
	testScanner(t, "true", tTRUE)
	testScanner(t, "false", tFALSE)
	testScanner(t, "(", tLPAREN)
	testScanner(t, ")", tRPAREN)
	testScanner(t, "[", tLBOXP)
	testScanner(t, "]", tRBOXP)
}
