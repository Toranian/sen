package lexer

import (
	// "fmt"
	"sen/token"
	"testing"
)

// func TestPositioning(t *testing.T) {
//
// 	tests := []struct {
// 		input       string
// 		expectedCol int
// 	}{
// 		{"\t", 3},
// 		{"let\n", 4},
// 		{"three\n", 6},
// 		{"let three\n", 10},
// 	}
//
// 	for i, tt := range tests {
// 		l := New(tt.input)
// 		tok := l.NextToken()
//
// 		if l.col != tt.expectedCol {
// 			t.Fatalf("Test[%d] (%q) - Column wrong. Expected=%d, got=%d",
// 				i, tok.Literal, tt.expectedCol, l.col)
// 		}
//
// 		fmt.Println()
//
// 	}
// }

func TestNextToken(t *testing.T) {
	input := `let three = 3
let six = 6

let mult = fn(x, y) {
	x * y
}

pub class Cat {}`

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
