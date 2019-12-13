package parser

import (
	"crawler/engine"
	"regexp"
)

const urlCityList string = "<a href=\"(http://www.zhenai.com/zhenghun/[0-9a-z]+)\"[^>]*>([^<]+)</a>"

func ParseCityList(contents []byte) engine.ParseResult {
	//* 0 - n ; + 1 - n
	re := regexp.MustCompile(urlCityList)
	//-1 表示读取所有的数据 返回 [][][]byte
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, urlLocation := range all {
		//存下地名
		result.Items = append(result.Items, string(urlLocation[2]))
		//存下URL
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(urlLocation[1]),
			ParserFunc: engine.NilParse, //下一个页面
		})
	}
	return result
}
