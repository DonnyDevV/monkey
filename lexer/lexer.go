// The lexer will be initialized with the source code and then tokenize it, also called tokenizer or scanner
// In a production environment we would probably initialize it with an io.Reader and a filename instead of a string
// to represent the source code. To also take line numbers and column numbers into account for more descriptive errors
// The lexer only supports ASCII. In order to support Unicode we need to change the index logic and use rune instead of byte
// for currentChar to support multiple bytes.

package lexer

import "github.com/DonnyDevV/monkey/token"

type Lexer struct {
	input        string
	currentIndex int
	readIndex    int
	currentChar  byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readIndex >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readIndex]
	}
	l.currentIndex = l.readIndex
	l.readIndex++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.currentChar {
	case '=':
		tok = newToken(token.ASSIGN, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '(':
		tok = newToken(token.OPAREN, l.currentChar)
	case ')':
		tok = newToken(token.CPAREN, l.currentChar)
	case '{':
		tok = newToken(token.OBRACE, l.currentChar)
	case '}':
		tok = newToken(token.CBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}

}
