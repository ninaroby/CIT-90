package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

type hotham int

func (hh hotham) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am a hot hamster! HAHAHAHAHAAAA")
}

func main() {
	var h hotdog
	var c hotcat
	var hh hotham

	http.Handle("/dog", h)
	http.Handle("/cat", c)
	http.Handle("/", hh)

	http.ListenAndServe(":8080", nil)
}

// change code to use DefaultServeMux
// add a route for hamsters
