package parser

import (
	"fmt"

	"github.com/jdodson3106/goXml2Json/internal/ast"
	"github.com/jdodson3106/goXml2Json/internal/lexer"
	"github.com/jdodson3106/goXml2Json/internal/token"
)

const (
	JSON = "json"
	XML  = "xml"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	fmt.Printf("New() Called: current tokens | %v, %v\n", p.currentToken, p.peekToken)
	return p
}

func (p *Parser) ParseDocument() *ast.Document {
	doc := &ast.Document{}
	doc.Elements = []ast.ElementNode{}

	for p.currentToken.Type != token.EOF {
		el := p.parseElement()
		if el != nil {
			doc.Elements = append(doc.Elements, el)
		}
		p.nextToken()
	}

	return doc
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseElement() ast.ElementNode {
	switch p.currentToken.Type {
	case token.TAG:
		return p.parseTagStatement()
	case token.KEY:
		return p.parseAttribute()
	default:
		return nil
	}
}

func (p *Parser) parseTagStatement() ast.ElementNode {
	tag := &ast.ElementTagNode{Token: p.currentToken}

	if !p.expectPeek(token.CLOSE_ANGLE) {
		return nil
	}

	return tag
}

func (p *Parser) parseAttribute() *ast.ElementAttributeNode {
	var quoteType token.TokenType

	// set the attribute key and makes sure the next token is an equal sign
	key := &ast.AttributeKeyNode{Token: p.currentToken, Value: p.currentToken.Literal}
	if !p.expectPeek(token.EQUAL) {
		p.errors = append(p.errors, fmt.Sprintf("Expected '=', got %v", p.currentToken.Type))
		return nil
	}
	p.nextToken()

	// make sure the next token in an opening single or double quote to hold the value
	if !p.expectPeek(token.QUOTE) && !p.expectPeek(token.SINGLE_QUOTE) {
		p.errors = append(p.errors, "XML element attribute values must be wrapped in quotes.")
		return nil
	}
	p.nextToken()
	quoteType = p.currentToken.Type // store the opening quote so we can make sure we have a valid match

	// confirm the next token is actually a value and use to construct the attribute value node
	if !p.expectPeek(token.VALUE) {
		p.errors = append(p.errors, fmt.Sprintf("Expected token.VALUE, got %v", p.currentToken.Type))
		return nil
	}
	p.nextToken()
	val := &ast.AttributeValueNode{Token: p.currentToken, Value: p.currentToken.Literal}

	// make sure there is a closing quote
	if !p.expectPeek(token.QUOTE) && !p.expectPeek(token.SINGLE_QUOTE) {
		p.errors = append(p.errors, "XML element attribute missing closing quote.")
		return nil
	}
	p.nextToken()

	// assert the closing and opening quotes match
	if p.currentToken.Type != quoteType {
		p.errors = append(p.errors, fmt.Sprintf("Mismatching quotes for value '%s'", val.Value))
		return nil
	}

	return &ast.ElementAttributeNode{Key: key, Value: val}
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
