package main

import (
	"hangman-api/router"
	"net/http"
)

func main() {
	app := router.NewRouter()

	defer http.ListenAndServe(":9000", app)
	println("-->")
	println("server started on localhost:9000")
	println("Play a game by using the CLI")
}
