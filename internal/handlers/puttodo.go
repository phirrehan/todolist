package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	db "todolist/internal/data"
)

func PutTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		description := r.FormValue("description")
		status := r.FormValue("status")
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing boolean: %v\n", err)
			return
		}
		db.EditTodo(id, description, statusBool)
		w.WriteHeader(http.StatusOK)
	}
}
