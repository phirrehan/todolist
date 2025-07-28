package main

import (
	"log"
	"net/http"
	"todolist/internal/handlers"
)

func main() {
	http.Handle("GET /static/", handlers.StaticHandler())

	http.HandleFunc("GET /", handlers.GetHomePage())
	http.HandleFunc("GET /todo", handlers.GetTodos())
	http.HandleFunc("POST /todo", handlers.PostTodo())
	http.HandleFunc("PUT /todo", handlers.PutTodo())
	http.HandleFunc("DELETE /todo", handlers.DeleteTodo())

	log.Fatal(http.ListenAndServe(":3000", nil))
}
