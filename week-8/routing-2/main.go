package main

import (
	"io"
	"net/http"
)

type dogHandler int

func (h dogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type catHandler int

func (h catHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}

func main() {
	var dog dogHandler
	var cat catHandler

	mux := http.NewServeMux()
	mux.Handle("/", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":8080", mux)
}

/*
if route ends in slash
/dog/
it includes anything beneath
if route ends in no-slash
/dog
it only includes that
*/
