package lexer

import (
	"testing"

	"gjko/token"
)

func TestNextToken(t *testing.T) {
	var input string = `let five = 5;
	let ten = 10;
	let add = fn(x,y) {
		x + y;
	};
	let result = add(five, ten);
	!-/*5;
	5< 10>5;`

	/* Me tha nje shok... konkretisht si nje "reviewer"
	qe kodi me siper eshte pa kuptim s'duhet te kete kuptim,
	nuk eshte puna e lexerit te kete kuptim, lexeri duhet thjesht te kthej stringun
	input ne Tokena dhe te dalloj mes numrave dhe operatorve etj, ne "1+2=5" mjafton
	qe kupton qe {5,1} jane TokenType
	INT, {+} eshte operatori PLUS dhe {=} ASSIGN, jo te dalloj errore */

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.LET, "let"}, {token.IDENT, "ten"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.LET, "let"}, {token.IDENT, "add"}, {token.ASSIGN, "="},
		{token.FUNCTION, "fn"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"},
		{token.LBRACE, "{"}, {token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"}, {token.RBRACE, "}"}, {token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"}, {token.ASSIGN, "="}, {token.IDENT, "add"},
		{token.LPAREN, "("}, {token.IDENT, "five"}, {token.COMMA, ","},
		{token.IDENT, "ten"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},
		{token.XLAM, "!"}, {token.MINUS, "-"}, {token.SLASH, "/"}, {token.ASTER, "*"}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.INT, "5"}, {token.LESST, "<"}, {token.INT, "10"}, {token.GREATT, ">"}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got %q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got %q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
