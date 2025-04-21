package userparam

type ResponseUser struct {
	userId uint
}

func NewResponseUser(userId uint) *ResponseUser {
	return &ResponseUser{userId: userId}
}

func (r *ResponseUser) GetUserId() uint {
	return r.userId
}
func (r *ResponseUser) SetUserId(userId uint) {
	r.userId = userId
}
