package ketchup

import "testing"

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedTokenType TokenType
		expectedLiteral   string
	}{
		{ASSIGN, "="},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("tests[%d] - incorrect token type. expected=%q, got=%q", i, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - incorrect literal. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
