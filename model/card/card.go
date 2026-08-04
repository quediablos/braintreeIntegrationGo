package card

var VaultCardRequest struct {
	PaymentMethodId string `json:"paymentMethodId"`
	BillingAddress  struct {
		Address1   string  `json:"address1"`
		Address2   string  `json:"address2"`
		City       *string `json:"city"`
		Country    string  `json:"country"`
		FirstName  string  `json:"firstName"`
		LastName   string  `json:"lastName"`
		PostalCode *string `json:"postalCode"`
		State      *string `json:"state"`
	} `json:"billingAddress"`
	RelatedAce   int     `json:"relatedAce"`
	Currency     *string `json:"currency"`
	DeviceData   *string `json:"deviceData"`
	CustomerData struct {
		AceId           int `json:"aceId"`
		PersonalDetails struct {
			Email     string  `json:"email"`
			FirstName *string `json:"firstName"`
			LastName  *string `json:"lastName"`
			Company   *string `json:"company"`
			Phone     struct {
				CountryCode     string  `json:"countryCode"`
				PhoneNumber     string  `json:"phoneNumber"`
				FullPhoneNumber *string `json:"fullPhoneNumber"`
			} `json:"phone"`
			UserName *string `json:"userName"`
		} `json:"personalDetails"`
		IpAddress           *string `json:"ipAddress"`
		PspClientId         *string `json:"pspClientId"`
		Phone               *string `json:"phone"`
		PspLegacyCustomerId *string `json:"pspLegacyCustomerId"`
		MopUid              int     `json:"mopUid"`
		PspCustomerId       *string `json:"pspCustomerId"`
		AccountCountry      *string `json:"accountCountry"`
		UserId              *string `json:"userId"`
	} `json:"customerData"`
	MerchantAccountId *string `json:"merchantAccountId"`
}
