package handler

import (
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

	fmt.Println("VendorRequest:", vendorReq)

}
