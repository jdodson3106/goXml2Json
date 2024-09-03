package tests

import (
	parser2 "github.com/jdodson3106/goXml2Json/internal/parser"
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
	input := string(loadDataFile(t, "tagDefTest.xml"))

	l, err := lexer.New(input, lexer.XML)
	require.NoError(t, err)

	parser := parser2.New(l)

	doc := parser.ParseDocument()
	require.NotNil(t, doc)

	require.Equal(t, 3, len(doc.Elements))

	tests := []ast.ElementTagNode{
		{
			Token: token.Token{
				Type:    token.TAG,
				Literal: "name",
			},
			Value: ast.ElementValueNode{
				Token: token.Token{
					Type:    token.VALUE,
					Literal: "Justin",
				},
				Value: "Justin",
			},
			EndToken: token.Token{
				Type:    token.TAG,
				Literal: "name",
			},
		},
		{
			Token: token.Token{
				Type:    token.TAG,
				Literal: "dob",
			},
			Value: ast.ElementValueNode{
				Token: token.Token{
					Type:    token.VALUE,
					Literal: "09-27-1989",
				},
				Value: "09-27-1989",
			},
			EndToken: token.Token{
				Type:    token.TAG,
				Literal: "dob",
			},
		},
		{
			Token: token.Token{
				Type:    token.TAG,
				Literal: "phone",
			},
			Value: ast.ElementValueNode{
				Token: token.Token{
					Type:    token.VALUE,
					Literal: "8675309",
				},
				Value: "8675309",
			},
			EndToken: token.Token{
				Type:    token.TAG,
				Literal: "phone",
			},
		},
	}

	for i, tt := range tests {
		el := doc.Elements[i]
		require.Equal(t, el, &tt)
	}
}

func TestAttributeDefinition(t *testing.T) {
	input := string(loadDataFile(t, "tagAttributeTest.xml"))
	l, err := lexer.New(input, lexer.XML)
	require.NoError(t, err)

	parser := parser2.New(l)

	doc := parser.ParseDocument()
	require.NotNil(t, doc)

	require.Equal(t, 3, len(doc.Elements))

	tests := []ast.ElementTagNode{
		{
			Token: token.Token{
				Type:    token.TAG,
				Literal: "name",
			},
			Attributes: []*ast.ElementAttributeNode{
				{
					Key: &ast.AttributeKeyNode{
						Token: token.Token{
							Type:    token.KEY,
							Literal: "value",
						},
						Value: "value",
					},
					Value: &ast.AttributeValueNode{
						Token: token.Token{
							Type:    token.VALUE,
							Literal: "Justin",
						},
						Value: "Justin",
					},
				},
			},
			EndToken: token.Token{
				Type:    token.CLOSE_ANGLE,
				Literal: ">",
			},
		},
		{
			Token: token.Token{
				Type:    token.TAG,
				Literal: "dob",
			},
			Attributes: []*ast.ElementAttributeNode{
				{
					Key: &ast.AttributeKeyNode{
						Token: token.Token{
							Type:    token.KEY,
							Literal: "value",
						},
						Value: "value",
					},
					Value: &ast.AttributeValueNode{
						Token: token.Token{
							Type:    token.VALUE,
							Literal: "09-27-1989",
						},
						Value: "09-27-1989",
					},
				},
			},
			EndToken: token.Token{
				Type:    token.TAG,
				Literal: "dob",
			},
		},
		{
			Token: token.Token{
				Type:    token.TAG,
				Literal: "ssn",
			},
			Attributes: []*ast.ElementAttributeNode{
				{
					Key: &ast.AttributeKeyNode{
						Token: token.Token{
							Type:    token.KEY,
							Literal: "value",
						},
						Value: "value",
					},
					Value: &ast.AttributeValueNode{
						Token: token.Token{
							Type:    token.VALUE,
							Literal: "999999999",
						},
						Value: "999999999",
					},
				},
			},
			EndToken: token.Token{
				Type:    token.CLOSE_ANGLE,
				Literal: ">",
			},
		},
	}

	for i, tt := range tests {
		el := doc.Elements[i]
		require.Equal(t, &tt, el)
	}
}

//func testTagDef(t *testing.T, el ast.ElementNode, elLit string, expected ast.ElementTagNode) bool {
//	if el.TokenLiteral() != elLit {
//		t.Errorf("el.TokenLiteral not '%s'. got=%q", elLit, el.TokenLiteral())
//		return false
//	}
//
//	tagEl, ok := el.(*ast.ElementTagNode)
//	if !ok {
//		t.Errorf("el not *ast.ElementTagNode. got=%T", el)
//		return false
//	}
//
//	if tagEl.Token.Literal != expected.expectedValue {
//		t.Errorf("tagEl.Token.Literal not '%s'. got=%s", expected.expectedValue, tagEl.Token.Literal)
//		return false
//	}
//
//	if tagEl.Token.Type != expected.expectedToken {
//		t.Errorf("tagEl.Token.Type not '%s'. got=%v", expected.expectedToken, tagEl.Token.Type)
//		return false
//	}
//
//	if tagEl.TokenLiteral() != expected.expectedValue {
//		t.Errorf("tagEl.TokenLiteral not '%s'. got=%s", expected.expectedValue, tagEl.TokenLiteral())
//		return false
//	}
//
//	return true
//}
