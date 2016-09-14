package main

type Parser struct {
    lexer Lexer
    current_token Token
}

func NewParser(l Lexer) Parser {
    return Parser{lexer: l, current_token: l.get_next_token()}
}

func (p *Parser) expr() int {
    result := p.term()

    for p.current_token.kind == "PLUS" || p.current_token.kind == "MINUS" {
        token := p.current_token
        if token.kind == PLUS {
            p.consume(PLUS)
            result = result + p.term()
        }
        if token.kind == MINUS {
            p.consume(MINUS)
            result = result - p.term()
        }
    }
    return result
}

func (p *Parser) term() int {
    token := p.current_token
    p.consume(INTEGER)
    return token.toInteger()
}

func (p *Parser) consume(token_type string) {
    if p.current_token.kind == token_type {
        p.current_token = p.lexer.get_next_token()
    } else {
        panic("Invalid syntax")
    }
}
