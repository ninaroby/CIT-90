package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("letter.gohtml")
	if err != nil {
		fmt.Println("There was an error parsing file", err)
	}

	friends := []string{"Alex", "Conner", "Ken", "Gentry", "Jeremy"}

	err = tpl.Execute(os.Stdout, friends)
	if err != nil {
		fmt.Println("There was an error executing template", err)
	}
}
