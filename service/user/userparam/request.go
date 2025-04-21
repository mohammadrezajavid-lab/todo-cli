package userparam

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
