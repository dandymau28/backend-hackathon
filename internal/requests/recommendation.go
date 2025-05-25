package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Recommendation struct {
	Prompt string `json:"prompt" validate:"required"`
}

func (r Recommendation) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Prompt, validation.Required),
	)
}

type MerchantDisplay struct {
	MerchantID    string `param:"merchantId" validate:"required"`
	TransactionID string `query:"transactionId" validate:"required"`
}

func (md MerchantDisplay) Validate() error {
	return validation.ValidateStruct(&md,
		validation.Field(&md.MerchantID, validation.Required),
		validation.Field(&md.TransactionID, validation.Required),
	)
}
