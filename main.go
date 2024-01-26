package main

import (
	"fmt"
	"github.com/SouzaBernardo/first-go-api/controllers/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	startMessage()
	routes.CreateRoutes()
	http.ListenAndServe(":8080", nil)
}
func startMessage() {
	fmt.Println("\n  /$$$$$$  /$$$$$$$  /$$$$$$       /$$$$$$$  /$$   /$$ /$$   /$$ /$$   /$$ /$$$$$$ /$$   /$$  /$$$$$$")
	fmt.Println("/$$__  $$| $$__  $$|_  $$_/      | $$__  $$| $$  | $$| $$$ | $$| $$$ | $$|_  $$_/| $$$ | $$ /$$__  $$")
	fmt.Println("| $$  \\ $$| $$  \\ $$  | $$        | $$  \\ $$| $$  | $$| $$$$| $$| $$$$| $$  | $$  | $$$$| $$| $$  \\__/")
	fmt.Println("| $$$$$$$$| $$$$$$$/  | $$        | $$$$$$$/| $$  | $$| $$ $$ $$| $$ $$ $$  | $$  | $$ $$ $$| $$ /$$$$")
	fmt.Println("| $$__  $$| $$____/   | $$        | $$__  $$| $$  | $$| $$  $$$$| $$  $$$$  | $$  | $$  $$$$| $$|_  $$")
	fmt.Println("| $$  | $$| $$        | $$        | $$  \\ $$| $$  | $$| $$\\  $$$| $$\\  $$$  | $$  | $$\\  $$$| $$  \\ $$")
	fmt.Println("| $$  | $$| $$       /$$$$$$      | $$  | $$|  $$$$$$/| $$ \\  $$| $$ \\  $$ /$$$$$$| $$ \\  $$|  $$$$$$/")
	fmt.Println("|__/  |__/|__/      |______/      |__/  |__/ \\______/ |__/  \\__/|__/  \\__/|______/|__/  \\__/ \\______/ ")
	fmt.Println("\n Access localhost:8080")
}
