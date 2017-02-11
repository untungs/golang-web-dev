package models

import (
	"errors"
	"github.com/satori/go.uuid"
)

type Udb map[string]User

var ErrUserNotFound = errors.New("user is not found")

func NewUDB() Udb {
	return Udb{}
}

// NewUser saves a new user to database and return it with the assigned Id
func (db Udb) NewUser(u User) User {
	u.Id = uuid.NewV4().String()
	db[u.Id] = u
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
