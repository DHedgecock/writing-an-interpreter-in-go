package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedTokenType    token.TokenType
		expectedTokenLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("tests[%d] - incorrect token type. expected=%q, got=%q", i, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedTokenLiteral {
			t.Fatalf("tests[%d] - incorrect token literal. expected=%q, got=%q", i, tt.expectedTokenLiteral, tok.Literal)
		}
	}
}
