package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func nina(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}
	res.Header().Set("content-type", "text/html; charset=UTF-8")

	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/happy", name)
	http.HandleFunc("/me", nina)
	http.ListenAndServe(":8080", nil)
}

func name(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "What's up")
}
