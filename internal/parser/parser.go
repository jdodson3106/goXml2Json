package parser

import (
	"fmt"

	models "github.com/jdodson3106/goXml2Json/internal"
)

const (
	X_TERMINATOR = '/'
	OPEN_ANGLE   = '<'
	CLOSE_ANGLE  = '>'
	OPEN_CURLY   = '{'
	CLOSE_CURLY  = '}'
	OPEN_SQUARE  = '['
	CLOSE_SQUARE = ']'
	COMMA        = ','
	COLON        = ':'
	EQUAL        = '='
	QUOTE        = '"'

	JSON = "json"
	XML  = "xml"
)

func parseObject(objType, obj string) ([]models.ParsedObject, error) {
	switch objType {
	case JSON:
		return parseJson(obj)
	case XML:
		return parseXml(obj)
	default:
		return nil, fmt.Errorf("unknown file type .%s", objType)
	}
}

func parseXml(xml string) ([]models.ParsedObject, error) {
	if xml == "" {
		return nil, fmt.Errorf("not object provided")
	}
	return []models.ParsedObject{&models.XmlObject{}}, nil
}

func parseJson(json string) ([]models.ParsedObject, error) {
	if json == "" {
		return nil, fmt.Errorf("not object provided")
	}
	return []models.ParsedObject{&models.JsonObject{}}, nil
}
