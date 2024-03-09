package ketchup

import "testing"

type nextTokenTest struct {
	expectedTokenType TokenType
	expectedLiteral   string
}

func assertNextTokenTests(t *testing.T, tests []nextTokenTest, input string) {
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

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []nextTokenTest{
		{ASSIGN, "="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{EOF, ""},
	}
	assertNextTokenTests(t, tests, input)

	input = "let five = 5; let ten = 10;"
	tests = []nextTokenTest{
		{LET, "let"},
		{IDENT, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "ten"},
		{ASSIGN, "="},
		{INT, "10"},
		{SEMICOLON, ";"},
	}
	assertNextTokenTests(t, tests, input)
}
