package routes

import (
	"net/http"
	"SouzaBernardo/first-go-api/controller"
)

func CreateRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/create", controller.CreateProducts)
}
