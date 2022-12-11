package main

import (
	"fmt"
	"net/url"
)

func main() {
	// URLを解析
	u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	fmt.Println(u.Scheme) // http
	fmt.Println(u.Host)   // example.com
	fmt.Println(u.Path)   // /search
	fmt.Println(u.RawQuery) // a=1&b=2
	fmt.Println(u.Fragment) // top

	fmt.Println(u.IsAbs()) // ture
	fmt.Println(u.Query()) // map[a:[1] b:[2]]

	// URLを生成
	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "google.com"
	// クエリーを設定
	q := url.Query()
	q.Set("q", "Golang")
	url.RawQuery = q.Encode()

	fmt.Println(url) // https:://google.com?q=Golang

}
