package request

type GraphQlQuery[T any] struct {
	Query         string              `json:"query"`
	Variables     GraphQlVariables[T] `json:"variables"`
	OperationName string              `json:"operationName"`
}

type GraphQlVariables[T any] struct {
	Input T `json:"input"`
}

type VaultCardInput struct {
	CustomerId      string                 `json:"customerId"`
	PaymentMethodId string                 `json:"paymentMethodId"`
	BillingAddress  *BillingAddress        `json:"billingAddress"`
	Verification    *VerificationVaultCard `json:"verification"`
}

type BillingAddress struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	CountryCode  string `json:"countryCode"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PostalCode   string `json:"postalCode"`
}

type VerificationVaultCard struct {
	MerchantAccountId string `json:"merchantAccountId"`
}
