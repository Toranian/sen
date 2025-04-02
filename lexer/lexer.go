package lexer

import (
	"sen/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // Current position of the ch (traces the readPosition char)
	readPosition int  // Current reading position (char directly ahead of position char)
	ch           rune // The current char. Rune is used so we can allow UTF-8
	row          int  // Row of the char
	col          int  // Col of the char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Get the first position of the input
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch), Pos: token.TokenPosition{Row: l.row, Col: l.col}}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l)
		}

	// Operators
	case '+':
		tok = newToken(token.PLUS, l.ch, l)
	case '-':
		tok = newToken(token.MINUS, l.ch, l)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l)
	case '/':

		if l.peekChar() == '/' {
			// ch := l.ch
			tok.Pos = token.TokenPosition{Row: l.row, Col: l.col}
			tok.Literal = l.readComment()
			tok.Type = token.COMMENT
			return tok

		} else {
			tok = newToken(token.SLASH, l.ch, l)
		}

	// Comparisons
	case '<':
		tok = newToken(token.LT, l.ch, l)
	case '>':
		tok = newToken(token.GT, l.ch, l)

	// Other
	case '(':
		tok = newToken(token.LPAREN, l.ch, l)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l)

	case '[':
		tok = newToken(token.LBRACKET, l.ch, l)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l)
	case ':':
		tok = newToken(token.COLON, l.ch, l)

	case ',':
		tok = newToken(token.COMMA, l.ch, l)

	case '\n':
		tok = newToken(token.NEWLINE, l.ch, l)
	case '\r':
		tok = newToken(token.CARRIAGE, l.ch, l)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

		// If it's none of our tokens, then it's gotta be an identifier!
	default:
		if unicode.IsLetter(l.ch) {
			tok.Pos = token.TokenPosition{Col: l.col, Row: l.row}
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // We've got the string, check to see if it's a keyword
			return tok
		} else if unicode.IsDigit(l.ch) {
			tok.Pos = token.TokenPosition{Col: l.col, Row: l.row}
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l)
		}

	}

	l.readChar()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {

	if l.readPosition >= len(l.input) {
		return 0
	} else {
		r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		return r
	}
}

func newToken(tokenType token.TokenType, ch rune, l *Lexer) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Pos: token.TokenPosition{Row: l.row, Col: l.col}}
}

// Read an identifier. Increment position until we encounter a non-letter char.
// Return the identifier.
func (l *Lexer) readIdentifier() string {
	position := l.position

	for unicode.IsLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readChar() {

	// Check first for the end of the file
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for the "NULL" character
	} else {

		if l.ch == '\n' || l.ch == '\r' {
			l.row++
			l.col = 1
		} else if l.ch == '\t' {
			l.col += 4
		} else {
			l.col++
		}

		// Move the char along to the next position
		r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r

	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readComment() string {

	position := l.position

	for l.ch != '\n' && l.ch != '\r' && l.ch != 0 {
		l.readChar()
	}

	return l.input[position:l.position]
}
