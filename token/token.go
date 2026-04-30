package token

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	I32       = "INT32"
	F32       = "FLOAT32"
	INTNUM    = "INTNUM"
	FLOATNUM  = "FLOATNUM"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "fn"
	RETURN    = "RETURN"
	RARROW    = ">"
	LARROW    = "<"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	EQUAL     = "EQ"
	NOT_EQUAL = "NOT_EQUAL"
	NONE      = "NONE"
)

type Token struct {
	Type    TokenType
	Literal string
}
