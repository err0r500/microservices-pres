package uc

import (
	"errors"
	"time"
)

type User struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type ReadWriter interface {
	NewUser(User) (id int, err error)
	SaveUser(User) error
	GetUserByID(id string) (User, error)
}

type Interactor struct {
	ReadWriter
}

func (i Interactor) CreateUser(name string, dateOfBirth time.Time) error {
	user := User{
		Name:        name,
		DateOfBirth: dateOfBirth,
	}

	if err := user.Check(); err != nil {
		return err
	}

	id, err := i.ReadWriter.NewUser(user)
	if err != nil {
		return err
	}
	if id == 0 {
		return errors.New("user creation failed without raising an error")
	}

	return nil
}

func (user User) Check() error {
	return nil
}
