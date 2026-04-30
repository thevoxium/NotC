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
		if l.peakChar() == '=' {
			tok = token.Token{Type: token.EQUAL, Literal: "=="}
			l.ReadChar()
		} else {
			tok = NewToken(token.ASSIGN, l.currentChar)
		}
	case '+':
		tok = NewToken(token.PLUS, l.currentChar)
	case '-':
		tok = NewToken(token.MINUS, l.currentChar)
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
	case '<':
		tok = NewToken(token.LARROW, l.currentChar)
	case '>':
		tok = NewToken(token.RARROW, l.currentChar)
	case '!':
		if l.peakChar() == '=' {
			tok = token.Token{Type: token.NOT_EQUAL, Literal: "!="}
			l.ReadChar()
		} else {
			tok = NewToken(token.BANG, l.currentChar)
		}
	case '*':
		tok = NewToken(token.ASTERISK, l.currentChar)
	case '/':
		tok = NewToken(token.SLASH, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:

		if isDigit(l.currentChar) {
			tok.Literal, tok.Type = l.readNum()
			return tok
		}

		literal := l.ReadUntilWhiteShapeOrSpecial()
		tok.Literal = literal

		typeCheck := isType(literal)
		if typeCheck != token.NONE {
			tok.Type = typeCheck
			return tok
		}

		switch literal {
		case "fn":
			tok.Type = token.FUNCTION
		case "return":
			tok.Type = token.RETURN
		case "true":
			tok.Type = token.TRUE
		case "false":
			tok.Type = token.FALSE
		case "if":
			tok.Type = token.IF
		case "else":
			tok.Type = token.ELSE
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
	case '>', '<', '=', '+', '-', ',', ';', '(', ')', '{', '}', '*', '/', '!', ' ', 0:
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

func isType(s string) token.TokenType {
	switch s {
	case "i32":
		return token.I32
	case "f32":
		return token.F32
	default:
		return token.NONE
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNum() (string, token.TokenType) {
	position := l.currentPosition
	hasDot := false
	for isDelimiter(l.currentChar) == false {
		if isDigit(l.currentChar) {
		} else if l.currentChar == '.' && !hasDot {
			hasDot = true
		} else {
			break
		}
		l.ReadChar()
	}

	literal := l.input[position:l.currentPosition]
	if hasDot {
		return literal, token.FLOATNUM
	}
	return literal, token.INTNUM
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
