package main

type Parser struct {
    lexer Lexer
    current_token Token
}

func NewParser(l Lexer) Parser {
    return Parser{lexer: l, current_token: l.get_next_token()}
}

func (p *Parser) term() int {
    token := p.current_token
    p.current_token = p.lexer.get_next_token()
    return token.toInteger()
}

func (p *Parser) expr() int {
    result := p.term()

    for p.current_token.kind == "PLUS" || p.current_token.kind == "MINUS" {
        token := p.current_token
        if token.kind == PLUS {
            p.current_token = p.lexer.get_next_token()
            result = result + p.term()
        }
        if token.kind == MINUS {
            p.current_token = p.lexer.get_next_token()
            result = result - p.term()
        }
    }
    return result
}
