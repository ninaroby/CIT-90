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
	// http.HandleFunc("/cool", cool)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./public"))))

	http.ListenAndServe(":8080", nil)
}

// func cool(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	io.WriteString(w, `<img src="assets/cool-cat.jpg"`)
// }
