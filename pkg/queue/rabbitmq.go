package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sendo/pkg/utils/response"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp.Channel
var RabbitConn *amqp.Connection

// setup rabbit mq channel
func SetupRabbitMQConnectionChannel() (*amqp.Connection, *amqp.Channel) {

	//dial
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"))

	conn, err := amqp.Dial(url)

	response.QueueFailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	response.QueueFailOnError(err, "Failed to open a channel")

	RabbitChannel = ch

	return RabbitConn, RabbitChannel

}

func Publish(ch *amqp.Channel, qName string, obj interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	serializedObj, err := json.Marshal(obj)
	err = ch.PublishWithContext(ctx, "", qName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        serializedObj,
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
