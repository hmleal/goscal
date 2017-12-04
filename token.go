package main

import "strconv"

const (
	INTEGER string = "INTEGER"
	PLUS    string = "PLUS"
	MINUS   string = "MINUS"
	MUL     string = "MUL"
	DIV     string = "DIV"
	LPAREN  string = "("
	RPAREN  string = ")"
	EOF     string = "EOF"
)

// Token basic struct
type Token struct {
	kind  string
	value string
}

// NewToken creates a new Token
func NewToken(kind, value string) Token {
	return Token{kind: kind, value: value}
}

func (t *Token) toInteger() int {
	i, _ := strconv.Atoi(t.value)
	return i
}
