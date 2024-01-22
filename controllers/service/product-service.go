package service

import (
	"SouzaBernardo/first-go-api/model/entity"
	"log"
	"net/http"
	"strconv"
)

func Create(r *http.Request) {
	if r.Method != "POST" {
		return
	}
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Cannot convert price")
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Println("Cannot convert quantity")
		return
	}
	entity.Create(name, description, price, quantity)
}

func Edit(r *http.Request) {
	if r.Method != "PUT" {
		return
	}
	productId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Cannot convert id")
		return
	}
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Cannot convert price")
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Println("Cannot convert quantity")
		return
	}
	entity.Edit(productId, name, description, price, quantity)
}

func Delete(r *http.Request) {
	productId := r.URL.Query().Get("id")
	entity.Delete(productId)
}
