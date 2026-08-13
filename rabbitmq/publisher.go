package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	VaultCardQueue = "vault-card"
	brokerURL      = "amqp://guest:guest@localhost:5672/"
)

type VaultCardPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewVaultCardPublisher() (*VaultCardPublisher, error) {
	conn, err := amqp.Dial(brokerURL)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: failed to connect: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("rabbitmq: failed to open channel: %w", err)
	}

	_, err = ch.QueueDeclare(
		VaultCardQueue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("rabbitmq: failed to declare queue: %w", err)
	}

	return &VaultCardPublisher{conn: conn, channel: ch}, nil
}

func (p *VaultCardPublisher) Publish(ctx context.Context, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("rabbitmq: failed to marshal payload: %w", err)
	}

	err = p.channel.PublishWithContext(
		ctx,
		"",             // default exchange
		VaultCardQueue, // routing key == queue name
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("rabbitmq: failed to publish message: %w", err)
	}

	fmt.Printf("RabbitMQ: published to queue %q\n", VaultCardQueue)
	return nil
}

func (p *VaultCardPublisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}
