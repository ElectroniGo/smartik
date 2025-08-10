package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Initializes a new RabbitMQ client connection
func NewRabbitMQClient(uri string) (*amqp.Connection, error) {
	client, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Publishes a message to a specified queue
func PublishMessage(client *amqp.Connection, queueName string, message []byte, messageContentType string) error {
	ch, err := client.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: messageContentType,
			Body:        message,
		})

	if err != nil {
		return err
	}

	return nil
}
