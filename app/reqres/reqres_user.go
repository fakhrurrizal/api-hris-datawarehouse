package reqres

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type GlobalUserRequest struct {
	Avatar   string `json:"avatar"`
	Fullname string `json:"fullname" validate:"required"`
	Gender   string `json:"gender"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (request GlobalUserRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Fullname, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
}
