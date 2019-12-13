package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

//抓取数据并转码
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status code : %d", resp.StatusCode)
	}
	//避免数据丢失
	newReader := bufio.NewReader(resp.Body)
	e := determineEncoding(newReader)
	//判断编码并转化为UTF8
	utf8Body := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Body)
}

//正则表达式匹配数据

// 转码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
