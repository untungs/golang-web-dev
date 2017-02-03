package main

import (
	"html/template"
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.HandleFunc("/", index)
	http.Handle("/pics/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
