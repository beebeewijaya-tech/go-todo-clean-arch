package models

type UserCreateResponse struct {
	ID    string `json:"id""`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	ID    string `json:"id""`
	Email string `json:"email"`
	Token string `json:"token"`
}
