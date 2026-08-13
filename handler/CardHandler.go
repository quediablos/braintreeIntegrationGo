package handler

import (
	"braintreeIntegrationGo/client"
	"braintreeIntegrationGo/model/card"
	"braintreeIntegrationGo/rabbitmq"
	"braintreeIntegrationGo/service"
	"braintreeIntegrationGo/validator"
	"encoding/json"
	"fmt"
	"net/http"
)

type CardHandler struct {
	BraintreeClient    client.BraintreeClientInterface
	CardValidator      validator.CardValidatorInterface
	CardAdapterService service.CardAdapterServiceInterface
	VaultCardPublisher *rabbitmq.VaultCardPublisher
}

func (h *CardHandler) Vault(w http.ResponseWriter, r *http.Request) {

	var body card.VaultCardRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO:need to return a json response.
	if _, err := h.CardValidator.Validate(body); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	fmt.Println("Request body:", body)

	//Adapt to the vendor request.
	vendorReq, err := h.CardAdapterService.AdaptVaultCardRequest(body)

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
	}

	responseNormalized, err := h.CardAdapterService.AdaptVaultCardResponse(clientRes.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Push to rabbitmq
	if h.VaultCardPublisher != nil {
		if err := h.VaultCardPublisher.Publish(r.Context(), responseNormalized); err != nil {
			fmt.Printf("Failed to publish vault-card event: %v\n", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseNormalized); err != nil {
		fmt.Printf("Failed to encode response: %v\n", err)
	}

}
