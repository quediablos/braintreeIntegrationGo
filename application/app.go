package application

import (
	"braintreeIntegrationGo/client"
	"braintreeIntegrationGo/rabbitmq"
	"braintreeIntegrationGo/service"
	"braintreeIntegrationGo/validator"
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	router              http.Handler
	braintreeClient     *client.BraintreeClient
	CardValidator       *validator.CardValidator
	CardAdapterService  *service.CardAdapterService
	vaultCardPublisher  *rabbitmq.VaultCardPublisher
	vaultCardConsumer   *rabbitmq.VaultCardConsumer
}

func New() *App {

	braintreeClient := client.NewBraintreeClient()
	CardValidator := validator.NewCardValidator()
	CardAdapterService := service.NewCardAdapterService()

	vaultCardPublisher, err := rabbitmq.NewVaultCardPublisher()
	if err != nil {
		fmt.Printf("Warning: RabbitMQ publisher unavailable: %v\n", err)
	}

	vaultCardConsumer, err := rabbitmq.NewVaultCardConsumer()
	if err != nil {
		fmt.Printf("Warning: RabbitMQ consumer unavailable: %v\n", err)
	}

	app := &App{
		braintreeClient:    braintreeClient,
		CardValidator:      CardValidator,
		CardAdapterService: CardAdapterService,
		vaultCardPublisher: vaultCardPublisher,
		vaultCardConsumer:  vaultCardConsumer,
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

	chConsumer := make(chan error, 1)
	if app.vaultCardConsumer != nil {
		go func() {
			if err := app.vaultCardConsumer.Start(ctx); err != nil {
				chConsumer <- fmt.Errorf("rabbitmq consumer error: %w", err)
			}
			close(chConsumer)
		}()
	} else {
		close(chConsumer)
	}

	select {
	case err := <-chServer:
		return err
	case err := <-chConsumer:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server.Shutdown(timeout)
	}

}
