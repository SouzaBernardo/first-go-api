package main

import (
	"SouzaBernardo/first-go-api/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	routes.CreateRoutes()
	http.ListenAndServe(":8080", nil)
}
