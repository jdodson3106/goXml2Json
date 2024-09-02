package lexer

import (
	"fmt"

	"github.com/jdodson3106/goXml2Json/internal/token"
)

const (
	JSON = "json"
	XML  = "xml"
)

type Lexer struct {
	lexType         string
	input           string
	currentPosition int // current char in the input
	nextPosition    int // next position in the input
	lastRead        byte
	ch              byte // current char being read
}

func New(input, lexType string) (*Lexer, error) {
	if lexType != JSON && lexType != XML {
		return nil, fmt.Errorf("invalid lexer type %s", lexType)
	}

	l := &Lexer{input: input, lexType: lexType}
	l.readChar()
	return l, nil
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0 // set the current char to 0 (ASCII NULL value)
	} else {
		l.lastRead = l.ch
		l.ch = l.input[l.nextPosition]
	}
	l.currentPosition = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.eatWhitespace()

	if l.lexType == JSON {
		t = l.nextJsonToken()
	} else {
		t = l.nextXmlToken()
	}

	// there was no type specific token
	if t.Type == "" {
		switch l.ch {
		case ',':
			t = newToken(token.COMMA, l.ch)
		case ':':
			t = newToken(token.COLON, l.ch)
		case '=':
			t = newToken(token.EQUAL, l.ch)
		case '\'':
			t = newToken(token.SINGLE_QUOTE, l.ch)
		case '"':
			t = newToken(token.QUOTE, l.ch)
		case 0:
			t.Literal = ""
			t.Type = token.EOF
		default:
			if isAlphaNumeric(l.ch) {
				t = l.readIdentifier()
				return t
			} else {
				t = newToken(token.ILLEGAL, l.ch)
			}
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) nextXmlToken() token.Token {
	var t token.Token

	switch l.ch {
	case '<':
		t = newToken(token.OPEN_ANGLE, l.ch)
	case '>':
		t = newToken(token.CLOSE_ANGLE, l.ch)
	case '/':
		t = newToken(token.XML_TERMINATOR, l.ch)
	}

	return t
}

func (l *Lexer) nextJsonToken() token.Token {
	var t token.Token

	switch l.ch {
	case '{':
		t = newToken(token.OPEN_CURLY, l.ch)
	case '}':
		t = newToken(token.CLOSE_CURLY, l.ch)
	case '[':
		t = newToken(token.OPEN_SQUARE, l.ch)
	case ']':
		t = newToken(token.CLOSE_SQUARE, l.ch)
	}
	return t
}

/*
readIdentifier - determines is the current read is a TAG, KEY, or VALUE
and reads the value into the appropriate TokenType
*/
func (l *Lexer) readIdentifier() token.Token {
	var tok token.Token
	lr := string(l.lastRead)

	switch lr {
	case token.OPEN_ANGLE:
		tok.Type = token.TAG
	case token.QUOTE:
		fallthrough
	case token.CLOSE_ANGLE:
		tok.Type = token.VALUE
	case token.XML_TERMINATOR:
		next := string(l.input[l.nextPosition])
		if next != token.CLOSE_ANGLE {
			tok.Type = token.TAG
		}
	default:
		tok.Type = token.KEY
	}

	pos := l.currentPosition
	for isAlphaNumeric(l.ch) {
		l.readChar()
	}

	tok.Literal = l.input[pos:l.currentPosition]
	return tok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isAlphaNumeric(ch byte) bool {
	isAN := (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '-' || ch == '.'
	return isAN
}
