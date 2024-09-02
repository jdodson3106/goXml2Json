package ast

import (
	"strings"

	"github.com/jdodson3106/goXml2Json/internal/token"
)

type Node interface {
	TokenLiteral() string
}

type ElementNode interface {
	Node
	elementNode()
}

type AttributeNode interface {
	Node
	attributeNode()
}

type ElementValueNode interface {
	Node
	elementValueNode()
}

// Document: the root node of all xml files to be parsed
type Document struct {
	Elements []ElementNode
}

func (d *Document) TokenLiteral() string {
	if len(d.Elements) > 0 {
		return d.Elements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ElementTagNode: the base node of all XML elements
type ElementTagNode struct {

	// Toke the actual token and literal of the ElementNode
	// this will usually be a "TAG" token with the tag
	//	name in the literal property
	Token token.Token

	// Attributes contains all the potential key/value attributes
	// that may be present on a given Element
	Attributes []*ElementAttributeNode

	// Elements are pointers to all the children Element tags
	Elements []*ElementTagNode

	// Value is the value of the element.
	// Typically this will be nil if the Elements property is
	// not nil and vice versa
	Value ElementValueNode
}

func (e *ElementTagNode) elementNode()         {}
func (e *ElementTagNode) TokenLiteral() string { return e.Token.Literal }

// ElementAttributeNode: represents a key/value pair of attributes on an xml element
type ElementAttributeNode struct {
	// Key is a pointer to the AttributeKeyNode that
	// contains the Token and string value of the attribute's key
	Key *AttributeKeyNode

	// Value is a pointer to an AttributeValueNode that
	// contains the Token and string value of the attribute's value
	Value *AttributeValueNode
}

func (e *ElementAttributeNode) elementAttributeNode() {}
func (e *ElementAttributeNode) TokenLiteral() string {
	var builder strings.Builder
	builder.WriteString("Key=")
	builder.WriteString(e.Key.Token.Literal)
	builder.WriteString(" Value=")
	builder.WriteString(e.Value.Token.Literal)
	return builder.String()
}

// AttributeKeyNode: holds the Token and string value of the key an element atrribute
type AttributeKeyNode struct {
	Token token.Token
	Value string
}

func (a *AttributeKeyNode) attributeNode()       {}
func (a *AttributeKeyNode) TokenLiteral() string { return a.Token.Literal }

// AttributeValueNode: holds the Token and string value of the value on an element atrribute
type AttributeValueNode struct {
	Token token.Token
	Value string
}

func (a *AttributeValueNode) attributeNode()       {}
func (a *AttributeValueNode) TokenLiteral() string { return a.Token.Literal }
