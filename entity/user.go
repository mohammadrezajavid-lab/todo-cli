package entity

import (
	"encoding/json"
	"fmt"
)

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

func (u *User) SetId(id uint) {
	u.id = id
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
	return fmt.Sprintf("ID: %d, Email: %s, Name: %s\n",
		u.GetId(),
		u.GetEmail(),
		u.GetName(),
	)
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"id":       u.GetId(),
		"name":     u.GetName(),
		"email":    u.GetEmail(),
		"password": u.GetPassword(),
	})
}
