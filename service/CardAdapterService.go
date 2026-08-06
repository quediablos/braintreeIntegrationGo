package service

import (
	"braintreeIntegrationGo/clientmodel/request"
	"braintreeIntegrationGo/model/card"
)

type CardAdapterService struct {
}

func (h *CardAdapterService) AdaptVaultCardRequest(r card.VaultCardRequest) (request.VaultCardInput, error) {

	vaultCardInput := request.VaultCardInput{
		PaymentMethodId: r.PaymentMethodId,
		CustomerId:      r.CustomerData.PspCustomerId,
		BillingAddress: &request.BillingAddress{
			AddressLine1: r.BillingAddress.Address1,
			AddressLine2: r.BillingAddress.Address2,
			CountryCode:  r.BillingAddress.Country,
			FirstName:    r.BillingAddress.FirstName,
			LastName:     r.BillingAddress.LastName,
			PostalCode:   r.BillingAddress.PostalCode,
		},
		Verification: &request.VerificationVaultCard{
			MerchantAccountId: r.MerchantAccountId,
		},
	}

	return vaultCardInput, nil
}
