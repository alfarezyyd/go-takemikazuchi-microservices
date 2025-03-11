package configs

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

// RabbitMQ struct
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

// NewRabbitMQ - Inisialisasi koneksi RabbitMQ
func NewRabbitMQ(url, queue string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Deklarasi Queue (jika tidak menggunakan exchange)
	_, err = ch.QueueDeclare(
		queue, // Nama queue
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn, ch, queue}, nil
}

// Publish - Fungsi untuk mengirim pesan ke RabbitMQ
func (r *RabbitMQ) Publish(data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",      // Exchange (kosong jika langsung ke queue)
		r.queue, // Routing key (queue name)
		false,   // Mandatory
		false,   // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	return err
}

// Close - Fungsi untuk menutup koneksi
func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}
