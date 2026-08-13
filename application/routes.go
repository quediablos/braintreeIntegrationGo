package application

import (
	"braintreeIntegrationGo/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) loadRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/cards", app.loadOrderRoutes)
	app.router = router
}

func (app *App) loadOrderRoutes(router chi.Router) {
	cardHandler := &handler.CardHandler{
		BraintreeClient:    app.braintreeClient,
		CardValidator:      app.CardValidator,
		CardAdapterService: app.CardAdapterService,
		VaultCardPublisher: app.vaultCardPublisher,
	}

	router.Post("/", cardHandler.Vault)
}
