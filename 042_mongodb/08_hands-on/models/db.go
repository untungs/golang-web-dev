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
	dbf := dbFile()
	defer dbf.Close()

	fi, err := dbf.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	bs := make([]byte, fi.Size())

	if _, err := dbf.Read(bs); err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(bs))

	if err := json.Unmarshal(bs, &udb); err != nil {
		log.Fatalln(err)
	}
	return udb
}

// NewUser saves a new user to database and return it with the assigned Id
func (db Udb) NewUser(u User) User {
	u.Id = uuid.NewV4().String()
	db[u.Id] = u

	bs, err := json.Marshal(db)
	if err != nil {
		log.Fatalln(err)
	}

	dbf := dbFile()

	if n, err := dbf.Write(bs); err != nil {
		log.Fatalln(err, "- byte written:", n)
	}
	dbf.Close()

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
	return nil
}
