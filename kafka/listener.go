package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	ExampleTopic  = "example_topic"
	consumerGroup = "braintree-integration-group"
	brokerAddress = "localhost:9092"
)

type ExampleTopicListener struct {
	reader *kafka.Reader
}

func NewExampleTopicListener() *ExampleTopicListener {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		GroupID: consumerGroup,
		Topic:   ExampleTopic,
	})

	return &ExampleTopicListener{reader: reader}
}

func (l *ExampleTopicListener) Start(ctx context.Context) error {
	fmt.Printf("Kafka listener started on topic: %s\n", ExampleTopic)

	for {
		msg, err := l.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				fmt.Println("Kafka listener shutting down")
				return l.reader.Close()
			}
			fmt.Printf("Kafka read error: %v\n", err)
			continue
		}

		l.handle(msg)
	}
}

func (l *ExampleTopicListener) handle(msg kafka.Message) {
	fmt.Printf("Received message — topic: %s, partition: %d, offset: %d, key: %s, value: %s\n",
		msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
}
