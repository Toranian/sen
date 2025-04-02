package lexer

import (
	"fmt"
	"sen/token"
	"testing"
)

func TestPositioning(t *testing.T) {

	tests := []struct {
		input       string
		expectedCol int
	}{
		{"\tHey", 7},
		{"let", 3},
		{"three", 5},
		{"\n", 1},
	}

	for i, tt := range tests {
		l := New(tt.input)
		tok := l.NextToken()

		if l.col != tt.expectedCol {
			t.Fatalf("Test[%d] (%q) - Column wrong. Expected=%d, got=%d",
				i, tok.Literal, tt.expectedCol, l.col)
		}

		fmt.Println()
	}
}

func TestNextToken(t *testing.T) {
	input := `let three = 3
let six = 6

let mult = fn(x, y) {
	x * y
}

pub class Cat {}
==
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "three"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.NEWLINE, "\n"},

		// Second line
		{token.LET, "let"},
		{token.IDENT, "six"},
		{token.ASSIGN, "="},
		{token.INT, "6"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},

		// Function
		{token.LET, "let"},
		{token.IDENT, "mult"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "x"},
		{token.ASTERISK, "*"},
		{token.IDENT, "y"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},

		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.PUB, "pub"},
		{token.CLASS, "class"},
		{token.IDENT, "Cat"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.EQ, "=="},
		{token.NEWLINE, "\n"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - TokenType wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - Literal wrong. Expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestTokenPos(t *testing.T) {
	input := `let three = 3
let three = 3
// comment`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedCol     int
		expectedRow     int
	}{
		{token.LET, "let", 1, 1},
		{token.IDENT, "three", 5, 1},
		{token.ASSIGN, "=", 11, 1},
		{token.INT, "3", 13, 1},
		{token.NEWLINE, "\n", 14, 1},

		// Second line
		{token.LET, "let", 1, 2},
		{token.IDENT, "three", 5, 2},
		{token.ASSIGN, "=", 11, 2},
		{token.INT, "3", 13, 2},
		// Comment line
		{token.NEWLINE, "\n", 14, 2},
		{token.COMMENT, "// comment", 1, 3},
		{token.EOF, "", 0, 0},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - TokenType wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - Literal wrong. Expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Pos.Col != tt.expectedCol {
			t.Fatalf("Test[%d] (%q) - Token column wrong. Expected=%d, got=%d",
				i, tok.Literal, tt.expectedCol, tok.Pos.Col)
		}

		if tok.Pos.Row != tt.expectedRow {
			t.Fatalf("Test[%d] (%q) - Token row wrong. Expected=%d, got=%d",
				i, tok.Literal, tt.expectedRow, tok.Pos.Row)
		}
	}
}

func TestLineEndings(t *testing.T) {

	input := `
return 5
return 10
return 50`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.INT, "10"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.INT, "50"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - TokenType wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - Literal wrong. Expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestTokens(t *testing.T) {

	input := `= == != +-*/ ! < > () {} [] : ,`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.BANG, "!"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.COLON, ":"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - TokenType wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - Literal wrong. Expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
