package application

import (
	"braintreeIntegrationGo/client"
	"braintreeIntegrationGo/service"
	"braintreeIntegrationGo/validator"
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	router             http.Handler
	braintreeClient    *client.BraintreeClient
	CardValidator      *validator.CardValidator
	CardAdapterService *service.CardAdapterService
}

func New() *App {

	braintreeClient := client.NewBraintreeClient()
	CardValidator := validator.NewCardValidator()
	CardAdapterService := service.NewCardAdapterService()

	app := &App{
		braintreeClient:    braintreeClient,
		CardValidator:      CardValidator,
		CardAdapterService: CardAdapterService,
	}

	app.loadRoutes()

	return app
}

func (app *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":3000",
		Handler: app.router,
	}

	chServer := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			chServer <- fmt.Errorf("start http server error: %w", err)
		}

		close(chServer)

		fmt.Println("started http server")
	}()

	select {
	case err := <-chServer:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server.Shutdown(timeout)
	}

}
