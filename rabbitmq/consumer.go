package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type VaultCardConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewVaultCardConsumer() (*VaultCardConsumer, error) {
	conn, err := amqp.Dial(brokerURL)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq consumer: failed to connect: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("rabbitmq consumer: failed to open channel: %w", err)
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
		return nil, fmt.Errorf("rabbitmq consumer: failed to declare queue: %w", err)
	}

	return &VaultCardConsumer{conn: conn, channel: ch}, nil
}

func (c *VaultCardConsumer) Start(ctx context.Context) error {
	deliveries, err := c.channel.Consume(
		VaultCardQueue,
		"",    // consumer tag — auto-generated
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		return fmt.Errorf("rabbitmq consumer: failed to start consuming: %w", err)
	}

	fmt.Printf("RabbitMQ consumer started on queue %q\n", VaultCardQueue)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("RabbitMQ consumer shutting down")
			return c.Close()
		case msg, ok := <-deliveries:
			if !ok {
				return fmt.Errorf("rabbitmq consumer: delivery channel closed")
			}
			c.handle(msg)
		}
	}
}

func (c *VaultCardConsumer) handle(msg amqp.Delivery) {
	var payload map[string]any
	if err := json.Unmarshal(msg.Body, &payload); err != nil {
		fmt.Printf("RabbitMQ consumer: failed to unmarshal message: %v\n", err)
		msg.Nack(false, false)
		return
	}

	fmt.Printf("RabbitMQ consumer: received from queue %q — %v\n", VaultCardQueue, payload)

	msg.Ack(false)
}

func (c *VaultCardConsumer) Close() error {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}
