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

// this functions constructs the root node of the AST (ast.Program) and iterates through every token with Parser.NextToken until EOF
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		//if it doesn't return nil, it returns an ast.Statement and it's added to the statement part of the AST root node.
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}
	//when there is nothing left to parse the *ast.Program node is returned
	return program
}

/* FIXME: this comment sucks */
// This function parses a single statement (🙄 duh)
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	//TODO: add other cases for other possible statement branches
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	//TODO: other expressions
	for !p.curTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		return false
	}
}
