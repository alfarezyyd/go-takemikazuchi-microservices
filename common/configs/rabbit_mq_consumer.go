package configs

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewConsumer(amqpURL, queueName string) (*RabbitMQConsumer, error) {
	// 1. Koneksi ke RabbitMQ
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// 2. Buka channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	// 3. Deklarasi Queue
	_, err = ch.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return &RabbitMQConsumer{
		conn:    conn,
		channel: ch,
		queue:   queueName,
	}, nil
}

// StartListening membaca pesan dari RabbitMQ
func (c *RabbitMQConsumer) StartListening(processMessage func(body []byte) error) {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,  // Auto-acknowledge
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// 4. Loop untuk membaca pesan
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// 5. Proses pesan
			err := processMessage(d.Body)
			if err != nil {
				log.Printf("Error processing message: %v", err)
			}
		}
	}()
	log.Println("Waiting for messages...")
}
