package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "This is the index page")
	})

	http.HandleFunc("/dog/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "woof woof 🐶")
	})

	http.HandleFunc("/me/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "I'm untungs")
	})

	http.ListenAndServe(":8080", nil)
}
