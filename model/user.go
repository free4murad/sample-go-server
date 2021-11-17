package model

type User struct {
	UserName string
	Email    string
	Password string
}

func (u *User) Validate() error {
	// Validate the user struct
	return nil
}
