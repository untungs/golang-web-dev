package main

import (
	"github.com/untungs/golang-web-dev/042_mongodb/10_hands-on/controllers"
	"net/http"
)

func main() {
	c := controllers.NewController()
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.Signup)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
