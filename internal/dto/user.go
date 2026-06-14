package dto

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
