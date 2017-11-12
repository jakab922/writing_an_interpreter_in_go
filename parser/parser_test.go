package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

type test_ident struct {
	expectedIdentifier string
}

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 3;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []test_ident{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// Since this is a let statement the token_literal should be "let"
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral no 'let'. got=%q", s.TokenLiteral())
		return false
	}

	// We check if it's really a let statement struct
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	/* We check if the name of the identifier is `name`.
	   The Name attribute of the let statement is an Identifier
	   and its Value is the name of the identifier. A bit confusing... */
	if letStmt.Name.Value != name {
		t.Errorf(
			"letStmt.Name.Value not '%s'. got=%s",
			name, letStmt.Name.Value)
		return false
	}

	// This is superfluous. We shouldn't really check this again...
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}
