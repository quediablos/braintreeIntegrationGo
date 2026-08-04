package service

import (
	"braintreeIntegrationGo/model/card"
	"braintreeIntegrationGo/vendormodel"
)

type CardAdapterService struct {
}

func (h *CardAdapterService) AdaptVaultCardRequest(request card.VaultCardRequest) (vendormodel.VaultCardInput, error) {

	vaultCardInput := vendormodel.VaultCardInput{
		PaymentMethodId: request.PaymentMethodId,
		BillingAddress: &vendormodel.BillingAddress{
			AddressLine1: request.BillingAddress.Address1,
			AddressLine2: request.BillingAddress.Address2,
			CountryCode:  request.BillingAddress.Country,
			FirstName:    request.BillingAddress.FirstName,
			LastName:     request.BillingAddress.LastName,
			PostalCode:   request.BillingAddress.PostalCode,
		},
		Verification: &vendormodel.VerificationVaultCard{},
	}

	return vaultCardInput, nil
}
