package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	ApplicationExchange   = "application_events"
	NotificationQueue     = "notification_queue"
	ApplicationCreatedKey = "application.created"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

var rabbit *RabbitMQ

type ApplicationMessage struct {
	ID         string `json:"id"`
	Text       string `json:"text"`
	UserName   string `json:"user_name"`
	UnitName   string `json:"unit_name"`
	CreateTime string `json:"create_time"`
}

func NewRabbitMQ() (*RabbitMQ, error) {
	// Get connection URI from environment
	uri := os.Getenv("RABBITMQ_URI")
	if uri == "" {
		// Default for local development
		uri = "amqp://user:password@localhost:5672/"
	}

	// Connect to RabbitMQ
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	// Initialize exchange and queues
	err = setupExchangesAndQueues(ch)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to setup exchanges and queues: %v", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

// setupExchangesAndQueues creates the exchanges and queues needed for the application
func setupExchangesAndQueues(ch *amqp.Channel) error {
	// Declare the application exchange
	err := ch.ExchangeDeclare(
		ApplicationExchange, // name
		"topic",             // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange: %v", err)
	}

	// Declare the notification queue
	_, err = ch.QueueDeclare(
		NotificationQueue, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	// Bind the queue to the exchange with the routing key
	err = ch.QueueBind(
		NotificationQueue,     // queue name
		ApplicationCreatedKey, // routing key
		ApplicationExchange,   // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind a queue: %v", err)
	}

	return nil
}

// InitRabbitMQ initializes the global RabbitMQ instance
func InitRabbitMQ() error {
	var err error
	rabbit, err = NewRabbitMQ()
	if err != nil {
		return err
	}

	return nil
}

// CloseRabbitMQ closes the RabbitMQ connection and channel
func CloseRabbitMQ() {
	if rabbit != nil {
		if rabbit.Channel != nil {
			rabbit.Channel.Close()
		}
		if rabbit.Conn != nil {
			rabbit.Conn.Close()
		}
	}
}

// PublishApplicationCreated publishes a message when an application is created
func PublishApplicationCreated(ctx context.Context, msg ApplicationMessage) error {
	if rabbit == nil {
		return fmt.Errorf("RabbitMQ not initialized")
	}

	// Convert message to JSON
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// Publish the message
	err = rabbit.Channel.PublishWithContext(
		ctx,
		ApplicationExchange,   // exchange
		ApplicationCreatedKey, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	log.Printf("Published application created message: %s", string(body))
	return nil
}

// ConsumeApplicationCreated consumes messages from the notification queue
func ConsumeApplicationCreated(ctx context.Context, handler func(ApplicationMessage) error) error {
	if rabbit == nil {
		return fmt.Errorf("RabbitMQ not initialized")
	}

	// Create a consumer
	msgs, err := rabbit.Channel.Consume(
		NotificationQueue, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	// Start a goroutine to consume messages
	go func() {
		for d := range msgs {
			var msg ApplicationMessage
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				d.Nack(false, true) // Nack the message, requeue it
				continue
			}

			// Handle the message
			err = handler(msg)
			if err != nil {
				log.Printf("Error handling message: %v", err)
				d.Nack(false, true) // Nack the message, requeue it
				continue
			}

			// Acknowledge the message
			d.Ack(false)
		}
	}()

	log.Println("Started consuming application created messages")
	return nil
}
