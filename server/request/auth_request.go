package request

import validation "github.com/go-ozzo/ozzo-validation"

type AuthRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required" example:"P@ssw0rd"`
}

func (a *AuthRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(a.FullName, validation.Required),
	)
}
