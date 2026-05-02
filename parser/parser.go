package parser

import (
	"fmt"
	"notc/ast"
	"notc/lexer"
	"notc/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.TokenType]int{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,
	token.LARROW:    LESSGREATER,
	token.RARROW:    LESSGREATER,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.SLASH:     PRODUCT,
	token.ASTERISK:  PRODUCT,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l              *lexer.Lexer
	currToken      token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INTNUM, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOATNUM, p.parseFloatLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) registerPrefix(tokType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokType] = fn
}

func (p *Parser) registerInfix(tokType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokType] = fn
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
	fmt.Printf("expected next token to be %s, got %s instead\n",
		t, p.peekToken.Type)
	return false
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
	case token.RETURN:
		return p.parseReturnStatements()
	default:
		return p.parseExpressionStatements()
	}
}

func (p *Parser) parseTypeStatements() *ast.TypeStatement {
	stmt := &ast.TypeStatement{Token: p.currToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.TypeName = &ast.Identifier{IdentName: p.currToken.Literal, Token: p.currToken}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenExpected(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatements() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	p.nextToken()

	for !p.currTokenExpected(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatements() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseExpression(order int) ast.Expression {
	prefixFn := p.prefixParseFns[p.currToken.Type]
	if prefixFn == nil {
		return nil
	}
	leftExp := prefixFn()
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{IdentName: p.currToken.Literal, Token: p.currToken}
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	f := &ast.FloatLiteral{Token: p.currToken}
	value, _ := strconv.ParseFloat(p.currToken.Literal, 32)
	f.Value = float32(value)
	return f
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	it := &ast.IntegerLiteral{Token: p.currToken}
	value, _ := strconv.ParseInt(p.currToken.Literal, 0, 64)
	it.Value = value
	return it
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	pf := &ast.PrefixExpression{Token: p.currToken, Operator: p.currToken.Literal}

	p.nextToken()
	right := p.parseExpression(PREFIX)
	pf.Right = right

	return pf
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
