package main

import (
	"fmt"
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
	fmt.Println(node.getName())
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
	if b.token.kind == PLUS {
		return b.left.visit(b.left) + b.right.visit(b.right)
	}
	if b.token.kind == MINUS {
		return b.left.visit(b.left) - b.right.visit(b.right)
	}
	if b.token.kind == MUL {
		return b.left.visit(b.left) * b.right.visit(b.right)
	}
	if b.token.kind == DIV {
		return b.left.visit(b.left) / b.right.visit(b.right)
	}
	return 0
}
