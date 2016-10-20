package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/things.jpeg", pic)
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/things.jpeg">`)
}

func pic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "things.jpeg")
}
