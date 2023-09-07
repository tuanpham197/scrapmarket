package queue

import (
	"context"
	"fmt"
	"log"
	"sendo/pkg/utils/response"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp.Channel
var RabbitConn *amqp.Connection

// setup rabbit mq channel
func SetupRabbbitMQConnectionChannel() (*amqp.Connection, *amqp.Channel) {
	fmt.Println("Setup RabbbitMQConnection")
	//dial
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", "guest", "guest", "localhost", "5672")

	conn, err := amqp.Dial(url)

	response.QueueFailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	response.QueueFailOnError(err, "Failed to open a channel")

	RabbitChannel = ch

	return RabbitConn, RabbitChannel

}

func Publish(ch *amqp.Channel, qName, text string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := ch.PublishWithContext(ctx, "", qName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(text),
	})
	response.QueueFailOnError(err, "Failed to publish a message")
}

func Consume(ch *amqp.Channel, qName string) {
	msgs, err := ch.Consume(
		qName, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	response.QueueFailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}
