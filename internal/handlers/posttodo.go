package handlers

import (
	"net/http"
	db "todolist/internal/data"
)

func PostTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		description := r.FormValue("description")
		db.AddTodo(description)
		w.WriteHeader(http.StatusCreated)
	}
}
