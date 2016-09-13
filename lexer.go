package main

import "unicode"

type Lexer struct {
    text         string
    position     int
    current_char rune
}

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
            token := Token{kind: INTEGER, value: string(l.current_char)}
            l.advance()
            return token
        }

        if string(l.current_char) == "+" {
            token := Token{kind: PLUS, value: string(l.current_char)}
            l.advance()
            return token
        }

        panic("Error")
    }

    return Token{kind: EOF, value: "EOF"}
}

func (l *Lexer) advance() {
    l.position++
    if l.position > len(l.text) - 1 {
        l.current_char = 0
    } else {
        l.current_char = rune(l.text[l.position])
    }
}

func (l *Lexer) skype_whitespace() {
    for l.current_char != 0 {
        if unicode.IsSpace(l.current_char) {
            l.advance()
        }
    }
}
