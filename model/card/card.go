package card

type VaultCardResponse struct {
	PaymentMethodId          string   `json:"paymentMethodId"`
	Usage                    string   `json:"usage"`
	AcquirerReference        string   `json:"acquirerReference"`
	AcquirerCode             string   `json:"acquirerCode"`
	CvvResponse              string   `json:"cvvResponse"`
	AvsPostalCodeResponse    string   `json:"avsPostalCodeResponse"`
	AvsStreetAddressResponse string   `json:"avsStreetAddressResponse"`
	CardBin                  string   `json:"cardBin"`
	CardSummary              string   `json:"cardSummary"`
	CardHolderName           string   `json:"cardHolderName"`
	CardIssuingCountry       string   `json:"cardIssuingCountry"`
	CardPaymentMethod        string   `json:"cardPaymentMethod"`
	ExpiryDate               string   `json:"expiryDate"`
	PaymentMethod            string   `json:"paymentMethod"`
	Status                   string   `json:"status"`
	Decision                 string   `json:"decision"`
	DecisionReasons          []string `json:"decisionReasons"`
	FundingSource            string   `json:"fundingSource"`
	CardIssuingBank          string   `json:"cardIssuingBank"`
	FraudScore               int      `json:"fraudScore"`
	UniqueNumberIdentifier   string   `json:"uniqueNumberIdentifier"`
	CardId                   string   `json:"cardId"`
	NormalisedCvvCode        string   `json:"normalisedCvvCode"`
	NormalisedAvsCode        string   `json:"normalisedAvsCode"`
	CardErrorCode            string   `json:"cardErrorCode"`
	PspReferenceId           string   `json:"pspReferenceId"`
	AuthCode                 string   `json:"authCode"`
	PaymentError             string   `json:"paymentError"`
	NetworkTransactionId     string   `json:"networkTransactionId"`
	NormalisedCvvCodeAsText  string   `json:"normalisedCvvCodeAsText"`
	NormalisedAvsCodeAsText  string   `json:"normalisedAvsCodeAsText"`
	TransactionLinkId        string   `json:"transactionLinkId"`
	DecisionReasonsSize      int      `json:"decisionReasonsSize"`
	DecisionReasonsIterator  string   `json:"decisionReasonsIterator"`
}

type VaultCardRequest struct {
	PaymentMethodId string `json:"paymentMethodId"`
	BillingAddress  struct {
		Address1   string `json:"address1"`
		Address2   string `json:"address2"`
		City       string `json:"city"`
		Country    string `json:"country"`
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		PostalCode string `json:"postalCode"`
		State      string `json:"state"`
	} `json:"billingAddress"`
	RelatedAce   int    `json:"relatedAce"`
	Currency     string `json:"currency"`
	DeviceData   string `json:"deviceData"`
	CustomerData struct {
		AceId           int `json:"aceId"`
		PersonalDetails struct {
			Email     string `json:"email"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Company   string `json:"company"`
			Phone     struct {
				CountryCode     string `json:"countryCode"`
				PhoneNumber     string `json:"phoneNumber"`
				FullPhoneNumber string `json:"fullPhoneNumber"`
			} `json:"phone"`
			UserName string `json:"userName"`
		} `json:"personalDetails"`
		IpAddress           string `json:"ipAddress"`
		PspClientId         string `json:"pspClientId"`
		Phone               string `json:"phone"`
		PspLegacyCustomerId string `json:"pspLegacyCustomerId"`
		MopUid              int    `json:"mopUid"`
		PspCustomerId       string `json:"pspCustomerId"`
		AccountCountry      string `json:"accountCountry"`
		UserId              string `json:"userId"`
	} `json:"customerData"`
	MerchantAccountId string `json:"merchantAccountId"`
}
