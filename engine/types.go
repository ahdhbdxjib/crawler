package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//返回体的结构
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
