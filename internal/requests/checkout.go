package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Checkout struct {
	CheckoutID string `param:"checkoutId" validate:"required"`
}

func (c Checkout) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.CheckoutID, validation.Required),
	)
}
