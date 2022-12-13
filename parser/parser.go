package parser

import (
	"gjko/ast"
	"gjko/lexer"
	"gjko/token"
)

type Parser struct {
	//pointer to an instance of the Lexer
	l *lexer.Lexer

	//the parser needs to look at the current token to know what to do and id curToken doesn't give enough info, it peek

	//the current token
	curToken token.Token
	//the next token
	peekToken token.Token
}

// start of parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//initializes curToken and peekToken by reading two tokens
	p.l.NextToken()
	p.l.NextToken()
	return p
}

// small helper that advances advances both curToken and peekToken
func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
