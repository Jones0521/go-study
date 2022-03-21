package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	addr := "0.0.0.0:8888"
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 请求行
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("method", r.Method)
		fmt.Println("URL", r.URL)
		fmt.Println("protocol", r.Proto)
		fmt.Println("RemoteAddr", r.RemoteAddr)
		fmt.Println(r.Host)

		header := r.Header
		fmt.Println(header.Get("User-Agent"))
		fmt.Println(header.Get("Token"))
		fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
		// 请求体
		fmt.Println("body:")
		io.Copy(os.Stdout, r.Body)
	})
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}
