package main

import (
	"net/http"
	"fmt"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("counter")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name: "counter",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)
	fmt.Fprintf(w, "number of visit: %d", count)
}