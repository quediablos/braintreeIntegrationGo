package validator

import (
	"braintreeIntegrationGo/model/card"
	"errors"
)

type CardValidator struct{}

func NewCardValidator() *CardValidator {
	return &CardValidator{}
}

type CardValidatorInterface interface {
	Validate(request card.VaultCardRequest) (bool, error)
}

func (h *CardValidator) Validate(request card.VaultCardRequest) (bool, error) {

	if request.PaymentMethodId == "" {
		return false, errors.New("paymentMethodId is required")
	}

	if request.BillingAddress.Address1 == "" {
		return false, errors.New("billingAddress.Address1 is required")
	}

	//TODO:implement the rest.

	return true, nil
}
