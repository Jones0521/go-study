package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		fmt.Println(strings.Repeat("-", 30))
		// 1. 提交数据方式
		/*
		   在URL中传递数据
		   url?argname1=argvalue1&argname2=argvalue2
		   例如curl -XGET http://127.0.0.1:8888/?x=1&y=2
		*/
		fmt.Println(r.URL)
		r.ParseForm()                // 解析参数
		fmt.Println(r.Form)          // 接收的参数类型都是string
		fmt.Println(r.Form.Get("x")) // 只返回第一个参数
		fmt.Println(r.Form["x"])     // 返回参数对应的字符串切片

		// 2. 通过body提交数据
		/*
		   curl -d “”
		   body 中数据格式 常用3种
		   application/x-www-form-urlencoded  例如a=b&c=d
		   multipart/form-data
		   application/json  在go中需要自己解码
		   其他类型
		   application/xml
		*/
		/*
			application/x-www-form-urlencoded
			Form 可以获取URL中的参数也可以获取Body中的参数
			PostForm 只能获取Body中的参数
		*/
		fmt.Println(r.PostForm)

		fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
	})

	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		// 提交数据的编码格式是什么
		// Request Header: Content-Type
		fmt.Println(r.Header)
		// Header获取Content-Type
		// json => jsonParser
		// x-www-form-urlencoded => ParseForm
		// xml => xmlParser
		// curl 指定 Content-Type 方式 curl -H “Content-Type: application/json”
		// 对于自定义类型需要获取body原始数据使用特定的解码器
		// 例如 curl -XPOST “127.0.0.1:8888/data/” -d "{\"x\":1}"
		//io.Copy(os.Stdout, r.Body)
		ctx, _ := ioutil.ReadAll(r.Body)
		var j map[string]interface{}
		json.Unmarshal(ctx, &j)
		fmt.Printf("%#V \n", j)

		fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
	})

	http.HandleFunc("/file/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024 * 1024) // 接受文件过程中最大使用的内存
		// url?
		// body k=v
		// body 文件内容
		// curl -XPOST -F "x=@go.mod" "127.0.0.1:8888/file/?a=1" -F "y=1" -F "z=2"
		fmt.Println(r.MultipartForm.File)
		fmt.Println(r.MultipartForm.Value)
		fmt.Println(r.Form)
		fmt.Println(r.PostForm)
		if fileHeaders, ok := r.MultipartForm.File["x"]; ok {
			for _, fileHeader := range fileHeaders {
				fmt.Println(fileHeader.Filename, fileHeader.Size)
				nfile, _ := os.Create("./file/" + fileHeader.Filename)
				file, _ := fileHeader.Open()
				io.Copy(nfile, file)
				file.Close()
				nfile.Close()
			}
		}

	})
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}
