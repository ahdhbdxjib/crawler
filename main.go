package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

}
