package handlers

import (
	"net/http"
	db "todolist/internal/data"
)

func DeleteTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		db.DeleteTodo(id)
		w.WriteHeader(http.StatusNoContent)
	}
}
