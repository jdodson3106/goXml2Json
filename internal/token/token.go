package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literals
	INT   = "INT"   // 1343456
	FLOAT = "FLOAT" // 3.1415926535
	BOOL  = "BOOL"

	// Identifiers
	TAG   = "TAG" // xml has tag names to parse (these will convert into json object names)
	KEY   = "KEY" // xml and json both have key/value pairs
	VALUE = "VALUE"

	// xml delimiter tokens
	XML_TERMINATOR = "/"
	OPEN_ANGLE     = "<"
	CLOSE_ANGLE    = ">"

	// json delimiter tokens
	OPEN_CURLY   = "{"
	CLOSE_CURLY  = "}"
	OPEN_SQUARE  = "["
	CLOSE_SQUARE = "]"

	// common tokens
	COMMA        = ","
	COLON        = ":"
	EQUAL        = "="
	QUOTE        = "\""
	SINGLE_QUOTE = "'"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
