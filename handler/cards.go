package handler

import (
	"braintreeIntegrationGo/model/card"
	"encoding/json"
	"fmt"
	"net/http"
)

type Card struct {
}

func (h *Card) Vault(w http.ResponseWriter, r *http.Request) {

	body := card.VaultCardRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Request body:", body)

}
