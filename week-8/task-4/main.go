package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("ninakey", "This is from Nina")
	fmt.Fprintln(res, "Any code you want in this func")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
