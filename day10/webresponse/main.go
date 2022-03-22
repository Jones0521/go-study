package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	addr := ":8888"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// html
		//fmt.Fprint(w, html)
		f, _ := os.Open("./template/index.html")
		io.Copy(w, f)
	})
	http.ListenAndServe(addr, nil)
}
