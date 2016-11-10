package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
}

func index(res http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)

	if err != nil {
		log.Fatalln("Template didn't show!", err)
	}
}
