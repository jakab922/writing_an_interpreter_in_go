package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var"},
					Value: "var",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "other_var"},
					Value: "other_var",
				},
			},
		},
	}

	if program.String() != "let var = other_var;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
