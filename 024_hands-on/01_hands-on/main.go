package main

import (
	"log"
	"net/http"
	"io"
	"html/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/t-rex/", trex)
	http.HandleFunc("/t-rex.png", trexImg)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "foo ran")
}

func trex(w http.ResponseWriter, r *http.Request)  {
	tpl := template.Must(template.ParseFiles("t-rex.gohtml"))
	tpl.Execute(w, nil)
}

func trexImg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "t-rex.png")
}