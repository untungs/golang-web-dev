package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/untungs/golang-web-dev/042_mongodb/08_hands-on/models"
	"net/http"
)

type UserController struct {
	db models.Udb
}

func NewUserController(udb models.Udb) *UserController {
	return &UserController{udb}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Fetch user
	u, err := uc.db.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintln(w, err.Error())
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// store the user in Udb
	u = uc.db.NewUser(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Delete user
	if err := uc.db.DeleteUser(id); err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user ", id, "\n")
}
