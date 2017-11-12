package lexer

import (
	"monkey/token"
	"testing"
)

type OToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func check_token_type(t *testing.T, i int, tok *token.Token, tt *OToken) {
	if tok.Type != tt.expectedType {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
	}
}

func check_token_literal(t *testing.T, i int, tok *token.Token, tt *OToken) {
	if tok.Literal != tt.expectedLiteral {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
	}
}

func check(t *testing.T, input string, tests []OToken) {
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		// t.Logf("The value of tok: %v", tok)
		check_token_type(t, i, &tok, &tt)
		check_token_literal(t, i, &tok, &tt)
	}
}

func TestNextTokenSimple(t *testing.T) {
	input := `=+(){},;`

	tests := []OToken{
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

	check(t, input, tests)
}

func TestNextTokenMoreComplex(t *testing.T) {
	input := `let five = 5;
let add = fn(x, y) {
    x + y;
};`

	tests := []OToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
	}

	check(t, input, tests)
}

func TestNextTokenMostComplex(t *testing.T) {
	input := `let five = 5;
let add = fn(x, y) {
    x + y;
};

5 == 6 < 12 !7 / 42;
if true != false { return; };
`

	tests := []OToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.EQ, "=="},
		{token.INT, "6"},
		{token.LT, "<"},
		{token.INT, "12"},
		{token.BANG, "!"},
		{token.INT, "7"},
		{token.SLASH, "/"},
		{token.INT, "42"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.TRUE, "true"},
		{token.NOT_EQ, "!="},
		{token.FALSE, "false"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
	}

	check(t, input, tests)
}
