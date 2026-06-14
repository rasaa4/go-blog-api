package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r RegisterUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username,
			validation.Required,
			validation.Length(3, 50),
		),
		validation.Field(&r.Password,
			validation.Required,
			validation.Length(6, 100),
		),
	)
}
func (r LoginUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
