package lexer

import (
	"fmt"
	"testing"

	"github.com/ghoshRitesh12/interpreter_in_go/token"
)

func TestNextToken(t *testing.T) {
	type testStruct struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	input := `=+(){}[],;`

	tests := []testStruct{
		{
			expectedType:    token.ASSIGNMENT,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.PLUS,
			expectedLiteral: "+",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
		},
		{
			expectedType:    token.LSQR,
			expectedLiteral: "[",
		},
		{
			expectedType:    token.RSQR,
			expectedLiteral: "]",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.EOF,
			expectedLiteral: "",
		},
	}

	l := NewLexer(input)

	fmt.Printf("%v\n", byte(0))

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type,
			)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}

	}
}
