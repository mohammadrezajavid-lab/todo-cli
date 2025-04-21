package deliveryparam

import "encoding/json"

type UserRequest struct {
	email    string
	password string
}

func NewUserRequest(email string, password string) *UserRequest {
	return &UserRequest{
		email:    email,
		password: password,
	}
}

func (ur *UserRequest) GetEmail() string {
	return ur.email
}
func (ur *UserRequest) GetPassword() string {
	return ur.password
}
func (ur *UserRequest) SetEmail(email string) {
	ur.email = email
}
func (ur *UserRequest) SetPassword(password string) {
	ur.password = password
}
func (ur *UserRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"email":    ur.GetEmail(),
		"password": ur.GetPassword(),
	})
}
func (ur *UserRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	ur.SetEmail(aux.Email)
	ur.SetPassword(aux.Password)

	return nil
}

type UserResponse struct {
	authenticateUserId uint
}

func NewUserResponse(authenticateUserId uint) *UserResponse {
	return &UserResponse{authenticateUserId: authenticateUserId}
}

func (r *UserResponse) GetAuthenticateUserId() uint {
	return r.authenticateUserId
}

func (r *UserResponse) SetAuthenticateUserId(authenticateUserId uint) {
	r.authenticateUserId = authenticateUserId
}

func (r *UserResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"authenticateUserId": r.GetAuthenticateUserId(),
	})
}
func (r *UserResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		AuthenticateUserId uint `json:"authenticateUserId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	r.SetAuthenticateUserId(aux.AuthenticateUserId)

	return nil
}
