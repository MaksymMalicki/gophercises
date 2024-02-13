package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Adventure struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Adventure

func DisplayStory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		story := readStory()
		vars := mux.Vars(r)
		adventureName, ok := vars["key"]
		if !ok {
			http.Error(w, "Key not found", http.StatusBadRequest)
			return
		}
		adventure, ok := story[adventureName]
		if !ok {
			http.Error(w, "Story not found", http.StatusNotFound)
			return
		}
		var tmplFile = "./templates/adventure.html" // Correct file extension
		tmpl, err := template.ParseFiles(tmplFile)  // Parse template file
		if err != nil {
			http.Error(w, "Internal server error, error parsing", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, adventure)
		if err != nil {
			http.Error(w, "Internal server error, error executing", http.StatusInternalServerError)
			return
		}
	}
}

func readStory() Story {
	file, err := os.Open("./gopher.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	story := Story{}
	json.Unmarshal(byteValue, &story)
	return story
}
