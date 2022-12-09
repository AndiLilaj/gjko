package ast

import "gjko/token"

//this is the most basic building block
type Node interface {
	//the raw string-form of the token
	TokenLiteral() string
}

//there's gonna be a bunch of these in a program, it has raw token literals and the statement node
type Statement interface {
	Node
	statementNode()
}

//this is a type of statent it's goint to keep the token literals and the expression node
type Expression interface {
	Node
	expressionNode()
}

// the root node of every tree
type Program struct {
	//the Program contains every statement
	Statements []Statement
}

//returns the string of the Program.Statement.Node.TokenLiteral
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token //the .LET token
	Name  *Identifier //to hold the identifier/name of the binding
	Value Expression  //for theexpression that produces the value
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string      //literal
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
