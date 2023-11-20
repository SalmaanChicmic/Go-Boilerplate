package request

import validation "github.com/go-ozzo/ozzo-validation"

type AwsTextMessagingRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
}

func (a AwsTextMessagingRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.PhoneNumber, validation.Required),
		validation.Field(&a.Message, validation.Required),
	)
}
