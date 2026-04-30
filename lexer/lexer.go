package lexer

import (
	"notc/token"
)

type Lexer struct {
	input           string
	currentPosition int
	readPosition    int
	currentChar     byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}

	l.currentPosition = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		tok = NewToken(token.ASSIGN, l.currentChar)
	case '+':
		tok = NewToken(token.PLUS, l.currentChar)
	case ',':
		tok = NewToken(token.COMMA, l.currentChar)
	case ';':
		tok = NewToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = NewToken(token.LPAREN, l.currentChar)
	case ')':
		tok = NewToken(token.RPAREN, l.currentChar)
	case '{':
		tok = NewToken(token.LBRACE, l.currentChar)
	case '}':
		tok = NewToken(token.RBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.currentChar) {
			tok.Type = token.NUM
			tok.Literal = l.readNum()
			return tok
		}

		literal := l.ReadUntilWhiteShapeOrSpecial()
		tok.Literal = literal

		if isType(literal) {
			tok.Type = token.I32
			return tok
		}

		switch literal {
		case "fn":
			tok.Type = token.FUNCTION
		case "->":
			tok.Type = token.RETURN_TYPE
		case "return":
			tok.Type = token.RETURN
		default:
			if isAlphaNumeric(literal) {
				tok.Type = token.IDENT
			} else {
				tok.Type = token.ILLEGAL
			}
		}
		return tok
	}

	l.ReadChar()
	return tok
}

func isDelimiter(ch byte) bool {
	switch ch {
	case '=', '+', ',', ';', '(', ')', '{', '}', ' ', 0:
		return true
	}
	return false
}

func (l *Lexer) ReadUntilWhiteShapeOrSpecial() string {
	position := l.currentPosition
	for !isDelimiter(l.currentChar) {
		l.ReadChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.ReadChar()
	}
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isAlphaNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if (ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '_' {
			continue
		}
		return false
	}
	return true
}

func isType(s string) bool {
	switch s {
	case "i32":
		return true
	default:
		return false
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNum() string {
	position := l.currentPosition
	for i := l.currentPosition; ; i++ {
		if !isDigit(l.currentChar) {
			break
		}
		l.ReadChar()
	}
	return l.input[position:l.currentPosition]
}
