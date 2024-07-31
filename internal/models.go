package internal

type ParsedObject interface {
	Parse(obj string) error
}

type XmlObject struct {
	Tag   string
	Group string
	Value string
	Text  string
	Id    string
}

func (x *XmlObject) Parse(obj string) error {
	return nil
}

func (x *XmlObject) ToJson(xObject string) (*JsonObject, error) {
	return nil, nil
}

type JsonObject struct {
	Key      string
	Value    interface{}
	dataType string
}

func (x *JsonObject) Parse(obj string) error {
	return nil
}
