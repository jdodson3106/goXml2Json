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

	// queue up the first two tokens into current and peek
	p.nextToken()
	p.nextToken()

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
	default:
		return nil
	}
}

func (p *Parser) parseTagStatement() ast.ElementNode {
	tag := &ast.ElementTagNode{Token: p.currentToken}

	// example: we come in with p.currentToken = {token.TAG, "tag"}
	// <tag key="value" />
	// <tag key="value"></tag>

	// parse all attributes from statement
	for p.expectPeek(token.KEY) {
		tag.Attributes = append(tag.Attributes, p.parseAttribute())
	}

	// this means there is no value, so the tag has an early termination like <tag />
	if p.expectPeek(token.XML_TERMINATOR) {
		// validate the last token is the '>' char
		if !p.expectPeek(token.CLOSE_ANGLE) {
			p.errors = append(p.errors, "missing closing angle at element tag termination")
			return nil
		}
		tag.EndToken = p.currentToken
		return tag
	}

	// if the next tag after the attrs is not a close angle then fail
	if !p.expectPeek(token.CLOSE_ANGLE) {
		p.errors = append(p.errors, "expected closing angle tag for tag")
		return nil
	}

	// example: we come in with p.currentToken = {token.TAG, "tag"}
	// <tag key="value"></tag>

	// if the next token is not a value then it must have children or nothing
	// this is the recursive bit of the recursive descent parser
	if !p.expectPeek(token.VALUE) && p.expectPeek(token.OPEN_ANGLE) {
		// the token is an empty element (<tag></tag>)
		if p.expectPeek(token.XML_TERMINATOR) {
			if !p.expectPeek(token.TAG) {
				p.errors = append(p.errors, "missing tag at element tag termination")
				return nil
			}
			tag.EndToken = p.currentToken
			return tag
		}

		child := p.parseTagStatement()
		if child != nil {
			tag.Elements = append(tag.Elements, &child)
		} else {
			p.errors = append(p.errors, "error parsing child element")
			return nil
		}
	}

	if p.currTokenIs(token.VALUE) {
		tag.Value = ast.ElementValueNode{
			Token: p.currentToken,
			Value: p.currentToken.Literal,
		}
	}

	// the next token should be the opening of the next tag
	if !p.expectPeek(token.OPEN_ANGLE) {
		p.errors = append(p.errors, "Invalid xml syntax. Expected open angle for element tag")
		return nil
	}

	if !p.expectPeek(token.XML_TERMINATOR) || !p.expectPeek(token.TAG) {
		p.errors = append(p.errors, "no closing tag for element.")
		return nil
	}
	tag.EndToken = p.currentToken
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

	// make sure the next token in an opening single or double quote to hold the value
	if !p.expectPeek(token.QUOTE) && !p.expectPeek(token.SINGLE_QUOTE) {
		p.errors = append(p.errors, "XML element attribute values must be wrapped in quotes.")
		return nil
	}
	quoteType = p.currentToken.Type // store the opening quote so we can make sure we have a valid match

	// confirm the next token is actually a value and use to construct the attribute value node
	if !p.expectPeek(token.VALUE) {
		p.errors = append(p.errors, fmt.Sprintf("Expected token.VALUE, got %v", p.currentToken.Type))
		return nil
	}
	val := &ast.AttributeValueNode{Token: p.currentToken, Value: p.currentToken.Literal}

	// make sure there is a closing quote
	if !p.expectPeek(token.QUOTE) && !p.expectPeek(token.SINGLE_QUOTE) {
		p.errors = append(p.errors, "XML element attribute missing closing quote.")
		return nil
	}

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
