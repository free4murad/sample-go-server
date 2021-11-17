package storage

import (
	"errors"
	"go-run/model"
)

type Service interface {
	PutUser(user model.User) error
	GetUser()
}

var (
	ErrUserExistsAlready = errors.New("user already exists")
)
