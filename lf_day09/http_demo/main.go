package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http 客户端 发请求示例

// getDemo get 请求示例
func getDemo() {
	resp, err := http.Get("https://www.luffycity.com/study/degree")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close() // 关闭
	// resp *http.Response
	// 使用ioutil库读取相应数据
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("内容为%s", b)
}

func getparms() {
	apiUrl := "http://127.0.0.1:9090/get"
	// URL param
	data := url.Values{}
	data.Set("name", "123")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func main() {
	//getDemo()
	getparms()
}
