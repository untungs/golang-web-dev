package models

type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

var DbUsers = map[string]User{} // user ID, user
