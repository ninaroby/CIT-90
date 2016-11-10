package main

import (
	"log"
	"net/http"
	"text/template"
)

type person struct {
	Fname     string
	Lname     string
	Subscribe bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)

	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	p1 := person{
		f,
		l,
		s,
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", p1)

	if err != nil {
		log.Fatalln(err)
	}
}
