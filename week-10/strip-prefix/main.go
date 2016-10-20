package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cool)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func cool(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img width="800px" src="/resources/cool-cat.jpg">`)
}
