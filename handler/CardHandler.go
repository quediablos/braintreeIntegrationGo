package handler

import (
	"braintreeIntegrationGo/client"
	"braintreeIntegrationGo/model/card"
	"braintreeIntegrationGo/service"
	"braintreeIntegrationGo/validator"
	"encoding/json"
	"fmt"
	"net/http"
)

type Card struct {
}

var cardValidator = new(validator.CardValidator)
var cardAdapterService = new(service.CardAdapterService)
var braintreeClient = new(client.BraintreeClient)

func (h *Card) Vault(w http.ResponseWriter, r *http.Request) {

	var body card.VaultCardRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO:need to return a json response.
	if _, err := cardValidator.Validate(body); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	fmt.Println("Request body:", body)

	//Adapt to the vendor request.
	vendorReq, err := cardAdapterService.AdaptVaultCardRequest(body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Call the braintree client.
	clientRes, err := braintreeClient.VaultCard(vendorReq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = clientRes

	fmt.Println("VendorRequest:", vendorReq)

}
