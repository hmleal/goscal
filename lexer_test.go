package main

import (
	"testing"
)

type testpair struct {
	input  string
	output Token
}

var tests = []testpair{
	{"1", Token{kind: INTEGER, value: "1"}},
	{"+", Token{kind: PLUS, value: "+"}},
	{"-", Token{kind: MINUS, value: "-"}},
	{"*", Token{kind: MUL, value: "*"}},
	{"/", Token{kind: DIV, value: "/"}},
	{"(", Token{kind: LPAREN, value: "("}},
	{")", Token{kind: RPAREN, value: ")"}},
	{":=", Token{kind: ASSIGN, value: ":="}},
	{";", Token{kind: SEMI, value: ";"}},
}

func TestGetNextToken(t *testing.T) {
	for _, pair := range tests {
		lexer := NewLexer(pair.input)
		token := lexer.getNextToken()

		if token.kind != pair.output.kind {
			t.Error(
				"For", pair.input,
				"expected", pair.output.kind,
				"got", token.kind,
			)
		}
	}
}
