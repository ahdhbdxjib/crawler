package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

func Run(seds ...Request) {
	var requests []Request
	for _, r := range seds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		//从队列中获取数据，再将队列切片
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching the URL : %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Error Url : %s ,ERROR : %v", r.Url, err)
			//发现错误不中断
			continue
		}
		parseResult := r.ParserFunc(body)
		//将所有的parseResult.Requests 添加到requests中
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			fmt.Printf("Item : %v \n", item)
		}

	}
}
