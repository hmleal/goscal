package main

import (
	"strconv"
)

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

// AST

// Node a simple interface for tree elements
type Node interface {
	getName() string
	visit(Node) int
}

// num{token: Token}
type num struct {
	token Token
}

func (n num) getName() string {
	return "num"
}

func (n num) visit(node Node) int {
	return n.token.toInteger()
}

// binOP{left: Token, token: Token, right: Token}
type binOP struct {
	left  Node
	token Token
	right Node
}

func (b binOP) getName() string {
	return "binOP"
}

func (b binOP) visit(node Node) int {
	switch b.token.kind {
	case PLUS:
		return b.left.visit(b.left) + b.right.visit(b.right)
	case MINUS:
		return b.left.visit(b.left) - b.right.visit(b.right)
	case MUL:
		return b.left.visit(b.left) * b.right.visit(b.right)
	case DIV:
		return b.left.visit(b.left) / b.right.visit(b.right)
	default:
		return 0
	}
}

type unaryOP struct {
	token Token
	expr  Node
}

func (u unaryOP) getName() string {
	return "unaryOP"
}

func (u unaryOP) visit(node Node) int {
	token := u.token.kind
	if token == PLUS {
		return +u.expr.visit(u.expr)
	}
	if token == MINUS {
		return -u.expr.visit(u.expr)
	}
	return 0
}
