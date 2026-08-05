package client

import (
	"braintreeIntegrationGo/clientmodel/request"
	"braintreeIntegrationGo/clientmodel/response"
	"net/http"
)

type BraintreeClient struct {
	httpClient *http.Client
}

func (cli *BraintreeClient) VaultCard(r request.VaultCardInput) (response.GraphQlResponse[response.VaultCreditCardResponse], error) {

	graphQlQuery := request.GraphQlQuery[request.VaultCardInput]{
		Query: "",
		GraphQlVariables: request.GraphQlVariables[request.VaultCardInput]{
			Input: r,
		},
		OperationName: "",
	}
	_ = graphQlQuery

	return response.GraphQlResponse[response.VaultCreditCardResponse]{}, nil
}
