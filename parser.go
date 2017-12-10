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
// factor: (PLUS | MINUS) factor | INTEGER | (LPAREN expr RPAREN)
func (p *Parser) expr() Node {
	result := p.term()

	for p.currentToken.kind == PLUS || p.currentToken.kind == MINUS {
		token := p.currentToken
		if token.kind == PLUS {
			p.consume(PLUS)
		}
		if token.kind == MINUS {
			p.consume(MINUS)
		}
		result = binOP{left: result, token: token, right: p.term()}
	}
	return result
}

func (p *Parser) term() Node {
	result := p.factor()

	for p.currentToken.kind == MUL || p.currentToken.kind == DIV {
		token := p.currentToken
		if token.kind == MUL {
			p.consume(MUL)
		}
		if token.kind == DIV {
			p.consume(DIV)
		}
		result = binOP{left: result, token: token, right: p.factor()}
	}
	return result
}

func (p *Parser) factor() Node {
	// (PLUS | MINUS) factor | INTEGER | (LPAREN expr RPAREN)
	token := p.currentToken

	if token.kind == PLUS {
		p.consume(PLUS)
		return unaryOP{token: token, expr: p.expr()}
	}
	if token.kind == MINUS {
		p.consume(MINUS)
		return unaryOP{token: token, expr: p.expr()}
	}
	if token.kind == INTEGER {
		p.consume(INTEGER)
		return num{token}
	}
	if token.kind == LPAREN {
		p.consume(LPAREN)
		result := p.expr()
		p.consume(RPAREN)
		return result
	}
	return num{Token{}}
}

func (p *Parser) consume(tokenType string) {
	if p.currentToken.kind == tokenType {
		p.currentToken = p.lexer.getNextToken()
	} else {
		panic("Invalid syntax")
	}
}

type Interpreter struct {
	parser Parser
}

// NewInterpreter creates a simple Interpreter struct
func (i *Interpreter) NewInterpreter(p Parser) Interpreter {
	return Interpreter{parser: p}
}

// Interpret parse a expression
func (i *Interpreter) Interpret() int {
	root := i.parser.expr()
	return root.visit(root)
}
