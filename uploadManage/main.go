package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// 根据请求方法的不同, 来做不同的处理
	// 如果是POST请求, 就提取用户提交的form表单的数据, 去数据库创建一行数据
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		// 往数据库中写
		err = createUser(username, password, email)
		if err != nil {
			w.WriteHeader(500)
		}
	} else {
		// 如果是GET请求, 就返回一个HTML页面, 供用户输入注册信息
		t, err := template.ParseFiles("./register.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		t.Execute(w, nil)
	}
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/register", registerHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动http server 失败")
	}
}
