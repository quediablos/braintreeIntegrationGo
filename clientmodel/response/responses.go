package response

type GraphQlResponse[T any] struct {
	data       T
	errors     []GraphQlError
	extensions GraphQlErrorExtension //TODO: is this needed because extensions already exists in GraphQlError
}

type GraphQlError struct {
	Message    string                  `json:"message"`
	Extensions []GraphQlErrorExtension `json:"extensions"`
}

type GraphQlErrorExtension struct {
	ErrorType  string `json:"errorType"`
	ErrorClass string `json:"errorClass"`
	LegacyCode string `json:"legacyCode"`
}

type VaultCreditCardResponse struct {
	Data struct {
		VaultCreditCard struct {
			PaymentMethod struct {
				Id    string `json:"id"`
				Usage string `json:"usage"`
			} `json:"paymentMethod"`
			Verification struct {
				Id                string `json:"id"`
				Status            string `json:"status"`
				ProcessorResponse struct {
					Message                     string `json:"message"`
					LegacyCode                  string `json:"legacyCode"`
					CvvResponse                 string `json:"cvvResponse"`
					AvsPostalCodeResponse       string `json:"avsPostalCodeResponse"`
					AvsStreetAddressResponse    string `json:"avsStreetAddressResponse"`
					MastercardTransactionLinkId string `json:"mastercardTransactionLinkId"`
				} `json:"processorResponse"`
				RiskData struct {
					Decision        string   `json:"decision"`
					DecisionReasons []string `json:"decisionReasons"`
				} `json:"riskData"`
				PaymentMethodSnapshot struct {
					Typename                string `json:"__typename"`
					NetworkTransactionId    string `json:"networkTransactionId"`
					AccountType             string `json:"accountType"`
					AcquirerReferenceNumber string `json:"acquirerReferenceNumber"`
				} `json:"paymentMethodSnapshot"`
				PaymentMethod struct {
					Id        string `json:"id"`
					LegacyId  string `json:"legacyId"`
					Usage     string `json:"usage"`
					CreatedAt string `json:"createdAt"`
					Details   struct {
						Typename  string `json:"__typename"`
						BrandCode string `json:"brandCode"`
						Last4     string `json:"last4"`
						Bin       string `json:"bin"`
						BinData   struct {
							IssuingBank       string `json:"issuingBank"`
							CountryOfIssuance string `json:"countryOfIssuance"`
						} `json:"binData"`
						ExpirationMonth        string  `json:"expirationMonth"`
						ExpirationYear         string  `json:"expirationYear"`
						CardholderName         *string `json:"cardholderName"`
						UniqueNumberIdentifier string  `json:"uniqueNumberIdentifier"`
						BillingAddress         struct {
							FullName     string `json:"fullName"`
							AddressLine1 string `json:"addressLine1"`
							AdminArea1   string `json:"adminArea1"`
							AdminArea2   string `json:"adminArea2"`
							PostalCode   string `json:"postalCode"`
							CountryCode  string `json:"countryCode"`
						} `json:"billingAddress"`
					} `json:"details"`
				} `json:"paymentMethod"`
			} `json:"verification"`
		} `json:"vaultCreditCard"`
	} `json:"data"`
}
