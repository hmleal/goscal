package main

import (
	"strings"
	"unicode"
)

type Lexer struct {
	text        string
	position    int
	currentChar rune
}

// NewLexer create's a Lexer instance
func NewLexer(text string) Lexer {
	return Lexer{text: text, position: 0, currentChar: rune(text[0])}
}

func (l *Lexer) getNextToken() Token {
	for l.currentChar != 0 {
		if unicode.IsSpace(l.currentChar) {
			l.advance()
			continue
		}

		if unicode.IsDigit(l.currentChar) {
			return NewToken(INTEGER, l.integer())
		}

		if string(l.currentChar) == ":" && string(l.peek()) == "=" {
			l.advance()
			l.advance()
			return NewToken(ASSIGN, ":=")
		}

		if string(l.currentChar) == ";" {
			l.advance()
			return NewToken(SEMI, ";")
		}

		if string(l.currentChar) == "+" {
			l.advance()
			return NewToken(PLUS, "+")
		}

		if string(l.currentChar) == "-" {
			l.advance()
			return NewToken(MINUS, "-")
		}

		if string(l.currentChar) == "*" {
			l.advance()
			return NewToken(MUL, "*")
		}

		if string(l.currentChar) == "/" {
			l.advance()
			return NewToken(DIV, "/")
		}

		if string(l.currentChar) == "(" {
			l.advance()
			return NewToken(LPAREN, "(")
		}

		if string(l.currentChar) == ")" {
			l.advance()
			return NewToken(RPAREN, ")")
		}

		panic("Invalid character")
	}

	return NewToken(EOF, "EOF")
}

func (l *Lexer) advance() {
	l.position++
	if l.position > len(l.text)-1 {
		l.currentChar = 0
	} else {
		l.currentChar = rune(l.text[l.position])
	}
}

func (l *Lexer) peek() rune {
	peekPosition := l.position + 1
	if peekPosition > len(l.text)-1 {
		return 0
	}
	return rune(l.text[peekPosition])
}

func (l *Lexer) integer() string {
	result := []string{}
	for unicode.IsDigit(l.currentChar) {
		result = append(result, string(l.currentChar))
		l.advance()
	}
	return strings.Join(result, "")
}
