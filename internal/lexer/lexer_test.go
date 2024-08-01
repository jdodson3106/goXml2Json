package lexer

import (
	"fmt"
	"testing"

	"github.com/jdodson3106/goXml2Json/internal/token"
	"github.com/stretchr/testify/require"
)

type TokenTestCase struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func runNextTokenChecks(lex *Lexer, testCases []TokenTestCase, t *testing.T) {
	for i, tc := range testCases {
		tok := lex.NextToken()

		fmt.Printf("[%d] checking case %+v :: tok=%+v\n", i, tc, tok)

		if tok.Type != tc.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", i, tc.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	xmlInput := `<name>Justin</name>`

	testCases := []TokenTestCase{
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
	runNextTokenChecks(lex, testCases, t)
	fmt.Println()
}

func TestAttributesNextToken(t *testing.T) {
	xmlInput := `<name category="given-name">Justin</name>`

	testCases := []TokenTestCase{
		{token.OPEN_ANGLE, "<"},
		{token.TAG, "name"},
		{token.KEY, "category"},
		{token.EQUAL, "="},
		{token.QUOTE, "\""},
		{token.VALUE, "given-name"},
		{token.QUOTE, "\""},
		{token.CLOSE_ANGLE, ">"},
		{token.VALUE, "Justin"},
		{token.OPEN_ANGLE, "<"},
		{token.XML_TERMINATOR, "/"},
		{token.TAG, "name"},
		{token.CLOSE_ANGLE, ">"},
	}

	lex, err := New(xmlInput, XML)
	require.NoError(t, err)
	runNextTokenChecks(lex, testCases, t)
	fmt.Println()
}
