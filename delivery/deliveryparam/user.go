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

type RegisterUserRequest struct {
	name     string
	email    string
	password string
}

func NewRegisterUserRequest(name, email, password string) *RegisterUserRequest {
	return &RegisterUserRequest{
		name:     name,
		email:    email,
		password: password,
	}
}

func (rur *RegisterUserRequest) GetName() string {
	return rur.name
}
func (rur *RegisterUserRequest) GetEmail() string {
	return rur.email
}
func (rur *RegisterUserRequest) GetPassword() string {
	return rur.password
}

func (rur *RegisterUserRequest) SetName(name string) {
	rur.name = name
}

func (rur *RegisterUserRequest) SetEmail(email string) {
	rur.email = email
}
func (rur *RegisterUserRequest) SetPassword(password string) {
	rur.password = password
}

func (rur *RegisterUserRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"name":     rur.GetName(),
		"email":    rur.GetEmail(),
		"password": rur.GetPassword(),
	})
}
func (rur *RegisterUserRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	rur.SetName(aux.Name)
	rur.SetEmail(aux.Email)
	rur.SetPassword(aux.Password)

	return nil
}

type RegisterUserResponse struct {
	email string
	error error
}

func NewRegisterUserResponse(email string, err error) *RegisterUserResponse {
	return &RegisterUserResponse{email: email, error: err}
}

func (rur *RegisterUserResponse) GetEmail() string {
	return rur.email
}
func (rur *RegisterUserResponse) GetError() error {
	return rur.error
}
func (rur *RegisterUserResponse) SetEmail(email string) {
	rur.email = email
}
func (rur *RegisterUserResponse) SetError(error error) {
	rur.error = error
}
func (rur *RegisterUserResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"email": rur.GetEmail(),
		"error": rur.GetError(),
	})
}
func (rur *RegisterUserResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		Email string `json:"email"`
		Error error  `json:"error"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	rur.SetEmail(aux.Email)
	rur.SetError(aux.Error)

	return nil
}
