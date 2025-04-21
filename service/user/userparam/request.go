package userparam

import "encoding/json"

type RequestUser struct {
	email    string
	password string
}

func NewRequestUser(email, password string) *RequestUser {
	return &RequestUser{
		email:    email,
		password: password,
	}
}

func (r *RequestUser) GetEmail() string {
	return r.email
}
func (r *RequestUser) SetEmail(email string) {
	r.email = email
}
func (r *RequestUser) GetPassword() string {
	return r.password
}
func (r *RequestUser) SetPassword(password string) {
	r.password = password
}

type RequestRegisterUser struct {
	name     string
	email    string
	password []uint8
}

func NewRequestRegisterUser(name, email string, password []uint8) *RequestRegisterUser {
	return &RequestRegisterUser{
		name:     name,
		email:    email,
		password: password,
	}
}

func (u RequestRegisterUser) GetName() string {
	return u.name
}
func (u RequestRegisterUser) GetEmail() string {
	return u.email
}
func (u RequestRegisterUser) GetPassword() []uint8 {
	return u.password
}

func (u RequestRegisterUser) SetName(name string) {
	u.name = name
}
func (u RequestRegisterUser) SetEmail(email string) {
	u.email = email
}
func (u RequestRegisterUser) SetPassword(password []uint8) {
	u.password = password
}

func (u RequestRegisterUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"name":     u.GetName(),
		"email":    u.GetEmail(),
		"password": u.GetPassword(),
	})
}

func (u RequestRegisterUser) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Password []uint8 `json:"password"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	u.SetName(aux.Name)
	u.SetEmail(aux.Email)
	u.SetPassword(aux.Password)

	return nil
}

type ResponseRegisterUser struct {
	email string
}

func NewResponseRegisterUser(email string) *ResponseRegisterUser {
	return &ResponseRegisterUser{
		email: email,
	}
}

func (ru *ResponseRegisterUser) GetEmail() string {
	return ru.email
}
func (ru *ResponseRegisterUser) SetEmail(email string) {
	ru.email = email
}

func (ru *ResponseRegisterUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"email": ru.GetEmail(),
	})
}

func (ru *ResponseRegisterUser) UnmarshalJSON(data []byte) error {
	var aux struct {
		Email string `json:"email"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	ru.SetEmail(aux.Email)

	return nil
}
