package ast

import (
	"bytes"
	"notc/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type TypeStatement struct {
	TypeName *Identifier
	Value    Expression
	Token    token.Token // type token like f32, i32
}

func (ts *TypeStatement) statementNode() {}
func (ts *TypeStatement) TokenLiteral() string {
	return ts.Token.Literal
}

func (ts *TypeStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ts.Token.Literal + " ")
	out.WriteString(ts.TypeName.String())

	if ts.Value != nil {
		out.WriteString(ts.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	IdentName string
	Token     token.Token
}

func (id *Identifier) expressionNode() {}
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}
func (id *Identifier) String() string {
	return id.IdentName
}

type ReturnStatement struct {
	ReturnValue Expression
	Token       token.Token
}

func (st *ReturnStatement) statementNode() {}
func (st *ReturnStatement) TokenLiteral() string {
	return st.Token.Literal
}

func (st *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(st.TokenLiteral() + " ")
	if st.ReturnValue != nil {
		out.WriteString(st.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
