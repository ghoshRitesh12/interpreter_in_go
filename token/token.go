package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// unknown token/character
	ILLEGAL = "ILLEGAL"
	// end of file
	EOF = "EOF"

	// identifiers(variables) + literals(constants)
	IDENT = "IDENT" // add, fooba, x, b
	INT   = "INT"   // 167723

	// Operators
	ASSIGNMENT = "="
	PLUS       = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LSQR      = "["
	RSQR      = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
