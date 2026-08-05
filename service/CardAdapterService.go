package service

import (
	"braintreeIntegrationGo/clientmodel/request"
	"braintreeIntegrationGo/model/card"
)

type CardAdapterService struct {
}

func (h *CardAdapterService) AdaptVaultCardRequest(request card.VaultCardRequest) (request.VaultCardInput, error) {

	vaultCardInput := request.VaultCardInput{
		PaymentMethodId: request.PaymentMethodId,
		BillingAddress: &request.BillingAddress{
			AddressLine1: request.BillingAddress.Address1,
			AddressLine2: request.BillingAddress.Address2,
			CountryCode:  request.BillingAddress.Country,
			FirstName:    request.BillingAddress.FirstName,
			LastName:     request.BillingAddress.LastName,
			PostalCode:   request.BillingAddress.PostalCode,
		},
		Verification: &request.VerificationVaultCard{},
	}

	return vaultCardInput, nil
}
