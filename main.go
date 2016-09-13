package main

import (
    "fmt"
    "bufio"
    "os"
)

type Parser struct {
    lexer Lexer
    current_token Token
}

func NewParser(l Lexer) Parser {
    return Parser{lexer: l, current_token: l.get_next_token()}
}

func (p *Parser) expr() {
    for p.current_token.value != "EOF" {
        fmt.Println(p.current_token.kind)
        p.current_token = p.lexer.get_next_token()
    }
}

type Interpreter struct {
}

func main() {
    fmt.Println("Welcome a implementation of Pascal in GoLang")

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">>> ")
        text, _ := reader.ReadString('\n')

        lexer := NewLexer(text)
        parser := NewParser(lexer)
        parser.expr()
    }
}
