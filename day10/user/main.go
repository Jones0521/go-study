package main

import (
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type User struct {
	Id   int64
	Name string
	Sex  bool
	Addr string
}

func main() {
	addr := ":8888"
	users := []*User{
		{1, "aaaa", true, "xxxxx"},
		{2, "bbbb", true, "xxxx2x"},
		{3, "cccc", true, "xxxx3x"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("template/user.html"))
		tpl.ExecuteTemplate(w, "user.html", users)

	})
	http.HandleFunc("/create/", func(w http.ResponseWriter, r *http.Request) {
		// a => 加载页面 GET
		// form => 添加数据, 跳转到用户列表页面 POST // 重定向
		if r.Method == "GET" {
			tpl := template.Must(template.ParseFiles("template/create.html"))
			tpl.ExecuteTemplate(w, "create.html", nil)
		} else {
			users = append(users, &User{
				time.Now().Unix(),
				r.FormValue("name"),
				r.FormValue("sex") == "1",
				r.FormValue("addr"),
			})
			http.Redirect(w, r, "/", 302)
		}
	})
	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		if id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
			nUsers := make([]*User, 0, len(users))
			for _, user := range users {
				if user.Id == id {
					continue
				}
				nUsers = append(nUsers, user)
			}
			users = nUsers
		}
		http.Redirect(w, r, "/", 302)
	})
	http.ListenAndServe(addr, nil)
}
