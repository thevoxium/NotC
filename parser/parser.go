package parser

import (
	"notc/ast"
	"notc/lexer"
	"notc/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.I32:
		return p.parseTypeStatements()
	case token.F32:
		return p.parseTypeStatements()
	default:
		return nil
	}
}

func (p *Parser) parseTypeStatements() *ast.TypeStatement {
	stmt := &ast.TypeStatement{Token: p.currToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.TypeName = &ast.Identifier{IdentName: p.currToken.Literal, Token: p.currToken}
	return stmt
}

func (p *Parser) currTokenExpected(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenExpected(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenExpected(t) {
		p.nextToken()
		return true
	}
	return false
}
