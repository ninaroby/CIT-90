package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", cool)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func cool(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprintln(w, "Go check your terminal!")
}
