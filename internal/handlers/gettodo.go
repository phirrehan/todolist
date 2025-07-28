package handlers

import (
	"encoding/json"
	"net/http"
	db "todolist/internal/data"
)

func GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := db.GetTodos()

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}
}
