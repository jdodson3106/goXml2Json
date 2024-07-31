package lexer

import (
	"testing"

	"github.com/jdodson3106/goXml2Json/internal/token"
	"github.com/stretchr/testify/require"
)

func TestNextToken(t *testing.T) {
	xmlInput := `<name>Justin</name>`

	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.OPEN_ANGLE, "<"},
		{token.TAG, "name"},
		{token.CLOSE_ANGLE, ">"},
		{token.VALUE, "Justin"},
		{token.OPEN_ANGLE, "<"},
		{token.XML_TERMINATOR, "/"},
		{token.TAG, "name"},
		{token.CLOSE_ANGLE, ">"},
	}

	lex, err := New(xmlInput, XML)
	require.NoError(t, err)

	for i, tc := range testCases {
		tok := lex.NextToken()

		if tok.Type != tc.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", i, tc.expectedLiteral, tok.Literal)
		}
	}
}
