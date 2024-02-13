package main

import (
	"net/http"

	cyoa "github.com/Maksymmalicki/gophercises/cyoa/code"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", cyoa.DisplayStory().ServeHTTP)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusFound)
	})
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
