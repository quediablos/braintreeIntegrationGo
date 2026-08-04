package main

import (
	"braintreeIntegrationGo/application"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {

	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Printf("Error starting application: %v\n", err)
	}
}
