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

type hamster int

func (hm hamster) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am a hamster!")
}

type wormhole int

func (w wormhole) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "The wormhole of the web!")
}

func main() {
	var h hotdog
	var c hotcat
	var hm hamster
	var w wormhole

	mux := http.NewServeMux()
	mux.Handle("/dog", h)
	mux.Handle("/cat/", c)
	mux.Handle("/hamster", hm)
	mux.Handle("/", hm)
	mux.Handle("/w", w)

	http.ListenAndServe(":8080", mux)
}

// create a route for hamster
