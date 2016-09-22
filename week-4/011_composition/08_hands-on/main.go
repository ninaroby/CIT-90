package main

import "html/template"

type menuItem struct {
	Name, Description, Meal string
	Price                   float64
}

type menuItems []menuItem

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	s := menuItems{

		menuItem{
			Name:        "Fish Fingers",
			Description: "Golden and delicious",
			Meal:        "Dinner",
			Price:       3.90,
		},

		menuItem{
			Name:        "Spag Bolognese",
			Description: "Tasty Tasty",
			Meal:        "Dinner",
			Price:       4.95,
		},

		menuItem{
			Name:        "Yoghurt & Fruit",
			Description: "Healthy and yummy",
			Meal:        "Breakfast",
			Price:       5.00,
		},
	}

	err := tpl.Execute(os.Stdoubt, s)
	if err != nil {
		log.fatalln(err)
	}
}
