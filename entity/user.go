package entity

import "fmt"

type User struct {
	id       uint
	name     string
	email    string
	password []uint8
}

func NewUser(id uint, name string, email string, password []uint8) *User {
	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}

// GetId Getter method
func (u *User) GetId() uint {
	return u.id
}

// GetPassword Getter method
func (u *User) GetPassword() []uint8 {
	return u.password
}

// GetEmail Getter method
func (u *User) GetEmail() string {
	return u.email
}

// GetName Getter method
func (u *User) GetName() string {
	return u.name
}

func (u *User) String() string {
	return fmt.Sprintf("ID: %d, Email: %s, Name: %s", u.GetId(), u.GetEmail(), u.GetName())
}
