package token

// We can use int or byte for better performance
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func CheckIdentOrKeyword(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
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
