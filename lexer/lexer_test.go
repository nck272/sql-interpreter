package lexer

import (
	"main/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		SELECT
			first_name,
			last_name,
			salary
		FROM employees
		WHERE salary > 3800;
	`

	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{token.SELECT, "SELECT"},
		{token.IDENT, "first_name"},
		{token.COMMA, ","},
		{token.IDENT, "last_name"},
		{token.COMMA, ","},
		{token.IDENT, "salary"},
		{token.FROM, "FROM"},
		{token.IDENT, "employees"},
		{token.WHERE, "WHERE"},
		{token.IDENT, "salary"},
		{token.GT, ">"},
		{token.INT, "3800"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
