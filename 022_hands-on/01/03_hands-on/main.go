package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse("{{.}}"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, r.URL.Path)
	handleError(w, err)
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "woof woof 🐶")
	handleError(w, err)
}

func me(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, template.HTML("I'm untungs"))
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
