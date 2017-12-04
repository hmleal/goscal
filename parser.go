package main

type Parser struct {
	lexer         Lexer
	current_token Token
}

// NewParser creates a parser
func NewParser(l Lexer) Parser {
	return Parser{lexer: l, current_token: l.get_next_token()}
}

func (p *Parser) expr() int {
	// expr  : term ((PLUS | MINUS) term)*
	// term  : factor ((MUL | DIV) factor)*
	// factor: INTEGER | (LPAREN expr RPAREN)
	result := p.term()

	for p.current_token.kind == PLUS || p.current_token.kind == MINUS {
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

func (p *Parser) factor() int {
	// factor: INTEGER | (LPAREN expr RPAREN)
	token := p.current_token
	result := 0
	if token.kind == INTEGER {
		p.consume(INTEGER)
		return token.toInteger()
	} else if token.kind == LPAREN {
		p.consume(LPAREN)
		result := p.expr()
		p.consume(RPAREN)
		return result
	}
	return result
}

func (p *Parser) term() int {
	// term : factor ((MUL | DIV) factor)*
	result := p.factor()

	for p.current_token.kind == MUL || p.current_token.kind == DIV {
		token := p.current_token
		if token.kind == MUL {
			p.consume(MUL)
			result = result * p.factor()
		}
		if token.kind == DIV {
			p.consume(DIV)
			result = result / p.factor()
		}
	}
	return result
}

func (p *Parser) consume(token_type string) {
	if p.current_token.kind == token_type {
		p.current_token = p.lexer.get_next_token()
	} else {
		panic("Invalid syntax")
	}
}
