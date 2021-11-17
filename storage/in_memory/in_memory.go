package in_memory

import (
	"go-run/model"
	"go-run/storage"
	"log"
)

type db struct {
	dbMap map[string]*model.User // This is not gorutine safe
}

func NewUserDB() storage.Service {
	d := db{dbMap: map[string]*model.User{}}
	s := storage.Service(&d)
	return s
}

func (d *db) PutUser(user model.User) error {
	if _, exists := d.dbMap[user.Email]; exists {
		log.Println("Error: User with this email already exists", user.Email)
		return storage.ErrUserExistsAlready
	}

	d.dbMap[user.Email] = &user
	log.Println("INFO: User added to database:", user.Email)
	return nil
}

func (d *db) GetUser() {
	// implement this
}
