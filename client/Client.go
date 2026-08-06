package client

import (
	"braintreeIntegrationGo/clientmodel/request"
	"braintreeIntegrationGo/clientmodel/response"
	"braintreeIntegrationGo/query"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BraintreeClient struct {
	httpClient    *http.Client
	apiUrl        string
	authorization string
}

func NewBraintreeClient() *BraintreeClient {
	return &BraintreeClient{
		httpClient:    &http.Client{},
		apiUrl:        "https://payments.sandbox.braintree-api.com/graphql",
		authorization: "c2ZtZnNmcmY3a3N2d2Rodjo5OThmZmI2N2M5ZGEyNjExZGIxNGU0ZWRiOTA0YTI2YQ==",
	}
}

func (h *BraintreeClient) VaultCard(r request.VaultCardInput) (response.GraphQlResponse[response.VaultCreditCardResponse], error) {

	graphQlQuery := request.GraphQlQuery[request.VaultCardInput]{
		Query: query.QueryVaultCard,
		Variables: request.GraphQlVariables[request.VaultCardInput]{
			Input: r,
		},
		OperationName: "",
	}

	jsonPayload, err := json.Marshal(graphQlQuery)
	if err != nil {
		fmt.Printf("Failed to marshal payload: %v\n", err)
		return response.GraphQlResponse[response.VaultCreditCardResponse]{}, err
	}

	req, err := http.NewRequest(http.MethodPost, h.apiUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Failed to build request: %v\n", err)
		return response.GraphQlResponse[response.VaultCreditCardResponse]{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic: "+h.authorization)
	req.Header.Set("Braintree-Version", time.Now().Format("2006-01-02"))

	braintreeResponse, err := h.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Failed Braintree Call: %v\n", err)
		return response.GraphQlResponse[response.VaultCreditCardResponse]{}, err
	}

	defer braintreeResponse.Body.Close()

	braintreeResponseBody, err := io.ReadAll(braintreeResponse.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)
		return response.GraphQlResponse[response.VaultCreditCardResponse]{}, err
	}

	var graphQlResponse response.GraphQlResponse[response.VaultCreditCardResponse]
	if err := json.Unmarshal(braintreeResponseBody, &graphQlResponse); err != nil {
		fmt.Printf("Failed to unmarshal response: %v\n", err)
		return response.GraphQlResponse[response.VaultCreditCardResponse]{}, err
	}

	return graphQlResponse, nil
}
