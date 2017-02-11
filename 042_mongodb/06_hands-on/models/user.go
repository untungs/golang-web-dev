package models

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}

// Id was of type string before
