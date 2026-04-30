package token

type TokenType string

const (
	ILLEGAL     = "ILLEGAL"
	EOF         = "EOF"
	IDENT       = "IDENT"
	I32         = "INT32"
	NUM         = "NUM"
	ASSIGN      = "="
	PLUS        = "+"
	COMMA       = ","
	SEMICOLON   = ";"
	LPAREN      = "("
	RPAREN      = ")"
	LBRACE      = "{"
	RBRACE      = "}"
	FUNCTION    = "fn"
	RETURN      = "RETURN"
	RETURN_TYPE = "->"
)

type Token struct {
	Type    TokenType
	Literal string
}
