package token

// We can use int or byte for better performance
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Seperators
	SEMICOLON = ";"
	COMMA     = ","

	OPAREN = "("
	CPAREN = ")"
	OBRACE = "{"
	CBRACE = "}"
)
