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

	mux := http.NewServeMux()
	mux.HandleFunc("/", h)
	mux.HandleFunc("/foo", f)
	http.ListenAndServe(":8080", mux)
}
