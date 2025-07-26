package main

import (
	"log"
	"net/http"
	"todolist/internal/data"
	"todolist/internal/handlers"
)

func main() {
	http.Handle("/static/", handlers.StaticHandler())
	http.HandleFunc("/", handlers.GetHomePage())

	log.Fatal(http.ListenAndServe(":3000", nil))
}
