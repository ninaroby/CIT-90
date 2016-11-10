package main

import (
	"io"
	"net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}

func f(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Foo Dog")
}

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/foo", f)
	http.HandleFunc("/cool", cool)

	http.ListenAndServe(":8080", nil)
}

func cool(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, req, "cool-cat.jpg")
}
