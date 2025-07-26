package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func GetHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./internal/templates/home.html")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while parsing html: %v\n", err)
			return
		}

		tmpl.Execute(w, nil)
	}
}
