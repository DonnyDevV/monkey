// The lexer will be initialized with the source code and then tokenize it (also called tokenizer or scanner)
// In a production environment we would probably initialize it with an io.Reader and a filename instead of a string
// to represent the source code. To also take line numbers and column numbers into account for more descriptive errors
// The lexer only supports ASCII. In order to support Unicode we need to change the index logic and use rune instead of byte
// for currentChar to support multiple bytes.

package lexer

import "github.com/DonnyDevV/monkey/token"

type Lexer struct {
	input     string
	index     int
	readIndex int
	ch        byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readIndex >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readIndex]
	}
	l.index = l.readIndex
	l.readIndex++
}

func (l *Lexer) readLetters() string {
	position := l.index
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.index]

}

func (l *Lexer) readNumber() string {
	position := l.index
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.index]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.OPAREN, l.ch)
	case ')':
		tok = newToken(token.CPAREN, l.ch)
	case '{':
		tok = newToken(token.OBRACE, l.ch)
	case '}':
		tok = newToken(token.CBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readLetters()
			tok.Type = token.CheckIdentOrKeyword(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}
