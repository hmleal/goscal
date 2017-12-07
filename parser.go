package main

type Parser struct {
	lexer        Lexer
	currentToken Token
}

// NewParser creates a parser
func NewParser(l Lexer) Parser {
	return Parser{lexer: l, currentToken: l.getNextToken()}
}

// expr Check the expression
// expr  : term   ((PLUS | MINUS) term)*
// term  : factor ((MUL | DIV) factor)*
// factor: INTEGER | (LPAREN expr RPAREN)
func (p *Parser) expr() int {
	result := p.term()

	for p.currentToken.kind == PLUS || p.currentToken.kind == MINUS {
		token := p.currentToken
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

// term : factor ((MUL | DIV) factor)*
func (p *Parser) term() int {
	result := p.factor()

	for p.currentToken.kind == MUL || p.currentToken.kind == DIV {
		token := p.currentToken
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

// factor: INTEGER | (LPAREN expr RPAREN)
func (p *Parser) factor() int {
	token := p.currentToken
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

func (p *Parser) consume(token_type string) {
	if p.currentToken.kind == token_type {
		p.currentToken = p.lexer.getNextToken()
	} else {
		panic("Invalid syntax")
	}
}
