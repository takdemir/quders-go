package rabbitmq

import (
	"context"
	"database/pkg/utils/common"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

var conn *amqp.Connection
var channel *amqp.Channel

func Connect() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_9015"))
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
}

func CreateChannel() {
	channel, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {

		}
	}(channel)
}

func Publish(message string, queueName string, durable bool, autoDelete bool) {
	queue, err := channel.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		false,
		false,
		nil,
	)
	common.FailOnError(err, "Failed to declare "+queueName+" queue on publish")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = channel.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	common.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", message)
}

func Consume(message string, queueName string, durable bool, autoDelete bool, consumerName string) {
	queue, err := channel.QueueDeclare(
		queueName,  // name
		durable,    // durable
		autoDelete, // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	common.FailOnError(err, "Failed to declare "+queueName+" queue on consume")
	messages, err := channel.Consume(
		queue.Name,
		consumerName,
		true,
		false,
		false,
		false,
		nil,
	)
	common.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	common.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", message)
}
