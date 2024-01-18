package controller

import (
	"SouzaBernardo/first-go-api/database/domain"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	products := domain.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)

}

func CreateProducts(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "CreateProducts", nil)

}

