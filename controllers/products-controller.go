package controllers

import (
	"SouzaBernardo/first-go-api/controllers/service"
	"SouzaBernardo/first-go-api/model/entity"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("view/*.html"))

const rootPath = "/"
const indexView = "IndexView"
const createProductView = "CreateProductView"
const editProductView = "EditProductView"

func IndexView(w http.ResponseWriter, r *http.Request) {
	products := entity.GetAll()
	_ = temp.ExecuteTemplate(w, indexView, products)
}

func CreateProductView(w http.ResponseWriter, r *http.Request) {
	_ = temp.ExecuteTemplate(w, createProductView, nil)
}

func EditProductView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := entity.Get(id)
	_ = temp.ExecuteTemplate(w, editProductView, product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	service.Create(r)
	http.Redirect(w, r, rootPath, 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	service.Delete(r)
	http.Redirect(w, r, rootPath, 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	service.Edit(r)
	http.Redirect(w, r, rootPath, 301)
}
