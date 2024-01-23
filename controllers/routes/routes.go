package routes

import (
	"github.com/SouzaBernardo/first-go-api/controllers"
	"net/http"
)

func CreateRoutes() {
	http.HandleFunc("/", controllers.IndexView)
	http.HandleFunc("/create", controllers.CreateProductView)
	http.HandleFunc("/edit", controllers.EditProductView)
	http.HandleFunc("/insert", controllers.Create)
	http.HandleFunc("/update", controllers.Edit)
	http.HandleFunc("/delete", controllers.Delete)
}
