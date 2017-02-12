package models

import (
	"errors"
	"github.com/satori/go.uuid"
	"os"
	"log"
	"encoding/json"
)

type Udb map[string]User

var ErrUserNotFound = errors.New("user is not found")

func dbFile() *os.File {
	f, err := os.OpenFile("udb.db", os.O_RDWR, os.ModePerm)
	if err != nil {
		if f, err = os.Create("udb.db"); err != nil {
			log.Fatalln(err)
		}
	}
	return f
}

func NewUDB() Udb {
	udb := Udb{}
	f := dbFile()
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&udb); err != nil {
		log.Println(err)
	}
	return udb
}

// NewUser saves a new user to database and return it with the assigned Id
func (db Udb) NewUser(u User) User {
	u.Id = uuid.NewV4().String()
	db[u.Id] = u

	f := dbFile()
	defer f.Close()

	if err := json.NewEncoder(f).Encode(db); err != nil {
		log.Println(err)
	}

	return u
}

func (db Udb) GetUser(id string) (User, error) {
	u, ok := db[id]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return u, nil
}

func (db Udb) DeleteUser(id string) error {
	if _, err := db.GetUser(id); err != nil {
		return err
	}
	delete(db, id)

	f := dbFile()
	defer f.Close()

	f.Truncate(0)
	if err := json.NewEncoder(f).Encode(db); err != nil {
		log.Println(err)
	}
	return nil
}
