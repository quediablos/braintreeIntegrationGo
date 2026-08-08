package service

import (
	"braintreeIntegrationGo/clientmodel/request"
	"braintreeIntegrationGo/clientmodel/response"
	"braintreeIntegrationGo/model/card"
)

type CardAdapterService struct {
}

func NewCardAdapterService() *CardAdapterService {
	return &CardAdapterService{}
}

type CardAdapterServiceInterface interface {
	AdaptVaultCardRequest(r card.VaultCardRequest) (request.VaultCardInput, error)
	AdaptVaultCardResponse(r response.VaultCreditCardResponse) (card.VaultCardResponse, error)
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

func (h *CardAdapterService) AdaptVaultCardResponse(r response.VaultCreditCardResponse) (card.VaultCardResponse, error) {

	//TODO:error and decline handling.

	v := r.VaultCreditCard
	ver := v.Verification
	details := ver.PaymentMethod.Details
	snapshot := ver.PaymentMethodSnapshot
	proc := ver.ProcessorResponse

	var cardHolderName string
	if details.CardholderName != nil {
		cardHolderName = *details.CardholderName
	}

	var expiryDate string
	if details.ExpirationMonth != "" && details.ExpirationYear != "" {
		expiryDate = details.ExpirationMonth + "/" + details.ExpirationYear
	}

	paymentMethod := details.BrandCode

	adapted := card.VaultCardResponse{
		PaymentMethodId:          v.PaymentMethod.Id,
		CardId:                   ver.PaymentMethod.Id,
		Usage:                    v.PaymentMethod.Usage,
		CardBin:                  details.Bin,
		CardSummary:              details.Last4,
		CardHolderName:           cardHolderName,
		ExpiryDate:               expiryDate,
		CardIssuingCountry:       details.BinData.CountryOfIssuance,
		CardIssuingBank:          details.BinData.IssuingBank,
		PaymentMethod:            paymentMethod,
		CardPaymentMethod:        paymentMethod,
		Status:                   ver.Status,
		FraudScore:               0,
		UniqueNumberIdentifier:   details.UniqueNumberIdentifier,
		FundingSource:            snapshot.AccountType,
		AcquirerReference:        snapshot.AcquirerReferenceNumber,
		AcquirerCode:             "",
		Decision:                 "COMPLETED",
		DecisionReasons:          ver.RiskData.DecisionReasons,
		DecisionReasonsSize:      len(ver.RiskData.DecisionReasons),
		PspReferenceId:           ver.Id,
		AuthCode:                 ver.PaymentMethod.LegacyId,
		CvvResponse:              proc.CvvResponse,
		AvsStreetAddressResponse: proc.AvsStreetAddressResponse,
		AvsPostalCodeResponse:    proc.AvsPostalCodeResponse,
		NormalisedCvvCode:        proc.CvvResponse,
		NormalisedAvsCode:        proc.AvsPostalCodeResponse,
		NormalisedCvvCodeAsText:  proc.CvvResponse,
		NormalisedAvsCodeAsText:  proc.AvsPostalCodeResponse,
		NetworkTransactionId:     snapshot.NetworkTransactionId,
	}

	return adapted, nil

}
