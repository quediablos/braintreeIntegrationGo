package client

import (
	"braintreeIntegrationGo/clientmodel/response"
	"braintreeIntegrationGo/model/card"
	"net/http"
)

type IntegrationClient struct {
	httpClient *http.Client
}

func (cli *IntegrationClient) VaultCard(request card.VaultCardRequest) (response.GraphQlResponse[response.VaultCreditCardResponse], error) {

	return response.GraphQlResponse[response.VaultCreditCardResponse]{}, nil
}
