package parser

import (
	"github.com/jdodson3106/goXml2Json/internal/token"
	"testing"

	"github.com/jdodson3106/goXml2Json/internal/ast"
	"github.com/jdodson3106/goXml2Json/internal/lexer"
	"github.com/stretchr/testify/require"
)

type expectedValue struct {
	expectedValue string
	expectedToken token.TokenType
}

type expectedAttribute struct {
	key   expectedValue
	value expectedValue
}

func TestTagDefinition(t *testing.T) {
	input := `
	<name>Justin</name>
	<dob>09-27-1989</dob>
	<phone>8675309</phone>
	`

	l, err := lexer.New(input, lexer.XML)
	require.NoError(t, err)

	parser := New(l)

	doc := parser.ParseDocument()
	require.NotNil(t, doc)

	require.Equal(t, len(doc.Elements), 6)

	tests := []expectedValue{
		{"name", token.TAG},
		{"name", token.TAG},
		{"dob", token.TAG},
		{"dob", token.TAG},
		{"phone", token.TAG},
		{"phone", token.TAG},
	}

	for i, tt := range tests {
		el := doc.Elements[i]
		if !testTagDef(t, el, el.TokenLiteral(), tt) {
			return
		}
	}
}

func testTagDef(t *testing.T, el ast.ElementNode, elLit string, expected expectedValue) bool {
	if el.TokenLiteral() != elLit {
		t.Errorf("el.TokenLiteral not '%s'. got=%q", elLit, el.TokenLiteral())
		return false
	}

	tagEl, ok := el.(*ast.ElementTagNode)
	if !ok {
		t.Errorf("el not *ast.ElementTagNode. got=%T", el)
		return false
	}

	if tagEl.Token.Literal != expected.expectedValue {
		t.Errorf("tagEl.Token.Literal not '%s'. got=%s", expected.expectedValue, tagEl.Token.Literal)
		return false
	}

	if tagEl.Token.Type != expected.expectedToken {
		t.Errorf("tagEl.Token.Type not '%s'. got=%v", expected.expectedToken, tagEl.Token.Type)
		return false
	}

	if tagEl.TokenLiteral() != expected.expectedValue {
		t.Errorf("tagEl.TokenLiteral not '%s'. got=%s", expected.expectedValue, tagEl.TokenLiteral())
		return false
	}

	return true
}

func TestParseAttributes(t *testing.T) {
	input := `
	<person name="Justin" dob="09-27-1989" phone="8675309" />
	<pet type="dog">Benji</pet>
	`
	l, err := lexer.New(input, lexer.XML)
	require.NoError(t, err)

	parser := New(l)

	doc := parser.ParseDocument()
	require.NotNil(t, doc)

	require.Equal(t, len(doc.Elements), 4)

	expectedAttributes := []expectedAttribute{
		{
			key:   expectedValue{"name", token.KEY},
			value: expectedValue{"Justin", token.VALUE},
		},
		{
			key:   expectedValue{"dob", token.KEY},
			value: expectedValue{"09-27-1989", token.VALUE},
		},
		{
			key:   expectedValue{"phone", token.KEY},
			value: expectedValue{"8675309", token.VALUE},
		},
		{
			key:   expectedValue{"type", token.KEY},
			value: expectedValue{"dog", token.VALUE},
		},
	}
}
