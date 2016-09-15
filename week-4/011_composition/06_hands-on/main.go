package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "West Hollywood",
			Address: "6250 Hollywood Blvd",
			City:    "Los Angeles",
			Zip:     "90028-5325",
			Region:  "California",
		},
		hotel{
			Name:    "Hyatt Regency",
			Address: "L St.",
			City:    "Sacramento",
			Zip:     "95814",
			Region:  "California",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
