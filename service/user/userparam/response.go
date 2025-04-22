package userparam

type ResponseUser struct {
	userId uint
	email  string
}

func NewResponseUser(userId uint, email string) *ResponseUser {
	return &ResponseUser{userId: userId, email: email}
}

func (r *ResponseUser) GetUserId() uint {
	return r.userId
}
func (r *ResponseUser) SetUserId(userId uint) {
	r.userId = userId
}
func (r *ResponseUser) GetEmail() string {
	return r.email
}
func (r *ResponseUser) SetEmail(email string) {
	r.email = email
}
