package main

import (
	"strings"
	"unicode"
)

type Lexer struct {
	text         string
	position     int
	current_char rune
}

// NewLexer create's a Lexer instance
func NewLexer(text string) Lexer {
	return Lexer{text: text, position: 0, current_char: rune(text[0])}
}

func (l *Lexer) get_next_token() Token {
	for l.current_char != 0 {
		if unicode.IsSpace(l.current_char) {
			l.advance()
			continue
		}

		if unicode.IsDigit(l.current_char) {
			return NewToken(INTEGER, l.integer())
		}

		if string(l.current_char) == "+" {
			token := NewToken(PLUS, string(l.current_char))
			l.advance()
			return token
		}

		if string(l.current_char) == "-" {
			token := NewToken(MINUS, string(l.current_char))
			l.advance()
			return token
		}

		if string(l.current_char) == "*" {
			token := NewToken(MUL, string(l.current_char))
			l.advance()
			return token
		}

		if string(l.current_char) == "/" {
			token := NewToken(DIV, string(l.current_char))
			l.advance()
			return token
		}

		if string(l.current_char) == "(" {
			token := NewToken(LPAREN, string(l.current_char))
			l.advance()
			return token
		}

		if string(l.current_char) == ")" {
			token := NewToken(RPAREN, string(l.current_char))
			l.advance()
			return token
		}

		panic("Invalid character")
	}

	return NewToken(EOF, "EOF")
}

func (l *Lexer) advance() {
	l.position++
	if l.position > len(l.text)-1 {
		l.current_char = 0
	} else {
		l.current_char = rune(l.text[l.position])
	}
}

func (l *Lexer) integer() string {
	result := []string{}
	for unicode.IsDigit(l.current_char) {
		result = append(result, string(l.current_char))
		l.advance()
	}
	return strings.Join(result, "")
}
