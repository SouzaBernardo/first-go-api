package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	products := []Product{
		{Name: "Shoe", Description: "Blue", Price: 25, Quantity: 5},
		{Name: "Tshirt", Description: "Purple", Price: 30, Quantity: 2},
		{Name: "Socke", Description: "White", Price: 15, Quantity: 2},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
