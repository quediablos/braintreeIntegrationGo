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

type CardHandler struct {
	BraintreeClient *client.BraintreeClient
}

// TODO:move these to the construction (just like BraintreeClient field).
var cardValidator = new(validator.CardValidator)
var cardAdapterService = new(service.CardAdapterService)

func (h *CardHandler) Vault(w http.ResponseWriter, r *http.Request) {

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
	clientRes, err := h.BraintreeClient.VaultCard(vendorReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Check for errors.
	if clientRes.Errors != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO:proper error handling
		return
	} else {
		//TODO:adapt the response
	}

	_ = clientRes

	fmt.Println("VendorRequest:", vendorReq)

}
