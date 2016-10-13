package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog dog")
}

func h(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "horses are happy!")
}

func i(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am the index!")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/horse", h)
	http.HandleFunc("/", i)

	http.ListenAndServe(":8080", nil)
}
