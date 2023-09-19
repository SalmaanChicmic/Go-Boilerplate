package request

import validation "github.com/go-ozzo/ozzo-validation"

type AuthRequest struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword" binding:"required" example:"11111111"`
	Password    string `json:"password" binding:"required" example:"11111111"`
}

func (a *AuthRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(a.FullName, validation.Required),
	)
}
