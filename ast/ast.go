package ast

import "notc/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// this returns the root node of every Program
// which is just a collection of Statements
// a program can not contain individual expression like 5+3
// an expression will have to be part of statment.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type TypeStatement struct {
	Name  *Identifier
	Value Expression
	Token token.Token // type token like f32, i32
}

func (ts *TypeStatement) statementNode() {}
func (ts *TypeStatement) TokenLiteral() string {
	return ts.Token.Literal
}

type Identifier struct {
	Value string
	Token token.Token
}

func (id *Identifier) expressionNode() {}
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}
