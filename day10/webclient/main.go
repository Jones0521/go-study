package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	response, err := http.Get("http://127.0.0.1:8888?x=1&y=2")
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, response.Body)

	response, err = http.Head("http://127.0.0.1:8888?x=1&y=2")
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, response.Body)
	values := url.Values{}
	values.Add("x", "1")
	values.Add("x", "2")
	values.Set("y", "1")
	values.Set("y", "2")
	response, err = http.PostForm("http://127.0.0.1:8888?x=1&y=2", values)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, response.Body)
	// application/json
	//f, _ := os.Open("a.json")
	reader := strings.NewReader(`{"a":"x"}`)
	response, err = http.Post("http://127.0.0.1:8888?x=1&y=2", "application/json", reader)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, response.Body)
}
