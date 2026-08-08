package test

import (
	"braintreeIntegrationGo/clientmodel/request"
	"braintreeIntegrationGo/clientmodel/response"
	"braintreeIntegrationGo/handler"
	"braintreeIntegrationGo/model/card"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockCardValidator struct {
	mock.Mock
}

type MockCardAdapterService struct {
	mock.Mock
}

type MockClient struct {
	mock.Mock
}

func (m *MockCardValidator) Validate(request card.VaultCardRequest) (bool, error) {
	args := m.Called(request)
	return args.Bool(0), args.Error(1)
}

func (m *MockCardAdapterService) AdaptVaultCardRequest(r card.VaultCardRequest) (request.VaultCardInput, error) {
	args := m.Called(r)
	return args.Get(0).(request.VaultCardInput), args.Error(1)
}

func (m *MockCardAdapterService) AdaptVaultCardResponse(r response.VaultCreditCardResponse) (card.VaultCardResponse, error) {
	args := m.Called(r)
	return args.Get(0).(card.VaultCardResponse), args.Error(1)
}

func (m *MockClient) VaultCard(r request.VaultCardInput) (response.GraphQlResponse[response.VaultCreditCardResponse], error) {
	args := m.Called(r)
	return args.Get(0).(response.GraphQlResponse[response.VaultCreditCardResponse]), args.Error(1)
}

func TestVault_success(t *testing.T) {

	mockValidator := new(MockCardValidator)
	mockValidator.On("Validate", mock.AnythingOfType("card.VaultCardRequest")).
		Return(true, nil)

	mockCardAdapterService := new(MockCardAdapterService)
	mockCardAdapterService.On("AdaptVaultCardRequest", mock.AnythingOfType("card.VaultCardRequest")).
		Return(request.VaultCardInput{}, nil)

	mockCardAdapterService.On("AdaptVaultCardResponse", mock.AnythingOfType("response.VaultCreditCardResponse")).
		Return(card.VaultCardResponse{}, nil)

	mockClient := new(MockClient)
	mockClient.On("VaultCard", mock.AnythingOfType("request.VaultCardInput")).
		Return(response.GraphQlResponse[response.VaultCreditCardResponse]{}, nil)

	cardHandler := handler.CardHandler{
		CardValidator:      mockValidator,
		CardAdapterService: mockCardAdapterService,
		BraintreeClient:    mockClient,
		// ...other fields
	}

	reqBody, _ := json.Marshal(card.VaultCardRequest{
		PaymentMethodId: "fake-valid-nonce",
	})

	//Mocks

	r := httptest.NewRequest(http.MethodPost, "/cards", bytes.NewBuffer(reqBody))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	cardHandler.Vault(w, r)

}
